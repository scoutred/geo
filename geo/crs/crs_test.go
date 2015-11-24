package crs

import (
	"testing"
)

func TestScale(t *testing.T) {
	scale := Scale(2)
	if scale != 1024 {
		t.Error("invalid scaling result")
	}
}

func TestZoom(t *testing.T) {
	zoom := Zoom(2)
	if zoom != -7 {
		t.Error("invalid zoom result")
	}
}
