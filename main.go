package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/poisnoir/spine-go"
)

func main() {

	namespace := flag.String("namespace", "rime", "spine namespace to join")
	name := flag.String("name", "r1-change", "publisher name")
	// port := flag.Int("port", 0, "iphone udp server port. by default it is set to random.")
	key := flag.String("key", "ppap", "spine namespace key")

	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	ns, err := spine.JointNamespace(*namespace, *key, logger)
	if err != nil {
		panic(err)
	}

	pub, err := spine.NewPublisher[[4][4]float64](ns, *name)
	if err != nil {
		panic(err)
	}

	worker, err := NewIphoneWorker(33047)
	if err != nil {
		panic(err)
	}

	fmt.Println("worker initialized at port: ", worker.Address())

	var delta [4][4]float64

	prevState := worker.GetData(context.Background())
	for {
		currentState := worker.GetData(context.Background())
		// currentState.Print()
		currentState.CreateDelta(&prevState, &delta)
		pub.Publish(delta)
		prevState = currentState
	}

}
