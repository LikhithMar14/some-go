package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrTruckNotFound = errors.New("truck not found")

type FleetManager interface {
	AddTruck(id string, cargo int) error
	GetTruck(id string) (*Truck, error)
	RemoveTruck(id string) error
	UpdateTruckCargo(id string, cargo int) error
}

type Truck struct {
	ID    string
	Cargo int
}

type truckManager struct {
	trucks map[string]*Truck
}

func (m *truckManager) AddTruck(id string, cargo int) error {
	if _, ok := m.trucks[id]; ok {
		return fmt.Errorf("truck %s already exists", id)
	}
	m.trucks[id] = &Truck{ID: id, Cargo: cargo}
	return nil
}

func (m *truckManager) GetTruck(id string) (*Truck, error) {
	truck, ok := m.trucks[id]
	if !ok {
		return nil, ErrTruckNotFound
	}
	return truck, nil
}

func (m *truckManager) RemoveTruck(id string) error {
	if _, ok := m.trucks[id]; !ok {
		return ErrTruckNotFound
	}
	delete(m.trucks, id)
	return nil
}

func (m *truckManager) UpdateTruckCargo(id string, cargo int) error {
	truck, ok := m.trucks[id]
	if !ok {
		return ErrTruckNotFound
	}
	truck.Cargo = cargo
	return nil
}

func NewTruckManager() *truckManager {
	return &truckManager{
		trucks: make(map[string]*Truck),
	}
}

func main() {
	manager := NewTruckManager()

	if err := manager.AddTruck("truck1", 100); err != nil {
		log.Fatalf("AddTruck failed: %v", err)
	}
	if err := manager.AddTruck("truck2", 200); err != nil {
		log.Fatalf("AddTruck failed: %v", err)
	}

	truck, err := manager.GetTruck("truck1")
	if err != nil {
		log.Fatalf("Error getting truck: %v", err)
	}
	fmt.Printf("Truck: %+v\n", truck)

	if err := manager.UpdateTruckCargo("truck1", 150); err != nil {
		log.Fatalf("UpdateTruckCargo failed: %v", err)
	}

	truck, err = manager.GetTruck("truck1")
	if err != nil {
		log.Fatalf("Error getting truck: %v", err)
	}
	fmt.Printf("Updated truck: %+v\n", truck)

	if err := manager.RemoveTruck("truck1"); err != nil {
		log.Fatalf("RemoveTruck failed: %v", err)
	}

	truck, err = manager.GetTruck("truck1")
	if err != nil {
		fmt.Printf("Truck after removal (expected error): %v\n", err)
	} else {
		fmt.Printf("Removed truck still exists: %+v\n", truck)
	}
}
