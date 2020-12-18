package dmdb

import (
	"database/sql/driver"
)

type Stmt struct {
	stmt driver.Stmt
}

func (c *Conn) Prepare(query string) (driver.Stmt, error) {
	stmt, err := c.conn.Prepare(query)
	if err != nil {
		return nil, err
	}
	return &Stmt{stmt: stmt}, nil
}

func (s *Stmt) NumInput() int {
	return s.stmt.NumInput()
}
func (s *Stmt) Close() error {
	return s.stmt.Close()
}
func (s *Stmt) Exec(args []driver.Value) (driver.Result, error) {
	return s.stmt.Exec(args)
}
func (s *Stmt) Query(args []driver.Value) (driver.Rows, error) {
	rows, err := s.stmt.Query(args)
	if err != nil {
		return nil, err
	}
	return &Rows{r: rows}, nil
}
