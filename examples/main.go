package main

import (
	"duit_example/internal"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	cors "github.com/itsjamie/gin-cors"
	"github.com/lesleysin/duit_go/pkg/duit_attributes"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_color"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_text_properties"
	"github.com/lesleysin/duit_go/pkg/duit_core"
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
	wsSession, err := upgrader.Upgrade(ginContext.Writer, ginContext.Request, nil)
	fmt.Println("Connected")
	if err != nil {
		log.Fatal(err)
	}

	//create form and send to client when connected
	view := internal.WsExample()
	wsSession.WriteMessage(websocket.TextMessage, view)
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

	eng.GET("/stack", func(ctx *gin.Context) {
		view := internal.StackExample()

		ctx.Data(200, "application/json", view)
	})

	eng.GET("/inputs", func(ctx *gin.Context) {
		fmt.Println("OK")
		view := internal.InputsExample()

		ctx.Data(200, "application/json", view)
	})

	eng.GET("/textInput1change", func(ctx *gin.Context) {
		fmt.Println(ctx.Request.URL)

		ctx.Data(200, "application/json", []byte("{}"))
	})

	eng.POST("/textInput2change", func(ctx *gin.Context) {
		var data interface{}

		err := ctx.BindJSON(&data)

		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(data)

		ctx.Data(200, "application/json", []byte("{}"))
	})

	eng.GET("/slider", func(ctx *gin.Context) {
		fmt.Println("Slider end change action")
		ctx.Data(200, "application/json", []byte("{}"))
	})

	eng.POST("/apply", func(ctx *gin.Context) {
		var data map[string]interface{}

		err := ctx.BindJSON(&data)

		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(data)

		var checkboxValue string
		var textColor duit_color.ColorString

		if data["checkbox"].(bool) == true {
			checkboxValue = "checked"
			textColor = "#58eb34"
		} else {
			checkboxValue = "unchecked"
			textColor = "#eb3734"
		}

		updates := map[string]interface{}{
			"text1": &duit_attributes.TextAttributes[duit_color.ColorString]{
				Data: data["value1"].(string),
			},
			"text2": &duit_attributes.TextAttributes[duit_color.ColorString]{
				Data: data["value2"].(string),
			},
			"text3": &duit_attributes.TextAttributes[duit_color.ColorString]{
				Data: checkboxValue,
				Style: &duit_text_properties.TextStyle[duit_color.ColorString]{
					Color: textColor,
				},
			},
		}

		var payload = duit_core.NewUpdateEvent(updates)

		val, err := json.Marshal(payload)

		if err != nil {
			fmt.Println(err.Error())
		}

		ctx.Data(200, "application/json", val)
	})

	eng.GET("/updateLayout", func(ctx *gin.Context) {
		data := duit_core.NewLayoutUpdateEvent(internal.UpdatePayload())

		val, err := json.Marshal(data)

		if err != nil {
			fmt.Println(err.Error())
		}

		ctx.Data(200, "application/json", val)
	})

	eng.GET("/updateExample", func(ctx *gin.Context) {
		view := internal.UpdateExample()

		ctx.Data(200, "application/json", view)
	})

	eng.GET("/gesture", func(ctx *gin.Context) {
		view := internal.GestureExample()

		ctx.Data(200, "application/json", view)
	})

	eng.GET("/tap", func(ctx *gin.Context) {
		fmt.Println("Tap")
		ctx.Data(200, "application/json", []byte("{}"))
	})

	eng.GET("/longPress", func(ctx *gin.Context) {
		fmt.Println("LongPress")
		ctx.Data(200, "application/json", []byte("{}"))
	})

	eng.GET("/transform", func(ctx *gin.Context) {
		view := internal.TransformExample()
		ctx.Data(200, "application/json", view)
	})

	eng.GET("/rich", func(ctx *gin.Context) {
		view := internal.RichTextExample()
		ctx.Data(200, "application/json", view)
	})

	eng.GET("/shoes", func(ctx *gin.Context) {
		view := internal.ComponentExample()
		ctx.Data(200, "application/json", view)
	})

	eng.GET("/send_component_update", func(ctx *gin.Context) {

		updates := map[string]interface{}{
			"shoes_card1": map[string]interface{}{
				"image":       "https://resizer.mail.ru/p/b0685008-5021-5715-a506-e73621c958f5/AQAFGR6McYd06E2eQ8J-GzknNO9eDHZcCIpgrre-K3SwrZ0QnmOBQVPyD6yWisBlYJYZ38YEM768jMY1M4m4NYbQkmM.jpg",
				"description": "Топовые и дорогие",
				"primary":     "Ботильоны",
				"cost":        `$228`,
			},
		}

		var payload = duit_core.NewUpdateEvent(updates)

		val, err := json.Marshal(payload)

		if err != nil {
			fmt.Println(err.Error())
		}

		ctx.Data(200, "application/json", val)
	})

	eng.GET("/scrollview", func(ctx *gin.Context) {
		view := internal.ScrollviewExample()
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
