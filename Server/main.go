package main

import (
	"fmt"
	"time"

	"github.com/Imamachi-n/BioRxivGo/Server/route"
	"github.com/gin-gonic/gin"
)

const PORT = ":9000"

func main() {
	// Disable Console Color, you don't need console color when writing the logs to file.
	// gin.DisableConsoleColor()

	// Logging to a file
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)

	// Default With the Logger and Recovery middleware already attached
	router := gin.New()
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())

	// Simple group: api
	api := router.Group("/api")
	{
		api.GET("/articles", route.GetArticlesAll)
		api.GET("/ping", route.Ping)
		api.GET("/user/:name", route.GetName)           // /user/naoto
		api.GET("/user/:name/*action", route.GetAction) // /user/naoto/kick
		api.GET("/welcome", route.GetWelcome)           // /welcome?firstname=Naoto&lastname=Imamachi

		api.POST("/login", route.PostEcho)
	}

	router.Run(PORT) // listen and serve on 0.0.0.0:8080
}
