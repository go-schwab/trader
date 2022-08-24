Major updates are h1, minors are h3

###### Main --> Stable (08/24/2022)

### 0.8.3

Changed `option` package to use `float64` instead of `string`, and changed InTheMoney to return a bool now.
- Moved all tests into `test`
- Other minor updates to accompany the above changes

###### Main --> Stable (08/07/2022)

### 0.8.2

**Important update**
Changed all structs dealing with numbers to return `float64` instead of `string`. This is for a number of reasons, but primarily of which is that I was working on [go-algotrader](https://github.com/samjtro/go-algotrader) and realized that it was a pain to convert each string to float64 just to do calculations. Therefore, I am making the decision to migrate this library over to convert those values natively. This should make it much easier to use the library for building models; conversion from `string` to `float64` shouldn't be a requirement to use this library.
- Updated timestamp function to use ISO 8601
- Instead of using .APIKEY, this library now utilizes viper to marshal a `config.env` file located at your `$HOME` (`~`)
- I have updated packages in all structs except `option` from `string` values for numbers to `float64`, as described above

### 0.8.1

Added concurrency to `Handler` function... 
- Centralized `keySearch` & `Handler` functions 
- Minor other updates

# 0.8.0

- Cross-platform bug fixes; `Handler` now works again on Linux/BSD/MacOS

### 0.7.9

- `account` package rewrites, bug fixes & more

### 0.7.8

- Cross-platform bug fixes; `Handler` works again on Windows
- Testing complete for all working packages
- Updates to `Handler` function efficiency
- Moved tests into package directories

### 0.7.7

- Cross-platform bug fixes; `KeySearch` works again on MacOS
- `TestData` updates
- Every function now returns errors rather than using `log.Fatalf`

### 0.7.6

- Bug fix in `Handler` - `~/` was not working for `.APIKEY` lookup - Brought back `KeySearch`
- `RealTime` + `PriceHistory` now return errors

### 0.7.5

- Standard practice is now to put `.APIKEY` in `~/` dir
- Tests added
- More go-simplecheck efficiency updates
- Update to syntax of all major packages structs - No longer all caps for struct fields
- Documentation updates

### 0.7.3-0.7.4

- Handler function returns errors
- Documentation updates
- Increases in efficiency

### 0.7.2

- Documentation updates
- Increases in efficiency
- Addressing staticcheck suggestions
- (**beta**) Handler function returns Error Code

### 0.7.1

- Documentation updates
- Cyclomatic complexity revisions

# v0.7.0

- Error-handling
- Ability to place `.APIKEY` anywhere within Parent (`../`) directory
- Organizational refresh
- Refactoring
- Documentation updates