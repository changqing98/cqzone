package handler

import (
    "github.com/changqing98/cqzone/user/infrastructure/constant"
    "github.com/changqing98/cqzone/user/infrastructure/utils/logger"
    "github.com/gin-gonic/gin"
)

func Recover(c *gin.Context) {
    defer func() {
        if r := recover(); r != nil {
            //封装通用json返回
            //c.JSON(handler.StatusOK, Result.Fail(errorToString(r)))
            //Result.Fail不是本例的重点，因此用下面代码代替
            var err *constant.Error
            err, isErr := r.(*constant.Error)
            if !isErr {
                err = constant.InternalError
                logger.Error(r)
            }
            c.JSON(err.Code / 100, err)
            //终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
            c.Abort()
        }
    }()
    //加载完 defer recover，继续后续接口调用
    c.Next()
}
