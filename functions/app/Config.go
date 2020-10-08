package app

import "github.com/Jblew/cairparavel/functions/app/lib/messenger"

// Config is functions main config
type Config struct {
	ProjectID string
	Messenger messenger.Config
}
