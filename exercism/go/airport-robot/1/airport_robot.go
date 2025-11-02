package airportrobot

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(string) string
}

// German language.
type German struct{}

// LanguageName returns German.
func (de German) LanguageName() string {
	return "German"
}

// Greet accepts a name and returns a string for the robot to say.
func (de German) Greet(name string) string {
	greeting := fmt.Sprintf("Hallo %s!", name)
	return greeting
}

// Italian language.
type Italian struct{}

// LanguageName returns Italian.
func (it Italian) LanguageName() string {
	return "Italian"
}

// Greet accepts a name and returns a string for the robot to say.
func (it Italian) Greet(name string) string {
	greeting := fmt.Sprintf("Ciao %s!", name)
	return greeting
}

// Portuguese language.
type Portuguese struct{}

// LanguageName returns Portuguese.
func (pt Portuguese) LanguageName() string {
	return "Portuguese"
}

// Greet accepts a name and returns a string for the robot to say.
func (pt Portuguese) Greet(name string) string {
	greeting := fmt.Sprintf("Olá %s!", name)
	return greeting
}

// SayHello accepts the name of the visitor and a language in which to repsond.
func SayHello(name string, g Greeter) string {
	greeting := fmt.Sprintf("I can speak %s: %s", g.LanguageName(), g.Greet(name))
	return greeting
}

func main() {
	fmt.Println(SayHello("Dietrich", German{}))
}
