package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role     string `gorm:"type:varchar(255);not null"`
	Company  string `gorm:"type:varchar(255);not null"`
	Location string `gorm:"type:varchar(255);not null"`
	Remote   bool   `gorm:"not null"`
	Link     string `gorm:"type:varchar(255);not null"`
	Salary   int64
	Closed   bool `gorm:"not null;default:false"`
}

type OpeningResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deteledAt,omitempty"`
	Role      string    `json:"role"`
	Company   string    `json:"company"`
	Location  string    `json:"location"`
	Remote    bool      `json:"remote"`
	Link      string    `json:"link"`
	Salary    int64     `json:"salary"`
	Closed    bool      `json:"closed"`
}
