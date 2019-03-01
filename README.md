[![GoDoc](https://godoc.org/github.com/Apfel/OpenHMD-GO?status.svg)](https://godoc.org/github.com/Apfel/OpenHMD-GO)
[![Go Report Card](https://goreportcard.com/badge/github.com/Apfel/OpenHMD-GO)](https://goreportcard.com/report/github.com/Apfel/OpenHMD-GO)

# OpenHMD-GO
[OpenHMD](http://www.openhmd.net/) API bindings for [Golang](https://golang.org/).

```
go get github.com/Apfel/OpenHMD-GO
```

#### Note: This requires you to have OpenHMD installed. [Click here for help.](http://www.openhmd.net/index.php/download/)

### Example Code
```go
package main

import (
	"log"
	"time"

	openhmd "github.com/Apfel/OpenHMD-GO"
)

func main() {
	context := openhmd.Create()
	if context == nil {
		log.Fatalln("Context couldn't be opened.\n")
	}

	if count := context.Probe(); count == 0 {
		log.Fatalln("No devices.\n")
	} else {
		log.Printf("Device count: %d\n", count)
	}

	// define your device's ID here
	id := 0

	device := context.ListOpenDevice(id)
	if device == nil || len(context.GetError()) != 0 {
		log.Fatalf("Device couldn't be opened. Error: %s\n", context.GetError())
	} else {
		log.Printf("Opened device %s, vendor is %s. ID: %s\n", context.ListGetString(id, openhmd.StringValueProduct),
			context.ListGetString(id, openhmd.StringValueVendor), context.ListGetString(id, openhmd.StringValuePath))
	}

	c, rot := device.GetFloat(openhmd.FloatValueRotationQuat, 4)
	if c != openhmd.StatusCodeOkay {
		log.Fatalf("Rotation - Error code %d\n", c)
	}

	c, pos := device.GetFloat(openhmd.FloatValuePositionVector, 3)
	if c != openhmd.StatusCodeOkay {
		log.Fatalf("Position - Error code %d\n", c)
	}

	for 1 == 1 {
		log.Printf("Rotation: %f %f %f %f\nPosition: %f %f %f\n", rot[0], rot[1], rot[2], rot[3], pos[0], pos[1], pos[2])
		time.Sleep(time.Millisecond * 100)
	}
}

```
