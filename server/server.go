package server

import (
	"fmt"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
	quieres "github.com/nexsabre/mikropoker-go/db"
	"github.com/nexsabre/mikropoker-go/schema"
	"gorm.io/gorm"
)

const SESSION_ID = "session_id"
const SESSION = "session"
const USER = "user"

func Start(db *gorm.DB) {
	r := gin.Default()

	// SESSIONS
	r.GET("/s", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"sessions": quieres.GetSessions(db)})
	})

	// - reveal points
	r.PATCH("/s/:session_id", func(ctx *gin.Context) {
		session_reveal := &schema.SessionReveal{}
		session_id := ctx.Param(SESSION_ID)

		if err := ctx.Bind(&session_reveal); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
		}
		sessionStatus := quieres.RevealSession(db, Atoi(session_id), session_reveal.Reveal)
		ctx.JSON(http.StatusOK, gin.H{SESSION: sessionStatus})
	})

	// - create session
	r.POST("/s", func(ctx *gin.Context) {
		session_create := schema.SessionCreate{}

		if err := ctx.Bind(&session_create); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		newSession := quieres.CreateSession(db, session_create)

		ctx.JSON(http.StatusCreated, gin.H{SESSION: newSession})
	})

	r.GET("/s/:session_id", func(ctx *gin.Context) {
		session_id := ctx.Param(SESSION_ID)
		ctx.JSON(http.StatusOK, gin.H{SESSION: quieres.GetSession(db, Atoi(session_id))})
	})

	// USERS
	r.POST("/u/:session_id", func(ctx *gin.Context) {
		session_id := ctx.Param(SESSION_ID)
		points_add := schema.UserPoints{}

		if err := ctx.Bind(&points_add); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		userPoints := quieres.UserPoints(db, Atoi(session_id), points_add)
		ctx.JSON(http.StatusCreated, gin.H{USER: userPoints})
	})

	r.Run()
}

func Atoi(str string) int {
	intVar, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("Cannot conver %s into int\n", str)
	}
	return intVar
}
