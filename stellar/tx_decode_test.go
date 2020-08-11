package stellar

import (
	"github.com/blocktree/openwallet/log"
	hClient "github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/txnbuild"
	"net/http"
	"strconv"
	"testing"
)
const node_url="https://stellar-horizon.satoshipay.io/"
const network_param="Public Global Stellar Network ; September 2015"
//获取下一个账户的seq
func getAccountSeq(account string) int64 {
	client := hClient.Client{HorizonURL: node_url, HTTP: http.DefaultClient}
	accountDetail, _ := client.AccountDetail(hClient.AccountRequest{AccountID: account})
	seq, _ := strconv.ParseInt(accountDetail.Sequence, 10, 64)
	return seq
}

//获取下一个账户的seq
func TestExistsAccount(t *testing.T) {
	client := hClient.Client{HorizonURL: node_url, HTTP: http.DefaultClient}
	accountDetail, err := client.AccountDetail(hClient.AccountRequest{AccountID: "GDJQNT7SMTRZPQBVB46R54J454F2ZCKUJ37YVW6KYD2P2Y2FPLK5ANGQ1"})
	if err != nil {
		panic(err)
	}
	log.Info(accountDetail)
}

//发送交易
func postTransaction(tx *txnbuild.Transaction) string {
	client := hClient.Client{HorizonURL: node_url, HTTP: http.DefaultClient}
	txBase64, err := tx.Base64()
	log.Info(txBase64)
	if err != nil {
		panic(err)
	}
	a, err := client.SubmitTransaction(tx)
	if err != nil {
		panic(err)
	}
	log.Info(a.Hash)
	txId := a.Hash
	return txId
}

//测试创建交易
func TestBuildPayTransaction(t *testing.T) {
	//privateKey, _ := hex.DecodeString("407f243d382e6754b14307efcde040b6f88904973e1ef306d93743d098d0f945")
	//nexSeq := getAccountSeq("GDJQNT7SMTRZPQBVB46R54J454F2ZCKUJ37YVW6KYD2P2Y2FPLK5ANGQ")
	//txSourceAccount := NewSimpleAccount("GDJQNT7SMTRZPQBVB46R54J454F2ZCKUJ37YVW6KYD2P2Y2FPLK5ANGQ", nexSeq)
	//
	//payment := txnbuild.Payment{
	//	Destination: "GD2FAPCGEX5L63FENNSN6ZZPIL7NEHVPV2NCTGCALY7HG7KHLDW3QOMM",
	//	Amount:      "0.234",
	//	Asset:       txnbuild.NativeAsset{},
	//}
	//
	//tx := txnbuild.Transaction{
	//	BaseFee:       10000,
	//	SourceAccount: &txSourceAccount,
	//	Operations:    []txnbuild.Operation{&payment},
	//	Timebounds:    txnbuild.NewInfiniteTimeout(),
	//	Network:       "SAAK5654--ARM-NETWORK--BHC3SQOHPO2GGI--BY-B.A.P--CNEMJQCWPTA--RUBY-AND-BLOCKCHAIN--3KECMPY5L7W--THANKYOU-CS--S542ZHDVHLFV",
	//}
	//tx.Build()
	//if tx.XdrEnvelope == nil {
	//	tx.XdrEnvelope = &xdr.TransactionEnvelope{}
	//	tx.XdrEnvelope.Tx = tx.XdrTransaction
	//}
	//hash, err := tx.Hash()
	//if err != nil {
	//	panic(err)
	//}
	//log.Info("hash=", hash)
	//sig, err := txsigner.Default.SignTransactionHash(hash[:], privateKey, owcrypt.ECC_CURVE_ED25519)
	//log.Info("sig=", sig, "len=", len(sig))
	//kp, _ := keypair.Parse("GDJQNT7SMTRZPQBVB46R54J454F2ZCKUJ37YVW6KYD2P2Y2FPLK5ANGQ")
	//xdrSig := xdr.DecoratedSignature{
	//	Hint:      xdr.SignatureHint(kp.Hint()),
	//	Signature: xdr.Signature(sig),
	//}
	//tx.XdrEnvelope.Signatures = append(tx.XdrEnvelope.Signatures, xdrSig)
	//
	//postTransaction(tx)
}

