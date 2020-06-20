build:
	GO111MODULE=on go build -o grswallet .

build-all: build
	GO111MODULE=on go build ./cmd/sweepaccount
	GO111MODULE=on go build ./cmd/dropwtxmgr

install: build
	cp grswallet `go env GOPATH`/bin

reset-mod:
	git checkout go.mod go.sum

test:
	GO111MODULE=on go test ./...
	cd wallet/txauthor; GO111MODULE=on go test ./...
	cd wallet/txrules/; GO111MODULE=on go test ./...
	cd wallet/txsizes/; GO111MODULE=on go test ./...
	cd walletdb; GO111MODULE=on go test ./...
	cd wtxmgr; GO111MODULE=on go test ./...

clean:
	rm -f grswallet btcwallet dropwtxmgr sweepaccount
