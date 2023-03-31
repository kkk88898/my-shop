package mysql8

import (
	"database/sql"
	"golang.org/x/exp/slog"
	"log"
	"time"
)

const (
	_defaultConnAttempts = 3
	_defaultConnTimeout  = time.Second
)

type DBConnString string

type mysql8 struct {
	connAttempts int
	connTimeout  time.Duration
	db           *sql.DB
}

var _ DBEngine = (*mysql8)(nil)

func NewMysql8DB(url DBConnString) (DBEngine, error) {
	slog.Info("CONN", "connect string", url)

	mysql8 := &mysql8{
		connAttempts: _defaultConnAttempts,
		connTimeout:  _defaultConnTimeout,
	}

	var err error
	for mysql8.connAttempts > 0 {
		mysql8.db, err = sql.Open("mysql8", string(url))
		if err != nil {
			break
		}

		log.Printf("Postgres is trying to connect, attempts left: %d", mysql8.connAttempts)

		time.Sleep(mysql8.connTimeout)

		mysql8.connAttempts--
	}
	slog.Info("📰 connected to mysql8 🎉")

	return mysql8, nil
}
func (m *mysql8) Configure(opts ...Option) DBEngine {
	for _, opt := range opts {
		opt(m)
	}
	return m
}

func (m *mysql8) GetDB() *sql.DB {
	return m.db
}

func (m *mysql8) Close() {
	if m.db != nil {
		m.db.Close()
	}
}
