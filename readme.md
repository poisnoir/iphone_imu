# iPhone IMU Publisher
This repository establishes a UDP endpoint to receive real-time Inertial Measurement Unit (IMU) data from an iPhone. The received data is subsequently serialized and forwarded to internal services using Spine.

## Spine Configuration
### Publisher
Default Publisher Name: `iphone_imu`
output struct:
```go
	[4][4]float64
````
Data RepresentationThe output is a 4*4 transformation matrix representing the iPhone's spatial orientation and position.
💡 Note: The code calibrates dynamically by setting the very first received data frame as the origin (0,0,0 reference point).



## CommandLine Arguments
namespace
name
port
key
