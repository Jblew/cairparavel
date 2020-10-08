package messenger

// Init initializes the app
func Init(config Config) *Messenger {
	return &Messenger{
		Config: config,
	}
}
