package iTransact

import (
	"log"
	"strings"
)

func (recurring RecurUpdate) Charge() (*SendRecurUpdate, error) {
	newTransaction, err := SendRecurringRequest(recurring)
	if err != nil {
		log.Fatalln("Failed create recurring transaction\n", err)
		return nil, err
	}
	return &newTransaction.GatewayInterface.RecurUpdate, err
}

func (transx PostAuthTransaction) Charge() (*iTransactResponse, error) {
	newTransaction, err := SendTransactionRequest(transx)
	if err != nil {
		log.Fatalln("Failed to POST auth charge to card\n", err)
		return nil, err
	}
	return newTransaction, err
}

func (transx AuthTransaction) Charge() (*iTransactResponse, error) {
	newTransaction, err := SendTransactionRequest(transx)
	if err != nil {
		log.Fatalln("Failed to AUTH card\n", err)
		return nil, err
	}
	return newTransaction, err
}

func (transx TranCredTransaction) Charge() (*iTransactResponse, error) {
	newTransaction, err := SendTransactionRequest(transx)
	if err != nil {
		log.Fatalln("Failed to credit a card\n", err)
		return nil, err
	}
	return newTransaction, err
}

func (transx TranForceTransaction) Charge() (*iTransactResponse, error) {
	newTransaction, err := SendTransactionRequest(transx)
	if err != nil {
		log.Fatalln("Failed to force transaction\n", err)
		return nil, err
	}
	return newTransaction, err
}

func (transx TranRetryTransaction) Charge() (*iTransactResponse, error) {
	newTransaction, err := SendTransactionRequest(transx)
	if err != nil {
		log.Fatalln("Failed to retry transaction\n", err)
		return nil, err
	}
	return newTransaction, err
}

func (transx VoidTransaction) Void() (*iTransactResponse, error) {
	newTransaction, err := SendTransactionRequest(transx)
	if err != nil {
		log.Fatalln("Failed to charge card\n", err)
		return nil, err
	}
	return newTransaction, err
}

func (transx TranRefundTransaction) Refund() (*iTransactResponse, error) {
	newTransaction, err := SendTransactionRequest(transx)
	if err != nil {
		log.Fatalln("Failed to charge card\n", err)
		return nil, err
	}
	return newTransaction, err
}

func (transx CreditTransaction) Credit() (*iTransactResponse, error) {
	newTransaction, err := SendTransactionRequest(transx)
	if err != nil {
		log.Fatalln("Failed to charge card\n", err)
		return nil, err
	}
	return newTransaction, err
}

func (resp iTransactResponse) Approved() bool {
	if (strings.ToLower(resp.GatewayInterface.TransactionResponse.TransactionResult.Status)) == "ok" {
		return true
	}
	return false
}

func (resp iTransactResponse) Failed() bool {
	if (strings.ToLower(resp.GatewayInterface.TransactionResponse.TransactionResult.Status)) == "fail" {
		return true
	}
	return false
}

func (resp iTransactResponse) ErrorMessage() string {
	errors := resp.GatewayInterface.TransactionResponse.TransactionResult.ErrorMessage
	errorCategory := resp.GatewayInterface.TransactionResponse.TransactionResult.ErrorCategory
	if errors != "" {
		return errorCategory + " - " + errors
	}
	return "None Found"
}

func (resp iTransactResponse) TransactionID() string {
	if !resp.Failed() {
		return resp.GatewayInterface.TransactionResponse.TransactionResult.XID
	}
	return "error"
}

func (resp iTransactResponse) Total() string {
	if !resp.Failed() {
		return resp.GatewayInterface.TransactionResponse.TransactionResult.Total
	}
	return "error"
}

func (resp RunBatchCloseResponse) Amount() int {
	return len(resp.GatewayInterface.BatchCloseResponse.BatchList.Batch)
}
