package messenger

// Init initializes the app
func Init(config Config) (*Messenger, error) {
	return &Messenger{
		Config: config,
	}, nil
}
