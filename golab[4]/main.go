package main

import "fmt"

// 1. Struct Creation:
// Defining a struct 'Plane' with fields: Manufacturer, Model, and Capacity.
type Plane struct {
    Manufacturer string
    Model        string
    Capacity     int
}

// 2. Method with Receiver:
// Adding a method 'Details' to the 'Plane' struct that prints the plane's details.
func (p Plane) Details() {
    fmt.Printf("Manufacturer: %s, Model: %s, Capacity: %d\n", p.Manufacturer, p.Model, p.Capacity)
}

// 3. Pointer Receiver:
// Adding a method 'UpdateCapacity' to the 'Plane' struct to update its capacity.
// Using a pointer receiver to make changes persist.
func (p *Plane) UpdateCapacity(newCapacity int) {
    p.Capacity = newCapacity
}

// 4. Interface:
// Defining an interface 'Aircraft' with a method 'Details'.
type Aircraft interface {
    Details()
}

// The 'Plane' struct implicitly implements the 'Aircraft' interface by having a 'Details' method.

// 5. Empty Interface:
// Writing a function 'DescribeAircraft' that takes an empty interface as an argument.
// It prints out the type and value of the argument.
func DescribeAircraft(i interface{}) {
    fmt.Printf("Value: %v, Type: %T\n", i, i)
}

func main() {
    // Instantiating a 'Plane'
    myPlane := Plane{"Boeing", "747", 416}

    // Updating the capacity of 'myPlane'
    myPlane.UpdateCapacity(500)

    // 6. Putting it All Together:
    // Creating a slice of 'Aircraft' and adding different 'Plane' instances.
    var aircrafts []Aircraft = []Aircraft{
        Plane{"Airbus", "A380", 853},
        Plane{"Boeing", "777", 396},
        myPlane, // Adding the updated 'myPlane'
    }

    // Iterating over the slice, calling the 'Details' method, and using 'DescribeAircraft'.
    for _, a := range aircrafts {
        a.Details()
        DescribeAircraft(a)
    }
}
