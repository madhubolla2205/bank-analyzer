package processor

import (
    "strings"
    "github.com/madhubolla2205/bank-analyzer/pkg/models"
)

type Categorized struct {
    Category string
    Amount   float64
}


var data = make(map[string]float64)

func CategorizeTransactions(txns []models.Transaction) []Categorized {
    for _, t := range txns {
        category := categorize(t.Description)
        data[category] += t.Amount
    }

    var result []Categorized
    for k, v := range data {
        result = append(result, Categorized{Category: k, Amount: v})
    }
    return result
}

func categorize(description string) string {
    desc := strings.ToLower(description)
    switch {
    case strings.Contains(desc, "amazon"):
        return "Shopping"
    case strings.Contains(desc, "electricity"):
        return "Utilities"
    case strings.Contains(desc, "salary"):
        return "Income"
    default:
        return "Others"
    }
}

func StoreData(cat []Categorized) {
    // already stored during CategorizeTransactions
}

func GetData() map[string]float64 {
    return data
}

