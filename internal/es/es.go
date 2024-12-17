package es

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"net/http"
	"search-nova/internal/config"
	"search-nova/internal/constant"
	"search-nova/internal/logger"
)

var (
	E *es
)

func init() {
	var err error
	E, err = new()
	if err != nil {
		logger.L.Fatalf("Fatal error es: %v\n", err)
	}
}

func new() (*es, error) {
	e := &es{}
	var err error
	e.client, err = elasticsearch.NewClient(elasticsearch.Config{
		Addresses: config.C.GetStringSlice(constant.ElasticsearchAddresses),
		Username:  config.C.GetString(constant.ElasticsearchUsername),
		Password:  config.C.GetString(constant.ElasticsearchPassword),
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	})
	if err != nil {
		return nil, err
	}
	_, err = e.client.Ping()
	if err != nil {
		return nil, err
	}
	e.index = config.C.GetString(constant.ElasticsearchIndex)
	return e, nil
}

type es struct {
	client *elasticsearch.Client
	index  string
}

func (e *es) IndexDoc(data []byte) error {
	resp, err := e.client.Index(e.index, bytes.NewReader(data))
	if err != nil {
		return err
	}
	if resp.IsError() {
		return errors.New(resp.String())
	}
	return nil
}

func (e *es) Search(body []byte) (*SearchResponse, error) {
	req := esapi.SearchRequest{
		Index: []string{e.index},
		Body:  bytes.NewReader(body),
	}
	resp, err := req.Do(context.Background(), e.client)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, errors.New(resp.String())
	}
	defer resp.Body.Close()
	var searchResponse SearchResponse
	err = json.NewDecoder(resp.Body).Decode(&searchResponse)
	if err != nil {
		return nil, err
	}
	return &searchResponse, err
}

type SearchResponse struct {
	Hits struct {
		Hits []struct {
			Source struct {
				Id int64 `json:"id"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}
