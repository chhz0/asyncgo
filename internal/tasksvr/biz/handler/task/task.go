package task

import (
	"github.com/chhz0/asyncgo/internal/tasksvr/biz/model/request"
	"github.com/chhz0/asyncgo/internal/tasksvr/biz/service"
	"github.com/gin-gonic/gin"
)

func TaskCreate(ctx *gin.Context) {
	var err error
	var req request.TaskCreateReq
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := service.NewTaskCreateService(ctx).Run(&req)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"task_id": resp})
}
