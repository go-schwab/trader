# go-tda

view this project at [pkg.go.dev](https://pkg.go.dev/github.com/samjtro/go-tda)  
LATEST WORKING VERSION: v0.4.7

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

4. use the `go get github.com/samjtro/go-tda/<package_name>` command to import whatever package you'd like to use
- `data`: contains `RealTime` and `PriceHistory`; used for getting either a RealTime quote of a ticker, or a long-term PriceHistory dataframe of a stock from tda (most common use-case)
- `movers`: contains `Get`; returns a list of movers for the day by index & direction
- `option`: contains `Single` & a number of other more advanced strategy spread functions; returns an option chain for your desired ticker, strike, etc.
- `instrument`: contains `Search` & `Get`; returns information on a desired ticker or CUSIP
- `trade` & `account` are both unusable as of right now, the rest are obviously very WIP  

5. use whatever package you wish in whatever file you wish; if you are using package control yourself, the .APIKEY file must be in each directory (alternatively, you may obviously pull a new branch and use the "../.APIKEY" path rather than the ".APIKEY" path in the handler.go file)  

read the documentation for proper function usage, most are straight forward (as described above) but some require some pretty specific input to get working correctly. if you still have a question, or something isn't quite working right, either file an issue or a pull request on the repo OR send me an email @ samjtro@protonmail.com

## what can i do with this project?

like previously mentioned, the goal is for you to use this in a wide variety of capacities. do what you wish with this project, but...  

see the license; it is permissive, there are guidelines for proper reproduction & crediting :)
