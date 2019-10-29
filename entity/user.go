package entity

import (
	"io/ioutil"
	"encoding/json"
	"os"
)

type User struct {
	Username string
	Password string
	Email string
	Telephone string
}

func ReadUser(filepath string)([]User, error) {
	var users []User
	str, err:=ioutil.ReadFile(filepath)
	if err!=nil {
		return users, err
	}
	jsonStr:=string(str)
	json.Unmarshal([]byte(jsonStr), &users)
	return users,nil
}

func WriteUser(filepath string, users []User) (error) {
	data, err:=json.Marshal(users);
	if err==nil {
		return ioutil.WriteFile(filepath,[]byte(data), os.ModeAppend)
	}
	return err
}