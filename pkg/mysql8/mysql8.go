package mysql8

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/exp/slog"
	"log"
	"myshop/cmd/order/config"
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

func NewMysql8DB(cfg *config.Config) (DBEngine, error) {
	slog.Info("CONN", "connect string", cfg.MYSQL8.DsnURL)

	mysql8 := &mysql8{
		connAttempts: _defaultConnAttempts,
		connTimeout:  _defaultConnTimeout,
	}

	var err error
	for mysql8.connAttempts > 0 {
		mysql8.db, err = sql.Open("mysql", cfg.MYSQL8.DsnURL)
		if err != nil {

			break
		}

		log.Printf("Mysql8 is trying to connect, attempts left: %d", mysql8.connAttempts)

		time.Sleep(mysql8.connTimeout)

		mysql8.connAttempts--
	}
	mysql8.db.SetMaxOpenConns(cfg.MYSQL8.PoolMax)                              //连接总数
	mysql8.db.SetMaxIdleConns(cfg.MYSQL8.IdleConnMax)                          //最大空闲连接
	mysql8.db.SetConnMaxIdleTime(time.Duration(cfg.MYSQL8.MaxIdleTime * 1000)) //空闲状态最大生命后期
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
