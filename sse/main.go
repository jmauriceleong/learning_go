package main

import (
	"fmt"
	"net/http"
	"time"
)

func sseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","text/event-stream")
	w.Header().Set("Cache-Control","no-cache")
	w.Header().Set("Connection","keep-alive")

	w.Header().Set("Access-Control-Allow-Origin","*")

	clientGone := r.Context().Done()

	rc := http.NewResponseController(w)
	t  := time.NewTicker(time.Second)
	defer t.Stop()
	for {
		select{
		case <- clientGone:
			fmt.Println("Client disconnected")
			return
		case <- t.C:
			_, err := fmt.Fprintf(w, "data: The time is %s\n\n", time.Now().Format(time.UnixDate))
			if err!=nil{
				return
			}
			err = rc.Flush()
			if err!=nil{
				return
			}
		}

	}
}

func main(){
	http.HandleFunc("/events",sseHandler)
	fmt.Println("server is running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil{
		fmt.Println(err.Error())
	}
}