package apps

import (
	"github.com/Jblew/cairparavel/functions/app/apps/notificationsapp"
	"github.com/Jblew/cairparavel/functions/app/apps/usersapp"
	"github.com/golobby/container/pkg/container"
)

// Bind to IoC
func Bind(container *container.Container) {
	usersapp.Bind(container)
	notificationsapp.Bind(container)
}
