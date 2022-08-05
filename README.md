# go-tda
[![Go Reference](https://pkg.go.dev/badge/github.com/samjtro/go-tda.svg)](https://pkg.go.dev/github.com/samjtro/go-tda)
[![Go Report Card](https://goreportcard.com/badge/github.com/samjtro/go-tda)](https://goreportcard.com/report/github.com/samjtro/go-tda)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)

Latest: [v0.8.1](https://github.com/samjtro/go-tda/tree/main) | Stable*: [v0.8.0](https://github.com/samjtro/go-tda/tree/stable)
- *(Suggested Pre-v1.0.0)

## what is this project?

this is a go implementation of a td-ameritrade API hook. the goal of this project is to incentivize others to build algorithmic trading models on top of these packages. the purpose for this projects existence is really speed; a lot of td-ameritrade hooks are built in Python, which is fine, but they are so unbelievably slow. when you have to wait multiple seconds to make a request, you know it's bad. we have built an incredibly light handler function that routes all calls in the library, making speeds predictably stable. this, on top of well optimized and efficient Marshaling of JSON to custom Structs, means that you can make up to 15 requests per second if you wish.

- the average response times for all functions in this library is 152ms (weighted average of over 100,000 test requests)
- a light function like RealTime can achieve request times as low as 80ms

## how can i use this project?

### quick start

0. go to [developer.tdameritrade.com](https://developer.tdameritrade.com/) and register for an account
- optionally, if you'd like to trade using the `trade` package, you will have to register for a td-ameritrade brokerage account - this is NOT neccesary, as you can use the data from tda to trade on any platform you wish  

1. create an app on td ameritrade developer at [https://developer.tdameritrade.com/user/me/apps](https://developer.tdameritrade.com/user/me/apps), or by going to the My Apps tab while logged in  

2. go to your app (once it has been approved), and get the api key under the "Consumer Key" section  
- create a file called `.APIKEY` and store your api key you generated in that file (with NO extra lines) 
- move that file to your home directory (`~/`)

3. `go get github.com/samjtro/go-tda`

- you're now ready to go! import the library by package ( `github.com/samjtro/go-tda/data` for the data package, for instance )
- if you have any questions, check the [go reference](https://pkg.go.dev/github.com/samjtro/go-tda); or, scroll down to the code samples below

### package details

- `data`: contains `RealTime` and `PriceHistory`; used for getting either a RealTime quote of a ticker, or a long-term PriceHistory dataframe of a stock from tda (most common use-case)
- `movers`: contains `Get`; returns a list of movers for the day by index & direction
- `option`: contains `Single`; returns Option Chains of your desired parameters
- `instrument`: contains `Fundamental` & `Get`; returns information on a desired ticker or CUSIP
- `account` will contain account monitoring and trading functions but is not functional as of right now   

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

fundamental, err = instrument.Fundamental("AAPL")

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

## roadmap to v1.0.0

in order to have a fully functional v1.0.0 release, we must implement:

- [ ] account package functionality, including trading
- [ ] Strategy option spreads
- [x] custom structs for every package for easy navigation
- [x] "pandas dataframe-esque" struct for data package

## what can i do with this project?

like previously mentioned, the goal is for you to use this in a wide variety of capacities. do what you wish with this project, but...  

see the license; it is permissive, there are guidelines for proper reproduction & crediting :) 