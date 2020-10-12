package lib

import (
	"github.com/Jblew/cairparavel/functions/app/lib/messenger"
	"github.com/Jblew/ioccontainer/pkg/ioccontainer"
)

// Bind to IoC
func Bind(container *ioccontainer.Container) {
	container.Singleton(func(config messenger.Config) messenger.Messenger {
		return *messenger.Init(config)
	})
}
