package listener

import (
	"encoding/json"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/danyukod/wallet-core-event-listener/internal/usecase/create_balance"
	"github.com/danyukod/wallet-core-event-listener/pkg/kafka"
)

type CreateBalanceKafkaListener struct {
	Kafka   *kafka.Consumer
	Usecase *create_balance.CreateBalanceUseCase
}

type CreateBalanceKafkaListenerDTO struct {
	AccountIDFrom        string  `json:"account_id_from"`
	AccountIDTo          string  `json:"account_id_to"`
	BalanceAccountIDFrom float64 `json:"balance_account_id_from"`
	BalanceAccountIDTo   float64 `json:"balance_account_id_to"`
}

func NewCreateBalanceKafkaListener(kafka *kafka.Consumer, usecase *create_balance.CreateBalanceUseCase) *CreateBalanceKafkaListener {
	return &CreateBalanceKafkaListener{
		Kafka:   kafka,
		Usecase: usecase,
	}
}

func (l *CreateBalanceKafkaListener) Listen() {
	kafkaMessageChan := make(chan *ckafka.Message)
	go l.Kafka.Consume(kafkaMessageChan)
	for msg := range kafkaMessageChan {
		var dto CreateBalanceKafkaListenerDTO
		err := json.Unmarshal(msg.Value, &dto)
		if err != nil {
			continue
		}
		balanceFrom := create_balance.CreateBalanceInputDTO{
			AccountID: dto.AccountIDFrom,
			Balance:   dto.BalanceAccountIDFrom,
		}

		balanceTo := create_balance.CreateBalanceInputDTO{
			AccountID: dto.AccountIDTo,
			Balance:   dto.BalanceAccountIDTo,
		}
		err = l.Usecase.Execute(balanceFrom)
		if err != nil {
			return
		}
		err = l.Usecase.Execute(balanceTo)
		if err != nil {
			return
		}
	}

}
