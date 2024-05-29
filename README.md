# go-trade
[![Go Reference](https://pkg.go.dev/badge/github.com/samjtro/go-trade.svg)](https://pkg.go.dev/github.com/samjtro/go-trade)
[![Go Report Card](https://goreportcard.com/badge/github.com/samjtro/go-trade)](https://goreportcard.com/report/github.com/samjtro/go-trade)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)

Stable Release: [v0.8.7.1](https://github.com/samjtro/go-trade/tree/stable)

[Roadmap to v1.0.0](https://github.com/samjtro/go-trade/blob/main/TODO.md)

[Changelog](https://github.com/samjtro/go-tradtradee/blob/main/CHANGELOG.md)

Built, maintained by [@samjtro](https://github.com/samjtro)

## what is this?

This project is the successor to [go-tda](https://github.com/samjtro/go-tda), a project I made both as a learning experience, as well as an attempt to provide a counterweight to the primarily Python-based algotrading sphere. Golang is SO much better for algorithmic trading, and this package is an attempt to prove out that case.

This is WIP - I have gotten the Schwab oAuth Flow, and Handler function, functional. I need to migrate the custom structs I created for TDA over to their Schwab counterparts now, and clean up a lot of the garbage that was leftover.

If you want to contribute - go for it! There is no contribution guide, just a simple golden rule: If it ain't broke, don't fix it. If your contribution breaks other functionality in the library, don't contribute it.

## what can i do with this?

### quick start

0. Go to developer.schwab.com, create an account.
1. Create ~/.foo/trade
2. go get github.com/samjtro/go-trade
3. Follow the instructions to generate authcode, access tokens.
4. $$$

### package details

#### schwab/data
#### schwab/account
#### utils

### code samples

#### schwab/data

```
candles, err := data.GetCandles(ticker)
handle(err)

quote, err := data.GetQuote(ticker)
handle(err)

priceHistory, err := data.GetPriceHistory(ticker, periodType, period, frequencyType, frequency, startDate, endDate)
handle(err)
```
