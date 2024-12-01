package main

import (
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"
)

var (
	inFile  string
	outFile string
)

func init() {
	flag.StringVar(&inFile, "in", "", "The path to the input KML file")
	flag.StringVar(&outFile, "out", "", "The path to output the converted GPX file to")
}

func main() {
	flag.Parse()
	if inFile == "" || outFile == "" {
		slog.Error("both input ('-in') and output ('-out') arguments must be specified.")
		os.Exit(1)
	}
	inF, err := os.OpenFile(inFile, os.O_RDONLY, 0744)
	if err != nil {
		slog.Error(fmt.Errorf("error opening input file %q: %w", inFile, err).Error())
		os.Exit(1)
	}
	defer inF.Close()
	outF, err := os.OpenFile(outFile, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0744)
	if err != nil {
		slog.Error(fmt.Errorf("error opening output file %q: %w", outFile, err).Error())
		os.Exit(1)
	}
	defer outF.Close()
	inBytes, err := io.ReadAll(inF)
	if err != nil {
		slog.Error(fmt.Errorf("error reading input file %q: %w", inFile, err).Error())
		os.Exit(1)
	}
	gpxdata, err := Convert(inBytes)
	if err != nil {
		slog.Error(fmt.Errorf("error converting data: %w", err).Error())
		os.Exit(1)
	}
	err = gpxdata.Write(outF)
	if err != nil {
		slog.Error(fmt.Errorf("error writing to output file %q: %w", outFile, err).Error())
		os.Exit(1)
	}
}
