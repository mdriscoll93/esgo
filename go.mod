module esgo

go 1.19

replace github.com/btcsuite/btcd/btcec/v2 => github.com/decred/dcrd/dcrec/secp256k1/v4 v4.3.0

require (
	github.com/btcsuite/btcd v0.20.1-beta
	github.com/btcsuite/btcutil v1.0.2
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.3.0
)

require (
	github.com/btcsuite/btclog v0.0.0-20170628155309-84c8d2346e9f // indirect
	golang.org/x/crypto v0.0.0-20200115085410-6d4e4cb37c7d // indirect
)
