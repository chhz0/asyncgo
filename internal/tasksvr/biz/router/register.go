package router

import (
	"github.com/chhz0/asyncgo/internal/tasksvr/biz/router/task"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	task.Register(r)
}
