package main

import (
	"fmt"
	"sort"
	"time"
)

// Book /////////////////////////////////////////
type Book struct {
	Title  string
	Author string
	Year   int
	Copies int
}

func (b *Book) GetDescription() string {
	return fmt.Sprintf("Title: %v; Author: %v; Year: %v; Copies: %v.\n", b.Title, b.Author, b.Year, b.Copies)
}

// UserBook User ///////////////////////////////////////
type UserBook struct {
	Book
	RentalDate time.Time
}

// User ////////////////////////////////////////////////////////
type User struct {
	id          int
	FirstName   string
	SecondName  string
	BirthDate   time.Time
	RentalBooks []UserBook
}

// Library ////////////////////////////////////////////////
type Library struct {
	Books []Book
}

func (l *Library) AddBook(book Book) {
	l.Books = append(l.Books, book)
}

func (l *Library) BorrowBook(title string) error {
	for _, book := range l.Books {
		if book.Title == title {
			if book.Copies > 0 {
				book.Copies--
				return nil
			} else {
				return fmt.Errorf("number of copies of the book %v = %v\n", book.Title, book.Copies)
			}
		}
	}

	return fmt.Errorf("there is no book with the title = %v in the library", title)
}

func (l *Library) GiveBookForUser(title string, user User) error {
	err := l.BorrowBook(title)
	if err != nil {
		return err
	}

	for _, book := range l.Books {
		if book.Title == title {
			user.RentalBooks = append(user.RentalBooks, UserBook{
				Book: Book{ //Решил возвращать не просто переменную book, а в таком виде,
					// чтоб количество выданных копий было 1 (упрощенная модель, пользователь не может взять больше одного экземпляра книги на себя)
					Title:  book.Title,
					Author: book.Author,
					Year:   book.Year,
					Copies: 1,
				},
				RentalDate: time.Now(),
			})
		}
	}

	return nil
}

func (l *Library) ReturnBook(title string) {
	for _, book := range l.Books {
		if book.Title == title {
			book.Copies++
		}
	}

}

// NotificationOfDelay ///////////////////////////////////////////////////////////////////////
func (u *User) NotificationOfDelay() string { // Как рациональней реализовать систему просрочки? Метод для проверки каждого читателя
	// или же слайс читателей в функцию. Или же в библиотеке хранить слайс с читателя и через метод библиотеки проверять (возвращать мапу[Юзер]список просроченных книг)?
	delayMessage := fmt.Sprintf("User id = %v it has the following delays:\n", u.id)
	var Notifications []string
	for _, userBook := range u.RentalBooks {
		currentDate := time.Now()
		duration := currentDate.Sub(userBook.RentalDate)
		dayPassed := int(duration.Hours() / 24)
		if dayPassed > 30 {
			Notifications = append(Notifications, fmt.Sprintf("-- Author: %v; Title: %v; Year: %v; The delay period (days): %v.\n", userBook.Author,
				userBook.Title, userBook.Year, dayPassed))
		}
	}

	if len(Notifications) > 0 {
		for _, notification := range Notifications {
			delayMessage += notification
		}
		return delayMessage

	} else {
		delayMessage = fmt.Sprintf("User id = %v has no delays:\n", u.id)
		return delayMessage
	}
}

// SortBooksByQuantity //////////////////////////////////////////////////////////////////////////////
func (l *Library) SortBooksByQuantity() {
	sort.Slice(l.Books, func(i, j int) bool {
		return l.Books[i].Copies > l.Books[j].Copies
	})
}

// GetTopPopularBooks ///////////////////////////////////////////////////////////////////////////////////////
func (l *Library) GetTopPopularBooks(limit int) []Book {
	l.SortBooksByQuantity()
	for i, book := range l.Books {
		if book.Copies <= limit {
			return l.Books[i:]
		}
	}

	return []Book{}
}

func main() {
	Pasha := User{
		id:         1,
		FirstName:  "Paul",
		SecondName: "Reznikov",
		BirthDate:  time.Date(1999, time.November, 1, 0, 0, 0, 0, time.UTC),
		RentalBooks: []UserBook{
			{
				Book: Book{
					Title:  "Зеленая миля",
					Author: "Стивен Кинг",
					Year:   1997,
					Copies: 1,
				},
				RentalDate: time.Date(2024, time.September, 19, 0, 0, 0, 0, time.UTC),
			},
			{
				Book: Book{
					Title:  "Властелин Колец",
					Author: "Джон Р. Р. Толкин",
					Year:   1955,
					Copies: 1,
				},
				RentalDate: time.Date(2024, time.October, 20, 0, 0, 0, 0, time.UTC),
			},
			{
				Book: Book{
					Title:  "Граф Монте-Кристо",
					Author: "Александр Дюма",
					Year:   1846,
					Copies: 1,
				},
				RentalDate: time.Date(2024, time.September, 4, 0, 0, 0, 0, time.UTC),
			},
		},
	}

	fmt.Println(Pasha.NotificationOfDelay())

	library := Library{}

	library.AddBook(Book{Title: "Зеленая миля",
		Author: "Стивен Кинг",
		Year:   1997,
		Copies: 5})

	library.AddBook(Book{Title: "Властелин Колец",
		Author: "Джон Р. Р. Толкин",
		Year:   1955,
		Copies: 2})

	library.AddBook(Book{Title: "Граф Монте-Кристо",
		Author: "Александр Дюма",
		Year:   1846,
		Copies: 7})

	library.AddBook(Book{Title: "Буря мечей",
		Author: "Джордж Мартин",
		Year:   2000,
		Copies: 1})

	fmt.Println(library.Books)
	library.SortBooksByQuantity()
	fmt.Println(library.Books)

	books := library.GetTopPopularBooks(2)
	fmt.Println(books)

}
