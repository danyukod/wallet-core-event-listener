package main

import (
	"database/sql"
	"fmt"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/danyukod/wallet-core-event-listener/internal/database"
	"github.com/danyukod/wallet-core-event-listener/internal/event/listener"
	"github.com/danyukod/wallet-core-event-listener/internal/usecase/create_balance"
	"github.com/danyukod/wallet-core-event-listener/internal/usecase/get_balance"
	"github.com/danyukod/wallet-core-event-listener/internal/web"
	"github.com/danyukod/wallet-core-event-listener/internal/web/webserver"
	"github.com/danyukod/wallet-core-event-listener/pkg/kafka"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql", "3306", "wallet"))
	if err != nil {
		panic(err)
	}

	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	}

	balanceDb := database.NewBalanceDB(db)

	kafkaConsumer := kafka.NewConsumer(&configMap, []string{"balances"})
	createBalanceUseCase := create_balance.NewCreateBalanceUseCase(balanceDb)
	kafkaListener := listener.NewCreateBalanceKafkaListener(kafkaConsumer, createBalanceUseCase)
	kafkaListener.Listen()

	webserver := webserver.NewWebServer(":8080")
	getBalanceUseCase := get_balance.NewGetBalanceUseCase(balanceDb)
	balanceHandler := web.NewWebBalanceHandler(*getBalanceUseCase)

	webserver.AddHandler("/balances/{account_id}", balanceHandler.GetBalance)

	fmt.Println("Server is running")
	webserver.Start()
}
