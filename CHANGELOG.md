Major updates are h1, minors are h2

## 0.7.5

- Documentation updates
- Handler function no longer uses KeySearch() function - Standard practice is now to put .APIKEY in ~/ dir
- Tests added
- More go-simplecheck efficiency updates
- Update to syntax of all major packages structs - No longer all caps for struct fields

## 0.7.3-0.7.4

- Documentation updates
- Handler function checks for error codes, log.Fatalf upon unsuccessful request
- Some minor efficiency updates

## 0.7.2

- Documentation updates
- Increases in efficiency
- Addressing staticcheck suggestions
- (**beta**) Handler function returns Error Code

## 0.7.1

- Documentation updates
- Minor cyclomatic complexity revisions

# v0.7.0

- Error-handling, slightly complex
- Ability to place `.APIKEY` anywhere within Parent (../) directory
- Organizational refresh
- Trade package (60%)
- Refactoring
- Documentation updates