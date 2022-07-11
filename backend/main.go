package main

import (
	"net/http"

	"github.com/6675-voting-system/voting-system/backend/contract"
	data_access "github.com/6675-voting-system/voting-system/backend/data-access"
	"github.com/6675-voting-system/voting-system/backend/handler"
	"github.com/6675-voting-system/voting-system/backend/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	data_access.InitDB()
	contract.SetupContract()
	handler.InitAdminEvents()
	r := setupRouter()
	// Listen and Server in localhost:8080
	r.Run(":8080")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("adminsession", store))

	r.Use(CORSMiddleware())

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST("/", handler.LoginHandler)

	user := r.Group("/user")
	user.POST("/vote-details/", handler.UserVoteDetailsContractHandler)
	user.POST("/vote/", handler.UserVoteSubmitContractHandler)

	r.POST("/admin_register", handler.AdminRegisterHandler)
	r.POST("/admin_login", handler.AdminLoginHandler)
	admin := r.Group("/admin")
	admin.Use(middleware.Authentication())
	admin.POST("/detail", handler.AdminDetailContractHandler)
	admin.POST("/create_event", handler.AdminCreateEventContractHandler)
	admin.POST("/end_event", handler.AdminEndEventContractHandler)
	admin.POST("/get_event", handler.AdminGetEventContractHandler)
	admin.POST("/download", handler.DownloadHandler)

	return r
}
