# go-trade
[![Go Reference](https://pkg.go.dev/badge/github.com/samjtro/go-trade.svg)](https://pkg.go.dev/github.com/samjtro/go-trade)
[![Go Report Card](https://goreportcard.com/badge/github.com/samjtro/go-trade)](https://goreportcard.com/report/github.com/samjtro/go-trade)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)

Latest: [v0.8.4](https://github.com/samjtro/go-trade/tree/schwab-migration) | Stable*: [v0.8.3](https://github.com/samjtro/go-trade/tree/stable)
- *(Suggested Pre-v1.0.0)

[Roadmap to v1.0.0](https://github.com/samjtro/go-trade/blob/main/TODO.md)

[Changelog](https://github.com/samjtro/go-tradtradee/blob/main/CHANGELOG.md)

## what is this?

the fastest td-ameritrade hook out there. purpose-built for real-time trading, weighted average request time is 152ms with the RealTime() quote function at 80ms. (per 100k test requests)

built, maintained by [@samjtro](https://github.com/samjtro).

### 2024 update

ameritrade was bought by schwab, who have taken over the developer program. you can register a dev account with schwab at [developer.schwab.com](https://developer.schwab.com). if you have an existing td-ameritrade project, you are unaffected by this change.

after i have completed the schwab integration, the last step before v1.0.0 is laying out a roadmap for integrating with other platforms. the reason for the change in name from 'go-tda' to 'go-trade' is because i want this to be a general purpose go integration for individual developers to hook into any brokerage account for algotrading. plus, i'd like this project to be a starting point for go ai-trading research.

## how can i use this project?

### quick start

0. go to your app (once it has been approved), and get the api key under the "Consumer Key" section  
- create a file called `tda-config.env`
- move that file to your `$HOME` directory (`~/`)
- edit the file and add the following information in the format:

```
APIKEY=Your_APIKEY_Here
UTC_DIFF=+0:00 // This is a placeholder; for MST, you would use -06:00, etc. It is your Difference from UTC time
```

1. `go get github.com/samjtro/go-trade`

- you're now ready to go! import the library by package ( `github.com/samjtro/go-tda/data` for the data package, for instance )
- if you have any questions, check the [go reference](https://pkg.go.dev/github.com/samjtro/go-trade); or, scroll down to the code samples below

### package details

- `data`: contains `RealTime` and `PriceHistory`; used for getting either a RealTime quote of a ticker, or a long-term PriceHistory dataframe of a stock from tda (most common use-case)
- `movers`: contains `Get`; returns a list of movers for the day by index & direction
- `option`: contains `Single`; returns Option Chains of your desired parameters
- `instrument`: contains `Fundamental` & `Get`; returns information on a desired ticker or CUSIP
- WIP: `account` will contain account monitoring and trading functions   

if you still have a question about something after checking the go reference and code samples, or something isn't quite working right, either file an issue or a pull request on the repo OR send me an email @ samjtro@proton.me

### code samples

#### data package

```
quote, err := data.RealTime("AAPL")

if err != nil {
        panic(err)
}

df, err := data.PriceHistory("AAPL", "month", "1", "daily", "1")

if err != nil {
        panic(err)
}
```

#### instrument package

```
simple, err := instrument.Simple("AAPL")

if err != nil {
	panic(err)
}

fundamental, err := instrument.Fundamental("AAPL")

if err != nil {
	panic(err)
}
```

#### movers package

```
movers, err := movers.Get("$DJI", "up", "percent")

if err != nil {
	panic(err)
}
```

#### option package

```
single, err := option.Single("AAPL", "ALL", "ALL", "15", "2022-09-20")

if err != nil {
	panic(err)
}
```

## what can i do with this project?

like previously mentioned, the goal is for you to use this in a wide variety of capacities. do what you wish with this project, but...  

see the license; it is permissive, there are guidelines for proper reproduction & crediting :) 
