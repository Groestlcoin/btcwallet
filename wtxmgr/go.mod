module github.com/btcsuite/btcwallet/wtxmgr

go 1.12

require (
	github.com/btcsuite/btcd v0.0.0-20190824003749-130ea5bddde3
	github.com/btcsuite/btclog v0.0.0-20170628155309-84c8d2346e9f
	github.com/btcsuite/btcutil v0.0.0-20190425235716-9e5f4b9a998d
	github.com/btcsuite/btcwallet/walletdb v1.2.0
)

replace (
	github.com/btcsuite/btcd => github.com/Groestlcoin/grsd v0.20.1-grs
	github.com/btcsuite/btcutil => github.com/Groestlcoin/grsutil v0.5.0-grsd-0-8
	github.com/btcsuite/btcwallet/walletdb => github.com/Groestlcoin/grswallet/walletdb v1.2.0-grs
)
