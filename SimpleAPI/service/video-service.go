package service

import "github.com/arya2004/GoLearn/SimpleAPI/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func New() VideoService{
	return &videoService{}
}

func (service *videoService) Save(vid entity.Video) entity.Video{
	service.videos = append(service.videos, vid)
	return vid
}

func (service *videoService) FindAll() []entity.Video{
	return service.videos
}