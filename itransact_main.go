package iTransact

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
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

func RunBatchClose() (*RunBatchCloseResponse, error) {
	transx := BatchClose{TransactionControl: TransactionControl{TestMode: "TRUE", SendMerchantEmail: "TRUE"}}
	batch, err := SendBatchCloseRequest(transx)
	return batch, err
}

func SendRecurringRequest(input interface{}) (*RecurUpdateReponse, error) {
	var err error
	output, err := SendToiTransact(input)
	if err != nil {
		log.Fatalln("Error sending recurring update transaction\n", err)
		return nil, err
	}
	var dat RecurUpdateReponse
	err = xml.Unmarshal(output, &dat.GatewayInterface)
	if err != nil {
		log.Fatalln("Error update response was invalid\n", err)
		return nil, err
	}
	return &dat, err
}

func SendTransactionRequest(input interface{}) (*iTransactResponse, error) {
	var err error
	output, err := SendToiTransact(input)
	if err != nil {
		log.Fatalln("Error sending transaction\n", err)
		return nil, err
	}
	var dat iTransactResponse
	err = xml.Unmarshal(output, &dat.GatewayInterface)
	if err != nil {
		log.Fatalln("Error response was invalid\n", err)
		return nil, err
	}
	return &dat, err
}

func SendBatchCloseRequest(input interface{}) (*RunBatchCloseResponse, error) {
	var err error
	output, err := SendToiTransact(input)
	if err != nil {
		log.Fatalln("Error closing batch\n", err)
		return nil, err
	}
	var dat RunBatchCloseResponse
	err = xml.Unmarshal(output, &dat.GatewayInterface)
	if err != nil {
		log.Fatalln("Error could not parse the batch close request\n", err)
		return nil, err
	}
	return &dat, err
}

func SendToiTransact(input interface{}) ([]byte, error) {
	var err error
	marshalAction, err := xml.Marshal(input)
	if err != nil {
		log.Fatalln("Error formatting XML to send to iTransact\n", err)
		return nil, err
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
		log.Fatalln("Error could not fomat XML for authentication\n", err)
		return nil, err
	}
	compiledMarshal := "<?xml version=\"1.0\"?><GatewayInterface>" + string(marshalCreds) + message + "</GatewayInterface>"
	//fmt.Println(compiledMarshal)
	req, err := http.NewRequest("POST", EndPoint, bytes.NewBuffer([]byte(compiledMarshal)))
	req.Header.Set("Content-Type", "text/xml")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("Error sending data to iTransact servers\n", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	if err != nil {
		log.Fatalln("Error could not read response from iTransact\n", err)
		return nil, err
	}
	return body, err
}
