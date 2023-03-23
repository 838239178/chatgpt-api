package main

import (
	"chatgpt-api/api"
	"chatgpt-api/config"
	"flag"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func init() {
	_ = config.APIKey()
}

func main() {
	portPtr := flag.String("port", "8080", "serve port")
	flag.Parse()

	engine := gin.Default()

	engine.Use(sessions.Sessions("chatgpt-api", cookie.NewStore([]byte("cookie-secret"))))

	new(api.HealthAPI).RegisterRoute(engine)
	v1 := engine.Group("v1")
	new(api.ChatAPI).RegisterRoute(v1)

	_ = engine.Run(":" + *portPtr)
}
