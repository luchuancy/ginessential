package dto

import "cyul.stu0323/ginessential/model"

// 前端返回信息
type UserDto struct {
	Name      string `json: name`
	Telephone string `json: telephone`
}

func ToUserDto(user model.User) UserDto {
	return UserDto {
		Name: user.Name,
		Telephone: user.Telephone,
	}
}