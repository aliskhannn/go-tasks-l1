package main

import "fmt"

// Встраивание структур.

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action
// от родительской структуры Human (аналог наследования).

// Structure representing a human
type human struct {
	FirstName string
	LastName  string
	Age       int
	Gender    string
	Height    float64
	Weight    float64
	IsMarried bool
}

// Full name of the human
func (h human) fullName() string {
	return fmt.Sprintf("%s %s", h.FirstName, h.LastName)
}

// Check if the human is an adult
func (h human) isAdult() bool {
	return h.Age >= 18
}

// Body Mass Index (BMI) calculation
func (h human) bmi() float64 {
	if h.Height <= 0 {
		return 0
	}
	return h.Weight / (h.Height * h.Height)
}

// Marital status of the human
func (h human) maritalStatus() string {
	if h.IsMarried {
		return "Married"
	}

	return "Single"
}

// Structure representing an action performed by a human
type action struct {
	human
	Activity string
	Duration int // in minutes
}

func main() {
	h := human{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
		Gender:    "male",
		Height:    1.75, // in meters
		Weight:    70.0, // in kilograms
		IsMarried: true,
	}

	a := action{
		human:    h,
		Activity: "Running",
		Duration: 30,
	}

	fmt.Printf("Full Name: %s\n", a.fullName())
	fmt.Printf("Is Adult: %t\n", a.isAdult())
	fmt.Printf("Gender: %s\n", a.Gender)
	fmt.Printf("BMI: %.2f\n", a.bmi())
	fmt.Printf("Marital Status: %s\n", a.maritalStatus())
	fmt.Printf("Activity: %s for %d minutes\n", a.Activity, a.Duration)
}
