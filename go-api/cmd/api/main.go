package main

import (
	"net/http"

	"github.com/KelsonTeixeira/go/go-api/database"
	"github.com/KelsonTeixeira/go/go-api/internal"
	"github.com/KelsonTeixeira/go/go-api/internal/post"
	"github.com/gin-gonic/gin"
)

func main() {
	connectionString := "postgresql://posts:mypassword@db:5432/posts"
	conn, err := database.NewConnection(connectionString)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	repo := post.Repository{
		Conn: conn,
	}

	repo = repo

	g := gin.Default()
	g.POST("/posts", func(ctx *gin.Context) {
		var post internal.Post
		if err := ctx.BindJSON(&post); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := repo.Insert(post); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err.Error(),
			})
			return
		}

	})

	g.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	g.Run(":3000")
}
