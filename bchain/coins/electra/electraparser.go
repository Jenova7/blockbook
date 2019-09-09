package electra

import (
	//"blockbook/bchain"
	"blockbook/bchain/coins/btc"
	//"blockbook/bchain/coins/utils"
	//"bytes"

	"github.com/martinboehm/btcd/wire"
	"github.com/martinboehm/btcutil/chaincfg"
)

//magic numbers
const (
	MainnetMagic wire.BitcoinNet = 0Xb5a2f1b4
)

//chain parameters
var (
	MainNetParams chaincfg.Params
)

func init() {
	MainNetParams = chaincfg.MainNetParams
	MainNetParams.Net = MainnetMagic

	MainNetParams.PubKeyHashAddrID = []byte{33} // 0x21 
	MainNetParams.ScriptHashAddrID = []byte{40} // 0x28 
	MainNetParams.PrivateKeyID = []byte{161} // 0xA1
		
	err := chaincfg.Register(&MainNetParams)
	if err != nil {
		panic(err)
	}
}

// ElectraParser handle
type ElectraParser struct {
	*btc.BitcoinParser
}

// NewElectraParser returns new ElectraParser instance
func NewElectraParser(params *chaincfg.Params, c *btc.Configuration) *ElectraParser {
	return &ElectraParser{BitcoinParser: btc.NewBitcoinParser(params, c)}
}

// GetChainParams contains network parameters
func GetChainParams(chain string) *chaincfg.Params {
	switch chain {
	default:
		return &MainNetParams
	}	
}
