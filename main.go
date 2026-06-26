package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	contentDb "arrraichu/portfolio-server/db"
	content "arrraichu/portfolio-server/internal"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func getContentTypes(ctx *gin.Context) {
	types, err := contentDb.FetchContentTypes(db)
	if err != nil {
		log.Printf("!! ERROR !! %s\n", err.Error())
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"ok": false, "error": err.Error()},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		types,
	)
}

func getAllContent(ctx *gin.Context) {
	path := ctx.Query("page_path")
	if path == "" {
		err := fmt.Errorf("Request must contain a page_path query parameter.")
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"ok": false, "error": err.Error()},
		)
		return
	}

	posts, err := contentDb.FetchContent(db, path)
	if err != nil {
		log.Printf("!! ERROR !! %s\n", err.Error())
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"ok": false, "error": err.Error()},
		)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"ok": true, "error": nil, "posts": posts})
}

func postContent(ctx *gin.Context) {
	var err error

	var input content.Content
	if err = ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": err.Error()})
		return
	}

	switch input.Type {
	case "title":
		err = contentDb.PostTitle(db, input)
	case "disclaimer":
		err = contentDb.PostDisclaimer(db, input)
	case "text":
		err = contentDb.PostTextContent(db, input)
	case "text_buttons":
		err = contentDb.PostTextButtonsContent(db, input)
	default:
		err = fmt.Errorf("Unknown content type: %s", input.Type)
	}

	if err != nil {
		log.Printf("%v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": err.Error()})
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{"ok": true, "error": nil},
	)
}

func main() {
	var err error

	dbsrc := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 connect_timeout=5 sslmode=prefer",
		os.Getenv("DBHOST"),
		os.Getenv("DBUSER"),
		os.Getenv("DBPASS"),
		os.Getenv("DBNAME"),
	)
	db, err = sqlx.Open("postgres", dbsrc)
	if err != nil {
		log.Fatal(err)
	}
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
