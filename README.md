# go-tda

view this project at [pkg.go.dev](https://pkg.go.dev/github.com/samjtro/go-tda)

## why tda?

td-ameritrade has the latest available market data, and the most in depth. if you want the best indicators, you use tda.

## what is this project?

this is a go implementation of a td-ameritrade API hook. the goal of this project is to incentivize others to build algorithmic trading models.

## how can i use this project?

0. first, go to [developer.tdameritrade.com](https://developer.tdameritrade.com/) and register for an account
- optionally, if you'd like to trade using the `trade` package, you will have to register for a td-ameritrade brokerage account
- this is NOT neccesary, as you can use the data from tda to trade on any platform you wish
1. next, use the `go get github.com/samjtro/go-tda/<package_name>` command to import whatever package you'd like to use
- quotes: contains RealTime and PriceHistory; used for getting either a RealTime quote of a ticker, or a long-term PriceHistory dataframe of a stock from tda (most common use-case)
- movers: contains Get; returns a list of movers for the day by index & direction
- option: contains Single & a number of other more advanced strategy spread functions; returns an option chain for your desired ticker, strike, etc.
- instrument: contains Search & Get; returns information on a desired ticker or CUSIP
- trade & account are both unusable as of right now, the rest are obviously very WIP

## what can i do with this project

like previously mentioned, the goal is for you to use this in a wide variety of capacities. see the license, however, it is permissive and there are guidelines for reproduction & proper crediting.
