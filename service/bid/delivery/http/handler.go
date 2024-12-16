package http

import (
	"github.com/gin-gonic/gin"
	"github.com/leehai1107/bipbip/pkg/websocket"
	"github.com/leehai1107/bipbip/service/bid/usecase"
)

type IHandler interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	ServeWS(ctx *gin.Context)
}

type Handler struct {
	usecase usecase.IUserUsecase
}

func NewHandler(usecase usecase.IUserUsecase) IHandler {
	return &Handler{
		usecase: usecase,
	}
}

func (h *Handler) Login(ctx *gin.Context) {
	// TODO
}

func (h *Handler) Register(ctx *gin.Context) {
	// TODO
}

func (h *Handler) ServeWS(ctx *gin.Context) {
	roomId := ctx.Param("roomId")
	// TODO: Test websocket handler
	websocket.ServeWs(ctx, roomId)
}
