package controller

import (
	"strconv"

	"github.com/arya2004/GoLearn/SimpleAPI/entity"
	"github.com/arya2004/GoLearn/SimpleAPI/service"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error

	ShowAll(ctx *gin.Context)
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

func (c *controller)Update(ctx *gin.Context) error{
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	video.ID = id

	
	c.service.Update(video)
	return nil

}

func (c *controller) Delete(ctx *gin.Context) error{
	var video entity.Video
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	video.ID = id
	c.service.Delete(video)
	return nil

}


func (c *controller) ShowAll(ctx *gin.Context){
	videos := c.service.FindAll()
	data := gin.H{
		"title" : "Video Page",
		"videos" : videos,
	}
	ctx.HTML(200, "index.html",data)
}