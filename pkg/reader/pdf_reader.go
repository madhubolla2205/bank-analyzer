package reader

import (
    "github.com/ledongthuc/pdf"
    "github.com/madhubolla2205/bank-analyzer/pkg/models"
    "io"
    "strings"
    "strconv"
)

func ParsePDF(filePath string) ([]models.Transaction, error) {
    f, r, err := pdf.Open(filePath)
    if (err != nil) {
        return nil, err
    }
    defer f.Close()

    var transactions []models.Transaction
    rtext, err := r.GetPlainText()
    if (err != nil) {
        return nil, err
    }

    var buf strings.Builder
    io.Copy(&buf, rtext)
    text := buf.String()

    lines := strings.Split(text, "\n")
    for _, line := range lines {
        fields := strings.Fields(line)
        if len(fields) < 3 {
            continue
        }

        amount, _ := strconv.ParseFloat(fields[len(fields)-1], 64)
        description := strings.Join(fields[1:len(fields)-1], " ")

        transactions = append(transactions, models.Transaction{
            Date:        fields[0],
            Description: description,
            Amount:      amount,
        })
    }

    return transactions, nil
}

