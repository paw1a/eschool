package domain

import "database/sql"

type User struct {
	ID        int64          `json:"id" db:"id"`
	Name      string         `json:"name" db:"name"`
	Surname   string         `json:"surname" db:"surname"`
	Phone     sql.NullString `json:"phone" db:"phone"`
	City      sql.NullString `json:"city" db:"city"`
	AvatarUrl sql.NullString `json:"avatar_url" db:"avatar_url"`
	Email     string         `json:"email" db:"email"`
	Password  string         `json:"-" db:"password"`
}

type UserInfo struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
}
