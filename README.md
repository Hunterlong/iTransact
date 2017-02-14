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

## Transaction Requests

#### :white_check_mark: AuthTransaction
```go
newTransaction := iTransact.AuthTransaction{
		Total:       "25.98",
		Description: "Order #2384 - Docker Jacket",
		CustomerData: iTransact.CustomerData{
			Email:           "info@socialeck.com",
			CustId:          "83842",
			BillingAddress:  myAddress,
			ShippingAddress: myAddress,
		},
		AccountInfo: iTransact.CardAccount{
		iTransact.CreditCard{
			AccountNumber:   "5454545454545454",
			ExpirationMonth: "05",
			ExpirationYear:  "2022",
			CVVNumber:       "394",
		}},
	}

	response := newTransaction.Charge()

	if response.Approved() {

	}
```

#### :white_check_mark: BatchClose
```go
batch := iTransact.RunBatchClose()
fmt.Println("Closed", batch.Amount(), "Transactions")
```

:white_check_mark: CreditTransaction
```go
refundTransaction := iTransact.CreditTransaction{
		Total:       "25.98",
		CustomerData: iTransact.CustomerData{
			Email:           "info@socialeck.com",
			CustId:          "83842",
			BillingAddress:  myAddress,
			ShippingAddress: myAddress,
		},
		AccountInfo: iTransact.CardAccount{
		CreditCard{
			AccountNumber:   "5454545454545454",
			ExpirationMonth: "05",
			ExpirationYear:  "2022",
			CVVNumber:       "394",
		}},
	}

	response := refundTransaction.Credit()

	if response.Approved() {
        fmt.Println("Approved!")
	}
```

:white_check_mark: PostAuthTransaction
```go
newTransaction := iTransact.PostAuthTransaction{
		OperationXID: "383838383",
		//Total: "25.98",
	}

	response := newTransaction.Charge()

	if response.Approved() {
        fmt.Println("Approved!")
	}
```

:white_check_mark: TranCredTransaction
```go
newTransaction := iTransact.TranCredTransaction{
		OperationXID: "383838383",
		Total: "25.98",
	}

	response := newTransaction.Charge()

	if response.Approved() {
        fmt.Println("Approved!")
	}
```

:white_check_mark: TranForceTransaction
```go
newTransaction := iTransact.TranForceTransaction{
		OperationXID: "383838383",
		ApprovalCode: "173833",
	}

	response := newTransaction.Charge()

	if response.Approved() {
         fmt.Println("Approved!")
	}
```

:white_check_mark: TranRetryTransaction
```go
newTransaction := iTransact.TranRetryTransaction{
		OperationXID: "383838383",
		Total: "25.98",
		Description: "Trying transaction again",
	}

	response := newTransaction.Charge()

	if response.Approved() {
        fmt.Println("Approved!")
	}
```

:white_check_mark: VoidTransaction
``` go
    voidTransaction := iTransact.VoidTransaction{
		OperationXID: "383838383",
	}

	response := voidTransaction.Void()

	if response.Approved() {
	    fmt.Println("Approved!")
	}
```

:white_check_mark: TranRefundTransaction
```go
newTransaction := iTransact.TranRefundTransaction{
		OperationXID: "383838383",
	}

	response := newTransaction.Refund()

	if response.Approved() {
         fmt.Println("Approved!")
	}
```

#### Recurring Payments
:white_medium_square: AuthTransaction

:white_medium_square: RecurUpdate

:white_medium_square: RecurDetails

