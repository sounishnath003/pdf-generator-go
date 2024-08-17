
.PHONY: build run install-deps test-pdfgen

install-deps:
	clear

build:
	go build -o dist/gooferr cmd/main.go 

run: build
	dist/gooferr

test-pdfgen:
	rm output.pdf
	/Applications/Google\ Chrome.app/Contents/MacOS/Google\ Chrome --headless --disable-gpu --print-to-pdf https://sounishnath003.github.io/