package drivers

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/ughvj/takayoshi/config"
)

type MysqlDriver struct {
	db *sql.DB
}

func NewMysqlDriver() (*MysqlDriver, error) {
	conf := config.Loader.Get()

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}
	c := mysql.Config{
		DBName:    conf.Db.Name,
		User:      conf.Db.User,
		Passwd:    conf.Db.Pass,
		Addr:      conf.Db.Addr,
		Net:       conf.Db.Net,
		ParseTime: true,
		Collation: conf.Db.Collation,
		Loc:       jst,
	}
	db, err := sql.Open(conf.Db.Ms, c.FormatDSN())
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
