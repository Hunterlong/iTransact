package iTransact

import "strings"

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

func (resp iTransactResponse) AuthAmount() string {
	if !resp.Failed() {
		return resp.GatewayInterface.TransactionResponse.TransactionResult.AuthAmount
	}
	return "error"
}

