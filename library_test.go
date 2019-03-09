package openhmd

import "testing"

func TestCreateContext(t *testing.T) {
	if ctx := CreateContext(); ctx == nil {
		t.Fatal("Got nil Context.")
	}
}

func TestProbe(t *testing.T) {
	if ctx := CreateContext(); ctx == nil {
		t.Fatal("Got nil Context.")
	} else if ctx.Probe() == 0 {
		t.Fatal("Got invalid count (0)")
	}
}

func TestCreateSettings(t *testing.T) {
	if ctx := CreateContext(); ctx == nil {
		t.Fatal("Got nil Context.")
	} else if settings := ctx.CreateSettings(); settings == nil {
		t.Fatal("Got nil Settings.")
	}
}

func TestListOpenDevice(t *testing.T) {
	if ctx := CreateContext(); ctx == nil {
		t.Fatal("Got nil Context.")
	} else if ctx.Probe() == 0 {
		t.Fatal("Got invalid count (0)")
	} else if dev := ctx.ListOpenDevice(0); dev == nil {
		t.Fatal("Got nil Device.")
	}
}

func TestListOpenDeviceSettings(t *testing.T) {
	if ctx := CreateContext(); ctx == nil {
		t.Fatal("Got nil Context.")
	} else if settings := ctx.CreateSettings(); settings == nil {
		t.Fatal("Got nil Settings.")
	} else if dev := ctx.ListOpenDeviceSettings(0, settings); dev == nil {
		t.Fatal("Got nil Device.")
	}
}

// TODO: Implement GetFloat/Int and SetFloat/Int as well as SetData
