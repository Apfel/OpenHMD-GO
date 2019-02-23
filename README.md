# OpenHMD-GO
[OpenHMD](http://www.openhmd.net/) API bindings for [Golang](https://golang.org/).

## Note: This requires you to have OpenHMD installed. [Click here for help.](http://www.openhmd.net/index.php/download/)

```
go get github.com/Apfel/OpenHMD-GO
```

### Example Code
##### This code needs to be updated, so it won't work for now.
```go
package main

import (
	"fmt"

	OpenHMD "github.com/Apfel/OpenHMD-GO"
)

func main() {
	ctx := OpenHMD.Create()
	fmt.Printf("Context: %v\n", ctx)

	numDevices := ctx.Probe()
	fmt.Printf("numDevices (Probe): %d\n", numDevices)

	dev := ctx.ListOpenDevice(0)
	fmt.Printf("Device: %v\n", dev)

	ctx.Update()

	var value int
	rot := dev.GetFloat(1, value)
	fmt.Printf("Float: %d | Value: %d\n", rot, value)
}
```
