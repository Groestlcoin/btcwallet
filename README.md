grswallet
=========

grswallet is a daemon handling Groestlcoin wallet functionality for a single
user.  It acts as both an RPC client to
[grsd](https://github.com/Groestlcoin/grsd) and an RPC server for wallet
clients and legacy RPC applications.

grswallet is based on [btcwallet](https://github.com/btcsuite/btcwallet) and
stays very similar to it.

## Requirements

[Go](http://golang.org) 1.11 or newer.

## Installation

- Install Go according to the installation instructions here:
  http://golang.org/doc/install

- Ensure Go was installed properly and is a supported version:

```bash
$ go version
$ go env GOROOT GOPATH
```

NOTE: The `GOROOT` and `GOPATH` above must not be the same path.  It is
recommended that `GOPATH` is set to a directory in your home directory such as
`~/goprojects` to avoid write permission issues.  It is also recommended to add
`$GOPATH/bin` to your `PATH` at this point.

- Run the following commands to obtain grswallet, all dependencies, and install it:

```bash
$ git clone https://github.com/Groestlcoin/grswallet
$ cd grswallet
$ make install
```

NOTE: Do not use `go install` command, because it will install the binary under wrong names.

- grswallet will now be installed in `$GOPATH/bin`.  If you did not already add
  the bin directory to your system path during Go installation, we recommend
  you do so now.

- If you do not have GNU make in Windows just type commands from build section of Makefile:

```
C:\grswallet> go build -o grswallet .
```

`grswallet.exe` can be copied to any directory you like.

## Updating

- Run the following commands to update grswallet, all dependencies, and install it:

```bash
$ cd grswallet
$ git pull
$ make install
```

## Getting Started

The following instructions detail how to get started with grswallet connecting
to a localhost grsd.  Commands should be run in `cmd.exe` or PowerShell on
Windows, or any terminal emulator on unix-like systems.

- Run the following command to start grsd:

```
grsd -u rpcuser -P rpcpass
```

- Run the following command to create a wallet:

```
grswallet -u rpcuser -P rpcpass --create
```

- Run the following command to start grswallet:

```
grswallet -u rpcuser -P rpcpass
```

If everything appears to be working, it is recommended at this point to copy
the sample grsd (sample-grsd.conf) and grswallet (sample-grswallet.conf)
configurations and update with your RPC username and password.

- Use grsctl to control grswallet and grsd:

```
grsctl --wallet getinfo
```

Note how grsctl gains additional powers when you specify `--wallet` flag.

## Ports

| - | Bitcoin mainnet | Groestlcoin mainnet | Bitcoin testnet | Groestlcoin testnet
 ---------------------- | ---- | ---- | ----- | -----
**Wallet/Original RPC** | 8332 | 1441 | 18332 | 17766
**P2P RPC**             | 8333 | 1331 | 18333 | 17777
**btcd/grsd RPC**       | 8334 | 1444 | 18334 | 17764

## License

grswallet is licensed under the liberal ISC License.
