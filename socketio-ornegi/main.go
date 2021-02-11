package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

type Context struct {
	Name string
	ID   int
}

func main() {
	server := socketio.NewServer(nil)

	// contextTest := Context{"Merhaba", 123}

	server.OnConnect("/", func(s socketio.Conn) error {
		log.Println("[SERVER] Connected", s.ID())
		roomName := string("newRoom" + s.ID())
		s.Join(roomName)
		fmt.Println("Rooms:", s.Rooms())
		return nil
	})

	server.OnEvent("/", "message", func(s socketio.Conn, msg string) {
		fmt.Printf("Servera gelen msg -> ID: %v, Msg: %v\n", s.ID(), msg)
		s.Emit("serverMsg", msg+" :)")
		rooms := server.Rooms("/")
		fmt.Println("Rooms -> ", rooms, rooms[1])
		server.BroadcastToRoom("/", "newRoom", "bc", "Client - "+s.ID()+" "+msg)
	})

	// server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
	// 	fmt.Println("notice:", msg)
	// 	s.Emit("reply", "have "+msg)
	// })

	// server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
	// 	s.SetContext(msg)
	// 	return "recv " + msg
	// })

	// server.OnEvent("/", "bye", func(s socketio.Conn) string {
	// 	last := s.Context().(string)
	// 	s.Emit("bye", last)
	// 	s.Close()
	// 	return last
	// })

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Printf("Disconnected -> ID: %v, Reason: %s\n", s.ID(), reason)
	})

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))

	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
