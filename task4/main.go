package main

import "fmt"

type Peter interface {
	Speak()
	Move()
	Feed(food string) string
}

// Dog ///////////////////////////////////////////
type Dog struct {
	HungerLevel    int
	TirednessLevel int
}

func (d *Dog) Speak() {
	fmt.Println("Woof-Woof")
}

func (d *Dog) Move() {
	if d.HungerLevel > 50 {
		d.Speak()
		return
	}

	d.HungerLevel += 10
	fmt.Printf("Собака двигается, уровень голода = %v\n", d.HungerLevel)
}

func (d *Dog) Feed(food string) string {
	if food == "meat" {
		d.HungerLevel = 0
		return fmt.Sprintf("Собака сыта, уровень голода = %v\n\n", d.HungerLevel)
	}

	return fmt.Sprintf("Еда не подходит, уровень голода = %v\n\n", d.HungerLevel)
}

func (d *Dog) Play() {
	if d.TirednessLevel > 70 {
		fmt.Println("Собака устала, играть не будет!")
		return
	}

	d.TirednessLevel += 10
	fmt.Printf("Собака играет, уровень усталости = %v\n", d.TirednessLevel)
}

// Cat ///////////////////////////////////////////
type Cat struct {
	HungerLevel    int
	TirednessLevel int
}

func (c *Cat) Speak() {
	fmt.Println("Meow-Meow")
}

func (c *Cat) Move() {
	if c.HungerLevel > 50 {
		c.Speak()
		return
	}
	c.HungerLevel += 10
	fmt.Printf("Кошка двигается, уровень голода = %v\n", c.HungerLevel)
}

func (c *Cat) Feed(food string) string {
	if food == "fish" {
		c.HungerLevel = 0
		return fmt.Sprintf("Кошка сыта, уровень голода = %v", c.HungerLevel)
	}

	return fmt.Sprintf("Еда не подходит, уровень голода = %v", c.HungerLevel)
}

func (c *Cat) Play() {
	if c.TirednessLevel > 70 {
		fmt.Println("Кошка устала, играть не будет!")
		return
	}

	c.TirednessLevel += 10
	fmt.Printf("Кошка играет, уровень усталости = %v\n", c.TirednessLevel)
}

func main() {
	cat := Cat{
		41,
		61,
	}

	cat.Move()
	cat.Move()
	fmt.Println(cat.Feed("meat"))
	fmt.Println(cat.Feed("fish"))
	cat.Play()
	cat.Play()

}
