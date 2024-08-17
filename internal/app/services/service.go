package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/ozwin/interview-assignment-sip/internal/app/dal"
	"github.com/ozwin/interview-assignment-sip/internal/app/models"
	"github.com/ozwin/interview-assignment-sip/internal/configs"
)

type TransactionService struct {
	store *dal.KeyValueStore[string, models.Transaction]
}

func NewTransactionService() *TransactionService {
	transactions, err := readTransactionsFromFile(configs.FileName)
	if err != nil {
		panic(fmt.Sprintf("failed to initialized the store: %v", err))
	}
	//initalize store capacity
	store := dal.Initialize[string, models.Transaction](len(*transactions))
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
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("error while reading file: %v", err.Error())
	}
	rawTransactions := strings.Split(string(data), "}")
	transactions := make(models.Transactions, 0, len(rawTransactions))
	for _, rawTransaction := range rawTransactions {
		if strings.TrimSpace(rawTransaction) == "" {
			continue
		}
		data := strings.TrimSpace(rawTransaction)
		//Removed it while splitting based on }
		//since objects weren't seperated with a delimiter in the file
		data += "}"
		var transaction models.Transaction
		if err := json.Unmarshal([]byte(data), &transaction); err != nil {
			//log and forget for now
			log.Fatalf("error while parsing string to transaction object: %v", err)
		}
		transactions = append(transactions, transaction)
	}
	return &transactions, nil
}
