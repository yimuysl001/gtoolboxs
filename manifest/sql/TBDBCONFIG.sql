
/*
mysql	mysql:root:12345678@tcp(127.0.0.1:3306)/test?loc=Local&parseTime=true	mysql
mariadb	mariadb:root:12345678@tcp(127.0.0.1:3306)/test?loc=Local&parseTime=true	mysql
tidb	tidb:root:12345678@tcp(127.0.0.1:3306)/test?loc=Local&parseTime=true	mysql
pgsql	pgsql:root:12345678@tcp(127.0.0.1:5432)/test	pq
mssql   mssql:root:12345678@tcp(127.0.0.1:1433)/test?encrypt=disable   go-mssqldb
sqlite	sqlite::@file(/var/data/db.sqlite3)  (可以使用相对路径，如: db.sqlite3)	go-sqlite3
oracle	oracle:root:12345678@tcp(127.0.0.1:5432)/test	go-oci8
clickhouse	clickhouse:root:12345678@tcp(127.0.0.1:9000)/test  clickhouse-go
dm	dm:root:12345678@tcp(127.0.0.1:5236)/test  dm


*/


create table  TBDBCONFIG {
    ID int,  --唯一id
	CYWBM  varchar(20), -- 业务编码
	STATUS int   , -- 状态 1 启用
	Host  varchar(20), --链接ip
	Port  varchar(20), --端口
	User   varchar(20), -- 用户名
	Pass varchar(20), --密码
	Name   varchar(20), --数据库
	Type  varchar(20), --数据库类型
	Link   text,  --链接
	Extra  text, --附加属性
	Role   varchar(20), --主从  master（写入）  slave（读取）
    CBZ varchar(255) --备注
}