package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)
type contextKey string
var userIDKey contextKey = "userID"
type Truck interface {
	LoadCargo() error
	UnloadCargo() error
}
type NormalTruck struct {
	id    string
	cargo int
}
type ElectricTruck struct {
	id      string
	cargo   int
	battery float64
}

var (
	ErrorNotImplemented = errors.New("not implemented")
	ErrorTruckNotFound  = errors.New("truck not found")
)

func processTruck(ctx context.Context, truck Truck) error {
	fmt.Printf("Processing truck: %+v \n", truck)


	// ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	// defer cancel()

	// delay := 3*time.Second
	// select{
	// case <-time.After(delay):
	// 	fmt.Printf("Hit Time: %v \n", truck)
	// 	break
	// case <-ctx.Done():
	// 	fmt.Printf("Timed out truck: %v \n", truck)
	// 	return ctx.Err()
	// }

	userID := ctx.Value(userIDKey)
	fmt.Printf("User ID: %v \n", userID)
	time.Sleep(1 * time.Second)
	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("error loading cargo: %w", err)
	}
	fmt.Println(time.Now())
	fmt.Printf("Finished proccessing truck %v \n", truck)

	return nil
}

func (t *NormalTruck) LoadCargo() error {
	t.cargo += 100
	return nil
}
func (t *NormalTruck) UnloadCargo() error {
	t.cargo -= 100
	return nil
}
func (t *ElectricTruck) LoadCargo() error {
	t.cargo += 100
	t.battery -= 10
	return nil
}
func (t *ElectricTruck) UnloadCargo() error {
	t.cargo -= 100
	return nil
}
func processFleet(ctx context.Context, fleet []Truck) error {
	var wg sync.WaitGroup
	errChan := make(chan error, len(fleet))

	for _, t := range fleet {
		wg.Add(1)
		go func(t Truck) {
			defer wg.Done()
			if err := processTruck(ctx, t); err != nil {
				log.Printf("Error processing truck: %v\n", err)
				errChan <- err
			}
		}(t)
	}

	wg.Wait()
	close(errChan) // Move close here, after all goroutines are done

	for err := range errChan {
		return err // return the first error
	}

	return nil
}

func main() {
	ctx := context.Background()
	//context is immutable
	//Here we are creating a new context from the previous context
	ctx = context.WithValue(ctx, userIDKey, "123")

	start := time.Now()
	fleet := []Truck{
		&NormalTruck{id: "Truck-1", cargo: 100},
		&ElectricTruck{id: "eTruck-1", battery: 100, cargo: 100},
		&NormalTruck{id: "Truck-2", cargo: 100},
		&ElectricTruck{id: "eTruck-2", battery: 100, cargo: 100},
		&NormalTruck{id: "Truck-3", cargo: 100},
		&ElectricTruck{id: "eTruck-3", battery: 100, cargo: 100},
	}
	processFleet(ctx, fleet)
	fmt.Println("Finished processing fleet")
	fmt.Println(time.Since(start))

}
