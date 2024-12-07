package task

import (
	"github.com/chhz0/asyncgo/internal/tasksvr/biz/handler/task"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	v1 := r.Group("/v1", rootMw()...)

	v1.POST("create", append(_taskMw(), task.TaskCreate)...)
}
