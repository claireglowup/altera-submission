package model

type Shoes struct {
	ShoesId   string `json:"-"`
	NoId      int    `json:"no_id" gorm:"autoIncrement"`
	Name      string `json:"name"`
	Size      int    `json:"size"`
	AddAt     string `json:"add_at"`
	UpdatedAt string `json:"updated_at,omitempty"`
}
