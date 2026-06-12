package main

import (
	"fmt"
	"math"
	"os"
	"text/tabwriter"
)

type IphoneOutput struct {
	Seq int64 `json:"seq"`

	// Motion
	MotionTime float64 `json:"motionTime"`
	Roll       float64 `json:"roll"`
	Pitch      float64 `json:"pitch"`
	Yaw        float64 `json:"yaw"`

	// Quaternion
	QuatX float64 `json:"quatX"`
	QuatY float64 `json:"quatY"`
	QuatZ float64 `json:"quatZ"`
	QuatW float64 `json:"quatW"`

	// Rotation Matrix
	Rot11 float64 `json:"rot11"`
	Rot12 float64 `json:"rot12"`
	Rot13 float64 `json:"rot13"`
	Rot21 float64 `json:"rot21"`
	Rot22 float64 `json:"rot22"`
	Rot23 float64 `json:"rot23"`
	Rot31 float64 `json:"rot31"`
	Rot32 float64 `json:"rot32"`
	Rot33 float64 `json:"rot33"`

	// Gravity
	GravityX float64 `json:"gravityX"`
	GravityY float64 `json:"gravityY"`
	GravityZ float64 `json:"gravityZ"`

	// Accelerometer
	AccelTime float64 `json:"accelTime"`
	AccelX    float64 `json:"accelX"`
	AccelY    float64 `json:"accelY"`
	AccelZ    float64 `json:"accelZ"`

	// Gyroscope
	GyroTime float64 `json:"gyroTime"`
	GyroX    float64 `json:"gyroX"`
	GyroY    float64 `json:"gyroY"`
	GyroZ    float64 `json:"gyroZ"`

	// Magnetometer
	MagTime float64 `json:"magTime"`
	MagX    float64 `json:"magX"`
	MagY    float64 `json:"magY"`
	MagZ    float64 `json:"magZ"`

	// Location
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (data *IphoneOutput) Print() {
	// Initialize tabwriter to align columns with spaces
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	fmt.Fprintln(w, "METRIC\tVALUE")
	fmt.Fprintln(w, "------\t-----")
	fmt.Fprintf(w, "Sequence\t%d\n", data.Seq)

	fmt.Fprintln(w, "\n[Motion]\t")
	fmt.Fprintf(w, "  Time\t%.4f\n", data.MotionTime)
	fmt.Fprintf(w, "  Roll\t%.4f\n", data.Roll)
	fmt.Fprintf(w, "  Pitch\t%.4f\n", data.Pitch)
	fmt.Fprintf(w, "  Yaw\t%.4f\n", data.Yaw)

	fmt.Fprintln(w, "\n[Quaternion]\t")
	fmt.Fprintf(w, "  X / Y / Z / W\t%.4f, %.4f, %.4f, %.4f\n", data.QuatX, data.QuatY, data.QuatZ, data.QuatW)

	fmt.Fprintln(w, "\n[Rotation Matrix]\t")
	fmt.Fprintf(w, "  Row 1\t[%.3f, %.3f, %.3f]\n", data.Rot11, data.Rot12, data.Rot13)
	fmt.Fprintf(w, "  Row 2\t[%.3f, %.3f, %.3f]\n", data.Rot21, data.Rot22, data.Rot23)
	fmt.Fprintf(w, "  Row 3\t[%.3f, %.3f, %.3f]\n", data.Rot31, data.Rot32, data.Rot33)

	fmt.Fprintln(w, "\n[Gravity]\t")
	fmt.Fprintf(w, "  X / Y / Z\t%.4f, %.4f, %.4f\n", data.GravityX, data.GravityY, data.GravityZ)

	fmt.Fprintln(w, "\n[Accelerometer]\t")
	fmt.Fprintf(w, "  Time\t%.4f\n", data.AccelTime)
	fmt.Fprintf(w, "  X / Y / Z\t%.4f, %.4f, %.4f\n", data.AccelX, data.AccelY, data.AccelZ)

	fmt.Fprintln(w, "\n[Gyroscope]\t")
	fmt.Fprintf(w, "  Time\t%.4f\n", data.GyroTime)
	fmt.Fprintf(w, "  X / Y / Z\t%.4f, %.4f, %.4f\n", data.GyroX, data.GyroY, data.GyroZ)

	fmt.Fprintln(w, "\n[Magnetometer]\t")
	fmt.Fprintf(w, "  Time\t%.4f\n", data.MagTime)
	fmt.Fprintf(w, "  X / Y / Z\t%.4f, %.4f, %.4f\n", data.MagX, data.MagY, data.MagZ)

	fmt.Fprintln(w, "\n[Location]\t")
	fmt.Fprintf(w, "  Latitude\t%.6f\n", data.Latitude)
	fmt.Fprintf(w, "  Longitude\t%.6f\n", data.Longitude)

	// Flush the buffer to print everything out perfectly aligned
	w.Flush()
}

func (io *IphoneOutput) CreateDelta(prevState *IphoneOutput, result *[4][4]float64) {

	result[3][3] = 1

	var deltaYawSin, deltaYawCos float64
	var deltaRollSin, deltaRollCos float64
	var deltaPitchSin, deltaPitchCos float64

	deltaYaw := io.Yaw - prevState.Yaw
	// if deltaYaw < 1 && deltaYaw > -1 {
	// 	deltaYaw = 0
	// }

	deltaRoll := io.Roll - prevState.Roll
	// if deltaRoll < 1 && deltaRoll > -1 {
	// 	deltaRoll = 0
	// }

	deltaPitch := io.Pitch - prevState.Pitch
	// if deltaPitch < 1 && deltaPitch > -1 {
	// 	deltaPitch = 0
	// }

	deltaYawSin, deltaYawCos = math.Sincos(degToRad(deltaYaw))
	deltaRollSin, deltaRollCos = math.Sincos(degToRad(deltaRoll))
	deltaPitchSin, deltaPitchCos = math.Sincos(degToRad(deltaPitch))

	fmt.Println(deltaPitch)
	fmt.Println(deltaRoll)
	fmt.Println(deltaYaw)

	// Rotation angle
	result[0][0] = deltaYawCos * deltaPitchCos
	result[0][1] = deltaYawSin * deltaPitchCos // <- Swapped from [1][0]
	result[0][2] = -deltaPitchSin              // <- Swapped from [2][0]

	// Row 1
	result[1][0] = (deltaYawCos * deltaPitchSin * deltaRollSin) - (deltaYawSin * deltaRollCos) // <- Swapped from [0][1]
	result[1][1] = (deltaYawSin * deltaPitchSin * deltaRollSin) + (deltaYawCos * deltaRollCos)
	result[1][2] = deltaPitchCos * deltaRollSin // <- Swapped from [2][1]

	// Row 2
	result[2][0] = (deltaYawCos * deltaPitchSin * deltaRollCos) + (deltaYawSin * deltaRollSin) // <- Swapped from [0][2]
	result[2][1] = (deltaYawSin * deltaPitchSin * deltaRollCos) - (deltaYawCos * deltaRollSin) // <- Swapped from [1][2]
	result[2][2] = deltaPitchCos * deltaRollCos

}

func degToRad(deg float64) float64 {
	return deg * (math.Pi / 180)
}
