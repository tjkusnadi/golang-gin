package repositories

import (
	"encoding/json"
	"io"

	"golang-gin/base"
)

type ElasticRepository struct {
	es     base.Elastic
	logger base.Logger
}

type ElasticVersion struct {
	Number string `json:"number"`
}

type ElasticInfo struct {
	Version ElasticVersion `json:"version"`
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

func (er ElasticRepository) GetInfo() (ElasticInfo, error) {
	res, err := er.es.ElasticClient.Info()

	var elasticInfo ElasticInfo

	if err != nil {
		er.logger.Zap.Error("[ERROR] GetInfo Elastic", err)
		return ElasticInfo{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		er.logger.Zap.Error("[ERROR] GetInfo Elastic", err)
		return ElasticInfo{}, err
	}

	json.Unmarshal(body, &elasticInfo)

	return ElasticInfo{
		Version: elasticInfo.Version,
	}, nil
}
