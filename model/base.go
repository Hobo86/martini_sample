package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"

	. "../common"
	"time"
)

var orm *xorm.Engine

func SetEngine() *xorm.Engine {
	Log.Info("db initializing...")
	var err error
	username := Cfg.MustValue("db", "username", "root")
	password := Cfg.MustValue("db", "password", "meet771027")
	dbName := Cfg.MustValue("db", "db_name", "im")
	orm, err = xorm.NewEngine("mysql", username+":"+password+"@/"+dbName+"?charset=utf8")
	PanicIf(err)
	orm.TZLocation = time.Local
	orm.Logger = xorm.NewSimpleLogger(Log.GetWriter())
	return orm
}