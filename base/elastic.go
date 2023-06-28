package base

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
)

type Elastic struct {
	ElasticClient *elasticsearch.TypedClient
}

func NewElasticsearch(env Env, logger Logger) Elastic {
	cfg := elasticsearch.Config{
		Addresses: []string{
			env.ElasticHost,
		},
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MaxVersion:         tls.VersionTLS11,
				InsecureSkipVerify: true,
			},
		},
	}

	es, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		logger.Zap.Error("Error creating the client: %s", err)
	}

	res := es.Info()

	response, err := res.Do(context.Background())

	if err != nil {
		logger.Zap.Error("Error creating the client: %s", err)
	}

	logger.Zap.Info("Connection success:", response.Version)

	return Elastic{
		ElasticClient: es,
	}
}
