Major updates are h1, minors are h3

### 0.7.8

- Cross-platform functionality bug fixes; Handler function works again on Windows
- Testing complete for all working packages
- Updates to Handler function efficiency
- Moved tests into package directories

### 0.7.7

- Cross-platform functionality bug fixes; KeySearch function works again on MacOS
- Added additional test for Data
- Every function now returns errors rather than using log.Fatalf

### 0.7.6

- Bug fix in Handler function - ~ was not working for .APIKEY lookup - Brought back KeySearch
- TestData updates
- RealTime + PriceHistory now return errors

### 0.7.5

- Documentation updates
- Handler function no longer uses KeySearch() function - Standard practice is now to put .APIKEY in ~/ dir
- Tests added
- More go-simplecheck efficiency updates
- Update to syntax of all major packages structs - No longer all caps for struct fields

### 0.7.3-0.7.4

- Documentation updates
- Handler function checks for error codes, log.Fatalf upon unsuccessful request
- Some minor efficiency updates

### 0.7.2

- Documentation updates
- Increases in efficiency
- Addressing staticcheck suggestions
- (**beta**) Handler function returns Error Code

### 0.7.1

- Documentation updates
- Minor cyclomatic complexity revisions

# v0.7.0

- Error-handling, slightly complex
- Ability to place `.APIKEY` anywhere within Parent (../) directory
- Organizational refresh
- Trade package (60%)
- Refactoring
- Documentation updates