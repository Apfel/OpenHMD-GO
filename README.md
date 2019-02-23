[![GoDoc](https://godoc.org/github.com/Apfel/OpenHMD?status.svg)](https://godoc.org/github.com/Apfel/OpenHMD)
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
	"fmt"
	"time"

	OpenHMD "github.com/Apfel/OpenHMD-GO"
)

func main() {
	ctx := OpenHMD.Create()
	fmt.Printf("Context: %v\n", ctx)

	numDevices := ctx.Probe()
	fmt.Printf("Device count: %d\n", numDevices)

	// Enter the ID for your device here
	id := 0

	dev := ctx.ListOpenDevice(id)
	fmt.Printf("Device Product: %s - Vendor: %s\n", ctx.ListGetString(id, OpenHMD.StringValueProduct), ctx.ListGetString(id, OpenHMD.StringValueVendor))

	// Simple while loop
	for 1 == 1 {
		// Make sure you update your Context as well
		ctx.Update()

		var value int

		// Replace with whatever you want
		if status, value := dev.GetFloat(OpenHMD.FloatValueRotationQuat); status == 0 {
			fmt.Printf("Value: %d\n", value)
		} else {
			fmt.Printf("Error code: %d", status)
		}

		time.Sleep(1 * time.Second)
	}
}
```
