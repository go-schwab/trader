NOTICE: We'll be switching over to GPL2.0 upon the release of v1.0.0.

# go-trade
[![Go Reference](https://pkg.go.dev/badge/github.com/samjtro/go-schwab-traderapi.svg)](https://pkg.go.dev/github.com/samjtro/go-schwab-traderapi)
[![Go Report Card](https://goreportcard.com/badge/github.com/samjtro/go-schwab-traderapi)](https://goreportcard.com/report/github.com/samjtro/go-schwab-traderapi)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)

Stable Release: [v0.8.7.1](https://github.com/samjtro/go-schwab-traderapi/tree/stable)
Built, maintained by [@samjtro](https://github.com/samjtro)

## What is this?

A pure-go-native* hook for the Schwab Trader API. This is WIP, changing very rapidly. ATM, a list of working functionality is provided in the code samples below. I tried to maintain a CHANGELOG, but this was fruitless as I barely have the time to be doing as much work as I already am doing on this project.

This project is the successor to [go-tda](https://github.com/samjtro/go-tda), a project I made both as a learning experience, as well as an attempt to provide a counterweight to the primarily Python-based algotrading sphere. Golang is SO much better for algorithmic trading, and this package is an attempt to prove out that case.

If you want to contribute - go for it! There is no contribution guide, just a simple golden rule: If it ain't broke, don't fix it. If your contribution breaks other functionality in the library, don't contribute it.

*spf13/viper will be gone by v1.0.0

## What can i do with this?

### Quick start

0. Go to developer.schwab.com, create an account
1. go get github.com/samjtro/go-schwab-traderapi
2. Follow the code samples below

### Package details

#### market-data
#### accounts-trading
#### utils

### Code samples

#### market-data

```
import (
    schwab "github.com/samjtro/go-trade"
)

movers, err := schwab.GetMovers("$index")
handle(err)

fundamental, err := schwab.SearchInstrumentFundamental("symbol")
handle(err)
```

#### accounts-trading

## Roadmap to v 1.0.0

- [ ] v0.9.0: Migrate working functionality to Schwab
    * [x] oAuth Flow (Retrieve, store tokens; refresh)
    * [x] Endpoints
    * [ ] Fully Functional: Funcs, Structs, Etc.
        * [x] movers
        * [x] data
        * [x] instrument
        * [x] pricehistory
        * [x] realtime
- [ ] v0.9.5: Account and Trading API
    * [ ] account.go
        * [ ] Transaction History
        * [ ] User Info & Prefs
        * [ ] Watchlist
    * [ ] trade.go
        * [ ] Trading
- [ ] v0.9.7: Performance Testing
- [ ] v0.9.9: Finish test package & integrate CI for new PRs
- [ ] v1.1.0: Finish option

## Copyright notice

Copyright 2022-2024 github.com/samjtro

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
