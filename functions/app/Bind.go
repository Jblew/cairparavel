package app

import (
	"github.com/Jblew/cairparavel/functions/app/apps"
	"github.com/golobby/container/pkg/container"
)

// Bind to IoC
func Bind(container *container.Container) {
	apps.Bind(container)
}
