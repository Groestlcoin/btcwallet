build:
	GO111MODULE=on go build -o grswallet .

install: build
	cp grswallet `go env GOPATH`/bin

reset-mod:
	git checkout go.mod go.sum

clean:
	rm -f grswallet btcwallet
