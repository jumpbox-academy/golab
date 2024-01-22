package main

import "fmt"

// 1 Type Declaration:
type Temperature float64
var roomTemp Temperature = 23.5

func main() {

// 4. Array:
var weekTemps [7]Temperature
// 5. Slice:
//weekendTemps := weekTemps[5:7] // Assuming index 5 and 6 are Saturday and Sunday
// 3. Zero Value:
var defaultTemp Temperature
fmt.Println("Default Temperature:", defaultTemp)

// 7. Map
tempByCity := make(map[string]Temperature)
tempByCity["New York"] = 15.0
tempByCity["Los Angeles"] = 20.0
tempByCity["Chicago"] = 10.0

for city, temp := range tempByCity {
    fmt.Printf("Temperature in %s is %.2f\n", city, temp)
	}

// 8. Put it all together
// Increase temperature of New York
if nyTemp, ok := tempByCity["New York"]; ok {
    increaseTemp(&nyTemp, 2.0)
    tempByCity["New York"] = nyTemp
}

// Update weekTemps and recalculate the average
weekTemps = [7]Temperature{10.0, 12.0, 15.0, 17.0, 20.0, 22.0, 25.0}
fmt.Println("Average weekly temperature:", averageTemp(weekTemps[:]...))

}

// 2 Pointer:
func increaseTemp(t *Temperature, amount float64) {
    *t += Temperature(amount)
}

// 6. Variadic Function:
func averageTemp(temps ...Temperature) Temperature {
    var sum Temperature
    for _, t := range temps {
        sum += t
    }
    return sum / Temperature(len(temps))
}
