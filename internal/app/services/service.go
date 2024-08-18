package services

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ozwin/interview-assignment-sip/internal/app/dal"
	"github.com/ozwin/interview-assignment-sip/internal/app/models"
	"github.com/ozwin/interview-assignment-sip/internal/configs"
)

type TransactionService struct {
	store dal.Store[string, models.Transaction]
}

func NewTransactionService() *TransactionService {
	transactions, err := readTransactionsFromFile(configs.FileName)
	if err != nil {
		panic(fmt.Sprintf("failed to initialized the store: %v", err))
	}
	//initalize store capacity
	store := dal.InitializeKeyValueStore[string, models.Transaction](len(*transactions))
	for _, transaction := range *transactions {
		store.Set(transaction.Address, transaction)
	}
	return &TransactionService{
		store: store,
	}
}

func (ts TransactionService) FindTransactionByAddress(address string) (*models.Transaction, bool) {
	transaction, exists := ts.store.Get(address)
	return &transaction, exists
}

func (ts TransactionService) HandleRequest(request string) string {
	transaction, exists := ts.FindTransactionByAddress(strings.TrimSpace(request))
	if !exists {
		return ""
	}
	response, _ := transaction.ToString()
	return response

}

func readTransactionsFromFile(fileName string) (*models.Transactions, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	transactions := make(models.Transactions, 0)
	decoder := json.NewDecoder(file)
	for {
		var transaction models.Transaction
		if err := decoder.Decode(&transaction); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return &transactions, nil
}
