package server

import (
    "net/http"
    "strings"
    "github.com/madhubolla2205/bank-analyzer/pkg/models"
    "github.com/gin-gonic/gin"
    "github.com/madhubolla2205/bank-analyzer/pkg/reader"
    "github.com/madhubolla2205/bank-analyzer/pkg/processor"
)

func SetupServer() *gin.Engine {
    r := gin.Default()
    r.LoadHTMLGlob("static/*")
    r.Static("/uploads", "./uploads")

    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", nil)
    })

    r.POST("/upload", func(c *gin.Context) {
        file, _ := c.FormFile("file")
        savePath := "./uploads/" + file.Filename
        c.SaveUploadedFile(file, savePath)

        var txns []models.Transaction
        var err error

        if strings.HasSuffix(file.Filename, ".csv") {
            txns, err = reader.ParseCSV(savePath)
        } else if strings.HasSuffix(file.Filename, ".pdf") {
            txns, err = reader.ParsePDF(savePath)
        } else {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported file type"})
            return
        }

        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        categorized := processor.CategorizeTransactions(txns)
        processor.StoreData(categorized)

        c.Redirect(http.StatusMovedPermanently, "/")
    })

    r.GET("/data", func(c *gin.Context) {
        c.JSON(http.StatusOK, processor.GetData())
    })

    return r
}

