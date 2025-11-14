package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UUID = uuid.UUID

var ParseUUID = uuid.Parse

func NewUUID() uuid.UUID {
	return uuid.New()
}

var ErrRecordNotFound = gorm.ErrRecordNotFound
