package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/shixinshuiyou/mayo/app/user/pojo/bo"
	"github.com/shixinshuiyou/mayo/tool/resp"
)

func UserLogin(ctx *gin.Context) {
	var form bo.UserRegisterBo
	if err := ctx.ShouldBindQuery(&form); err != nil {
		resp.FailWithMsg(ctx, err)
		return
	}
	ctx.AbortWithStatusJSON(0, "login success")
}

func UserLogout(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(0, "logout success")
}

func UserRegister(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(0, "register success")
}

func UserInfo(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(0, "logout success")
}
