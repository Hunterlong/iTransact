package iTransact

import (
	"os"
	"testing"
)

var myAddress = Address{
	Address1:  "1111 awesome ave.",
	Address2:  "apt 321",
	FirstName: "Hunter",
	LastName:  "Long",
	City:      "awesome mountain",
	State:     "UT",
	Zip:       "43999",
	Country:   "USA",
	Phone:     "4555555555",
}

func TestSetAPIInfo(t *testing.T) {
	iTransactUsername := os.Getenv("ITRANSACT_USER")
	iTransactAPIPass := os.Getenv("ITRANSACT_APIPASS")
	iTargetGateway := os.Getenv("ITRANSACT_GATEWAY")
	SetAPIInfo(iTransactUsername, iTransactAPIPass, iTargetGateway, "test")
	t.Log("iTransact API Keys have been set! \n")
}

func TestAuthTransaction(t *testing.T) {

	t.Log("Creating a new transaction with a Total and Description with Credit Card\n")

	newTransaction := &AuthTransaction{
		Total:       "25.98",
		Description: "Order #2384 - Docker Jacket",
		CustomerData: CustomerData{
			Email:           "info@socialeck.com",
			CustId:          "83842",
			BillingAddress:  myAddress,
			ShippingAddress: myAddress,
		},
		AccountInfo: CardAccount{CreditCard{
			AccountNumber:   "5454545454545454",
			ExpirationMonth: "05",
			ExpirationYear:  "2022",
			CVVNumber:       "394",
		}},
	}

	response := newTransaction.Charge()

	if response.Approved() {
		t.Log("Transaction Approved")
		t.Log("Transaction ID: ", response.TransactionID())
		t.Log("Transaction Amount: ", response.Total())
	} else {

		if response.Failed() {
			t.Log(response.ErrorMessage())
		}

		t.Log("Transaction Declined")
		t.Fail()
	}

}

func TestItemsAuthTransaction(t *testing.T) {

	t.Log("Creating a new transaction with Items inside an Order")

	myOrder := []Item{
		{
			Description: "Docker Jacket",
			Cost:        "25.98",
			Qty:         "1",
		},
	}

	newTransaction := &AuthTransaction{
		CustomerData: CustomerData{
			Email:           "info@socialeck.com",
			CustId:          "83842",
			BillingAddress:  myAddress,
			ShippingAddress: myAddress,
		},
		AccountInfo: CardAccount{CreditCard{
			AccountNumber:   "5454545454545454",
			ExpirationMonth: "09",
			ExpirationYear:  "2021",
			CVVNumber:       "314",
		}},
		OrderItems: &Items{myOrder},
		//Total:      "25.98",
		//Description: "Order #2384 - Docker Jacket",
	}

	response := newTransaction.Charge()

	if response.Approved() {
		t.Log("Transaction Approved")
		t.Log("Transaction ID: ", response.TransactionID())
		t.Log("Transaction Amount: ", response.Total())
	} else {

		if response.Failed() {
			t.Log(response.ErrorMessage())
		}

		t.Log("Transaction Declined")
		t.Fail()
	}

}

func TestCloseBatch(t *testing.T) {
	batch := RunBatchClose()

	t.Log("Closed", batch.Amount(), "Transactions")

}



func TestCreditTransaction(t *testing.T) {

	t.Log("Refunding a transaction\n")

	refundTransaction := &CreditTransaction{
		Total:       "25.98",
		CustomerData: CustomerData{
			Email:           "info@socialeck.com",
			CustId:          "83842",
			BillingAddress:  myAddress,
			ShippingAddress: myAddress,
		},
		AccountInfo: CardAccount{CreditCard{
			AccountNumber:   "5454545454545454",
			ExpirationMonth: "05",
			ExpirationYear:  "2022",
			CVVNumber:       "394",
		}},
	}

	response := refundTransaction.Credit()

	if response.Approved() {
		t.Log("Transaction Approved")
		t.Log("Transaction ID: ", response.TransactionID())
		t.Log("Transaction Amount: ", response.Total())
	} else {

		if response.Failed() {
			t.Log(response.ErrorMessage())
		}

		t.Log("Transaction Declined")
		t.Fail()
	}

}


func TestPostAuthTransaction(t *testing.T) {


	newTransaction := PostAuthTransaction{
		OperationXID: "383838383",
		//Total: "25.98",
	}

	response := newTransaction.Charge()

	if response.Approved() {
		t.Log("Transaction Approved")
		t.Log("Transaction ID: ", response.TransactionID())
		t.Log("Transaction Amount: ", response.Total())
	} else {

		if response.Failed() {
			t.Log(response.ErrorMessage())
		}

		t.Log("Transaction Declined")
		t.Fail()
	}



}


func TestTranCredTransaction(t *testing.T) {

	newTransaction := TranCredTransaction{
		OperationXID: "383838383",
		Total: "25.98",
	}

	response := newTransaction.Charge()

	if response.Approved() {
		t.Log("Transaction Approved")
		t.Log("Transaction ID: ", response.TransactionID())
		t.Log("Transaction Amount: ", response.Total())
	} else {

		if response.Failed() {
			t.Log(response.ErrorMessage())
		}

		t.Log("Transaction Declined")
		t.Fail()
	}

}


func TestTranForceTransaction(t *testing.T) {

	newTransaction := TranForceTransaction{
		OperationXID: "383838383",
		ApprovalCode: "173833",
	}

	response := newTransaction.Charge()

	if response.Approved() {
		t.Log("Transaction Approved")
		t.Log("Transaction ID: ", response.TransactionID())
		t.Log("Transaction Amount: ", response.Total())
	} else {

		if response.Failed() {
			t.Log(response.ErrorMessage())
		}

		t.Log("Transaction Declined")
		t.Fail()
	}

}


func TestTranRetryTransaction(t *testing.T) {

	newTransaction := TranRetryTransaction{
		OperationXID: "383838383",
		Total: "25.98",
		Description: "Trying transaction again",
	}

	response := newTransaction.Charge()

	if response.Approved() {
		t.Log("Transaction Approved")
		t.Log("Transaction ID: ", response.TransactionID())
		t.Log("Transaction Amount: ", response.Total())
	} else {

		if response.Failed() {
			t.Log(response.ErrorMessage())
		}

		t.Log("Transaction Declined")
		t.Fail()
	}

}


func TestVoidTransaction(t *testing.T) {

	newTransaction := VoidTransaction{
		OperationXID: "383838383",
	}

	response := newTransaction.Void()

	if response.Approved() {
		t.Log("Transaction Approved")
		t.Log("Transaction ID: ", response.TransactionID())
		t.Log("Transaction Amount: ", response.Total())
	} else {

		if response.Failed() {
			t.Log(response.ErrorMessage())
		}

		t.Log("Transaction Declined")
		t.Fail()
	}

}


func TestTranRefundTransaction(t *testing.T) {

	newTransaction := TranRefundTransaction{
		OperationXID: "383838383",
	}

	response := newTransaction.Refund()

	if response.Approved() {
		t.Log("Transaction Approved")
		t.Log("Transaction ID: ", response.TransactionID())
		t.Log("Transaction Amount: ", response.Total())
	} else {

		if response.Failed() {
			t.Log(response.ErrorMessage())
		}

		t.Log("Transaction Declined")
		t.Fail()
	}

}