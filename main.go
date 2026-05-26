package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/poisnoir/spine-go"
)

func main() {

	namespace := flag.String("namespace", "rime", "spine namespace to join")
	name := flag.String("name", "goal", "publisher name")
	port := flag.Int("port", 0, "iphone udp server port. by default it is set to random.")
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

	worker, err := NewIphoneWorker(*port)
	if err != nil {
		panic(err)
	}
	fmt.Println("worker initialized at port: $d", worker.Port())

	var goal [4][4]float64

	for {
		pub.Publish(goal)
		fmt.Println("hello")
	}

}
