package models

import (
	"encoding/json"
)

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

func (t *Transaction) ToString() (string, error) {
	//skipping error checking fr now
	bytes, _ := json.Marshal(t)
	return string(bytes), nil
}

type Transactions []Transaction
