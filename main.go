package main

import (
	"flag"
	"fmt"
	"github.com/caarlos0/spin"
	"gopkg.in/cheggaaa/pb.v1"
	"os"
	"path/filepath"
	"time"
)

var appVersion = "2.0.0"

type state struct {
	step  int
	total int
	sent  int
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s: [flags] <firmware file>\n", os.Args[0])
		flag.PrintDefaults()
	}
	version := flag.Bool("version", false, "print the version and exit")
	flag.Parse()

	if *version {
		fmt.Println(fmt.Sprintf("wally-cli v%s", appVersion))
		os.Exit(0)
	}

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(2)
	}

	path := flag.Arg(0)
	extension := filepath.Ext(path)
	if extension != ".bin" && extension != ".hex" {
		fmt.Println("Please provide a valid firmware file: a .hex file (ErgoDox EZ) or a .bin file (Moonlander / Planck EZ)")
		os.Exit(2)
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("The file path you specified does not exist:", path)
		os.Exit(1)
	}

	spinner := spin.New("%s Press the reset button of your keyboard.")
	spinner.Start()
	spinnerStopped := false

	var progress *pb.ProgressBar
	progressStarted := false

	s := state{step: 0, total: 0, sent: 0}
	if extension == ".bin" {
		go dfuFlash(path, &s)
	}
	if extension == ".hex" {
		go teensyFlash(path, &s)
	}

	for s.step != 2 {
		time.Sleep(500 * time.Millisecond)
		if s.step > 0 {
			if spinnerStopped == false {
				spinner.Stop()
				spinnerStopped = true
			}
			if progressStarted == false {
				progressStarted = true
				progress = pb.StartNew(s.total)
			}
			progress.Set(s.sent)
		}
	}
	progress.Finish()
	fmt.Println("Your keyboard was successfully flashed and rebooted. Enjoy the new firmware!")
}
