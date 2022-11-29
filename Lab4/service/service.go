package service

import (
	"encoding/json"
	"minimal/ent"
	"minimal/mqtt"
	"minimal/repository"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type WaterService struct {
	repo repository.WaterRepository
	man  mqtt.MqttManager
}

func NewWaterService(repo repository.WaterRepository, man mqtt.MqttManager) WaterService {
	return WaterService{
		repo: repo,
		man:  man,
	}
}

func (s WaterService) Subscribe(topic string) error {

	onMessageReceived := func(client MQTT.Client, message MQTT.Message) {
		payload := struct {
			Liters float64 `json:"liters"`
		}{}

		err := json.Unmarshal(message.Payload(), &payload)
		if err != nil {
			print(err)
		}

		err = s.repo.CreateRegister(payload.Liters, topic)
		if err != nil {
			print(err)
		}
	}

	s.man.Subscribe(topic, onMessageReceived)

	return nil
}

func (s WaterService) GetRegisters(topic string) ([]*ent.Water, error) {
	registers, err := s.repo.GetRegisters(topic)
	if err != nil {
		return nil, err
	}

	return registers, nil
}
