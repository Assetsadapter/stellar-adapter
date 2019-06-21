package algorand

import (
	"encoding/json"
	"fmt"

	"github.com/algorand/go-algorand-sdk/client/algod/models"
	"github.com/blocktree/openwallet/common"
	"github.com/blocktree/openwallet/crypto"
	"github.com/blocktree/openwallet/openwallet"
)

type Block struct {
	Hash             string
	CurrentProtocol  string
	PrevBlockHash    string
	TransactionsRoot string
	Proposer         string
	Time             int64
	Height           uint64
	Transactions     []string
}

func NewBlock(block models.Block) *Block {
	obj := Block{}
	//解析json
	obj.Hash = block.Hash
	obj.CurrentProtocol = block.CurrentProtocol
	obj.PrevBlockHash = block.PreviousBlockHash
	obj.TransactionsRoot = block.TransactionsRoot
	obj.Height = block.Round
	obj.Proposer = block.Proposer
	obj.Time = block.Timestamp

	txs := make([]string, 0)
	for _, t := range block.Transactions.Transactions {
		tx, _ := json.Marshal(t)
		txs = append(txs, string(tx))
	}
	obj.Transactions = txs

	return &obj
}

//BlockHeader 区块链头
func (b *Block) BlockHeader(symbol string) *openwallet.BlockHeader {

	obj := openwallet.BlockHeader{}
	//解析json
	obj.Hash = b.Hash
	//obj.Confirmations = b.Confirmations
	//obj.Merkleroot = b.TransactionMerkleRoot
	obj.Previousblockhash = b.PrevBlockHash
	obj.Height = b.Height
	// obj.Version = uint64(b.Version)
	obj.Time = uint64(obj.Time)
	obj.Symbol = symbol

	return &obj
}

type Transaction struct {
	Type        string
	BlockHash   string
	BlockHeight uint64
	BlockTime   int64
	GenesisID   string
	Fee         uint64
	TxID        string
	From        string
	Note        []byte
	Payment     *models.PaymentTransactionType
}

func NewTransaction(tx models.Transaction) *Transaction {
	obj := Transaction{}
	obj.Type = string(tx.Type)
	obj.BlockHeight = tx.ConfirmedRound
	obj.GenesisID = tx.GenesisID
	obj.Fee = tx.Fee
	obj.TxID = tx.TxID
	obj.From = tx.From
	obj.Payment = tx.Payment
	return &obj
}

//UnscanRecords 扫描失败的区块及交易
type UnscanRecord struct {
	ID          string `storm:"id"` // primary key
	BlockHeight uint64
	TxID        string
	Reason      string
}

func NewUnscanRecord(height uint64, txID, reason string) *UnscanRecord {
	obj := UnscanRecord{}
	obj.BlockHeight = height
	obj.TxID = txID
	obj.Reason = reason
	obj.ID = common.Bytes2Hex(crypto.SHA256([]byte(fmt.Sprintf("%d_%s", height, txID))))
	return &obj
}
