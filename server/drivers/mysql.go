package drivers

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
)

type MysqlDriver struct {
	db *sql.DB
}

func NewMysqlDriver() (*MysqlDriver, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}
	c := mysql.Config{
		DBName:    "takamori",
		User:      "takamori",
		Passwd:    "password",
		Addr:      "db:3306",
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}
	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		return nil, err
	}
	return &MysqlDriver{db}, nil
}

func (d *MysqlDriver) Use() *sql.DB {
	return d.db
}

func (d *MysqlDriver) Close() {
	d.db.Close()
}
