package service

import "github.com/saravase/golang_echo_api/entity"

type PlantService interface {
	Save(entity.Plant) entity.Plant
	FindAll() []entity.Plant
}

type service struct {
	plants []entity.Plant
}

func New() PlantService {
	return &service{
		plants: []entity.Plant{},
	}
}

func (service *service) Save(plant entity.Plant) entity.Plant {
	service.plants = append(service.plants, plant)
	return plant
}

func (service *service) FindAll() []entity.Plant {
	return service.plants
}
