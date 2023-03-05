package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	First_Name string
	Last_Name  string
	Email      string
	Password   string
	RoleID     uint
	Role       Role `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Role struct {
	gorm.Model
	Title       string
	Description string
	Approvers   MultiString `gorm:"type:VARCHAR(255)"`
	Status      string
}

type Permission struct {
	gorm.Model
	Title       string
	Description string
	Status      string
}

type Role_Permission struct {
	gorm.Model
	RoleID       uint
	Role         Role `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PermissionID uint
	Permission   Permission `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type AccessRequest struct {
	gorm.Model
	Requestor     string
	Justification string
	StatusID      uint
	RequestStatus RequestStatus `gorm:"foreignKey:StatusID; constraint:OnDelete:SET NULL;"`
	RoleID        uint
	Role          Role
}

type RequestStatus struct {
	gorm.Model
	Result          string
	Total_Approvers int
	Approve         int
	Deny            int
	Comment         string
}
