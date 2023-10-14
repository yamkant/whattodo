package models

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primary_key;column:id"`
	Name      string    `json:"name" gorm:"column:name"`
	KakaoID   uint64    `json:"kakao_id" gorm:"unique"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at default:CURRENT_TIMESTAMP()"`
}

func (User) TableName() string {
	return "users"
}

type UserAddDTO struct {
	KakaoID uint64 `json:"kakao_id" gorm:"column:kakao_id"`
}
