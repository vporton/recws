package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var addr = flag.Int("port", 8081, "http port")

var upgrader = websocket.Upgrader{} // use default options

func home(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", *addr), nil))
}
