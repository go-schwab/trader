# go-tda
[![Go Reference](https://pkg.go.dev/badge/github.com/samjtro/go-tda.svg)](https://pkg.go.dev/github.com/samjtro/go-tda)![Go Report Card](https://goreportcard.com/badge/github.com/samjtro/go-tda) 
 
LATEST WORKING VERSION: v0.5.5

## why tda?

td-ameritrade has the latest & most in-depth available market data. if you want the best indicators, you use tda. as well, tda has the easiest available option chain data. there is almost no other platform that allows you to fetch option chains with as much precision as tda.

## what is this project?

this is a go implementation of a td-ameritrade API hook. the goal of this project is to incentivize others to build algorithmic trading models on top of these packages.

## how can i use this project?

0. go to [developer.tdameritrade.com](https://developer.tdameritrade.com/) and register for an account
- optionally, if you'd like to trade using the `trade` package, you will have to register for a td-ameritrade brokerage account
- this is NOT neccesary, as you can use the data from tda to trade on any platform you wish  

1. create an app on td ameritrade developer at [this link](https://developer.tdameritrade.com/user/me/apps), or by going to the My Apps tab while logged in  

2. go to your app (once it has been approved), and get the api key under the "Consumer Key" section  

3. create a file called `.APIKEY` and store your api key you generated in that file (with NO extra lines)  

4. `go get github.com/samjtro/go-tda` 

5. move the .APIKEY file into the dir you call from (alternatively, you may obviously pull a new branch and use the "../.APIKEY" path rather than the ".APIKEY" path in the handler.go file)   

package details:  

- `data`: contains `RealTime` and `PriceHistory`; used for getting either a RealTime quote of a ticker, or a long-term PriceHistory dataframe of a stock from tda (most common use-case)
- `movers`: contains `Get`; returns a list of movers for the day by index & direction
- `option`: contains `Single`, `Analytical`, `Covered` & `Butterfly`; `Single` returns a Single contract Option Chains, the rest return strategy chains
- `instrument`: contains `Fundamental` & `Get`; returns information on a desired ticker or CUSIP
- `account` are will contain account monitoring and trading functions but is not functional as of right now, the rest of the packages are still obviously very WIP   

read the documentation for proper function usage, most are straight forward (as described above) but some require some pretty specific input to get working correctly. if you still have a question, or something isn't quite working right, either file an issue or a pull request on the repo OR send me an email @ samjtro@protonmail.com

## roadmap to v1.0.0

by 06-01-2022:

[ ] implement custom structs for each major package
- data,option,movers,instrument,account

[ ] implement trading functionality for account package

[ ] finish up Analytical, Covered & Butterfly spread implementations

## what can i do with this project?

like previously mentioned, the goal is for you to use this in a wide variety of capacities. do what you wish with this project, but...  

see the license; it is permissive, there are guidelines for proper reproduction & crediting :)
