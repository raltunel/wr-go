package views

import (
	"github.com/CrocSwap/graphcache-go/cache"
	"github.com/CrocSwap/graphcache-go/loader"
	"github.com/CrocSwap/graphcache-go/types"
)

type IViews interface {
	QueryCreatedContracts(chaninId types.ChainId) CreatedContractsResponse
}

type Views struct {
	Cache   *cache.MemoryCache
	OnChain *loader.OnChainLoader
}

const MAX_POOL_POSITIONS = 100
