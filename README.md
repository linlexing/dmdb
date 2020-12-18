# dmdb
达梦go驱动

该项目包装了odbc驱动，以解决达梦数据库string类型数据编码问题。

## 解决的问题

windows平台中，达梦数据库使用odbc驱动时，所有的字符串类型数据，将按照gbk编码返回，需要额外编写代码转换回utf-8，否则go语言中将是乱码。该项目拦截了Scan操作，自动对所有的`[]byte`类型进行编码转换。