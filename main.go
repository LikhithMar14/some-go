package main

import (
	"errors"
	"fmt"
	"log"
)
type Truck interface{
	LoadCargo() error
	UnloadCargo() error

}
type NormalTruck struct{
	id string
	cargo int
}
type ElectricTruck struct{
	id string
	cargo int
	battery float64
}
var (
	ErrorNotImplemented = errors.New("not implemented")
	ErrorTruckNotFound = errors.New("truck not found")
)
func processTruck(truck Truck) error {
	fmt.Printf("Processing truck: %+v \n",truck)
	if err := truck.LoadCargo(); err != nil{
		return fmt.Errorf("error loading cargo: %w",err)
	}

	return nil
}

func(t *NormalTruck) LoadCargo() error{
	t.cargo += 100
	return nil
}
func(t *NormalTruck) UnloadCargo() error{
	t.cargo -= 100
	return nil
}
func(t *ElectricTruck) LoadCargo() error{
	t.cargo += 100
	t.battery -= 10
	return nil
}
func(t *ElectricTruck) UnloadCargo() error{
	t.cargo -= 100
	return nil
}
func main(){
	nt := NormalTruck{id : "Truck-1",cargo : 100}
	et := ElectricTruck{id : "eTruck-1",battery : 100,cargo : 100}
	err := processTruck(&nt)
	// can be replaced with any interfce{} == any
	person := make(map[string]interface{},0)
	person["name"] = "John"
	person["age"] = 20
	person["city"] = "New York"
	
	age, ok := person["age"].(int)
	if !ok{
		log.Fatalf("Age does not exist")
	}
	log.Println(age)

	if err != nil{
		log.Fatalf("Error processing truck: %v\n",err)
	}
	err = processTruck(&et)
	if err != nil{
		log.Fatalf("Error processing truck: %v\n",err)
	}
	log.Println(&nt)
	log.Println(&et)
	
}