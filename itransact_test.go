package iTransact

import (
	"testing"
	"os"
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
	SetAPIInfo(iTransactUsername,iTransactAPIPass,iTargetGateway, "test")
	t.Log("iTransact API Keys have been set! \n")
}

func TestAuthTransaction(t *testing.T) {

	t.Log("Creating a new transaction with a Total and Description with Credit Card\n")

	tranz := &AuthTransaction{
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
		Total:       "25.98",
		Description: "Order #2384 - Docker Jacket",
	}

	newTransaction, _ := SendRequest(tranz)

	if newTransaction.Approved() {
		t.Log("Transaction Approved")
		t.Log("Transaction ID: ", newTransaction.TransactionID())
		t.Log("Transaction Amount: ", newTransaction.AuthAmount())
	} else {

		if newTransaction.Failed() {
			t.Log(newTransaction.ErrorMessage())
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

	tranz := &AuthTransaction{
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

	newTransaction, _ := SendRequest(tranz)

	if newTransaction.Approved() {
		t.Log("Transaction Approved")
		t.Log("Transaction ID: ", newTransaction.TransactionID())
		t.Log("Transaction Amount: ", newTransaction.AuthAmount())
	} else {

		if newTransaction.Failed() {
			t.Log(newTransaction.ErrorMessage())
		}

		t.Log("Transaction Declined")
		t.Fail()
	}

}
