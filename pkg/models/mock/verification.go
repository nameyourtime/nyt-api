package mock

import (
	"nameyourtime.com/api/pkg/models"
	"time"
)

var mockVerificationCode = &models.VerificationCode{
	UserID:  "test_user_id",
	Code:    "test_code",
	CodeExp: time.Now().Add(time.Hour),
}

type VerificationModel struct {
}

func (m *VerificationModel) Create(code models.VerificationCode) (string, error) {
	return mockVerificationCode.UserID, nil
}
