package types

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"gopkg.in/olivere/elastic.v7"
)

const ElasticIndexStudents = "students"

type Student struct {
	Name         string  `json:"name"`
	Age          int64   `json:"age"`
	AverageScore float64 `json:"average"`
}

type StudentsServiceIndex interface {
	Add(ctx context.Context, student *Student) (*elastic.IndexResponse, error)
	FindByName(ctx context.Context, name string) ([]Student, error)
}

type studentsServiceIndex struct {
	esClient *elastic.Client
}

func NewStudentsServiceIndex(c *elastic.Client) StudentsServiceIndex {

	exists, err := c.IndexExists(ElasticIndexStudents).Do(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	if !exists {
		log.Println("Index not exists!!!")
	}

	return &studentsServiceIndex{esClient: c}
}

func (s *studentsServiceIndex) Add(ctx context.Context, student *Student) (*elastic.IndexResponse, error) {
	data, err := json.Marshal(student)
	if err != nil {
		return nil, err
	}
	return s.
		esClient.
		Index().
		Index(ElasticIndexStudents).
		BodyJson(string(data)).
		Do(ctx)
}

func (s *studentsServiceIndex) FindByName(ctx context.Context, name string) ([]Student, error) {

	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewMatchQuery("name", name))
	query, err := searchSource.Source()
	if err != nil {
		return nil, err
	}
	b, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}
	fmt.Println("Query: ", string(b))

	result, err := s.esClient.
		Search().
		Index(ElasticIndexStudents).
		SearchSource(searchSource).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%+v\n", result)

	var students []Student

	for _, hit := range result.Hits.Hits {
		var student Student
		err := json.Unmarshal(hit.Source, &student)
		if err != nil {
			log.Println(err.Error())
		}
		students = append(students, student)
	}

	return students, nil
}
