# Password Generator
This API has been created for the sole purpose of learning how Go works with echo framework.

## How to use

With the server initialized, you must make a GET request with the parameters lenght and difficulty.


|Key|Description|Example|
|--|--|--|
|lenght|The number of characters you want to use in the password|10|
|difficulty|The difficulty you wish to use in the password, this will vary in the characters that will be added in the password, 1 = easy, 2 = hard, if the difficulty parameter is not added, it will use 1 by default.|2|

Practical example

Input

`https://password-generator-api.fzbian.repl.co/api?lenght=15&difficulty=2`

Output

```json
{
    "password": "A[]y!l5$H7]J]LT"
}
```


## Running

[![Run on Repl.it](https://repl.it/badge/github/fzbian/password-generator-api)](https://repl.it/github/fzbian/password-generator-api)

The code imports the packages `encoding/json`, `fmt`, `math/rand`, `net/http`, `strconv`.

To start the server you only have to start the following command
```
go run main.go
```

### Acknowledgements
Thanks [AndresXLP](https://www.github.com/AndresXLP) for guiding me in my learning process
