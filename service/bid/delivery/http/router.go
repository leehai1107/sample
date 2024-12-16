package http

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/leehai1107/bipbip/pkg/apiwrapper"
	"github.com/leehai1107/bipbip/pkg/logger"
)

type Router interface {
	Register(routerGroup gin.IRouter)
}

type routerImpl struct {
	handler IHandler
}

func NewRouter(
	handler IHandler,
) Router {
	return &routerImpl{
		handler: handler,
	}
}

func (p *routerImpl) Register(r gin.IRouter) {
	lg := logger.EnhanceWith(context.Background())
	lg.Infow("RegisterRouterStart!")
	//routes for apis
	api := r.Group("api/v1")
	{
		api.GET("/ping", apiwrapper.Wrap(func(c *gin.Context) *apiwrapper.Response {
			return apiwrapper.SuccessWithDataResponse(time.Now())
		}))
	}

	userApi := r.Group("api/v1/user")
	{
		userApi.POST("/login", p.handler.Login)
		userApi.POST("/register", p.handler.Register)
	}

	websocket := r.Group("ws")
	{
		websocket.GET("/:roomId", p.handler.ServeWS)
	}
}
