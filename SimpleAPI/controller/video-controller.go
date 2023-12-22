package controller

import (
	"github.com/arya2004/GoLearn/SimpleAPI/entity"
	"github.com/arya2004/GoLearn/SimpleAPI/service"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController{
	return &controller{
		service: service,
	}
}



func (c *controller) FindAll() []entity.Video{
	return c.service.FindAll()
}
func (c* controller) Save(ctx *gin.Context) error{
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	c.service.Save(video)
	return nil
}