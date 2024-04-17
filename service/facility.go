package service

import (
	"backend/backend"
	"backend/model"
	"fmt"
	"log"
)

func GetAllFacilities() ([]model.Facility, error) {
	facilities, err := backend.PGBackend.SelectAllFacilities()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	fmt.Println("Service fetched all facilities")
	return facilities, nil
}
