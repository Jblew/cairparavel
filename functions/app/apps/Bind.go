package apps

import (
	"github.com/golobby/container/pkg/container"
	"github.com/Jblew/cairparavel/functions/app/apps/usersapp"

)

// Bind to IoC
func Bind(container *container.Container) {
	usersapp.Bind(container)
}
