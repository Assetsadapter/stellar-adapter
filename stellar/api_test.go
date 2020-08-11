package stellar

import (
	"fmt"
	"github.com/blocktree/openwallet/log"
	"github.com/shopspring/decimal"
	hClient "github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/protocols/horizon/operations"
	"github.com/stellar/go/txnbuild"
	"net/http"
	"testing"
)

func TestApi(t *testing.T) {
	//api :=&Client{BaseURL:"https://testnet-horizon.triamnetwork.com/"}
	//resp,err := api.Get("/ledgers/1000",nil)
	//if err != nil {
	//	panic(err)
	//}
	//log.Info(resp)
	a, _ := decimal.NewFromString("")
	log.Info("value=", a.String())
}

func TestApiLedgers(t *testing.T) {

	// Use the default pubnet client
	client := hClient.Client{HorizonURL: "https://testnet-horizon.triamnetwork.com/", HTTP: http.DefaultClient}

	// Create an account request
	//hClient.LedgerRequest{Limit:1}
	// Load the account detail from the network
	ledger, err := client.Ledgers(hClient.LedgerRequest{Limit: 1, Order: "desc"})
	if err != nil {
		fmt.Println(err)
		return
	}
	// Account contains information about the stellar account
	fmt.Println(ledger.Embedded.Records[0].Hash)
	fmt.Println(ledger.Embedded.Records[0].Sequence)
}

func TestApiRoot(t *testing.T) {

	// Use the default pubnet client
	client := hClient.Client{HorizonURL: "https://testnet-horizon.triamnetwork.com/", HTTP: http.DefaultClient}

	root, err := client.Root()
	if err != nil {
		log.Error(err)
	}
	log.Info(root.CoreSequence)
	log.Info(root.HorizonSequence)
}

func TestApiLedgerDetal(t *testing.T) {
	// Use the default pubnet client
	client := hClient.Client{HorizonURL: "https://testnet-horizon.triamnetwork.com/", HTTP: http.DefaultClient}

	ledge, err := client.LedgerDetail(111)
	if err != nil {
		log.Error(err)
	}
	log.Info(ledge.Hash)
	log.Info(ledge.Sequence)
}

//提取账本 payment 交易
func TestApiGetLedgerPayment(t *testing.T) {
	// Use the default pubnet client
	client := hClient.Client{HorizonURL: "https://testnet-horizon.triamnetwork.com/", HTTP: http.DefaultClient}

	ledge, err := client.Operations(hClient.OperationRequest{ForLedger: 9958186})
	if err != nil {
		log.Error(err)
	}
	log.Info(ledge.Embedded)
	for _, op := range ledge.Embedded.Records {
		//log.Infof("from=%s to=%s amount=%s")
		opType := op.GetType()
		if opType == "payment" {
			payOP, OK := op.(operations.Payment)
			if !OK {
				panic("not payment op")
			}
			log.Infof("from=%s to=%s amount=%s code=%s", payOP.From, payOP.To, payOP.Amount, payOP.Code)
			log.Info(payOP.TransactionHash)
		}

	}
}
func newKeypair(seed string) *keypair.Full {
	myKeypair, _ := keypair.Parse(seed)
	return myKeypair.(*keypair.Full)
}

func newKeypair0() *keypair.Full {
	// Address: GDQNY3PBOJOKYZSRMK2S7LHHGWZIUISD4QORETLMXEWXBI7KFZZMKTL3
	return newKeypair("SBPQUZ6G4FZNWFHKUWC5BEYWF6R52E3SEP7R3GWYSM2XTKGF5LNTWW4R")
}

// NewSimpleAccount is a factory method that creates a SimpleAccount from "accountID" and "sequence".
func NewSimpleAccount(accountID string, sequence int64) txnbuild.SimpleAccount {
	return txnbuild.SimpleAccount{accountID, sequence}
}
func newKeypair1() *keypair.Full {
	// Address: GAS4V4O2B7DW5T7IQRPEEVCRXMDZESKISR7DVIGKZQYYV3OSQ5SH5LVP
	return newKeypair("SBMSVD4KKELKGZXHBUQTIROWUAPQASDX7KEJITARP4VMZ6KLUHOGPTYW")
}
func buildSignEncode(t *testing.T, tx txnbuild.Transaction, kps ...*keypair.Full) string {

	return ""
}
func TestAllowTrustMultSigners(t *testing.T) {


}
