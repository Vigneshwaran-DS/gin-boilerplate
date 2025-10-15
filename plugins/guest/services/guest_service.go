package services

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"time"

	"gin-boilerplate/plugins/guest/models"
	guestutils "gin-boilerplate/plugins/guest/utils"
	"gorm.io/gorm"
)

type GuestService struct {
	db *gorm.DB
}

func NewGuestService(db *gorm.DB) *GuestService {
	return &GuestService{db: db}
}

// CreateGuest creates a new guest user with auto-generated credentials
func (s *GuestService) CreateGuest(deviceID, deviceInfo, ipAddress string) (*models.Guest, string, error) {
	// Generate unique guest UID
	guestUID := s.generateGuestUID()

	// Generate guest token
	guestToken := s.generateToken()

	now := time.Now()
	guest := &models.Guest{
		GuestUID:     guestUID,
		GuestToken:   guestToken,
		DeviceID:     deviceID,
		DeviceInfo:   deviceInfo,
		IPAddress:    ipAddress,
		LastActiveAt: &now,
		IsUpgraded:   false,
	}

	if err := s.db.Create(guest).Error; err != nil {
		return nil, "", err
	}

	// Generate JWT token for guest using guest-specific JWT utility
	token, err := guestutils.GenerateGuestToken(guest.ID, guest.GuestUID)
	if err != nil {
		return nil, "", err
	}

	return guest, token, nil
}

// GetGuestByUID retrieves a guest by their UID
func (s *GuestService) GetGuestByUID(guestUID string) (*models.Guest, error) {
	var guest models.Guest
	err := s.db.Where("guest_uid = ? AND is_upgraded = ?", guestUID, false).First(&guest).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("guest not found")
		}
		return nil, err
	}
	return &guest, nil
}

// UpdateLastActive updates the last active timestamp for a guest
func (s *GuestService) UpdateLastActive(guestUID string) error {
	now := time.Now()
	return s.db.Model(&models.Guest{}).
		Where("guest_uid = ?", guestUID).
		Update("last_active_at", now).Error
}

// CleanupInactiveGuests removes guests that haven't been active for the specified duration
func (s *GuestService) CleanupInactiveGuests(inactiveDuration time.Duration) (int64, error) {
	cutoffTime := time.Now().Add(-inactiveDuration)

	result := s.db.Where("last_active_at < ? AND is_upgraded = ?", cutoffTime, false).
		Delete(&models.Guest{})

	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}

// UpgradeToUser marks a guest as upgraded when they register as a full user
func (s *GuestService) UpgradeToUser(guestUID string, userID uint) error {
	return s.db.Model(&models.Guest{}).
		Where("guest_uid = ?", guestUID).
		Updates(map[string]interface{}{
			"is_upgraded":      true,
			"upgraded_user_id": userID,
		}).Error
}

// generateGuestUID generates a unique guest UID
func (s *GuestService) generateGuestUID() string {
	// Generate 12 random bytes (24 hex characters)
	randomBytes := make([]byte, 12)
	rand.Read(randomBytes)
	return "guest_" + hex.EncodeToString(randomBytes)
}

// generateToken generates a secure random token
func (s *GuestService) generateToken() string {
	bytes := make([]byte, 32)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
