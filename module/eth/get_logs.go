package eth

import (
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

// Logs requests the logs of the given ethereum.FilterQuery q.
func Logs(q ethereum.FilterQuery) *LogsFactory {
	return &LogsFactory{filterQuery: q}
}

type LogsFactory struct {
	// args
	filterQuery ethereum.FilterQuery

	// returns
	result  []types.Log
	returns *[]types.Log
}

func (f *LogsFactory) Returns(logs *[]types.Log) *LogsFactory {
	f.returns = logs
	return f
}

// CreateRequest implements the core.RequestCreator interface.
func (f *LogsFactory) CreateRequest() (rpc.BatchElem, error) {
	arg, err := toFilterArg(f.filterQuery)
	if err != nil {
		return rpc.BatchElem{}, err
	}

	return rpc.BatchElem{
		Method: "eth_getLogs",
		Args:   []interface{}{arg},
		Result: &f.result,
	}, nil
}

// HandleResponse implements the core.ResponseHandler interface.
func (f *LogsFactory) HandleResponse(elem rpc.BatchElem) error {
	if err := elem.Error; err != nil {
		return err
	}
	*f.returns = f.result
	return nil
}
