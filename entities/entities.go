package entities

import (
	"fmt"
)

// User is a user's info struct
type User struct {
	Id       string `json:"id"` // Định nghĩa tên trường hiển thị trong JSON
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Url struct {
	ID       string `json:"id"`
	UserName string `json:"username"`
	LongUrl  string `json:"longUrl"`
	ShortUrl string `json:"shortUrl"`
}

func (user User) ToString() string {
	return fmt.Sprintf("id: %s\nName: %s\nPassword: %s\n", user.Id, user.Name, user.Password)
}
