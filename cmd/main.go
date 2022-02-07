package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
_ "github.com/go-sql-driver/mysql"
	repository2 "github.com/odhs/go-course/infra/repository"
	usecase2 "github.com/odhs/go-course/usecase"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	db, err := sql.Open(driverName: "mysql", dataSourceName:"root:root@tcp(mysql:3306)/db_cursos")
	if err != nil {
		log.Fatalln(err)
	}
	repository := repository2.CourseMySQLRepository{Db: db}
	usecase := usecase2.CreateCourse{repository: repository}

	var msgChan = make(chan *ckafka.Message)
	configMapConsumer := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9094",
		"group.id": "appgo",
	} 
	topics := []string{"courses"}
	consumer := kafka.NewConsumer(configMapConsumer, topics)

	go consumer.Consume(msgChan)

	for msg := range(msgChan){
		var input usecase2.CreateCourseInputDto
		json.Unmarshal(msg.Value, &input)
		output, err := usecase.Execute(input)
		if err != nil{
			fmt.Println("Error: ", err)
		} else {
			fmt.Println(output)
		}
	}
}

// {"name":"Go com Kafka", "description": "Curso Implementação Go com Kafka", "status":"Pending"}