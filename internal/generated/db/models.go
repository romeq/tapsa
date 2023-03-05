// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type TokenType string

const (
	TokenTypeRefresh        TokenType = "refresh"
	TokenTypeAuthentication TokenType = "authentication"
)

func (e *TokenType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = TokenType(s)
	case string:
		*e = TokenType(s)
	default:
		return fmt.Errorf("unsupported scan type for TokenType: %T", src)
	}
	return nil
}

type NullTokenType struct {
	TokenType TokenType
	Valid     bool // Valid is true if TokenType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullTokenType) Scan(value interface{}) error {
	if value == nil {
		ns.TokenType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.TokenType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullTokenType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.TokenType), nil
}

type Account struct {
	AccountID      uuid.UUID `json:"account_id"`
	Username       string    `json:"username"`
	Password       string    `json:"password"`
	RegisterDate   time.Time `json:"register_date"`
	LastLogin      time.Time `json:"last_login"`
	ActivityPoints int32     `json:"activity_points"`
}

type AccountSession struct {
	SessionID  string    `json:"session_id"`
	AccountID  uuid.UUID `json:"account_id"`
	StartDate  time.Time `json:"start_date"`
	ExpireDate time.Time `json:"expire_date"`
	TokenType  TokenType `json:"token_type"`
}

type Feedback struct {
	ID      int32          `json:"id"`
	Comment sql.NullString `json:"comment"`
	Boxes   string         `json:"boxes"`
}

type File struct {
	FileUuid     string         `json:"file_uuid"`
	Title        sql.NullString `json:"title"`
	Passwdhash   sql.NullString `json:"passwdhash"`
	AccessToken  string         `json:"access_token"`
	Encrypted    bool           `json:"encrypted"`
	FileSize     int32          `json:"file_size"`
	EncryptionIv []byte         `json:"encryption_iv"`
	UploadDate   time.Time      `json:"upload_date"`
	LastSeen     time.Time      `json:"last_seen"`
	Viewcount    int32          `json:"viewcount"`
	FileHash     string         `json:"file_hash"`
}

type FileToAccount struct {
	AccountID uuid.UUID `json:"account_id"`
	FileUuid  string    `json:"file_uuid"`
}

type PeerBan struct {
	PeerID string `json:"peer_id"`
}

type Report struct {
	FileUuid string `json:"file_uuid"`
	Reason   string `json:"reason"`
}