//测试创建change trust交易
func TestBuildTrustTransaction(t *testing.T) {
	//privateKey, _ := hex.DecodeString("389d28fb68ad49af90c7d6ec062c6e7d2e1aa0fa1dd11309060b388c2c72fc56")
	//nexSeq := getAccountSeq("GCEFX6JKWIJSN5SR5R3JB5CPDZIPT2YCQZYVKSQFMWJQ6LGSSKFJM2CT")
	//txSourceAccount := NewSimpleAccount("GCEFX6JKWIJSN5SR5R3JB5CPDZIPT2YCQZYVKSQFMWJQ6LGSSKFJM2CT", nexSeq)
	//
	//changeTrust := txnbuild.ChangeTrust{
	//	Line: txnbuild.CreditAsset{"WGX", "GCSJMVRD43JLFR7GXG7EESMFVB53BHGODTZ2EY7C63I7EGW3JWAG2F6L"},
	//}
	//
	//tx := txnbuild.TransactionParams{
	//	BaseFee:       10000,
	//	SourceAccount: &txSourceAccount,
	//	Operations:    []txnbuild.Operation{&changeTrust},
	//	Timebounds:    txnbuild.NewInfiniteTimeout(),
	//}
	//
	//
	//hash, err :=
	//if err != nil {
	//	panic(err)
	//}
	//log.Info("hash=", hash)
	//sig, err := txsigner.Default.SignTransactionHash(hash[:], privateKey, owcrypt.ECC_CURVE_ED25519)
	//log.Info("sig=", sig, "len=", len(sig))
	//kp, _ := keypair.Parse("GCEFX6JKWIJSN5SR5R3JB5CPDZIPT2YCQZYVKSQFMWJQ6LGSSKFJM2CT")
	//xdrSig := xdr.DecoratedSignature{
	//	Hint:      xdr.SignatureHint(kp.Hint()),
	//	Signature: xdr.Signature(sig),
	//}
	//tx.XdrEnvelope.Signatures = append(tx.XdrEnvelope.Signatures, xdrSig)
	//postTransaction(tx)
}

func TestBuildCreateAccountTransaction(t *testing.T) {
	//privateKey, _ := hex.DecodeString("389d28fb68ad49af90c7d6ec062c6e7d2e1aa0fa1dd11309060b388c2c72fc56")
	seq := getAccountSeq("GA7J7XW6AVYYMMZHKKYPSRTBLJ46QPTZHZ46ADB6KAXG26ZT6URDCMWJ")
	txSourceAccount := NewSimpleAccount("GA7J7XW6AVYYMMZHKKYPSRTBLJ46QPTZHZ46ADB6KAXG26ZT6URDCMWJ",seq+1)

	createAccount := txnbuild.CreateAccount{
		Destination: "GBHL6SH67VE6TMSR2HI2ZINPWYDWSRPNXD7MNFPEPM5OZYVCHU7RUSVY",
		Amount:      "1",
	}

	txParam := txnbuild.TransactionParams{
		BaseFee:       10000,
		SourceAccount: &txSourceAccount,
		Operations:    []txnbuild.Operation{&createAccount},
		Timebounds:    txnbuild.NewInfiniteTimeout(),
	}
	tx,err := txnbuild.NewTransaction(txParam)

	hash, err := tx.Hash(network_param)
	if err != nil {
		panic(err)
	}
	log.Info("hash=", hash)
	keyPair := newKeypair("SC473A2F474MZ5ZFF2DT2YRCXADZZFTPA5I227GCBRGPNVN4FTEFJQHG")
	if err != nil {
		panic(err)
	}
	txSign ,err :=tx.Sign(network_param,keyPair)
	if err != nil {
		panic(err)
	}
	txId:=postTransaction(txSign)
	log.Info("txId = ",txId)
}