package delivery

import (
	"back/cmd/docs"
	"back/internal/delivery/handlers"
	"back/internal/repository/user"
	userserv "back/internal/service/user"
	"back/pkg/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Start(db *sqlx.DB, log *log.Logs) {
	r := gin.Default()
	r.ForwardedByClientIP = true
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userRouter := r.Group("/user")

	userRepo := user.InitUserRepository(db)
	userService := userserv.InitUserService(userRepo, log)
	userHandler := handlers.InitUserHandler(userService)

	userRouter.POST("/create", userHandler.CreateUser)
	userRouter.GET("/:id", userHandler.GetUser)
	userRouter.POST("/login", userHandler.LoginUser)

	if err := r.Run("0.0.0.0:8080"); err != nil {
		panic(fmt.Sprintf("error running client: %v", err.Error()))
	}
}
