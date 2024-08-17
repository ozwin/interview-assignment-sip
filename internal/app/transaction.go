package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"
)

type Transactions []Transaction

type Transaction struct {
	Address      string   `json:"addressOfRecord"`
	TenantId     string   `json:"tenantId"`
	Uri          string   `json:"uri"`
	Contact      string   `json:"contact"`
	Path         []string `json:"path"`
	Source       string   `json:"source"`
	Target       string   `json:"target"`
	UserAgent    string   `json:"userAgent"`
	RawUserAgent string   `json:"rawUserAgent"`
	Created      string   `json:"created"`
	LineId       string   `json:"lineId"`
}

func ReadTransactionsFromFile(fileName string) (*Transactions, error) {
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
	transactions := make(Transactions, 0, len(rawTransactions))
	for _, rawTransaction := range rawTransactions {
		if rawTransaction == "" {
			continue
		}
		data := strings.TrimSpace(rawTransaction)
		//Removed it while splitting based on }
		//since objects weren't seperated with a delimiter in the file
		data += "}"
		var transaction Transaction
		if err := json.Unmarshal([]byte(data), &transaction); err != nil {
			//log and forget for now
			log.Fatalf("error while parsing string to transaction object: %v", err)
		}
		transactions = append(transactions, transaction)
	}
	return &transactions, nil
}

func (ts Transactions) FindByAddress(address string) *Transaction {
	for index, t := range ts {
		if t.Address == address {
			return &ts[index]
		}
	}
	return nil
}
