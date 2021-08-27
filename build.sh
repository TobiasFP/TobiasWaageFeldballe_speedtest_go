GOOS=js GOARCH=wasm go build wasm/main.go
mv main test/main.wasm
~/workspace/go/bin/gopherjs build js/main.go
mv main.js test/speedtest.js