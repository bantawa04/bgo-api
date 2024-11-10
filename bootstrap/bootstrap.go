package bootstrap

import (
	"context"
	"github.com/bantawa04/bgo-api/config"
	"github.com/bantawa04/bgo-api/controllers"
	"github.com/bantawa04/bgo-api/routes"

	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	controllers.Module,
	routes.Module,
	config.Module,

	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	handler config.Router,
	routes routes.Routes,
	env config.Env,
	logger config.Logger,
	database config.Database,
) {

	appStop := func(context.Context) error {
		logger.Zap.Info("Stopping Application")
		conn, _ := database.DB.DB()
		_ = conn.Close()
		return nil
	}

	//if utils.IsCli() {
	//	lifecycle.Append(fx.Hook{
	//		OnStart: func(context.Context) error {
	//			logger.Zap.Info("Starting boilerplate cli Application")
	//			logger.Zap.Info("------ 🤖 Boilerplate 🤖 (CLI) ------")
	//			go cliApp.Start()
	//			return nil
	//		},
	//		OnStop: appStop,
	//	})
	//
	//	return
	//}

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Zap.Info("Starting Application")
			logger.Zap.Info("------------------------")
			logger.Zap.Info("------ Boilerplate 📺 ------")
			logger.Zap.Info("------------------------")

			go func() {
				//if env.Environment != "production" && env.HOST != "" {
				//	logger.Zap.Info("Setting Swagger Host...")
				//	docs.SwaggerInfo.Host = env.HOST
				//}

				//if env.Environment == "production" {
				//	logger.Zap.Info("Migrating DB schema...")
				//	migrations.Migrate()
				//}
				//middlewares.Setup()
				//routes.Setup()
				//logger.Zap.Info("🌱 seeding data...")
				//seeds.Run()

				if env.ServerPort == "" {
					_ = handler.Gin.Run()
				} else {
					_ = handler.Gin.Run(":" + env.ServerPort)
				}
			}()
			return nil
		},
		OnStop: appStop,
	})
}
