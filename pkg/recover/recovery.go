package recover

import (
	"github.com/gin-gonic/gin"
	"github.com/leehai1107/bipbip/pkg/apiwrapper"
	"github.com/leehai1107/bipbip/pkg/errors"
	"github.com/leehai1107/bipbip/pkg/logger"
)

func RPanic(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			logger.EnhanceWith(c.Request.Context()).Errorf("method %v, path %v, err %v",
				c.Request.Method,
				c.Request.URL.EscapedPath(),
				err,
			)
			apiwrapper.Abort(c, &apiwrapper.Response{Error: errors.InternalServerError.New()})
		}
	}()

	c.Next()
}
