package main

import(
	"flag"
)

var port *int = flag.Int("port", 25565, "The port the server will run on.")

func main(){
	flag.Parse()
}