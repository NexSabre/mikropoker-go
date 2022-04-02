package server

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	quieres "github.com/nexsabre/mikropoker-go/db"
	"github.com/nexsabre/mikropoker-go/schema"
	"gorm.io/gorm"
)

const SESSION_ID = "session_id"
const SESSION = "session"

type TPORT struct {
	port string `json:"port"`
}

func Start(db *gorm.DB) {
	r := gin.Default()
	r.Static("/assets", "./dist/assets/")
	r.StaticFile("/", "./dist/index.html")
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:3000", "https://mikropoker.com", "https://www.mikropoker.com", "https://mipo.app", "https://www.mipo.app"},
		AllowMethods:  []string{"GET", "DELETE", "POST", "PUT", "PATCH"},
		AllowHeaders:  []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}))

	go h.run()

	// websocket for session_id
	r.GET("/ws/:room_id", func(c *gin.Context) {
		roomId := Atoi(c.Param("room_id"))
		serveWs(c.Writer, c.Request, strconv.Itoa(roomId))
	})

	r.GET("/port", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, &TPORT{port: os.Getenv("PORT")})
	})

	// SESSIONS
	r.GET("/s", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, quieres.GetSessions(db))
	})

	// - create session
	r.POST("/s", func(ctx *gin.Context) {
		session_create := schema.SessionCreate{}

		if err := ctx.Bind(&session_create); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		newSession := quieres.CreateSession(db, session_create)

		ctx.JSON(http.StatusCreated, newSession)
	})

	r.GET("/s/:session_id", func(ctx *gin.Context) {
		session_id := ctx.Param(SESSION_ID)
		ctx.JSON(http.StatusOK, quieres.GetSession(db, Atoi(session_id)))
	})

	// - restart session
	r.DELETE("/s/:session_id", func(ctx *gin.Context) {
		sessionID := ctx.Param(SESSION_ID)

		quieres.RestartSession(db, Atoi(sessionID))
		ctx.JSON(http.StatusNoContent, gin.H{})
	})

	// - reveal points
	r.PATCH("/s/:session_id", func(ctx *gin.Context) {
		sessionReveal := &schema.SessionReveal{}
		sessionId := ctx.Param(SESSION_ID)

		fmt.Printf("%+v", sessionReveal)
		if err := ctx.Bind(&sessionReveal); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
		}
		sessionStatus := quieres.RevealSession(db, Atoi(sessionId), sessionReveal.Reveal)
		ctx.JSON(http.StatusOK, sessionStatus)
	})

	r.POST("/s/:session_id", func(ctx *gin.Context) {
		session_id := ctx.Param(SESSION_ID)
		points_add := schema.UserPoints{}

		if err := ctx.Bind(&points_add); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		userPoints := quieres.UserPoints(db, Atoi(session_id), points_add)
		ctx.JSON(http.StatusCreated, userPoints)
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
