package client

import "gopkg.in/olivere/elastic.v7"

const esURL = "http://localhost:9200"

func NewClient() (*elastic.Client, error) {
	return elastic.NewClient(
		elastic.SetURL(esURL),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
	)
}
