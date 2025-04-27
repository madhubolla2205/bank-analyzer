package reader

import (
    "github.com/madhubolla2205/bank-analyzer/pkg/models"
    "encoding/csv"
    "os"
    "strconv"
)
func ParseCSV(filePath string) ([]models.Transaction, error) {
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

    var transactions []models.Transaction
    for i, record := range records {
        if i == 0 {
            continue // Skip header
        }
        amount, _ := strconv.ParseFloat(record[2], 64)
        transactions = append(transactions, models.Transaction{
            Date:        record[0],
            Description: record[1],
            Amount:      amount,
        })
    }

    return transactions, nil
}

