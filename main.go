package main

import(
	"flag"
	"./ourcraft/Server"
)

var lowmem *bool = flag.Bool("lowmem", false, "Should the server use aggressive chunk unloading.")
var maxplayers *int = flag.Int("maxplayers", 32, "Maximum number of players.")

func main(){
	flag.Parse()
	server := Server.NewServer(25565, *lowmem, *maxplayers)
	server.Setup()
}