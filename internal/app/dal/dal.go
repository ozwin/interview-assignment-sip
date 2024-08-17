package dal

type KeyValueStore[K comparable, V any] struct {
	store map[K]V
}

func Initialize[K comparable, V any](capacity int) *KeyValueStore[K, V] {
	return &KeyValueStore[K, V]{
		store: make(map[K]V, capacity),
	}
}

func (kv *KeyValueStore[K, V]) Set(key K, value V) {
	kv.store[key] = value
}

func (kv *KeyValueStore[K, V]) Get(key K) (V, bool) {
	value, exists := kv.store[key]
	return value, exists
}

// func ReadTransactionsFromFile(fileName string) (*Transactions, error) {
// 	file, err := os.Open(fileName)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()
// 	data, err := io.ReadAll(file)
// 	if err != nil {
// 		log.Fatalf("error while reading file: %v", err.Error())
// 	}
// 	rawTransactions := strings.Split(string(data), "}")
// 	transactions := make(Transactions, 0, len(rawTransactions))
// 	for _, rawTransaction := range rawTransactions {
// 		if rawTransaction == "" {
// 			continue
// 		}
// 		data := strings.TrimSpace(rawTransaction)
// 		//Removed it while splitting based on }
// 		//since objects weren't seperated with a delimiter in the file
// 		data += "}"
// 		var transaction models.Transaction
// 		if err := json.Unmarshal([]byte(data), &transaction); err != nil {
// 			//log and forget for now
// 			log.Fatalf("error while parsing string to transaction object: %v", err)
// 		}
// 		transactions = append(transactions, transaction)
// 	}
// 	return &transactions, nil
// }

// func (ts Transactions) FindByAddress(address string) *models.Transaction {
// 	for index, t := range ts {
// 		if t.Address == address {
// 			return &ts[index]
// 		}
// 	}
// 	return nil
// }
