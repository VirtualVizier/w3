package eth

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
)

// GasPrice requests the current gas price in wei.
func GasPrice() *GasPriceFactory {
	return &GasPriceFactory{}
}

type GasPriceFactory struct {
	// returns
	result  hexutil.Big
	returns *big.Int
}

func (f *GasPriceFactory) Returns(gasPrice *big.Int) *GasPriceFactory {
	f.returns = gasPrice
	return f
}

// CreateRequest implements the core.RequestCreator interface.
func (f *GasPriceFactory) CreateRequest() (rpc.BatchElem, error) {
	return rpc.BatchElem{
		Method: "eth_gasPrice",
		Result: &f.result,
	}, nil
}

// HandleResponse implements the core.ResponseHandler interface.
func (f *GasPriceFactory) HandleResponse(elem rpc.BatchElem) error {
	if err := elem.Error; err != nil {
		return err
	}
	f.returns.Set((*big.Int)(&f.result))
	return nil
}
