package main

import "fmt"

type human struct {
	FirstName string
	LastName  string
	Age       int
	Gender    string
	Height    float64
	Weight    float64
	IsMarried bool
}

// Full name of the human / Полное имя человека
func (h human) fullName() string {
	return fmt.Sprintf("%s %s", h.FirstName, h.LastName)
}

// Check if the human is an adult / Проверка, является ли человек взрослым
func (h human) isAdult() bool {
	return h.Age >= 18
}

// Body Mass Index (BMI) calculation / Расчет индекса массы тела (ИМТ)
func (h human) bmi() float64 {
	if h.Height <= 0 {
		return 0
	}
	return h.Weight / (h.Height * h.Height)
}

// Marital status of the human / Семейное положение человека
func (h human) maritalStatus() string {
	if h.IsMarried {
		return "Married"
	}

	return "Single"
}

type action struct {
	human
	Activity string
	Duration int
}

func main() {
	h := human{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
		Gender:    "male",
		Height:    1.75,
		Weight:    70.0,
		IsMarried: true,
	}

	a := action{
		human:    h,
		Activity: "Running",
		Duration: 30, // in minutes
	}

	fmt.Printf("Full Name: %s\n", a.fullName())
	fmt.Printf("Is Adult: %t\n", a.isAdult())
	fmt.Printf("Gender: %s\n", a.Gender)
	fmt.Printf("BMI: %.2f\n", a.bmi())
	fmt.Printf("Marital Status: %s\n", a.maritalStatus())
	fmt.Printf("Activity: %s for %d minutes\n", a.Activity, a.Duration)
}
