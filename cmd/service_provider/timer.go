package service_provider

import (
	"github.com/Kiveri/wh-be/internal/pkg"
)

func (sp *ServiceProvider) getTimer() *pkg.Timer {
	if sp.timer == nil {
		sp.timer = pkg.NewTimer()
	}

	return sp.timer
}
