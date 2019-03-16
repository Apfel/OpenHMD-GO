# OpenGL Example
This is [OpenHMD's OpenGL example](https://github.com/OpenHMD/OpenHMD/tree/master/examples/opengl), ported to Golang using OpenHMD-GO.

### Note: this Example requires [g3n](http://g3n.rocks/) to work.

```go
package main

import (
	"runtime"
	"time"

	openhmd "github.com/Apfel/OpenHMD-GO"

	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/util/application"
)

func main() {
	checkcode := func(code openhmd.StatusCode) {
		switch code {
		case openhmd.StatusCodeOkay:
			return
		case openhmd.StatusCodeInvalidOperation:
			panic("InvalidOperation")
		case openhmd.StatusCodeInvalidParameter:
			panic("InvalidParameter")
		case openhmd.StatusCodeUnknownError:
			panic("UnknownError")
		case openhmd.StatusCodeUnsupported:
			panic("Unsupported")
		case openhmd.StatusCodeUserReserved:
			panic("UserReserved")
		}
	}

	runtime.LockOSThread()

	context := openhmd.CreateContext()
	if context == nil {
		panic("No Context")
	}

	if context.Probe() == 0 {
		panic("No Devices")
	}

	device := context.ListOpenDevice(0)
	if device == nil {
		panic("No Devices")
	}

	code, width := device.GetInt(openhmd.IntValueScreenHorizontalResolution, 1)
	checkcode(code)

	code, height := device.GetInt(openhmd.IntValueScreenVerticalResolution, 1)
	checkcode(code)

	app, err := application.Create(application.Options{
		Title:  "OpenHMD",
		Width:  int(width[0]),
		Height: int(height[0]),
	})

	if err != nil {
		panic(err)
	}

	geom := geometry.NewTorus(1, .4, 12, 32, math32.Pi*2)
	mat := material.NewPhong(math32.NewColor("Yellow"))
	torusMesh := graphic.NewMesh(geom, mat)
	app.Scene().Add(torusMesh)

	ambientLight := light.NewAmbient(&math32.Color{1.0, 1.0, 1.0}, 0.8)
	app.Scene().Add(ambientLight)

	pointLight := light.NewPoint(&math32.Color{1, 1, 1}, 5.0)
	pointLight.SetPosition(1, 0, 2)
	app.Scene().Add(pointLight)

	app.CameraPersp().SetPosition(0, 0, 3)

	for {
		context.Update()

		code, rot := device.GetFloat(openhmd.FloatValueRotationQuat, 4)
		checkcode(code)

		app.CameraPersp().SetRotation(rot[0], rot[1], rot[2])
		app.TimerManager.ProcessTimers()
		rendered, err := app.Renderer().Render(app.Camera())
		if err != nil {
			panic(err)
		}

		if rendered {
			app.Window().SwapBuffers()
		}

		time.Sleep(1 * time.Millisecond)
	}
}
```