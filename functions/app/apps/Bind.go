package apps

import (
	"github.com/Jblew/cairparavel/functions/app/apps/eventsapp"
	"github.com/Jblew/cairparavel/functions/app/apps/messengerapp"
	"github.com/Jblew/cairparavel/functions/app/apps/notificationsapp"
	"github.com/Jblew/cairparavel/functions/app/apps/usersapp"
	"github.com/Jblew/ioccontainer/pkg/ioccontainer"
)

// Bind to IoC
func Bind(container *ioccontainer.Container) {
	usersapp.Bind(container)
	notificationsapp.Bind(container)
	eventsapp.Bind(container)
	messengerapp.Bind(container)
}
