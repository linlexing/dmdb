# dmdb
达梦go驱动

该项目包装了odbc驱动，以解决达梦数据库string类型数据编码问题。

## 解决的问题

windows平台中，达梦数据库使用odbc驱动时，所有的字符串类型数据，将按照gbk编码返回，需要额外编写代码转换回utf-8，否则go语言中将是乱码。该项目拦截了Scan操作，自动对所有的`[]byte`类型进行编码转换。

## 示例

```go
package main

import (
	"database/sql"

	_ "github.com/linlexing/dmdb"

	_ "github.com/alexbrainman/odbc"
)

func main() {
	db1, err := sql.Open("odbc", "driver={DM8 ODBC DRIVER};server=localhost:5236;database=DMSERVER;uid=test;pwd=123456789;charset=UTF8;")
	if err != nil {
		panic(err)
	}
	defer db1.Close()
	db2, err := sql.Open("dmdb", "driver={DM8 ODBC DRIVER};server=localhost:5236;database=DMSERVER;uid=test;pwd=123456789;charset=UTF8;")
	if err != nil {
		panic(err)
	}
	defer db2.Close()

	var str1, str2 string
	if err := db1.QueryRow("select '测试'").Scan(&str1); err != nil {
		panic(err)
	}
	if err := db2.QueryRow("select '测试'").Scan(&str2); err != nil {
		panic(err)
	}
	println("odbc:", str1, ",dmdb:", str2)

}

```