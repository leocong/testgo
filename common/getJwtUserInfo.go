package common

import (
	"encoding/json"
)

type JwtUserInfo struct {
	UserId    int64    `json:"user_id"`
	Paths     []string `json:"paths"`
	RoleNames []string `json:"role_names"`
	UserName  string   `json:"user_name"`
}

func GetJwtUserInfo(payload any) (*JwtUserInfo, error) {
	// 将map转换为JSON字节码
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	// 将JSON字节码转换为JwtUserInfo对象
	var person *JwtUserInfo
	err = json.Unmarshal(jsonData, &person)
	if err != nil {
		return nil, err
	}
	return person, nil
}
