// Copyright 2016 Grant Haywood
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/alexflint/go-arg"
	"github.com/jponge/vertx-go-tcp-eventbus-bridge/eventbus"
	"log"
	"os"
	"fmt"
	"bufio"
	"encoding/json"
)

var conn struct {
  eb	*eventbus.EventBus
  dp	*eventbus.Dispatcher
}

func main() {
        r := bufio.NewReader(os.Stdin)

	var args struct {
	    Connect string	`arg:"-c,help:connect to host:port"`
	    Channel string	`arg:"-n,help:channel name"`
	    Listen bool		`arg:"-l,help:listen"`
	}
	arg.MustParse(&args)
	if args.Connect == "" {
		args.Connect = "localhost:7000"
	}
	var dp,eb = connect(args.Connect)
	conn.dp = dp
	conn.eb = eb
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		line, _, err := r.ReadLine()
		if err != nil {
			log.Fatal("readline failed(?!): ", err)
		}
		if len(line) > 0 {
			send(args.Channel, string(line))
		}
	}
        if args.Listen {
		listen(args.Channel)
	}

}

func send(Channel string, Message string) {
	conn.eb.Send(Channel,  nil, map[string]string{
				"source": "vxc",
				"data":   Message,
			})
}

func listen(Channel string) {
	ch, id, err := conn.dp.Register(Channel, 2)
	if err != nil || len(id) == 0 {
		log.Fatal("Registration failed: ", err)
	}

	for inMsg := range ch {
		out, jerr := json.Marshal(inMsg.Body)
		fmt.Println(string(out))
		if(jerr != nil){
			fmt.Println(jerr)
		}
	}
}

func connect(Connect string) (*eventbus.Dispatcher,*eventbus.EventBus) {
	eventBus, err := eventbus.NewEventBus(Connect)
	if err != nil {
		log.Fatal("Connection to the Vert.x bridge failed: ", err)
	}

	disp := eventbus.NewDispatcher(eventBus)
	disp.Start()
	return disp,eventBus
}
