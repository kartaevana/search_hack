package delivery

import (
	"back/cmd/docs"
	"back/internal/delivery/handlers"
	"back/internal/delivery/middleware"
	"back/internal/repository/approve"
	"back/internal/repository/form"
	"back/internal/repository/team"
	"back/internal/repository/user"
	approveserv "back/internal/service/approve"
	formserv "back/internal/service/form"
	teamserv "back/internal/service/team"
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
	formRouter.GET("/all", formHandler.GetAllForm)

	teamRouter := r.Group("/team")

	teamRepo := team.InitTeamRepository(db)
	teamService := teamserv.InitTeamService(teamRepo, log)
	teamHandler := handlers.InitTeamHandler(teamService)

	teamRouter.POST("/create", teamHandler.CreateTeam)
	teamRouter.GET("/:id", teamHandler.GetTeam)
	teamRouter.POST("/add/:id_team/:id_user", teamHandler.AddUserTeam)

	approveRouter := r.Group("/approve")

	approveRepo := approve.InitApproveRepository(db)
	approveService := approveserv.InitApproveService(approveRepo, log)
	approveHandler := handlers.InitApproveHandler(approveService)

	approveRouter.POST("/create", approveHandler.CreateApprove)
	approveRouter.POST("/reject", approveHandler.RejectApprove)
	approveRouter.GET("/all", approveHandler.GetAllApprove)

	if err := r.Run("0.0.0.0:8080"); err != nil {
		panic(fmt.Sprintf("error running client: %v", err.Error()))
	}
}
