package main

import (
	speedtest "github.com/TobiasFP/TobiasWaageFeldballe_speedtest_go/speedtest"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Module.Get("exports").Set("SpeedTest", speedtest.SpeedTest)
}
