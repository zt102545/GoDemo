package Common

import (
	"Demo/Models"
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
)

var DbEngine *Orm

type Orm struct {
	*xorm.Engine
}

func OrmEngine(cfg *Config) (*Orm,error) {
	conn := cfg.DataBase.Url
	engine, err := xorm.NewEngine("mysql", conn)
	if err != nil {
		return nil, err
	}
	//输出SQL
	engine.ShowSQL()
	//同步数据库结构
	err =engine.Sync2(new(Models.User))
	if err!=nil{
		return nil, err
	}


	orm := new(Orm)
	orm.Engine = engine
	DbEngine = orm
	return orm, err
}
