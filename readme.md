# Golang com Elasticsearch

## Rodar o Elasticsearch com Docker

```bash
docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" --rm --name elasticsearch -d docker.elastic.co/elasticsearch/elasticsearch:7.5.2

# test elasticsearch api
curl -X GET "localhost:9200/_cat/nodes?v&pretty"
```

Referência:
https://www.elastic.co/guide/en/elasticsearch/reference/7.5/docker.html

## Comandos

Adicionar um index:

```bash
curl -X PUT http://localhost:9200/students?pretty
```

Saída

```json
{
  "acknowledged" : true,
  "shards_acknowledged" : true,
  "index" : "students"
}
```

Adicionar dados ao index:

```bash
curl -H "Content-Type: application/json" -d "{\"name\":\"Alice\",\"age\":17,\"average_score\":81.1}" http://localhost:9200/students/doc?pretty
```

```json
{
  "_index" : "students",
  "_type" : "doc",
  "_id" : "yshojncBnR5yqnMsacCj",
  "_version" : 1,
  "result" : "created",
  "_shards" : {
    "total" : 2,
    "successful" : 1,
    "failed" : 0
  },
  "_seq_no" : 0,
  "_primary_term" : 1
}

````

Adicionar dados em massa a partir de um arquivo

```bash
curl -XPOST -H "Content-Type: application/json" localhost:9200/students/_bulk?pretty --data-binary @bulk.json
```

Saída

```json
{
  "took" : 7,
  "errors" : false,
  "items" : [
    {
      "index" : {
        "_index" : "students",
        "_type" : "_doc",
        "_id" : "xchOjncBnR5yqnMs98Ag",
        "_version" : 1,
        "result" : "created",
        "_shards" : {
          "total" : 2,
          "successful" : 1,
          "failed" : 0
        },
        "_seq_no" : 5,
        "_primary_term" : 1,
        "status" : 201
      }
    },
    {
      "index" : {
        "_index" : "students",
        "_type" : "_doc",
        "_id" : "xshOjncBnR5yqnMs98Ag",
        "_version" : 1,
        "result" : "created",
        "_shards" : {
          "total" : 2,
          "successful" : 1,
          "failed" : 0
        },
        "_seq_no" : 6,
        "_primary_term" : 1,
        "status" : 201
      }
    },
    {
      "index" : {
        "_index" : "students",
        "_type" : "_doc",
        "_id" : "x8hOjncBnR5yqnMs98Ag",
        "_version" : 1,
        "result" : "created",
        "_shards" : {
          "total" : 2,
          "successful" : 1,
          "failed" : 0
        },
        "_seq_no" : 7,
        "_primary_term" : 1,
        "status" : 201
      }
    },
    {
      "index" : {
        "_index" : "students",
        "_type" : "_doc",
        "_id" : "yMhOjncBnR5yqnMs98Ag",
        "_version" : 1,
        "result" : "created",
        "_shards" : {
          "total" : 2,
          "successful" : 1,
          "failed" : 0
        },
        "_seq_no" : 8,
        "_primary_term" : 1,
        "status" : 201
      }
    }
  ]
}
```

Fazer uma busca simples

```bash
curl -H "Content-Type: application/json" -d "{\"query\":{\"match\": {\"name\": \"doe\"}}}" http://localhost:9200/_search?pretty
```

Saída

```json
{
  "took" : 5,
  "timed_out" : false,
  "_shards" : {
    "total" : 1,
    "successful" : 1,
    "skipped" : 0,
    "failed" : 0
  },
  "hits" : {
    "total" : {
      "value" : 4,
      "relation" : "eq"
    },
    "max_score" : 0.6899493,
    "hits" : [
      {
        "_index" : "students",
        "_type" : "doc",
        "_id" : "wchOjncBnR5yqnMsCcCh",
        "_score" : 0.6899493,
        "_source" : {
          "name" : "john doe",
          "age" : 18,
          "average_score" : 77.7
        }
      },
      {
        "_index" : "students",
        "_type" : "doc",
        "_id" : "w8hOjncBnR5yqnMsCcCh",
        "_score" : 0.6899493,
        "_source" : {
          "name" : "mary doe",
          "age" : 18,
          "average_score" : 97.7
        }
      },
      {
        "_index" : "students",
        "_type" : "doc",
        "_id" : "xchOjncBnR5yqnMs98Ag",
        "_score" : 0.6899493,
        "_source" : {
          "name" : "john doe",
          "age" : 18,
          "average_score" : 77.7
        }
      },
      {
        "_index" : "students",
        "_type" : "doc",
        "_id" : "x8hOjncBnR5yqnMs98Ag",
        "_score" : 0.6899493,
        "_source" : {
          "name" : "mary doe",
          "age" : 18,
          "average_score" : 97.7
        }
      }
    ]
  }
}
```

Deletar Index

```
curl -X DELETE localhost:9200/students
```

Saída

```json
{"acknowledged":true}
```


## Instalar a biblioteca do Elasticsearch para Golang



```bash
go get -u -v gopkg.in/olivere/elastic.v7
```

https://gopkg.in/olivere/elastic.v7

### Referências

https://www.freecodecamp.org/news/go-elasticsearch/
