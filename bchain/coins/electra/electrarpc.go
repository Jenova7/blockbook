package electra

import (
	"blockbook/bchain"
	"blockbook/bchain/coins/btc"
	"encoding/json"

	"github.com/golang/glog"
)

// ElectraRPC is an interface to JSON-RPC bitcoind service.
type ElectraRPC struct {
	*btc.BitcoinRPC
}

// NewElectraRPC returns new ElectraRPC instance.
func NewElectraRPC(config json.RawMessage, pushHandler func(bchain.NotificationType)) (bchain.BlockChain, error) {
	b, err := btc.NewBitcoinRPC(config, pushHandler)
	if err != nil {
		return nil, err
	}

	s := &ElectraRPC{
		b.(*btc.BitcoinRPC),
	}
	s.RPCMarshaler = btc.JSONMarshalerV1{}
	s.ChainConfig.SupportsEstimateSmartFee = true

	return s, nil
}

// Initialize initializes ElectraRPC instance.
func (b *ElectraRPC) Initialize() error {
	ci, err := b.GetChainInfo()
	if err != nil {
		return err
	}
	chainName := ci.Chain

	glog.Info("Chain name ", chainName)
	params := GetChainParams(chainName)

	// always create parser
	b.Parser = NewElectraParser(params, b.ChainConfig)

	// parameters for getInfo request 
	//if params.Net == MainnetMagic {
		b.Testnet = false
		b.Network = "livenet"
	/*} else {
		b.Testnet = true
		b.Network = "testnet"
	}*/

	glog.Info("rpc: block chain ", params.Name)

	return nil
}