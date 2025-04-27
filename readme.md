# 📈 Bank Statement Analyzer (Go + Bootstrap)

A stylish Go web application to upload your bank statements (CSV/PDF), automatically categorize transactions, and view beautiful graphs! 🚀

---

## ✨ Features
- Upload Bank Statements (.csv and .pdf)
- Auto-categorize by description keywords (Amazon → Shopping, Rent, Utilities, etc.)
- Stylish frontend with Bootstrap 5
- Interactive Pie Chart and Line Chart (Chart.js)
- PDF parsing supported (basic text extraction)
- No database needed — fast and simple!

---

## 📁 Project Structure
bank-analyzer/
├── cmd/
│   └── main.go
├── pkg/
│   ├── reader/
│   │   ├── reader.go
│   │   └── pdf_reader.go
│   ├── processor/
│   │   └── categorize.go
│   ├── server/
│   │   └── routes.go
├── static/
│   └── index.html
├── uploads/                # (Empty folder where uploaded files will go)
├── sample/
│   ├── statement.csv
│   └── statement.pdf
├── go.mod
├── go.sum
├── README.md

