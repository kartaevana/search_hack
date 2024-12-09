package delivery

import (
	"back/cmd/docs"
	"back/internal/delivery/handlers"
	"back/internal/delivery/middleware"
	"back/internal/repository/form"
	"back/internal/repository/user"
	formserv "back/internal/service/form"
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
	middlewareStruct := middleware.InitMiddleware(log)
	r.Use(middlewareStruct.CORSMiddleware())

	userRouter := r.Group("/user")

	userRepo := user.InitUserRepository(db)
	userService := userserv.InitUserService(userRepo, log)
	userHandler := handlers.InitUserHandler(userService)

	userRouter.POST("/create", userHandler.CreateUser)
	userRouter.GET("/:id", userHandler.GetUser)
	userRouter.POST("/login", userHandler.LoginUser)

	formRouter := r.Group("/form")

	formRepo := form.InitFormRepository(db)
	formService := formserv.InitFormService(formRepo, log)
	formHandler := handlers.InitFormHandler(formService)

	formRouter.POST("/create", formHandler.CreateForm)
	formRouter.GET("/:id", formHandler.GetForm)

	if err := r.Run("0.0.0.0:8080"); err != nil {
		panic(fmt.Sprintf("error running client: %v", err.Error()))
	}
}
