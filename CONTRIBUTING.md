# go-schwab contribution rules
---

## testing

ALL contributions to our libraries must be tested. for go-schwab/trader, running the tests requires configuring the library, which in turn requires setting your environment variables.

follow the instructions in the README quick start to generate your tokens. after you've finished your contribution, before testing, go to `utils.go`, and uncomment lines 41 & 66-67, or the lines denoted "For Testing:".

```
41:
// "github.com/joho/godotenv"
```

```
66-67:
// err := godotenv.Load(findAllEnvFiles()...)
// isErrNil(err)
```

then, run `go test` in the trader directory. if your contribution does not pass, then don't contribute it!

if it does pass, then follow these steps:

## contributing

```
0. fork the repository
1. commit your changes
2. create a pr
```
