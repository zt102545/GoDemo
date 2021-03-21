package Controller

import (
	"Demo/Common"
	"Demo/Dao"
	"Demo/Dto"
	"Demo/Models"
	"github.com/gin-gonic/gin"
	"strconv"
)

type MySqlController struct {

}

func Select(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	rsp := new(Dto.Rsp)
	dUser := Dao.DaoUser{Orm: Common.DbEngine}
	result := dUser.Select(id)
	if result != nil {
		rsp.Code = 200
		rsp.Msg = "成功"
		rsp.Data = result
	} else {
		rsp.Code = 200
		rsp.Msg = "失敗"
	}
	ctx.JSON(200, rsp)
}

func Insert(ctx *gin.Context) {
	account := ctx.PostForm("account")
	passWord := ctx.PostForm("passWord")

	user := Models.User{Account: account, PassWord: passWord}
	dUser := Dao.DaoUser{Orm: Common.DbEngine}
	result := dUser.Insert(user)
	if result > 0 {
		ctx.JSON(200, map[string]interface{}{
			"message": "添加成功",
		})
	} else {
		ctx.JSON(200, map[string]interface{}{
			"message": "添加失败",
		})
	}
}

func Delete(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Query("id"), 10, 64)

	dUser:=Dao.DaoUser{Orm: Common.DbEngine}
	result:=dUser.DeleteById(id)
	if result>0 {
		ctx.JSON(200, map[string]interface{}{
			"message": "删除成功",
		})
	} else{
		ctx.JSON(200, map[string]interface{}{
			"message": "删除失败",
		})
	}
}

func Update(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.PostForm("id"), 10, 64)
	account := ctx.PostForm("account")
	passWord := ctx.PostForm("passWord")

	user := Models.User{Id: id, Account: account, PassWord: passWord}
	dUser := Dao.DaoUser{Orm: Common.DbEngine}
	result := dUser.UpdateUser(user)
	if result > 0 {
		ctx.JSON(200, map[string]interface{}{
			"message": "更新成功",
		})
	} else {
		ctx.JSON(200, map[string]interface{}{
			"message": "更新失败",
		})
	}
}