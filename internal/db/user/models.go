// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package user

import (
	"database/sql"
)

type User struct {
	ID        int64
	Username  string
	Password  string
	CreatedAt sql.NullTime
}
