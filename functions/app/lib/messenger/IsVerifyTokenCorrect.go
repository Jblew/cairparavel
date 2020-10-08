package messenger

// IsVerifyTokenCorrect verifues the FB messenger verify token
func (messenger *Messenger) IsVerifyTokenCorrect(verifyToken string) bool {
	return verifyToken == messenger.Config.VerifyToken
}
