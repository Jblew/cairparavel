package app

import (
	"github.com/Jblew/cairparavel/functions/app/apps"
	"github.com/Jblew/cairparavel/functions/app/lib"
	"github.com/Jblew/ioccontainer/pkg/ioccontainer"
)

// Bind to IoC
func Bind(container *ioccontainer.Container) {
	apps.Bind(container)
	lib.Bind(container)
}
