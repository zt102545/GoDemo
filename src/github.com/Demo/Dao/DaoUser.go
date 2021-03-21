package Dao

import (
	"Demo/Common"
	"Demo/Models"
	"log"
)

type DaoUser struct {
	*Common.Orm
}

func (d *DaoUser) Select(id int64) *Models.User{
	user :=new(Models.User)
	_,err:=d.Where("id =?",id).Get(user)
	if err!= nil{
		log.Printf("err:%+v", err.Error())
	}
	return user
}

func (d *DaoUser) Insert(m Models.User) int64{
	result,err:=d.InsertOne(m)
	if err!= nil{
		log.Printf("err:%+v", err.Error())
	}
	return result
}

func (d *DaoUser) DeleteById(id int64) int64{
	user:=new(Models.User)
	result,err:=d.ID(id).Delete(user)
	if err!=nil {
		log.Printf("err:%v", err)
	}
	return result
}

func (d *DaoUser) UpdateUser(m Models.User) int64{
	result,err:=d.ID(m.Id).Update(m)
	if err!= nil{
		log.Printf("err:%+v", err.Error())
	}
	return result
}