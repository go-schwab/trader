Major updates are h1, minors are h3

# 0.8.0 (WIP)

### 0.7.9 (WIP)

### 0.7.8

- Cross-platform functionality bug fixes; `Handler` works again on Windows
- Testing complete for all working packages
- Updates to `Handler` function efficiency
- Moved tests into package directories

### 0.7.7

- Cross-platform functionality bug fixes; `KeySearch` works again on MacOS
- `TestData` updates
- Every function now returns errors rather than using `log.Fatalf`

### 0.7.6

- Bug fix in `Handler` - `~/` was not working for .APIKEY lookup - Brought back `KeySearch`
- `RealTime` + `PriceHistory` now return errors

### 0.7.5

- Standard practice is now to put .APIKEY in `~/` dir
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