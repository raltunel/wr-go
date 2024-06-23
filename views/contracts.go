package views

import (
	"github.com/CrocSwap/graphcache-go/types"
)

type CreatedContractsResponse struct {
	ChainId   types.ChainId            `json:"chainId"`
	User      types.EthAddress         `json:"user"`
	Block     int64                    `json:"block"`
	Contracts []types.CreatedContracts `json:"tokens"`
}
