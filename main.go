package main

import (
	"flag"
	"log"
	"net/http"
	"socket/conn"
	"socket/push"
)

var addr = flag.String("addr", ":7070", "http service address")

func main() {
	flag.Parse()
	hub := conn.NewHub()
	go hub.Run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn.ServeWs(hub, w, r)
	})
	push:=push.NewPush(hub)//初始化推送实例
	http.HandleFunc("/push/msg",push.TestPush)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}