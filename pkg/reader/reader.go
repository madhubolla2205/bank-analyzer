package reader

import (
    "encoding/csv"
    "os"
    "strconv"
)

type Transaction struct {
    Date        string
    Description string
    Amount      float64
}

func ParseCSV(filePath string) ([]Transaction, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        return nil, err
    }

    var transactions []Transaction
    for i, record := range records {
        if i == 0 {
            continue // Skip header
        }
        amount, _ := strconv.ParseFloat(record[2], 64)
        transactions = append(transactions, Transaction{
            Date:        record[0],
            Description: record[1],
            Amount:      amount,
        })
    }

    return transactions, nil
}

