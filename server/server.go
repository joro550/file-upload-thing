package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
 var Db *gorm.DB

func RunServer(){
    Db, _ = ConnectToDatabase()

    r := gin.Default();
    r.LoadHTMLGlob("server/static/**/*.html")

    r.GET("/ping", func(ctx *gin.Context) {
        ctx.JSON(http.StatusTeapot, gin.H {
            "message": "pong",
        })
    })

    r.GET("/", func(c *gin.Context){
        c.HTML(http.StatusOK, "main.html", gin.H {
            "title":"Home",
            "file":"main.html",
            "model": 3,
        })
    })

    r.POST("/api/folder", func(ctx *gin.Context) {
        request := CreateFolderRequest{}

        err := ctx.ShouldBindJSON(&request)
        if err != nil {
            ctx.Status(http.StatusBadRequest)
            return
        }

        Db.Create(&Folder{ Name: request.Name})
        ctx.Status(http.StatusCreated)
    })

    r.Run("localhost:8080")
}


type CreateFolderRequest struct {
    Name string
}
