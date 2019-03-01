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
	"os"
	"strconv"
	"time"

	openhmd "github.com/Apfel/OpenHMD-GO"
)

var id int

func main() {
	log.Printf("OpenHMD-GO - Example")

	if len(os.Args) < 1 || os.Args[0] == "" {
		log.Fatalln("Please provide an device ID.")
	} else {
		id, err := strconv.Atoi(os.Args[0])
		if err != nil {
			log.Fatalf("Couldn't convert '%s' to an integer.\nError: %s\n", os.Args[0], err.Error())
		}
		log.Printf("Using ID %d.", id)
	}

	context := openhmd.Create()
	if context == nil {
		log.Fatalln("Context couldn't be opened.\n")
	}

	if count := context.Probe(); count == 0 {
		log.Fatalln("No devices, quitting...\n")
	} else {
		log.Printf("Found device(s). Device count: %d\n", count)
	}

	device := context.ListOpenDevice(id)
	if device == nil || len(context.GetError()) != 0 {
		log.Fatalf("Device with ID %d couldn't be opened. Error: %s\n", id, context.GetError())
	} else {
		log.Printf("Opened device %s, vendor is %s. ID: %s\n", context.ListGetString(id, openhmd.StringValueProduct),
			context.ListGetString(id, openhmd.StringValueVendor), context.ListGetString(id, openhmd.StringValuePath))
	}

	c, width := device.GetInt(openhmd.IntValueScreenVerticalResolution, 1)
	if c != openhmd.StatusCodeOkay {
		log.Fatalf("Fetching Device width error.\nStatus Code: %d\n\n", c)
	}

	c, height := device.GetInt(openhmd.IntValueScreenHorizontalResolution, 1)
	if c != openhmd.StatusCodeOkay {
		log.Fatalf("Fetching Device height error.\nStatus Code: %d\n", c)
	}

	c, class := device.GetInt(openhmd.IntValueDeviceClass, 1)
	if c != openhmd.StatusCodeOkay {
		log.Fatalf("Fetching Device class error.\nStatus Code: %d\n", c)
	}

	c, flags := device.GetInt(openhmd.IntValueDeviceFlags, 1)
	if c != openhmd.StatusCodeOkay {
		log.Fatalf("Fetching Device flags error.\nStatus Code: %d\n", c)
	}

	log.Printf("Device properties:\nResolution: %dx%d\nDevice class: %d\nDevice flags: %d\n", width, height, class, flags)

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