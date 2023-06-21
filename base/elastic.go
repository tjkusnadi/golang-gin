package base

import (
	"github.com/elastic/go-elasticsearch/v8"
)

type Elastic struct {
	elastic *elasticsearch.Client
}

func NewElasticsearch(logger Logger) Elastic {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		logger.Zap.Error("Error creating the client: %s", err)
	}

	res, err := es.Info()
	if err != nil {
		logger.Zap.Error("Error getting response: %s", err)
	}

	defer res.Body.Close()

	logger.Zap.Info("Connection success: %s", res)

	return Elastic{
		elastic: es,
	}
}