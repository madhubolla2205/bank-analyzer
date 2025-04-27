package reader

import (
    "github.com/ledongthuc/pdf"
    "strconv"
    "strings"
)

func ParsePDF(filePath string) ([]Transaction, error) {
    f, r, err := pdf.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    var transactions []Transaction
    rtext, err := r.GetPlainText()
    if err != nil {
        return nil, err
    }

    lines := strings.Split(rtext, "\n")
    for _, line := range lines {
        fields := strings.Fields(line)
        if len(fields) < 3 {
            continue
        }

        amount, _ := strconv.ParseFloat(fields[len(fields)-1], 64)
        description := strings.Join(fields[1:len(fields)-1], " ")

        transactions = append(transactions, Transaction{
            Date:        fields[0],
            Description: description,
            Amount:      amount,
        })
    }

    return transactions, nil
}

