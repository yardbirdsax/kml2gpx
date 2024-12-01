package main

import (
	"bytes"
	_ "embed"
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/twpayne/go-gpx"
)

//go:embed testdata/Cruise.kml
var inByte []byte

//go:embed testdata/Cruise.gpx
var wantByte []byte

func TestConvert(t *testing.T) {
	want := &gpx.GPX{}
	//nolint:errcheck // This is known data and should always decode
	xml.NewDecoder(bytes.NewBuffer(wantByte)).Decode(want)

	got, _ := Convert(inByte)

	assert.Equal(t, want, got, "got and wanted value differ")
}
