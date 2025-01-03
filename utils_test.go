package trader_test

import (
	"github.com/go-schwab/trader"
)

var (
	agent = trader.Initiate()
	an, _ = agent.GetAccountNumbers()
)
