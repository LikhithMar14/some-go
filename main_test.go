package main 

import (
	"context"
	"testing"
)

func TestMain(t *testing.T){
	t.Run("processTruck",func(t *testing.T){
		t.Run("Should load and unload a cargo truck",func(t *testing.T){
			nt := NormalTruck{id : "Truck-1",cargo : 100}
			et := ElectricTruck{id : "eTruck-1",battery : 100,cargo : 100}

			err := processTruck(context.Background(),&nt)
			if err != nil{
				t.Fatalf("Error processing truck: %v\n",err)
			}
			err = processTruck(context.Background(),&et)
			if err != nil{
				t.Fatalf("Error processing truck: %v\n",err)
			}
			if nt.cargo != 200{
				t.Fatalf("Normal truck %s has %d cargo, expected 200\n",nt.id,nt.cargo)
			}
			if et.cargo != 200{
				t.Fatalf("Electric truck %s has %d cargo, expected 200\n",et.id,et.cargo)
			}
			if et.battery != 90{
				t.Fatalf("Electric truck %s has %f battery, expected 90\n",et.id,et.battery)
			}
			
		})
	})
}