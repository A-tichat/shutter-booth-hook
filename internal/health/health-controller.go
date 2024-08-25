package health

import "git.akyoto.dev/go/web"

type HealthController struct {
	CheckHealth func(ctx web.Context) error
}

func New() *HealthController {
	return &HealthController{}
}

func CheckHealth(ctx web.Context) error {
	return ctx.String("OK")
}
