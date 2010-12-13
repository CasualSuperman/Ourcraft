package State

const(
	OFF = iota
	WORLDINIT
	IDLE
	RUNNING
	ERROR
)

type State int

func NewStateHandler() *State{
	t := new(State)
	*t = OFF
	return t
}