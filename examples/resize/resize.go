package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/davidbyttow/govips"
)

func run(inputFile, outputFile string) error {
	in, err := govips.NewImageFromFile(inputFile,
		govips.IntInput("access", int(govips.AccessSequential)))
	if err != nil {
		return err
	}

	interp, err := govips.NewInterpolator("nohalo")
	if err != nil {
		return err
	}

	out := in.Resize(0.2, govips.InterpolatorInput("interpolate", interp))

	out.WriteToFile(outputFile)

	return nil
}

var (
	flagIn  = flag.String("in", "", "file to load")
	flagOut = flag.String("out", "", "file to write out")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "resize -in input_file -out output_file")
	}
	flag.Parse()

	govips.Startup(nil)
	defer govips.Shutdown()

	err := run(*flagIn, *flagOut)
	if err != nil {
		os.Exit(1)
	}
}
