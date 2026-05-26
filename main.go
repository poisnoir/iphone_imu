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
	fmt.Println("worker initialized at port: $s", worker.Address())

	var goal [4][4]float64

	for {
		data := worker.GetData()

		goal[0][0] = data.Rot11
		goal[0][1] = data.Rot12
		goal[0][2] = data.Rot13

		goal[1][0] = data.Rot21
		goal[1][1] = data.Rot22
		goal[1][2] = data.Rot23

		goal[2][0] = data.Rot31
		goal[2][1] = data.Rot32
		goal[2][2] = data.Rot33

		goal[0][3] = 0.0 // X translation
		goal[1][3] = 0.0 // Y translation
		goal[2][3] = 0.0 // Z translation

		goal[3][0] = 0.0
		goal[3][1] = 0.0
		goal[3][2] = 0.0
		goal[3][3] = 1.0

		pub.Publish(goal)
	}

}
