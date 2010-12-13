package GUI

import(
	"bytes"
	"fmt"
	"http"
	"io"
	"io/ioutil"
//	"os"
	"./State"
	"strconv"
	"websocket"
)
var world chan int
var state *State.State
func Start(ch *chan int, st *State.State){
	world = *ch
	state = st
	go http.ListenAndServe(":25560", http.HandlerFunc(httpServe))
	go http.ListenAndServe(":25561", websocket.Handler(wssServe))
//	os.ForkExec("http://localhost:25560/index.oc", []string{}, []string{}, "", []*os.File{})
}

func httpServe(c http.ResponseWriter, r *http.Request){
	file, err := ioutil.ReadFile("ourcraft/web/"+r.URL.Path[1:])
	//fmt.Println("Incoming request for", r.URL.Path[1:])
	if err == nil{
		io.Copy(c, bytes.NewBuffer(file))
	}else{
		fmt.Println(err)
		fmt.Fprintf(c, "<h1>Page not found.</h1>")
	}
}

func wssServe(ws *websocket.Conn){
	fmt.Println("Socket connect")
	ws.Write([]byte(strconv.Itoa(int(*state))))
	b := make([]byte, 2)
	var quit bool
	for !quit{
		ws.Read(b)
		t, _ := strconv.Atoi(string(b))
		fmt.Println(t)
		switch t{
			case 0: func(){
					fmt.Println("Server Init")
					if(*state > 0){
						ws.Write([]byte(strconv.Itoa(int(*state))))
					}else{
						ws.Write([]byte("0"))
						b = make([]byte, 5)
						ws.Read(b)
						world, _ := strconv.Atob(string(b))
						fmt.Println(world)
						if(world){
						
						}else{
						
						}
					}
				}()
			default: func(){
					fmt.Println(string(b))
					quit = true
				}()
		}
		quit = true
	}
}