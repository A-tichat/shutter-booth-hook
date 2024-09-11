package healthCtr

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func CheckHealth(ctx context.Context, c *app.RequestContext) {
	c.String(consts.StatusOK, "OK")
}
