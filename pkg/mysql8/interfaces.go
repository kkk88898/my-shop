package mysql8

import "database/sql"

type DBEngine interface {
	GetDB() *sql.DB
	Configure(...Option) DBEngine
	Close()
}
