.PHONY: run debug

run:
	go run ./cmd/ondict/ -q=doctor 2>/dev/null

easy:
	go run ./cmd/ondict/ -q=doctor -m=true 2>/dev/null

debug:
	go run ./cmd/ondict/ -q=doctor 2>&1

build:
	go build ./cmd/ondict/

install:
	echo "TODO: install"
