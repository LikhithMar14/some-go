package main

import (
	"errors"
	"fmt"
	"log"
)

type Truck struct{
	id string
	cargo int
}

var (
	ErrorNotImplemented = errors.New("not implemented")
	ErrorTruckNotFound = errors.New("truck not found")
)
func processTruck(truck Truck) error {
	fmt.Printf("Processing Truck %s\n",truck.id)
	if err := truck.LoadCargo(); err != nil{
		return fmt.Errorf("error loading cargo: %w",err)
	}
	return nil
}

func(t *Truck) LoadCargo() error{
	return nil
}
func(t *Truck) UnloadCargo() error{
	return nil
}
func main(){
	trucks := []Truck{
		{id : "Truck-1"},
		{id : "Truck-2"},
		{id : "Truck-3"},
	}
	for _,truck := range trucks{
		fmt.Printf("Processing truck: %s\n",truck.id)
		err := processTruck(truck); if err != nil{
			log.Fatalf("Error processing truck: %v\n",err)
		}
	}
	
	
}