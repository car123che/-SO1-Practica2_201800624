package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "./helloworld"

	"fmt"
	"net/http"

)

const (
	defaultName = "Carlos Ché"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func handler1(w http.ResponseWriter, peticion *http.Request) {

	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	lista_id_juego := peticion.URL.Query()["juego"]
	lista_no_jugadores :=  peticion.URL.Query()["jugadores"]

	id_juego := lista_id_juego[0]
	no_jugadores := lista_no_jugadores[0]

	fmt.Println(id_juego)
	fmt.Println(no_jugadores)
	fmt.Fprintf(w, "Greeting: %s", r.GetMessage())
}


func main() {
	http.HandleFunc("/test", handler1)
	http.ListenAndServe(":8000", nil)
}
