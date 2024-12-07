package service

import (
	"github.com/chhz0/asyncgo/internal/tasksvr/biz/model/request"
	"github.com/gin-gonic/gin"
)

type TaskCreateService struct {
	ginCtx *gin.Context
}

func NewTaskCreateService(ctx *gin.Context) *TaskCreateService {
	return &TaskCreateService{ginCtx: ctx}
}

func (s *TaskCreateService) Run(req *request.TaskCreateReq) (taskId string, err error) {

	return "taskId", nil
}
