# go wrapper for schwab's trader-api

[![Go Reference](https://pkg.go.dev/badge/github.com/samjtro/schwab.svg)](https://pkg.go.dev/github.com/samjtro/schwab)
[![Go Report Card](https://goreportcard.com/badge/github.com/samjtro/schwab)](https://goreportcard.com/report/github.com/samjtro/schwab)
[![License](https://img.shields.io/badge/License-GPLv2-green)](LICENSE)

built by: [@samjtro](https://github.com/samjtro)

see: [CONTRIBUTING.md](https://github.com/go-schwab/trader/blob/main/CONTRIBUTING.md)

---

why should you use this project?

- lightning fast
- return structs are easily indexable
- easy to setup, easy to use (personal preference, i know - but trust me!)

## docs

### 0.0 quick start

#### mac & windows

0. go to <https://developer.schwab.com>, create an account, create an app, get app credentials from <https://developer.schwab.com/dashboard/apps>
1. create any file with the `.env` extension in your project directory (can also have multiple, if necessary), formatted as such:

```
APPKEY=KEY0 // App Key
SECRET=KEY1 // App Secret
CBURL=https://127.0.0.1 // App Callback URL
```

2. run the following command in your cwd to generate ssl certs for secure tls transmission of your bearer token:

```
openssl req -x509 -out localhost.crt -keyout localhost.key   -newkey rsa:2048 -nodes -sha256   -subj '/CN=localhost' -extensions EXT -config <( \
printf "[dn]\nCN=localhost\n[req]\ndistinguished_name = dn\n[EXT]\nsubjectAltName=DNS.1:localhost,IP:127.0.0.1\nkeyUsage=digitalSignature\nextendedKeyUsage=serverAuth")
```

3. `go get github.com/go-schwab/trader@v0.9.1`

#### linux

0. go to <https://developer.schwab.com>, create an account, create an app, get app credentials from <https://developer.schwab.com/dashboard/apps>
1. create any file with the `.env` extension in your project directory (can also have multiple, if necessary), formatted as such:

```
APPKEY=KEY0 // App Key
SECRET=KEY1 // App Secret
CBURL=https://127.0.0.1 // App Callback URL
```

2. `go get github.com/go-schwab/trader@v0.9.1`

### 0.1 agent

requests in this library are made through a `Handler()`, facilitated by an `Agent{}`. from here on out, the documentation assumes you have included the following code prior to making any requests:

```
import (
    "github.com/go-schwab/trader"
)

agent := trader.Initiate()
```

### 1.0 accessing market data

#### 1.1.0 price history

code samples:

```go
df, err := agent.GetPriceHistory("AAPL", "month", "1", "daily", "1", "", "")
check(err)
```

return:

`agent.GetPriceHistory` returns a []Candle, of the requested data period & frequency.
[struct details](https://pkg.go.dev/github.com/samjtro/schwab#Candle)

```
[{1714971600000 78569667 182.354 181.71 184.2 180.42} {1715058000000 77305771 183.45 182.4 184.9 181.32} {1715144400000 45057087 182.85 182.74 183.07 181.45} {1715230800000 48982972 182.56 184.57 184.66 182.11} {1715317200000 50759496 184.9 183.05 185.09 182.13} {1715576400000 72044809 185.435 186.28 187.1 184.62} {1715662800000 52393619 187.51 187.43 188.3 186.29} {1715749200000 70399988 187.91 189.72 190.65 187.37} {1715835600000 52845230 190.47 189.84 191.095 189.6601} {1715922000000 41282925 189.51 189.87 190.81 189.18} {1716181200000 44361275 189.325 191.04 191.9199 189.01} {1716267600000 42309401 191.09 192.35 192.73 190.9201} {1716354000000 34648547 192.265 190.9 192.8231 190.27} {1716440400000 51005924 190.98 186.88 191 186.625} {1716526800000 36326975 188.82 189.98 190.58 188.0404} {1716872400000 52280051 191.51 189.99 193 189.1} {1716958800000 53068016 189.61 190.29 192.247 189.51} {1717045200000 49947941 190.76 191.29 192.18 190.63} {1717131600000 75158277 191.44 192.25 192.57 189.91} {1717390800000 50080539 192.9 194.03 194.99 192.52} {1717477200000 47471445 194.635 194.35 195.32 193.0342} {1717563600000 54156785 195.4 195.87 196.9 194.87}]
```

#### 1.2.0 quotes

code samples:

```go
quote, err := agent.GetQuote("AAPL")
check(err)
```

return:

`agent.GetQuote` returns a real-time Quote of the asset's performance.
[struct details](https://pkg.go.dev/github.com/samjtro/schwab#Quote)

```
{EQUITY COE NBBO true 1973757747 AAPL 199.62 164.075 EDGX 195.75 5 1717687921950 XNAS 195.74 4 1717687920970 195.87 196.5 XADF 195.745 100 195.21 195.745 -0.125 -0.06381784 -0.125 -0.06381784 0 0 0 1717687921950 Normal 14237020 1717687921574}
```

#### 1.3.0 instruments

##### 1.3.1 simple

code samples:

```go
simple, err := agent.SearchInstrumentSimple("AAPL")
check(err)
```

return:

`agent.SearchInstrumentSimple` returns a SimpleInstrument for the desired asset.
[struct details](https://pkg.go.dev/github.com/samjtro/schwab#SimpleInstrument)

```
{037833100 AAPL Apple Inc NASDAQ EQUITY}
```

##### 1.3.2 fundamental

code samples:

```go
fundamental, err := agent.SearchInstrumentFundamental("AAPL")
check(err)
```

return:

`agent.SearchInstrumentFundamental` returns a FundamentalInstrument for the desired asset.
[struct details](https://pkg.go.dev/github.com/samjtro/schwab#FundamentalInstrument)

```
{AAPL 037833100 Apple Inc NASDAQ EQUITY 199.62 164.075 0.51454 1 2024-05-10 00:00:00.0 30.51106 96.708 35.44905 6.98872 23.01545 45.5858 0 26.3058 46.578 0 26.0443 147.2497 22.0738 47.00102 0.87464 1.0371 0 51.3642 140.9682 123.7714 6.41964 9.1372 0 0 -2.8005 -0.9016 0 1.5334082e+16 0 2.9801788367e+12 4.83737 0 0 0 0.25 2024-05-16 00:00:00.0 0 0 0 0 57760123 49229712 61209206 2024-05-02 00:00:00.0 4 6.13 47471445 2024-08-16 00:00:00.0 2024-08-12 00:00:00.0 0}
```

#### 1.4.0 movers

code samples:

```go
movers, err := agent.GetMovers("$DJI", "up", "percent")
check(err)
```

return:

`agent.GetMovers` returns a []Screener of the day's top movers.
[struct details](https://pkg.go.dev/github.com/samjtro/schwab#Screener)

```
[{AMZN Amazon.com Inc 2883278 183.2 183.2 58.33 4943302 5450 1} {MRK Merck & Co. Inc. 546846 127.78 127.78 11.06 4943302 5420 1} {PG Procter & Gamble 540105 169.93 169.93 10.93 4943302 8155 1} {INTC INTEL CORP 284653 30.83 30.83 5.76 4943302 3654 1} {AAPL Apple Inc 216203 218.24 218.24 4.37 4943302 6985 1} {MSFT Microsoft Corp 137598 426.73 426.73 2.78 4943302 7008 1} {KO The Coca-Cola Co 90896 66.83 66.83 1.84 4943302 1103 1} {DIS Walt Disney Co 56167 92.14 92.14 1.14 4943302 1884 1} {NKE Nike Inc B 33306 73.55 73.55 0.67 4943302 1209 1} {MMM 3M Co 20242 125.16 125.16 0.41 4943302 632 1}]
```

### 2.0 trading, accessing account data

#### 2.1.0 trading

to submit any trades in this library, one must use your encrypted account id. this as accessed by using the `agent.GetAccountNumbers()` function, which is then passed to the submission function. this is because there are use cases where you might want to change between multiple accounts while trading the same session.

```go
an, err := agent.GetAccountNumbers()
check(err)
```

the rest of the docs assume you want to use the first element of the `[]AccountNumbers` array returned. the encrypted id is stored in the HashValue element of the `AccountNumbers` struct: e.g. `an[0].HashValue`

##### 2.1.1.0 single-leg

suppose we wanted to submit a single-leg market order for the symbol "AAPL". this is as easy as:

```go
err = agent.SubmitSingleLegOrder(an[0].HashValue, CreateSingleLegOrder(OrderType("MARKET"), Session("NORMAL"), Duration("DAY"), Strategy("SINGLE"), Instruction("BUY"), Quantity(1.0), Instrument(SimpleOrderInstrument{
    Symbol:    "AAPL",
    AssetType: "EQUITY",
})))
check(err)
```

let's break this down, although it's fairly straight forward. `CreateSingleLegOrder` returns a `SingleLegOrder`, passed to `agent.SubmitSingleLegOrder` after the hash value of your encrypted id. `CreateSingleOrder` accepts an unknown amount of parameters setting the various elements for the order:

```
OrderType:
Session:
Duration:
Strategy:
Instruction:
Quantity:
Instrument:
```

the default behavior of CreateSingleLegOrder() assumes you are submitting an order with the following parameters:

```
OrderType: MARKET
Session: NORMAL
Duration: DAY
Strategy: SINGLE
```

meaning only `INSTRUCTION`, `QUANTITY` & `INSTRUMENT` are the only required directives. the above example can be simplified thusly:

```go
err = agent.SubmitSingleLegOrder(an[0].HashValue, CreateSingleLegOrder(Instruction("BUY"), Quantity(1.0), Instrument(SimpleOrderInstrument{
    Symbol:    "AAPL",
    AssetType: "EQUITY",
})))
check(err)
```

## WIP: DO NOT CROSS, DANGER DANGER

#### 2.2.0 accessing account data

##### 2.2.1.0

```go
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
