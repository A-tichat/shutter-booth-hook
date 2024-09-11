package health

import (
	"git.akyoto.dev/go/web"
)

type HealthController struct {
}

func New() HealthController {
	return HealthController{}
}

func (ctr HealthController) CheckHealth(c web.Context) error {
	return c.String("OK")
}
