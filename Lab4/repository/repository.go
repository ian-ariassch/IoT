package repository

import (
	"context"
	"fmt"
	"minimal/db"
	"minimal/ent"
	"minimal/ent/water"
)

type WaterRepository struct {
	Repo *ent.Client
}

func NewWaterRepository() WaterRepository {
	client := db.NewPostgresClient()

	return WaterRepository{
		Repo: client,
	}
}

func (r WaterRepository) CreateRegister(liters float64, topic string) error {
	_, err := r.Repo.Water.Create().SetLiters(liters).SetTopic(topic).Save(context.Background())

	if err != nil {
		fmt.Printf("failed creating register: %s", err.Error())
		return fmt.Errorf("failed creating register: %s", err.Error())
	}

	fmt.Println("create register query executed")
	return nil
}

func (r WaterRepository) GetRegisters(topic string) ([]*ent.Water, error) {
	registers, err := r.Repo.Water.Query().Where(water.TopicContains(topic)).All(context.Background())

	if err != nil {
		fmt.Printf("failed getting registers: %s", err.Error())
		return nil, fmt.Errorf("failed getting registers: %s", err.Error())
	}
	fmt.Println("get registers query executed")
	return registers, nil
}
