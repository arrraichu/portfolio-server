package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	dbPosts "arrraichu/portfolio-server/db"
	content "arrraichu/portfolio-server/internal"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

var db *sql.DB

func getContentTypes(ctx *gin.Context) {
	var contentTypes []content.ContentType

	rows, err := db.Query("SELECT code, label, reqs FROM content_type")
	if err != nil {
		panic(err) // todo: FIX ME
	}
	defer rows.Close()

	for rows.Next() {
		var c content.ContentType
		if err := rows.Scan(&c.Code, &c.Label, &c.Reqs); err != nil {
			panic(err) // todo: FIX ME
		}
		contentTypes = append(contentTypes, c)
	}

	ctx.IndentedJSON(http.StatusOK, contentTypes)
}

func getAllContent(ctx *gin.Context) {

}

func postContent(ctx *gin.Context) {
	var err error

	var input content.Content
	if err = ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": err.Error()})
		return
	}

	switch input.Type {
	case "text":
		err = dbPosts.PostTextContent(db, input)
	default:
		err = fmt.Errorf("Unknown content type: %s", input.Type)
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"ok": true, "error": nil})
}

func main() {
	cfg := pq.Config{
		Host:           os.Getenv("DBHOST"),
		Port:           5432,
		User:           os.Getenv("DBUSER"),
		Password:       os.Getenv("DBPASS"),
		Database:       os.Getenv("DBNAME"),
		ConnectTimeout: 5 * time.Second,
		SSLMode:        pq.SSLMode("prefer"),
	}
	c, err := pq.NewConnectorConfig(cfg)
	if err != nil {
		log.Fatal(err)
	}
	db = sql.OpenDB(c)
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.GET("/types", getContentTypes)
	router.GET("/content", getAllContent)
	router.POST("/content", postContent)

	log.Println("Server listening on port 8080!")
	log.Fatal(router.Run(":8080"))
}
