package services

import (
	"golang-gin/api/repositories"
)

type ElasticService struct {
	repository repositories.ElasticRepository
}

func NewElasticService(
	repository repositories.ElasticRepository,
) ElasticService {
	return ElasticService{
		repository: repository,
	}
}

func (es ElasticService) GetInfo() (repositories.ElasticInfo, error) {
	return es.repository.GetInfo()
}
