package model

import "time"

type User struct {
	UserCore
	Email     string     `json:"email" binding:"required,email,max=255"`
	Password  string     `json:"password" binding:"required,min=8,max=255"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

type Users []*User

type UserCore struct {
	ID           string `json:"id"`
	DisplayName  string `json:"displayName" binding:"required,max=32"`
	IconURL      string `json:"iconURL" binding:"omitempty,max=255,url"`
	Introduction string `json:"introduction" binding:"max=255"`
}
