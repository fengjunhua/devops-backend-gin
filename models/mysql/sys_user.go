package mysql

import "github.com/google/uuid"

type SysUser struct {
	UUID  uuid.UUID
	Username   string
	Password   string
	HeaderImg  string
}