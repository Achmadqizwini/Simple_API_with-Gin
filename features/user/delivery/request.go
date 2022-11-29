package delivery

import "be13/ca/features/user"

type UserRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Address  string `json:"address" form:"address"`
	Phone    string `gorm:"type:varchar(15)" json:"phone" form:"phone"`
	Role     string `json:"role" form:"role"`
}

func requestToCore(userInput UserRequest) user.Core {
	userCoreData := user.Core{
		Name:     userInput.Name,
		Email:    userInput.Email,
		Password: userInput.Password,
		Phone:    userInput.Phone,
		Address:  userInput.Address,
		Role:     userInput.Role,
	}
	return userCoreData
}
