package vehicle

import "fmt"

type VehicleController interface {
	StartEngine() string
	Drive(float64) string
	StopEngine() string
}

type VehicleType struct {
	Name       string
	EngineType string
}

type Car struct {
	VehicleType
}
type ElectricCar struct {
	VehicleType
	BatteryLevel string
}
type Truck struct {
	VehicleType
}

func (c *Car) StartEngine() string {
	fmt.Printf("Started engine of %s with engine type %s \n", c.Name, c.EngineType)
	return c.EngineType
}
func (c *Car) Drive(distance float64) string {
	fmt.Printf("The car %s is driving in distance %0.2f \n", c.Name, distance)
	return c.Name
}
func (c *Car) StopEngine() string {
	fmt.Printf("Stopped engine of %s with engine type %s \n", c.Name, c.EngineType)
	return c.EngineType
}

func (e *ElectricCar) StartEngine() string {
	fmt.Printf("Started engine of %s. Battery level is %s \n", e.Name, e.BatteryLevel)
	return e.BatteryLevel
}
func (e *ElectricCar) Drive(distance float64) string {
	fmt.Printf("The car %s is driving in distance %0.2f \n", e.Name, distance)
	return e.Name
}
func (e *ElectricCar) StopEngine() string {
	fmt.Printf("Stopped engine of %s. Battery level is %s \n", e.Name, e.BatteryLevel)
	return e.BatteryLevel
}

func (t *Truck) StartEngine() string {
	fmt.Printf("Started engine of %s with engine type %s \n", t.Name, t.EngineType)
	return t.EngineType
}
func (t *Truck) Drive(distance float64) string {
	fmt.Printf("The car %s is driving in distance %0.2f \n", t.Name, distance)
	return t.Name
}
func (t *Truck) StopEngine() string {
	fmt.Printf("Stopped engine of %s with engine type %s \n", t.Name, t.EngineType)
	return t.EngineType
}

func StartEngineOfAllVehicles(vehicles []VehicleController) {
	for _, v := range vehicles {
		v.StartEngine()
	}
}

func DriveToDistance(vehicles []VehicleController, distance float64) {
	for _, v := range vehicles {
		v.Drive(distance)
	}
}

func StopEngineOfAllVehicles(vehicles []VehicleController) {
	for _, v := range vehicles {
		v.StopEngine()
	}
}
