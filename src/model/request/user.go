package request

type UserRegister struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
}

type ReqUserShoes struct {
	Noid int    `json:"no_id"`
	Name string `json:"name" validate:"required"`
	Size int    `json:"size" validate:"required"`
}

type UpdateUserShoes struct {
	Name string `json:"name" validate:"required"`
	Size int    `json:"size" validate:"required"`
}
