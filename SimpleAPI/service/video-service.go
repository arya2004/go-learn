package service

import (
	"github.com/arya2004/GoLearn/SimpleAPI/entity"
	"github.com/arya2004/GoLearn/SimpleAPI/repository"
)

type VideoService interface {
	Save(entity.Video) entity.Video
	Update(entity.Video) error
	Delete(entity.Video) error
	FindAll() []entity.Video
}

type videoService struct {
	videoRepository repository.VideoRepository
}

func New(repo repository.VideoRepository) VideoService{
	return &videoService{
		videoRepository: repo,
	}
}

func (service *videoService) Save(vid entity.Video) entity.Video{
	service.videoRepository.Save(vid)
	return vid
}

func (service *videoService) FindAll() []entity.Video{
	return service.videoRepository.FindAll()
}

func (service *videoService) Update(video entity.Video) error {
	service.videoRepository.Update(video)
	return nil
}

func (service *videoService) Delete(video entity.Video) error {
	service.videoRepository.Delete(video)
	return nil
}
