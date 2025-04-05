package consumer

import (
	"FreeGarden/src/sensor_dht11/application/use_case"
	"FreeGarden/src/sensor_dht11/domain/entities"
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
)

type MQTTAdapter struct {
	useCase *use_case.Create_DHT11
}

func NewMQTTAdapter(useCase *use_case.Create_DHT11) *MQTTAdapter {
	return &MQTTAdapter{useCase: useCase}
}

func (adapter *MQTTAdapter) HandleMessage(client mqtt.Client, msg mqtt.Message) {
	var sensor entities.DHT11Sensor

	if err := json.Unmarshal(msg.Payload(), &sensor); err != nil {
		log.Printf("Error al deserializar los datos: %v\n", err)
		return
	}

	_, err := adapter.useCase.Execute(sensor)
	if err != nil {
		log.Printf("Error al procesar los datos: %v\n", err)
		return
	}

	log.Printf("Datos procesados correctamente: %+v\n", sensor)
}
