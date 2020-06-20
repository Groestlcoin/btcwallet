module github.com/btcsuite/btcwallet/wallet/txrules

go 1.12

require (
	github.com/btcsuite/btcd v0.0.0-20190824003749-130ea5bddde3
	github.com/btcsuite/btcutil v0.0.0-20190425235716-9e5f4b9a998d
)

replace (
	github.com/btcsuite/btcd => github.com/Groestlcoin/grsd v0.20.1-grs
	github.com/btcsuite/btcutil => github.com/Groestlcoin/grsutil v0.5.0-grs3
)
