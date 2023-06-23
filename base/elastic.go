package base

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
)

type Elastic struct {
	ElasticClient *elasticsearch.Client
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

	es, err := elasticsearch.NewClient(cfg)
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
		ElasticClient: es,
	}
}
