package main

import "fmt"

// Animal //////////////////////////////////////////////
type Animal struct {
	Name        string
	Species     string
	Sound       string
	TypeFood    string
	TiredLevel  int
	HungryLevel int
	Age         int
}

func (a *Animal) GetInfo() string {
	return fmt.Sprintf("Имя: %v; Вид: %v; Усталость: %v; Голод: %v; Возраст: %v.\n",
		a.Name, a.Species, a.TiredLevel, a.HungryLevel, a.Age)
}

func (a *Animal) Speak() {
	fmt.Println(a.Sound)
}

// Zoo /////////////////////////////////////////////////////////
type Zoo struct {
	Animals []Animal
}

func (z *Zoo) AddAnimal(animal Animal) {
	z.Animals = append(z.Animals, animal)
}

func (z *Zoo) RemoveAnimal(animal Animal) {
	for i := 0; i < len(z.Animals); i++ {
		if z.Animals[i] == animal {
			z.Animals = append(z.Animals[:i], z.Animals[i+1:]...)
			break
		}
	}
}

func (z *Zoo) FindOldestAnimal() Animal {
	maxAgeAnimal := Animal{}

	for _, animal := range z.Animals {
		if animal.Age > maxAgeAnimal.Age {
			maxAgeAnimal = animal
		}
	}

	return maxAgeAnimal
}

func (z *Zoo) FeedAllAnimals(food string) error {
	for i := range z.Animals {
		if z.Animals[i].TypeFood != food {
			return fmt.Errorf("%v не подходит животному с именем - %v,"+
				" которое относится к виду - %v, с типом еды - %v", food, z.Animals[i].Name, z.Animals[i].Species, z.Animals[i].TypeFood)
		}

		z.Animals[i].HungryLevel = 0
	}

	return nil
}

func (z *Zoo) CountBySpecies(species string) int {
	countSpecies := 0
	for i := range z.Animals {
		if z.Animals[i].Species == species {
			countSpecies++
		}
	}

	return countSpecies
}

func ScheduleEvent(name string, animals []Animal) string {
	event := fmt.Sprintf("В мероприятии \"%v\" будут участвовать следующие животные:\n", name)
	speciesMap := make(map[string]struct{})

	for _, animal := range animals {

		if _, ok := speciesMap[animal.Species]; !ok && animal.Age < 5 && animal.TiredLevel < 60 {
			speciesMap[animal.Species] = struct{}{}
			event += animal.GetInfo()
		}
	}

	return event
}

func main() {
	zoo := Zoo{Animals: []Animal{
		{
			"Дима",
			"Тигр",
			"ARARARRRR",
			"мясо",
			20,
			40,
			4,
		},
		{
			"Витя",
			"Олень",
			"FFFFAFAFFAFFF",
			"мясо",
			59,
			35,
			100,
		},
		{
			"Артур",
			"Лев",
			"ARARARRRR",
			"мясо",
			61,
			40,
			2,
		},
		{
			"Caша",
			"Лев",
			"ARARARRRR",
			"мясо",
			20,
			40,
			3,
		},
		{
			"Женя",
			"Тигр",
			"ARARARRRR",
			"мясо",
			20,
			50,
			2,
		},
	},
	}

	fmt.Println(zoo)
	zoo.AddAnimal(Animal{
		Name:        "Костя",
		Species:     "Лев",
		Sound:       "фывфывфывфы",
		TypeFood:    "мясо",
		TiredLevel:  35,
		HungryLevel: 20,
		Age:         1,
	})
	fmt.Println(zoo)
	zoo.RemoveAnimal(Animal{
		Name:        "Костя",
		Species:     "Лев",
		Sound:       "фывфывфывфы",
		TypeFood:    "мясо",
		TiredLevel:  35,
		HungryLevel: 20,
		Age:         1,
	})
	fmt.Println(zoo)
	fmt.Println(zoo.FindOldestAnimal())
	fmt.Println(zoo.FeedAllAnimals("мясо"))
	fmt.Println(zoo)
	fmt.Println(zoo)
	zoo.AddAnimal(Animal{
		Name:        "Костя",
		Species:     "Лев",
		Sound:       "фывфывфывфы",
		TypeFood:    "мясо",
		TiredLevel:  35,
		HungryLevel: 20,
		Age:         1,
	})
	fmt.Println(zoo)
	fmt.Println(zoo.CountBySpecies("Лев"))

	fmt.Println(ScheduleEvent("AAAAAAAAA", zoo.Animals))
}
