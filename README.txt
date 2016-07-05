This is a command line client for the Vert.x event bus, and will connect to the TCP event bus bridge

you can pipe a message into it, and it will send it to the channel you name

you can listen on a channel, and it will output the messages line by line

JSON is the ONLY supported format

plain strings are NOT supported




$ ./vxc -h
usage: vxc [--connect CONNECT] [--channel CHANNEL] [--listen]

options:
  --connect CONNECT, -c CONNECT
                         connect to host:port
  --channel CHANNEL, -n CHANNEL
                         channel name
  --listen, -l           listen
