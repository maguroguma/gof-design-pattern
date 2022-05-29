package frontend

// メディエータに統治・調整されるユーザ
type User struct {
	name     string
	chatroom *ChatRoom
}

func NewUser(name string, chatroom *ChatRoom) *User {
	res := new(User)
	res.name = name
	res.chatroom = chatroom
	return res
}

func (u *User) getName() string {
	return u.name
}

func (u *User) send(message string) string {
	return u.chatroom.logMessage(u.getName(), message)
}
