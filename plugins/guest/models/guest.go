package models

import (
	"time"

	"gorm.io/gorm"
)

type Guest struct {
	ID             uint           `gorm:"primarykey" json:"id"`
	GuestUID       string         `gorm:"uniqueIndex;size:64;not null" json:"guest_uid"`
	GuestToken     string         `gorm:"size:128;not null" json:"-"`
	DeviceID       string         `gorm:"size:64;index" json:"device_id"`
	DeviceInfo     string         `gorm:"type:text" json:"device_info"`
	IPAddress      string         `gorm:"size:45" json:"ip_address"`
	LastActiveAt   *time.Time     `json:"last_active_at"`
	IsUpgraded     bool           `gorm:"default:false;index" json:"is_upgraded"`
	UpgradedUserID *uint          `json:"upgraded_user_id,omitempty"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}
