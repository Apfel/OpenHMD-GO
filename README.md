[![GoDoc](https://godoc.org/github.com/Apfel/OpenHMD-GO?status.svg)](https://godoc.org/github.com/Apfel/OpenHMD-GO)
[![Go Report Card](https://goreportcard.com/badge/github.com/Apfel/OpenHMD-GO)](https://goreportcard.com/report/github.com/Apfel/OpenHMD-GO)
[![Build Status](https://travis-ci.org/Apfel/OpenHMD-GO.svg?branch=master)](https://travis-ci.org/Apfel/OpenHMD-GO)

# OpenHMD-GO
[OpenHMD](http://www.openhmd.net/) API bindings for [Golang](https://golang.org/).

```
go get github.com/Apfel/OpenHMD-GO
```

#### Note: This module requires OpenHMD. [Click here for help](http://www.openhmd.net/index.php/download/). This module also requires Go 1.12 or higher.

## Examples
### Simple Example
This is [OpenHMD's simple example](https://github.com/OpenHMD/OpenHMD/tree/master/examples/simple), ported to Golang using OpenHMD-GO.

```go
package main

import (
	"log"
	"os"
	"strconv"
	"time"

	openhmd "github.com/Apfel/OpenHMD-GO"
)

func main() {
	log.Printf("OpenHMD-GO - Simple Example")
	var id int

	if len(os.Args) < 2 || os.Args[1] == "" {
		log.Fatalln("Please provide an device ID.")
	} else {
		id, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("Couldn't convert '%s' to an integer.\nError: %s\n", os.Args[1], err.Error())
		}
		log.Printf("Using ID %d.", id)
	}

	context := openhmd.CreateContext()
	if context == nil {
		log.Fatalln("Context couldn't be opened.")
	}

	if count := context.Probe(); count == 0 {
		log.Fatalln("No devices, quitting...")
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

	log.Printf("Device properties:\nResolution: %dx%d\n", width[0], height[0]) // I do know that this is rather poorly designed, but whatever

	for 1 == 1 {
		c, rot := device.GetFloat(openhmd.FloatValueRotationQuat, 4)
		if c != openhmd.StatusCodeOkay {
			log.Fatalf("Rotation - Error code %d\n", c)
		}
		c, pos := device.GetFloat(openhmd.FloatValuePositionVector, 3)
		if c != openhmd.StatusCodeOkay {
			log.Fatalf("Position - Error code %d\n", c)
		}
		log.Printf("Rotation: %f %f %f %f\nPosition: %f %f %f", rot[0], rot[1], rot[2], rot[3], pos[0], pos[1], pos[2])
		time.Sleep(time.Millisecond * 100)
		context.Update()
	}
}
```

### OpenGL Example
This is [OpenHMD's OpenGL example](https://github.com/OpenHMD/OpenHMD/tree/master/examples/opengl), ported to Golang using OpenHMD-GO.

##### Note: this example requires [g3n](http://g3n.rocks/) to work.

```go
package main

import (
	"log"
	"os"
	"strconv"
	"time"

	openhmd "github.com/Apfel/OpenHMD-GO"

	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/util/application"
)

func checkcode(code openhmd.StatusCode) {
	switch code {
	case openhmd.StatusCodeOkay:
		return
	case openhmd.StatusCodeInvalidOperation:
		log.Fatalln("InvalidOperation")
	case openhmd.StatusCodeInvalidParameter:
		log.Fatalln("InvalidParameter")
	case openhmd.StatusCodeUnknownError:
		log.Fatalln("UnknownError")
	case openhmd.StatusCodeUnsupported:
		log.Fatalln("Unsupported")
	case openhmd.StatusCodeUserReserved:
		log.Fatalln("UserReserved")
	}
}

func main() {
	log.Printf("OpenHMD-GO - OpenGL Example")
	var id int

	if len(os.Args) < 2 || os.Args[1] == "" {
		log.Fatalln("Please provide an device ID.")
	} else {
		id, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("Couldn't convert '%s' to an integer.\nError: %s\n", os.Args[1], err.Error())
		}
		log.Printf("Using ID %d.", id)
	}

	context := openhmd.CreateContext()
	if context == nil {
		log.Fatalln("Context couldn't be opened.")
	}

	if count := context.Probe(); count == 0 {
		log.Fatalln("No devices, quitting...")
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

	code, width := device.GetInt(openhmd.IntValueScreenHorizontalResolution, 1)
	checkcode(code)

	code, height := device.GetInt(openhmd.IntValueScreenVerticalResolution, 1)
	checkcode(code)

	app, err := application.Create(application.Options{
		Title:  "OpenHMD - OpenGL example",
		Width:  int(width[0]),
		Height: int(height[0]),
	})

	if err != nil {
		log.Fatalf("App error: %s\n", err)
	}

	var (
		geom         = geometry.NewTorus(1, .4, 12, 32, math32.Pi*2)
		torusMesh1   = graphic.NewMesh(geom, material.NewPhong(math32.NewColor("Yellow")))
		torusMesh2   = graphic.NewMesh(geom, material.NewPhong(math32.NewColor("Red")))
		torusMesh3   = graphic.NewMesh(geom, material.NewPhong(math32.NewColor("Orange")))
		torusMesh4   = graphic.NewMesh(geom, material.NewPhong(math32.NewColor("DarkBlue")))
		torusMesh5   = graphic.NewMesh(geom, material.NewPhong(math32.NewColor("LightBlue")))
		torusMesh6   = graphic.NewMesh(geom, material.NewPhong(math32.NewColor("DarkGreen")))
		torusMesh7   = graphic.NewMesh(geom, material.NewPhong(math32.NewColor("LightGreen")))
		torusMesh8   = graphic.NewMesh(geom, material.NewPhong(math32.NewColor("Black")))
		torusMesh9   = graphic.NewMesh(geom, material.NewPhong(math32.NewColor("White")))
		torusMesh10  = graphic.NewMesh(geom, material.NewPhong(math32.NewColor("Purple")))
		torusMesh11  = graphic.NewMesh(geom, material.NewPhong(math32.NewColor("Pink")))
		torusMesh12  = graphic.NewMesh(geom, material.NewPhong(math32.NewColor("Gray")))
		lightColor   = &math32.Color{R: 1, G: 1, B: 1}
		ambientLight = light.NewAmbient(lightColor, 0.8)
		pointLight   = light.NewPoint(lightColor, 5.0)
		camera       = app.CameraPersp().Camera
	)

	torusMesh1.SetPosition(5, 0, 5)
	torusMesh2.SetPosition(-5, 0, 5)
	torusMesh3.SetPosition(5, 0, -5)
	torusMesh4.SetPosition(-5, 0, -5)
	torusMesh5.SetPosition(5, 0, 2.5)
	torusMesh6.SetPosition(-5, 0, 2.5)
	torusMesh7.SetPosition(5, 0, -2.5)
	torusMesh8.SetPosition(-5, 0, -2.5)
	torusMesh9.SetPosition(5, 0, 0)
	torusMesh10.SetPosition(-5, 0, 0)
	torusMesh11.SetPosition(0, 0, 5)
	torusMesh12.SetPosition(0, 0, -5)
	pointLight.SetPosition(1, 0, 2)

	app.Scene().Add(torusMesh1)
	app.Scene().Add(torusMesh2)
	app.Scene().Add(torusMesh3)
	app.Scene().Add(torusMesh4)
	app.Scene().Add(torusMesh5)
	app.Scene().Add(torusMesh6)
	app.Scene().Add(torusMesh7)
	app.Scene().Add(torusMesh8)
	app.Scene().Add(torusMesh9)
	app.Scene().Add(torusMesh10)
	app.Scene().Add(torusMesh11)
	app.Scene().Add(torusMesh12)
	app.Scene().Add(ambientLight)
	app.Scene().Add(pointLight)

	camera.SetRotationX(0)
	camera.SetRotationY(0)
	camera.SetRotationZ(0)
	camera.SetPositionX(0)
	camera.SetPositionY(0)
	camera.SetPositionZ(0)

	for {
		context.Update()

		code, rot := device.GetFloat(openhmd.FloatValueRotationQuat, 4)
		checkcode(code)
		code, pos := device.GetFloat(openhmd.FloatValuePositionVector, 3)
		checkcode(code)

		camera.SetRotationX(rot[0])
		camera.SetRotationY(rot[1])
		camera.SetRotationZ(rot[2])

		camera.SetPositionX(pos[0])
		camera.SetPositionY(pos[1])
		camera.SetPositionZ(pos[2])

		app.TimerManager.ProcessTimers()
		rendered, err := app.Renderer().Render(app.Camera())
		if err != nil {
			log.Fatalf("Error: %s\n", err)
		}

		if rendered {
			app.Window().SwapBuffers()
		}

		time.Sleep(1 * time.Millisecond)
	}
}
```
