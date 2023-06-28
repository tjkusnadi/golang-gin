package repositories

import (
	"context"

	"golang-gin/base"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type ElasticRepository struct {
	es     base.Elastic
	logger base.Logger
}

func NewElasticRepository(
	elastic base.Elastic,
	logger base.Logger,
) ElasticRepository {
	return ElasticRepository{
		es:     elastic,
		logger: logger,
	}
}

func (er ElasticRepository) GetInfo() (types.ElasticsearchVersionInfo, error) {
	res := er.es.ElasticClient.Info()

	body, err := res.Do(context.Background())

	if err != nil {
		er.logger.Zap.Error("[ERROR] GetInfo Elastic", err)
		return body.Version, err
	}

	return body.Version, nil
}
