package dmdb

import (
	"database/sql/driver"

	"github.com/axgle/mahonia"
)

type Rows struct {
	r driver.Rows
}

// 字符串解码函数，处理中文乱码
func convertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
func (r *Rows) Columns() []string {
	return r.r.Columns()
}

// linux不需要转换编码
func (r *Rows) Next(dest []driver.Value) error {
	if err := r.r.Next(dest); err != nil {
		return err
	}
	return nil
}
func (r *Rows) Close() error {
	return r.r.Close()
}

// ColumnTypeDatabaseTypeName 附加的
func (r *Rows) ColumnTypeDatabaseTypeName(index int) string {
	if tr, ok := r.r.(driver.RowsColumnTypeDatabaseTypeName); ok {
		return tr.ColumnTypeDatabaseTypeName(index)
	}
	return ""
}
