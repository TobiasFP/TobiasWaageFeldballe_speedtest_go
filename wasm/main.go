package main

import (
	"syscall/js"

	speedtest "github.com/TobiasFP/TobiasWaageFeldballe_speedtest_go/speedtest"
)

//     GOOS=js GOARCH=wasm go build wasm/main.go
func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("speedtest", js.FuncOf(SpeedTest))
	<-c
}

func SpeedTest(this js.Value, p []js.Value) interface{} {
	res := speedtest.SpeedTest(p[0].Int())
	return js.ValueOf(res)
}
