package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/leehai1107/bipbip/cmd/banner"
	"github.com/leehai1107/bipbip/di/apifx"
	"github.com/leehai1107/bipbip/di/dbfx"
	"github.com/leehai1107/bipbip/pkg/config"
	"github.com/leehai1107/bipbip/pkg/errors"
	"github.com/leehai1107/bipbip/pkg/graceful"
	"github.com/leehai1107/bipbip/pkg/infra"
	"github.com/leehai1107/bipbip/pkg/logger"
	"github.com/leehai1107/bipbip/pkg/middleware/cors"
	"github.com/leehai1107/bipbip/pkg/recover"
	"github.com/leehai1107/bipbip/pkg/swagger"
	"github.com/leehai1107/bipbip/pkg/utils/ginbuilder"
	"github.com/leehai1107/bipbip/pkg/utils/timeutils"
	"github.com/leehai1107/bipbip/pkg/websocket"
	"github.com/leehai1107/bipbip/service/bid/delivery/http"
	"github.com/leehai1107/bipbip/sql"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Command of Internal Service",
	Long:  "CLI used to manage internal apis, datas when users access.",
	Run: func(_ *cobra.Command, _ []string) {
		NewServer().Run()
	},
	Version: "1.0.0",
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Command of Database Migration",
	Long:  "CLI used to manage database migration.",
	Run: func(_ *cobra.Command, _ []string) {
		infra.InitPostgresql()
		err := sql.NewMigration(infra.GetDB()).Migrate()
		if err != nil {
			logger.Fatalf("failed to migrate database: %v", err)
		}
	},
	Version: "1.0.0",
}

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run() {
	app := fx.New(
		fx.Invoke(config.InitConfig),
		fx.Invoke(initLogger),
		fx.Invoke(errors.Initialize),
		fx.Invoke(timeutils.Init),
		fx.Invoke(infra.InitPostgresql),
		fx.Invoke(websocket.InitHub),
		//... add module here
		dbfx.Module,
		apifx.Module,
		fx.Provide(provideGinEngine),
		fx.Invoke(
			registerService,
			registerSwaggerHandler),
		fx.Invoke(startServer),
		fx.Invoke(banner.Print),
	)
	logger.Info("Server started!")
	app.Run()
}

func provideGinEngine() *gin.Engine {
	return ginbuilder.BaseBuilder().Build()
}

func registerService(
	g *gin.Engine,
	router http.Router,
) {
	internal := g.Group("/internal")
	internal.Use(
		recover.RPanic,
		cors.CorsCfg(config.ServerConfig().CorsProduction))
	router.Register(internal)
}

func registerSwaggerHandler(g *gin.Engine) {
	swaggerAPI := g.Group("/internal/swagger")
	swag := swagger.NewSwagger()
	swaggerAPI.Use(swag.SwaggerHandler(config.ServerConfig().Production))
	swag.Register(swaggerAPI)
}

func startServer(lifecycle fx.Lifecycle, g *gin.Engine) {
	gracefulService := graceful.NewService(graceful.WithStopTimeout(time.Second), graceful.WithWaitTime(time.Second))

	gracefulService.Register(g)
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				port := fmt.Sprintf("%d", config.ServerConfig().HTTPPort)
				logger.Info("run on port:", port)
				go gracefulService.StartServer(g, port)
				return nil
			},
			OnStop: func(context.Context) error {
				gracefulService.Close()
				infra.ClosePostgresql() // nolint
				return nil
			},
		},
	)
}
func initLogger() {
	logger.Initialize(config.ServerConfig().Logger)
}
