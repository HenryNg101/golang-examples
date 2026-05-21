package main

import "fmt"

// Products interfaces
type Button interface {
	Render()
}

type Checkbox interface {
	Render()
}

// Concrete products family
// Window family
type WindowsButton struct{}

func (WindowsButton) Render() {
	fmt.Println("Windows button")
}

type WindowsCheckbox struct{}

func (WindowsCheckbox) Render() {
	fmt.Println("Windows checkbox")
}

// Mac family
type MacButton struct{}

func (MacButton) Render() {
	fmt.Println("Mac button")
}

type MacCheckbox struct{}

func (MacCheckbox) Render() {
	fmt.Println("Mac checkbox")
}

// Abstract factory
type GUIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

// Concrete factories
// Factories for Window's set
type WindowsFactory struct{}

func (WindowsFactory) CreateButton() Button {
	return WindowsButton{}
}

func (WindowsFactory) CreateCheckbox() Checkbox {
	return WindowsCheckbox{}
}

// Factories for Mac's set
type MacFactory struct{}

func (MacFactory) CreateButton() Button {
	return MacButton{}
}

func (MacFactory) CreateCheckbox() Checkbox {
	return MacCheckbox{}
}

// Usage
// If you want Mac's facmily, just replace factory := WindowsFactory{} with factory := MacFactory{}
func main() {
	factory := WindowsFactory{}
	btn := factory.CreateButton()
	chk := factory.CreateCheckbox()

	btn.Render()
	chk.Render()
}
