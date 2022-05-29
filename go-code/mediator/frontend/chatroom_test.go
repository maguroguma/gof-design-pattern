package frontend

import "testing"

func TestChatRoom(t *testing.T) {
	chatroom := NewChatRoom()
	user := NewUser("Yokoyama", chatroom)

	msg := user.send("Hello")
	expected := "POST: Yokoyama send message 'Hello'"
	if msg != expected {
		t.Errorf("got %v, want %v", msg, expected)
	}
}
