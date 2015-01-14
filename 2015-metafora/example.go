package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/lytics/metafora"
	"github.com/lytics/metafora/embedded"
)

func main() {
	// START OMIT
	handlerFunc := metafora.SimpleHandler(func(task string, c <-chan bool) bool {
		select {
		case <-time.After(time.Duration(rand.Intn(10)) * time.Second):
			fmt.Println("\n>>>>>>>>>", task+" completed!\n")
			return true
		case <-c:
			fmt.Println("\n>>>>>>>>>", task+" told to stop\n")
			return false
		}
	})

	coord, client := embedded.NewEmbeddedPair("example-embedded")
	consumer, _ := metafora.NewConsumer(coord, handlerFunc, &metafora.DumbBalancer{})
	go consumer.Run()

	client.SubmitTask("task 1")
	client.SubmitTask("task 2")
	client.SubmitTask("task 3")

	<-time.After(2 * time.Second)
	fmt.Println("\n>>>>>>>>> Times up, shutting down.\n")
	consumer.Shutdown()
	// END OMIT
}
