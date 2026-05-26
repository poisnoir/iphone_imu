package main

type JSONData struct {
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
