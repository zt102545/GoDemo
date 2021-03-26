package Controller

import (
	"Demo/Common"
	"Demo/Dto"
	"github.com/gin-gonic/gin"
)

func Set(ctx *gin.Context) {
	account := ctx.Param("account")
	re := Common.Redis{Client: Common.RedisClient}
	err := re.Set("test_accout", account, 60)
	if err == nil {
		ctx.JSON(200, map[string]interface{}{
			"message": "添加成功",
		})
	} else {
		ctx.JSON(200, map[string]interface{}{
			"message": "添加失败",
		})
	}
}

func Get(ctx *gin.Context) {
	account := ctx.Param("account")
	re := Common.Redis{Client: Common.RedisClient}
	str, err := re.Get(account)
	rsp := new(Dto.Rsp)
	if err != nil {
		rsp.Code = 200
		rsp.Msg = "失敗"
		rsp.Data = err
	} else {
		rsp.Code = 200
		rsp.Msg = "成功"
		rsp.Data = str
	}
	ctx.JSON(200, rsp)
}

