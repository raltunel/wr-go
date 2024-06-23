package views

import (
	"github.com/CrocSwap/graphcache-go/types"
)

func (v *Views) QueryUserTokens(chainId types.ChainId, user types.EthAddress) UserTokensResponse {
	resp := UserTokensResponse{
		ChainId: chainId,
		User:    user,
		Block:   v.Cache.LatestBlock(chainId),
		Tokens:  make([]types.EthAddress, 0),
	}

	balances := v.Cache.RetrieveUserBalances(chainId, user)
	resp.Tokens = append(resp.Tokens, balances...)

	return resp
}
