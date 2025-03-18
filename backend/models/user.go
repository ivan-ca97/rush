package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id       uuid.UUID  `json:"id" gorm:"primaryKey"`
	Password string     `json:"password"`
	Username string     `json:"username"`
	EMail    string     `json:"email"`
	PersonId *uuid.UUID `json:"person_id"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Person      uuid.UUID    `gorm:"foreignKey:PersonId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Permissions []Permission `gorm:"many2many:user_permissions;"`
	Roles       []Role       `gorm:"many2many:user_roles;"`
}
