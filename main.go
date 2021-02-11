package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go-elasticsearch/client"
	"go-elasticsearch/types"
	"log"
)

func main() {
	esClient, err := client.NewClient()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("Elasticsearch conectado com sucesso!")

	studentsSvcIndex := types.NewStudentsServiceIndex(esClient)

	// AddingData(studentsSvcIndex)

	items, err := studentsSvcIndex.FindByName(context.Background(), "doe")
	if err != nil {
		log.Panicln(err)
	}

	for index := range items {
		Console(items[index])
	}

}

func AddingData(studentsSvcIndex types.StudentsServiceIndex) {

	student := &types.Student{
		Name:         "Gopher Doe",
		Age:          10,
		AverageScore: 99.9,
	}

	ctx := context.Background()

	response, err := studentsSvcIndex.Add(ctx, student)
	if err != nil {
		log.Panicln(err)
	}

	Console(response)
}

func Console(data interface{}) {
	b, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		log.Println(err.Error())
	} else {
		fmt.Println(string(b))
	}
}
