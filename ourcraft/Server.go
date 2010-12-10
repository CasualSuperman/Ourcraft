package Server

import(
	Gui "./Gui"
	"./Player"
	"./StateHandler"
	"./World"
)

type Server struct{
	Port int
	worlds []*World.World
	players []*Player.Player
	lowmem bool
}

func NewServer(port int, lowmem bool, maxplayers int) *Server{
	sh := StateHandler.NewStateHandler()
	world := make(chan int)
	Gui.Start(&world, sh)
	<-world
	return &Server{port,nil,make([]*Player.Player, 32, 0), lowmem}
}

func (s *Server) Setup(){
	s.worlds = make([]*World.World, 1)
}