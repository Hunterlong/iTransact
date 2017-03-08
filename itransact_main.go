package iTransact

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"fmt"
)

// your iTransact Username
var ITransactUsername string

// your iTransact API Password found in Account Settings
var ITransactAPIPass string

// iTransact Gateway ID number (5 digits)
var ITargetGateway string

// Endpoint for the iTransact Processing API
var EndPoint string = "https://secure.paymentclearing.com/cgi-bin/rc/xmltrans2.cgi"

// Test mode true or false
var TestMode bool = true

func SetAPIInfo(user string, pass string, gateway string, testMode string) {
	ITransactUsername = user
	ITransactAPIPass = pass
	ITargetGateway = gateway
	if testMode == "live" {
		TestMode = false
	}
}

func RunBatchClose() RunBatchCloseResponse {
	transx := BatchClose{TransactionControl: TransactionControl{TestMode: "TRUE", SendMerchantEmail: "TRUE"}}
	batch, _ := SendBatchCloseRequest(transx)
	return batch
}

func SendRecurringRequest(input interface{}) (RecurUpdateReponse, interface{}) {
	output := SendToiTransact(input)
	fmt.Println(string(output))
	var dat RecurUpdateReponse
	err := xml.Unmarshal(output, &dat.GatewayInterface)
	if err != nil {
		panic(err)
	}
	return dat, err
}

func SendTransactionRequest(input interface{}) (iTransactResponse, interface{}) {
	output := SendToiTransact(input)
	fmt.Println(string(output))
	var dat iTransactResponse
	err := xml.Unmarshal(output, &dat.GatewayInterface)
	if err != nil {
		panic(err)
	}
	return dat, err
}

func SendBatchCloseRequest(input interface{}) (RunBatchCloseResponse, interface{}) {
	output := SendToiTransact(input)
	var dat RunBatchCloseResponse
	err := xml.Unmarshal(output, &dat.GatewayInterface)
	if err != nil {
		panic(err)
	}
	return dat, err
}

func SendToiTransact(input interface{}) []byte {
	marshalAction, err := xml.Marshal(input)
	if err != nil {
		panic(err)
	}
	key := []byte(ITransactAPIPass)
	message := string(marshalAction)
	sig := hmac.New(sha1.New, key)
	sig.Write([]byte(message))
	payloadSig := base64.StdEncoding.EncodeToString(sig.Sum(nil))

	newCreds := APICredentials{
		Username:         ITransactUsername,
		TargetGateway:    ITargetGateway,
		PayloadSignature: payloadSig,
	}

	marshalCreds, err := xml.Marshal(newCreds)
	if err != nil {
		panic(err)
	}
	compiledMarshal := "<?xml version=\"1.0\"?><GatewayInterface>" + string(marshalCreds) + message + "</GatewayInterface>"
	fmt.Println(compiledMarshal)
	req, err := http.NewRequest("POST", EndPoint, bytes.NewBuffer([]byte(compiledMarshal)))
	req.Header.Set("Content-Type", "text/xml")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	return body
}
