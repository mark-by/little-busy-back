package redis

func (s Session) sessionKey(ID string) string {
	return "sessions:" + ID
}
