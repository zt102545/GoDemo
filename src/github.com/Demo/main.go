package main

import (
	"Demo/Common"
	"Demo/Routers"
	"log"
)

func main() {
	//读取配置文件
	cfg, err := Common.ParseConfig("./Config/app.json")
	if err != nil {
		log.Printf("err:%+v", err)
		return
	}
	//创建连接数据库对象
	_, err = Common.OrmEngine(cfg)
	if err != nil {
		log.Printf("Orm err:%+v", err)
		return
	}

	//创建redis连接对象
	err = Common.RedisConnect(cfg)
	if err != nil {
		log.Printf("Redis err:%+v", err)
		return
	}

	//注册路由
	app := Routers.RegisterRouter()
	app.Run(cfg.App.AppHost + ":" + cfg.App.AppPort)

}
