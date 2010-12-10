package StateHandler

import(

)

type Handler struct{
	State int
}

func NewStateHandler() *Handler{
	return &Handler{0}
}