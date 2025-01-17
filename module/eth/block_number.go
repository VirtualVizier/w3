package eth

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
)

// BlockNumber requests the number of the most recent block.
func BlockNumber() *BlockNumberFactory {
	return &BlockNumberFactory{}
}

type BlockNumberFactory struct {
	// returns
	result  hexutil.Big
	returns *big.Int
}

func (f *BlockNumberFactory) Returns(blockNumber *big.Int) *BlockNumberFactory {
	f.returns = blockNumber
	return f
}

// CreateRequest implements the core.RequestCreator interface.
func (f *BlockNumberFactory) CreateRequest() (rpc.BatchElem, error) {
	return rpc.BatchElem{
		Method: "eth_blockNumber",
		Result: &f.result,
	}, nil
}

// HandleResponse implements the core.ResponseHandler interface.
func (f *BlockNumberFactory) HandleResponse(elem rpc.BatchElem) error {
	if err := elem.Error; err != nil {
		return err
	}
	f.returns.Set((*big.Int)(&f.result))
	return nil
}
