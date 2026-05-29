package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"math"
	"os"

	"github.com/poisnoir/spine-go"
)

func main() {

	namespace := flag.String("namespace", "rime", "spine namespace to join")
	name := flag.String("name", "goal", "publisher name")
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

	worker, err := NewIphoneWorker(0)
	if err != nil {
		panic(err)
	}

	// fmt.Println(worker)
	fmt.Println("worker initialized at port: $s", worker.Address())

	var goal [4][4]float64

	for {
		data := worker.GetData(context.Background())

		ax := (data.GyroX * data.GyroTime / 1000) / 50000
		ay := (data.GyroY * data.GyroTime / 1000) / 50000
		az := (data.GyroZ * data.GyroTime / 1000) / 50000

		// Rotation angle (magnitude of the angle vector)
		theta := math.Sqrt(ax*ax + ay*ay + az*az)

		// Unit axis of rotation
		ux := ax / theta
		uy := ay / theta
		uz := az / theta

		s := math.Sin(theta)
		c := math.Cos(theta)
		t := 1 - c

		goal[0][0] = t*ux*ux + c
		goal[0][1] = t*ux*uy - s*uz
		goal[0][2] = t*ux*uz + s*uy

		goal[1][0] = t*ux*uy + s*uz
		goal[1][1] = t*uy*uy + c
		goal[1][2] = t*uy*uz - s*ux

		goal[2][0] = t*ux*uz - s*uy
		goal[2][1] = t*uy*uz + s*ux
		goal[2][2] = t*uz*uz + c

		goal[3][0] = 0.0
		goal[3][1] = 0.0
		goal[3][2] = 0.0
		goal[3][3] = 1.0

		goal[0][3] = (data.AccelX * data.AccelTime / 1000) / 50000
		goal[1][3] = (data.AccelY * data.AccelTime / 1000) / 50000
		// goal[2][3] = (data.AccelZ*data.AccelTime/1000 + 86) / 10000

		for _, row := range goal {
			fmt.Printf("[ ")
			for _, val := range row {
				fmt.Printf("%5.1f ", val) // Adjust %5.1f to change spacing/decimals
			}
			fmt.Printf("]\n")
		}

		pub.Publish(goal)
	}

}
