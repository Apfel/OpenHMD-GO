# OpenHMD-GO
[OpenHMD](http://www.openhmd.net/) API bindings for [Golang](https://golang.org/).

## Note: This requires you to have OpenHMD installed.

### Example Code
```go
package main

import (
	"fmt"

	OpenHMD "github.com/Apfel/OpenHMD-GO"
)

func main() {
	ctx := OpenHMD.Create()
	fmt.Println("Context: " + ctx)

	numDevices := OpenHMD.Probe(ctx)
	fmt.Println("numDevices (Probe): " + numDevices)

	dev := OpenHMD.ListOpenDevice(ctx, 0)
	fmt.Println("Device: " + dev)

	OpenHMD.Update(ctx)

	rot := OpenHMD.GetFloatDevice(dev, 1)
	for _, v := range rot {
		fmt.Println(f + ", ")
	}
}
```