package main

import (
	"fmt"
	_ "golangERP/routers"
	"golangERP/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func init() {
	dbType := beego.AppConfig.String("db_type")
	fmt.Println(dbType)
	//获得数据库参数，不同数据库可能存在没有值的情况没有的值nil
	dbAlias := beego.AppConfig.String(dbType + "::db_alias")
	dbName := beego.AppConfig.String(dbType + "::db_name")
	dbUser := beego.AppConfig.String(dbType + "::db_user")
	dbPwd := beego.AppConfig.String(dbType + "::db_pwd")
	dbPort := beego.AppConfig.String(dbType + "::db_port")
	dbHost := beego.AppConfig.String(dbType + "::db_host")
	orm.RegisterDriver(dbType, orm.DRPostgres)
	switch dbType {
	//数据库类型和数据库驱动名一致
	case "postgres":

		dbSslmode := beego.AppConfig.String(dbType + "::db_sslmode")
		dataSource := "user=" + dbUser + " password=" + dbPwd + " dbname=" + dbName + " host=" + dbHost + " port=" + dbPort + " sslmode=" + dbSslmode
		fmt.Print(dataSource)
		orm.RegisterDataBase(dbAlias, dbType, dataSource)

	case "mysql":
		dbCharset := beego.AppConfig.String(dbType + "db_charset")
		dataSource := dbUser + ":" + dbPwd + "@/" + dbName + "?charset=" + dbCharset
		orm.RegisterDataBase(dbAlias, dbType, dataSource)
	case "sqlite3":
		orm.RegisterDataBase(dbAlias, "sqlite3", dbName)

	}
	utils.LogOut("info", "使用数据库为:"+dbType)
	//重新运行时是否覆盖原表创建,false:不会删除原表,修改表信息时将会在原来的基础上修改，true删除原表重新创建
	coverDb, _ := beego.AppConfig.Bool("cover_db")

	//自动建表
	orm.RunSyncdb(dbAlias, coverDb, true)

	// 加载权限控制文件
	// LoadSecurity()
	// 初始化cache
	utils.InitCache()
}
func main() {
	beego.Run()
}
