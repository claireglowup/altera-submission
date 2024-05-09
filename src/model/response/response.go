package response

type Response struct {
	Status  int         `json:"status"`
	Message interface{} `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type UserShoes struct {
	Id        string      `json:"id"`
	Name      string      `json:"name"`
	Email     string      `json:"email" gorm:"unique"`
	RackShoes interface{} `json:"shoes"`
}
