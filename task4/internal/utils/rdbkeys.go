package util

import "fmt"

func GetAuthKey(sessionId string) string {
	authKey := fmt.Sprintf("auth_key:%s", sessionId)
	return authKey
}

func GetSessionIdKey(userName string) string {
	sessionIdKey := fmt.Sprintf("session_id:%s", userName)
	return sessionIdKey
}
