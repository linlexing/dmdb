package dmdb

import (
	"database/sql/driver"
)

type Conn struct {
	conn driver.Conn
}

func (d *Driver) Open(dsn string) (driver.Conn, error) {
	h, err := d.drv.Open(dsn)
	if err != nil {
		return nil, err
	}
	return &Conn{conn: h}, nil
}
func (c *Conn) Close() (err error) {
	return c.conn.Close()
}
func (c *Conn) Begin() (driver.Tx, error) {
	return c.conn.Begin()
}
