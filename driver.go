package dmdb

import (
	"database/sql"

	"github.com/linlexing/odbc"
)

var drv Driver

type Driver struct {
	drv *odbc.Driver
}

func (d *Driver) Close() error {
	return d.drv.Close()
}
func init() {
	db, err := sql.Open("odbc", "aaa")
	if err != nil {
		panic(err)
	}
	drv.drv = db.Driver().(*odbc.Driver)
	sql.Register("dmdb", &drv)
}
