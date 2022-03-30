package server

import (
	"github.com/gin-gonic/gin"
	quieres "github.com/nexsabre/mikropoker-go/db"
	"gorm.io/gorm"
)

func Start(db *gorm.DB) {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"objetcs": quieres.GetSessions(db)})
	})

	r.Run()
}
