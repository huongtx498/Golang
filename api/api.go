package api

import (
	"GOLANG/entities"
	"GOLANG/models"
	"encoding/json"
	"math/rand"
	"net/http"
)

func HashUrl(response http.ResponseWriter, request *http.Request) {
	urls, ok1 := request.URL.Query()["url"]
	userNames, ok2 := request.URL.Query()["username"]
	if !ok1 || !ok2 || len(urls) < 1 || len(userNames) < 1 {
		responseWithError(response, http.StatusBadRequest, "Url Param id is missing")
		return
	}
	userNameHash := models.HashString(userNames[0])
	urlHash := models.HashString(urls[0] + userNameHash + string(rand.Intn(1000000)))
	shortenUrl := models.EncodeString(urlHash)
	shorten := "http://misa/" + shortenUrl[:6]
	responseWithJSON(response, http.StatusOK, shorten)

}

func FindUser(response http.ResponseWriter, request *http.Request) {
	ids, ok := request.URL.Query()["id"]
	if !ok || len(ids) < 1 {
		responseWithError(response, http.StatusBadRequest, "Url Param id is missing")
		return
	}
	user, err := models.FindUser(ids[0])
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
		return
	}
	responseWithJSON(response, http.StatusOK, user)
}

func GetAll(response http.ResponseWriter, request *http.Request) {
	users := models.GetAllUser()
	responseWithJSON(response, http.StatusOK, users)
}

func CreateUser(response http.ResponseWriter, request *http.Request) {
	var user entities.User
	err := json.NewDecoder(request.Body).Decode(&user) // Decoder(*variable) đọc dữ liệu JSON và lưu vào user, truyền tham chiếu => truyền vào địa chỉ của biến; Encoder() ghi dữ liệu JSON
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
	} else {
		result := models.CreateUser(&user)
		if !result {
			responseWithError(response, http.StatusBadRequest, "Could not create user")
			return
		}
		responseWithJSON(response, http.StatusOK, user)
	}
}

func UpdateUser(response http.ResponseWriter, request *http.Request) {
	var user entities.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
	} else {
		result := models.UpdateUser(&user)
		if !result {
			responseWithError(response, http.StatusBadRequest, "Could not update user")
			return
		}
		responseWithJSON(response, http.StatusOK, "Update user successfully")
	}
}

func Delete(response http.ResponseWriter, request *http.Request) {
	ids, ok := request.URL.Query()["id"]
	if !ok || len(ids) < 1 {
		responseWithError(response, http.StatusBadRequest, "Url Param id is missing")
		return
	}
	result := models.DeleteUser(ids[0])
	if !result {
		responseWithError(response, http.StatusBadRequest, "Could not delete user")
		return
	}
	responseWithJSON(response, http.StatusOK, "Delete user successfully")
}

func responseWithError(response http.ResponseWriter, statusCode int, msg string) {
	responseWithJSON(response, statusCode, map[string]string{
		"error": msg,
	})
}

func responseWithJSON(response http.ResponseWriter, statusCode int, data interface{}) {
	result, _ := json.Marshal(data)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	response.Write(result)
}
