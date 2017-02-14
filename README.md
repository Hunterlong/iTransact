# iTransact :credit_card: for Go Language
[![Build Status](https://travis-ci.org/hunterlong/iTransact.svg?branch=master)](https://travis-ci.org/hunterlong/iTransact) [![Go Report Card](https://goreportcard.com/badge/github.com/hunterlong/iTransact)](https://goreportcard.com/report/github.com/hunterlong/iTransact)
### iTransact API Connections in Go Language
A simple golang package to run credit card transactions via iTransact. This package is not complete as you can tell below.
This package automatically creates the HMAC SHA1 payload required to use the API.

```go
go get github.com/hunterlong/iTransact
```

```go
go import "github.com/hunterlong/iTransact"

func main() {
    iTransactUsername := "my_username_account_38d9d2xqik9"
    iTransactAPIPass := "IUSHADF87A9AHF"
    iTargetGateway := "00000"

    iTransact.SetAPIInfo(iTransactUsername, iTransactAPIPass, iTargetGateway, "test")
}
```

#### Transaction Requests

:white_check_mark: AuthTransaction

:white_check_mark: BatchClose

:white_check_mark: CreditTransaction

:white_check_mark: PostAuthTransaction

:white_check_mark: TranCredTransaction

:white_check_mark: TranForceTransaction

:white_check_mark: TranRetryTransaction

:white_check_mark: VoidTransaction

:white_check_mark: TranRefundTransaction

#### Recurring Payments
:white_medium_square: AuthTransaction

:white_medium_square: RecurUpdate

:white_medium_square: RecurDetails

