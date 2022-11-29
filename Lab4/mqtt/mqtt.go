package mqtt

import (
	"crypto/tls"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
)

type MqttManager struct {
	c MQTT.Client
}

func NewMqttManager() *MqttManager {
	return &MqttManager{
		c: MQTT.NewClient(getClientOptions()),
	}
}

func (man MqttManager) CreateClient() error {
	if token := man.c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return nil
}

func (man MqttManager) Subscribe(topic string, onMessageReceived func(MQTT.Client, MQTT.Message)) error {
	if token := man.c.Subscribe(topic, byte(0), onMessageReceived); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}

func getClientOptions() *MQTT.ClientOptions {
	id := uuid.New()

	return MQTT.NewClientOptions().AddBroker(
		"tcp://tcp-mo5.mogenius.io:20242",
	).SetClientID(
		id.String(),
	).SetCleanSession(true).SetTLSConfig(
		&tls.Config{
			InsecureSkipVerify: true,
			ClientAuth:         tls.NoClientCert,
		},
	)
}
