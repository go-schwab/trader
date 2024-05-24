# go-trade
[![Go Reference](https://pkg.go.dev/badge/github.com/samjtro/go-trade.svg)](https://pkg.go.dev/github.com/samjtro/go-trade)
[![Go Report Card](https://goreportcard.com/badge/github.com/samjtro/go-trade)](https://goreportcard.com/report/github.com/samjtro/go-trade)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)

Stable*: [v0.8.3](https://github.com/samjtro/go-trade/tree/stable)
- *(Suggested Pre-v1.0.0)

[Roadmap to v1.0.0](https://github.com/samjtro/go-trade/blob/main/TODO.md)

[Changelog](https://github.com/samjtro/go-tradtradee/blob/main/CHANGELOG.md)

## what is this?

built, maintained by [@samjtro](https://github.com/samjtro).

## how do i use this?

### initialization

0. go to developer.schwab.com
1. create config.env
2. oAuth flow

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
