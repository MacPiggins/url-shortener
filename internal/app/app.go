package app

import (
	"fmt"
	"url-shortener/internal/database"
	"url-shortener/internal/database/noDB"
	"url-shortener/internal/database/postgres"
	"url-shortener/internal/database/redis"
	"url-shortener/internal/transport/rest"

	"github.com/gin-gonic/gin"
)

func Run(storage, link string) (err error) {
	var db database.Wrapper

	if storage == "redis" {
		db, err = redis.New(link)
	} else if storage == "postgres" {
		db, err = postgres.New(link)
	} else {
		db, err = noDB.New()
	}

	if err != nil {
		fmt.Println("unable to connect to DataBase: ", err)
		return err
	}

	defer db.Close()

	server := gin.Default()

	server.GET("/api/retrieve", func(ctx *gin.Context) {
		var token rest.Link
		token.Data = ctx.Query("token")
		link, err := rest.RetrieveLink(token.Data, db)
		if err != nil {
			fmt.Println("an error occured while retrieving link: ", err)
			ctx.IndentedJSON(500, "an error occured on the server")
			return
		}
		ctx.IndentedJSON(200, link)
	})

	server.POST("/api/create", func(ctx *gin.Context) {
		var link rest.Link
		err := ctx.ShouldBindJSON(&link)
		if err != nil {
			fmt.Println("error on binging json: ", err)
			ctx.IndentedJSON(500, "an error occured on the server")
			return
		}
		token, err := rest.CreateToken(link.Data, db)
		if err != nil {
			fmt.Println("an error occured while creating token: ", err)
			ctx.IndentedJSON(500, "an error occured on the server")
			return
		}

		ctx.IndentedJSON(200, token)
	})

	return server.Run()
}
