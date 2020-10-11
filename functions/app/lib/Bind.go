package lib

import (
	"github.com/Jblew/cairparavel/functions/app/lib/messenger"
	"github.com/golobby/container/pkg/container"
)

// Bind to IoC
func Bind(container container.Container) {
	container.Singleton(func(config messenger.Config) *messenger.Messenger {
		return messenger.Init(config)
	})
}
