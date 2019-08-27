package electra

import (
	"blockbook/bchain"
	"blockbook/bchain/coins/btc"
	"blockbook/bchain/coins/utils"
	"bytes"

	"github.com/martinboehm/btcd/wire"
	"github.com/martinboehm/btcutil/chaincfg"
)

const (
	MainnetMagic wire.BitcoinNet = 0Xb5a2f1b4
	//TestnetMagic wire.BitcoinNet = 0xe1e2e3e7
)

var (
	MainNetParams chaincfg.Params
	//TestNetParams chaincfg.Params
)

func init() {
	MainNetParams = chaincfg.MainNetParams
	MainNetParams.Net = MainnetMagic
	MainNetParams.PubKeyHashAddrID = []byte{33} // 0x21 
	MainNetParams.ScriptHashAddrID = []byte{40} // 0x28 
	MainNetParams.PrivateKeyID = []byte{161} // 0xA1

	/*TestNetParams = chaincfg.TestNet3Params
	TestNetParams.Net = TestnetMagic
	TestNetParams.PubKeyHashAddrID = []byte{92} // 0x5C	
	TestNetParams.ScriptHashAddrID = []byte{41} // 0x29
	TestNetParams.PrivateKeyID = []byte{220} // 0xDC*/
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
	if !chaincfg.IsRegistered(&MainNetParams) {
		err := chaincfg.Register(&MainNetParams)
		/*if err == nil {
			err = chaincfg.Register(&TestNetParams)
		}*/
		if err != nil {
			panic(err)
		}
	}
	/*switch chain {
	case "test":
		return &TestNetParams
	default:*/
		return &MainNetParams
	//}
}