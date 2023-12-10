package main

import (
	"duit_example/internal"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	cors "github.com/itsjamie/gin-cors"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Server struct {
	engine *gin.Engine
}

func NewServer() *Server {
	return &Server{
		engine: gin.New(),
	}
}

func wsHandler(ginContext *gin.Context) {
	_, err := upgrader.Upgrade(ginContext.Writer, ginContext.Request, nil)
	fmt.Println("Connected")
	if err != nil {
		log.Fatal(err)
	}

	//create form and send to client when connected
	// view := internal.Form1()
	// wsSession.WriteMessage(websocket.TextMessage, view)
}

func (server *Server) Run(addr string) error {
	server.engine.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE, OPTIONS",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          500 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))
	server.engine.Use(
		gin.Recovery(),
	)

	return server.engine.Run(":" + addr)
}

func (server *Server) Engine() *gin.Engine {
	return server.engine
}

func ConfigureRoutes(server *Server) {

	eng := server.engine

	eng.Use(func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, ResponseType, accept, origin, Cache-Control, X-Requested-With")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(200)
			return
		}
		ctx.Next()
	})

	eng.Any("/", func(ctx *gin.Context) {
		isWs := ctx.IsWebsocket()

		if isWs {
			wsHandler(ctx)
		} else {
			fmt.Println("Not websocket")
		}
	})

	eng.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "check",
		})
	})

	eng.GET("/decoratedbox", func(ctx *gin.Context) {
		view := internal.DecoratedBoxExample()
		ctx.Data(200, "application/json", view)
	})

	eng.GET("/img", func(ctx *gin.Context) {
		view := internal.ImageExample()

		ctx.Data(200, "application/json", view)
	})
}

func main() {
	app := NewServer()

	ConfigureRoutes(app)

	if err := app.Run("8999"); err != nil {
		fmt.Println(err.Error())
	}
}
