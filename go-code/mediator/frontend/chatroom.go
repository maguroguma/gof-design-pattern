package frontend

import "fmt"

// メディエータに相当
type ChatRoom struct{}

func NewChatRoom() *ChatRoom {
	return new(ChatRoom)
}

func (cr *ChatRoom) logMessage(userName, message string) string {
	logString := fmt.Sprintf("POST: %s send message '%s'", userName, message)
	fmt.Println(logString)
	return logString
}
