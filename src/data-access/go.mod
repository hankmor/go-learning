module koobyte.com/data-access

go 1.16

require (
	// 1.6.0报错，换成1.3.0正常： https://github.com/go-sql-driver/mysql/issues/828
	// could not use requested auth plugin 'mysql_native_password': this user requires mysql native password authentication.
	//           2021/07/18 01:20:19 this user requires mysql native password authentication.
	github.com/go-sql-driver/mysql v1.3.0
)
