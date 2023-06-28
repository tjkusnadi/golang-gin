package services

import (
	"golang-gin/api/repositories"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
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

func (es ElasticService) GetInfo() (types.ElasticsearchVersionInfo, error) {
	return es.repository.GetInfo()
}
