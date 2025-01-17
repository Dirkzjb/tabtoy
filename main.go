package main

import (
	"flag"
	"fmt"
	"github.com/Dirkzjb/tabtoy/build"
	"github.com/davyxu/golog"
	"github.com/pkg/profile"
	"os"
)

var log = golog.New("main")
var enableProfile = false

func main() {

	flag.Parse()

	// 版本
	if *paramVersion {
		build.Print()
		return
	}

	switch *paramMode {
	case "v3":

		type stopper interface {
			Stop()
		}

		var s stopper

		if enableProfile {
			s = profile.Start(profile.CPUProfile, profile.ProfilePath("."))
		}

		V3Entry()

		if s != nil {
			s.Stop()
		}
	case "exportorv2", "v2":
		V2Entry()
	case "v2tov3":
		V2ToV3Entry()
	default:
		fmt.Println("--mode not specify")
		os.Exit(1)
	}

}
