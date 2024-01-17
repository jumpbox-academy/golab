package main

import "fmt"

func main() {
	fmt.Println(TotalBirdCount([]int{1, 2, 3, 4, 5, 6, 7, 8,}))
	fmt.Println(BirdsInWeek([]int{1, 2, 3, 4, 5, 6, 7, 8,}, 2))
	fmt.Println(FixBirdCountLog([]int{1, 2, 3, 4, 5, 6, 7, 8}))
}
// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var birds int
	for i := 0; i < len(birdsPerDay); i++ {
		birds += birdsPerDay[i]
	}

	return birds
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var birds int
	for i := (week-1)*7; i < week*7; i++ {
		birds += birdsPerDay[i]
	}

	return birds
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i+=2 {
		birdsPerDay[i]++
	}

	return birdsPerDay
}