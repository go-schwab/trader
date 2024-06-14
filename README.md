# schwab
[![Go Reference](https://pkg.go.dev/badge/github.com/samjtro/schwab.svg)](https://pkg.go.dev/github.com/samjtro/schwab)
[![Go Report Card](https://goreportcard.com/badge/github.com/samjtro/schwab)](https://goreportcard.com/report/github.com/samjtro/schwab)
[![License](https://img.shields.io/badge/License-GPLv2-green)](LICENSE)

Built, maintained by [@samjtro](https://github.com/samjtro)

## Roadmap to v 1.0.0

- [x] Current Release: v0.0.1: Migrate working functionality to Schwab
    * [x] oAuth Flow (Retrieve, store tokens; refresh)
    * [x] Endpoints
    * [x] Markets and Data API - Full Access (minus option chains)
        * [x] movers
        * [x] data
        * [x] instrument
        * [x] pricehistory
        * [x] realtime
- [ ] v0.5.0: Account and Trading API - Full Access
    * [x] account
        * [x] Account Numbers, Account Info
        * [x] Transaction History
        * [x] User Info & Prefs
    * [ ] trade
        * [ ] Trading
- [ ] v0.9.0: Performance testing
- [ ] v1.0.0: Finish option chains

## What is this?

This project is WIP: The Market Data API is functional, and I am working on Accounts & Trading now.

Go is much faster than Python, has a far more robust standard library, and is compiled. Therefore, it should be the standard for algorithmic trading; yet it is not. This project is my attempt to create a countweight to the primarily Python-based algotrading scene.

Why should you use this library?

1. It's fast. Really fast. Average request time is ~200ms, while the lighter requests average ~140ms.
2. Super easy to setup. Plus, both the Bearer Token & Authentication Codes auto-refresh.

If you want to contribute - go for it! There is no contribution guide, just a simple golden rule: If it ain't broke, don't fix it. All contributions should be tested via `go test` before submission.

## What can i do with this?

### Quick start

0. Go to developer.schwab.com, create an account, create an app, get app credentials from https://developer.schwab.com/dashboard/apps
1. create `config.env` in your project directory, formatted as such:
```
APPKEY=KEY0 // App Key
SECRET=KEY1 // App Secret
CBURL=https://127.0.0.1 // App Callback URL
```
2. `go get github.com/samjtro/schwab@v0.0.1`

### Code samples

#### market-data

```
import (
    "github.com/samjtro/schwab"
)

agent := schwab.Initiate()

df, err := agent.GetPriceHistory("AAPL", "month", "1", "daily", "1", "", "")
check(err)

quote, err := agent.GetQuote("AAPL")
check(err)

simple, err := agent.SearchInstrumentSimple("AAPL")
check(err)

fundamental, err := agent.SearchInstrumentFundamental("AAPL")
check(err)

fmt.Println(df)
fmt.Println(quote)
fmt.Println(simple)
fmt.Println(fundamental)
```

Output:

```
[{1714971600000 78569667 182.354 181.71 184.2 180.42} {1715058000000 77305771 183.45 182.4 184.9 181.32} {1715144400000 45057087 182.85 182.74 183.07 181.45} {1715230800000 48982972 182.56 184.57 184.66 182.11} {1715317200000 50759496 184.9 183.05 185.09 182.13} {1715576400000 72044809 185.435 186.28 187.1 184.62} {1715662800000 52393619 187.51 187.43 188.3 186.29} {1715749200000 70399988 187.91 189.72 190.65 187.37} {1715835600000 52845230 190.47 189.84 191.095 189.6601} {1715922000000 41282925 189.51 189.87 190.81 189.18} {1716181200000 44361275 189.325 191.04 191.9199 189.01} {1716267600000 42309401 191.09 192.35 192.73 190.9201} {1716354000000 34648547 192.265 190.9 192.8231 190.27} {1716440400000 51005924 190.98 186.88 191 186.625} {1716526800000 36326975 188.82 189.98 190.58 188.0404} {1716872400000 52280051 191.51 189.99 193 189.1} {1716958800000 53068016 189.61 190.29 192.247 189.51} {1717045200000 49947941 190.76 191.29 192.18 190.63} {1717131600000 75158277 191.44 192.25 192.57 189.91} {1717390800000 50080539 192.9 194.03 194.99 192.52} {1717477200000 47471445 194.635 194.35 195.32 193.0342} {1717563600000 54156785 195.4 195.87 196.9 194.87}]
{EQUITY COE NBBO true 1973757747 AAPL 199.62 164.075 EDGX 195.75 5 1717687921950 XNAS 195.74 4 1717687920970 195.87 196.5 XADF 195.745 100 195.21 195.745 -0.125 -0.06381784 -0.125 -0.06381784 0 0 0 1717687921950 Normal 14237020 1717687921574}
{037833100 AAPL Apple Inc NASDAQ EQUITY}
{AAPL 037833100 Apple Inc NASDAQ EQUITY 199.62 164.075 0.51454 1 2024-05-10 00:00:00.0 30.51106 96.708 35.44905 6.98872 23.01545 45.5858 0 26.3058 46.578 0 26.0443 147.2497 22.0738 47.00102 0.87464 1.0371 0 51.3642 140.9682 123.7714 6.41964 9.1372 0 0 -2.8005 -0.9016 0 1.5334082e+16 0 2.9801788367e+12 4.83737 0 0 0 0.25 2024-05-16 00:00:00.0 0 0 0 0 57760123 49229712 61209206 2024-05-02 00:00:00.0 4 6.13 47471445 2024-08-16 00:00:00.0 2024-08-12 00:00:00.0 0}
```

#### accounts-trading

account functionality:

```
agent := Initiate()

an, err := agent.GetAccountNumbers()
check(err)
fmt.Println(an)

aca, err := agent.GetAccounts()
check(err)
fmt.Println(aca)

ac, err := agent.GetAccount(an[0].HashValue)
check(err)
fmt.Println(ac)

orders, err := agent.GetAllOrders("2023-06-12T00:00:00.000Z", "2024-06-12T00:00:00.000Z")
check(err)
```

create an order - trading functionality:

```
instrument := InstrumentRef{
	Symbol: "HLTH",
	Type:   "EQUITY",
}
leg0 := OrderLeg{
	Instruction: "BUY",
	Quantity:    1,
	Instrument:  instrument,
}
newMarketOrder := CreateMarketOrder(Session("NORMAL"), Duration("DAY"), Strategy("SINGLE"), Leg(leg0))
err = agent.Submit(an[0].HashValue, newMarketOrder)
check(err)
```
