package models

import (
	"GOLANG/entities"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
)

var (
	listUser = make([]*entities.User, 0) // make a slice with init len(listUser) = 0
)

func HashString(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	sha256_hash := hex.EncodeToString(h.Sum(nil))
	return sha256_hash
}

func EncodeString(s string) string {
	data := []byte(s)
	encode := base64.StdEncoding.EncodeToString(data)
	return encode
}

func CreateUser(user *entities.User) bool {
	if user.Id != "" && user.Name != "" && user.Password != "" {
		if userF, _ := FindUser(user.Id); userF == nil {
			listUser = append(listUser, user)
			return true
		}
	}
	return false
}

func UpdateUser(eUser *entities.User) bool {
	for index, user := range listUser {
		if user.Id == eUser.Id {
			listUser[index] = eUser
			return true
		}
	}
	return false
}

func FindUser(id string) (*entities.User, error) {
	for _, user := range listUser {
		if user.Id == id {
			return user, nil
		}
	}
	return nil, errors.New("User does not exist")
}

func DeleteUser(id string) bool {
	for index, user := range listUser {
		if user.Id == id {
			copy(listUser[index:], listUser[index+1:])
			listUser[len(listUser)-1] = &entities.User{}
			listUser = listUser[:len(listUser)-1]
			return true
		}
	}
	return false
}

func GetAllUser() []*entities.User {
	return listUser
}
