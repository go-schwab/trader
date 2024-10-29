# contributing to go-schwab/trader

contributing to this project is very easy! development happens on the [dev](https://github.com/go-schwab/trader/tree/dev) branch. the goal is for the main branch to remain essentially unchanged (barring library-breaking behavior) until the next semantic release, to keep things as stable as possible.

---

BEFORE DOING ANY OF THIS, YOU MUST TEST YOUR CHANGES BY RUNNING GO TEST.

IF YOUR CODE DOESNT PASS OUR CI TESTS, YOUR PR WILL NOT BE REVIEWED KINDLY :)

0. create a fork
1. commit your changes
2. create a pull request to the dev branch, preferably with the following description:

```
major | minor:
ref: [issue] #_ | [pr] v_._._
desc:
- ...
```
