build: clean
	GOOS=js GOARCH=wasm go build -o ./html/main.wasm ./pkg/main.go
	cp index.html ./html/index.html
	cp wasm_exec.js ./html/wasm_exec.js
	sed -i -e 's;</button>;</button>\n\t<div id=\"target\"></div>;' ./html/index.html

clean:
	rm -f ./html/*

serve:
	go run server.go -p 8000
