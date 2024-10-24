package main

import (
	"fmt"
	"math/rand"
)

// User //////////////////////////////////////////////////
type User struct {
	Username            string
	Password            string
	Role                string
	FailedLoginAttempts int
	StatusAccount       string
}

func (u *User) Authenticate(password string) {
	if u.Password == password && u.StatusAccount != "Locked" {
		fmt.Printf("Пользователь %v успешно авторизовался\n\n", u.Username)
		u.FailedLoginAttempts = 0
		return
	} else {
		u.BlockAccount()
		if u.StatusAccount == "Locked" {
			activation := ""
			fmt.Print("Хотите разблокировать учетную запись? (введите да/нет): ")
			_, err := fmt.Scan(&activation)
			if err != nil {
				panic("ошибка ввода")
			}

			if activation == "да" {
				err = u.ActivateAccount()
				if err != nil {
					fmt.Println(err)
				}
				return
			}

			return
		}

		fmt.Printf("Пользователь %v ввел некорректный пароль, попробуйте еще раз\nОсталось попыток входа - %v\n\n", u.Username, 3-u.FailedLoginAttempts)
		u.FailedLoginAttempts++
	}

}

func (u *User) UpdatePassword(newPassword string) error {
	if newPassword == u.Password || newPassword == "" {
		return fmt.Errorf("\nпароль %v невалиден, попробуйте другой пароль", newPassword)
	}

	u.Password = newPassword
	return nil
}

func (u *User) ResetPassword() error {
	newPassword := ""
	fmt.Print("\nВведите новый пароль, чтобы активировать учетную запись: ")
	_, err := fmt.Scan(&newPassword)
	if err != nil {
		return fmt.Errorf("oшибка ввода пароля, повторите попытку: %v", err)
	}

	err = u.UpdatePassword(newPassword)
	if err != nil {
		return err
	}

	fmt.Println("Пароль успешно сброшен и обновлен на новый!")

	return nil
}

func (u *User) HasAccess(resource string) bool {
	if u.Role == "админ" && (resource == "консоль администратора" || resource == "вебсайт") {
		return true
	} else if u.Role == "пользователь" && resource == "вебсайт" {
		return true
	} else {
		fmt.Printf("Подключение невозможно! Неопознанный ресурс - %v или некорректная роль - %v\n", resource, u.Role)
		return false
	}
}

func (u *User) BlockAccount() {
	if u.FailedLoginAttempts >= 3 {
		u.StatusAccount = "Locked"
		fmt.Printf("Учетная запись пользователя %v заблокирована!\nКоличество неудачных попыток входа %v\n", u.Username, u.FailedLoginAttempts+1)
	}
}

func (u *User) ActivateAccount() error {
	err := u.ResetPassword()
	if err != nil {
		return err
	}

	u.StatusAccount = "Activated"
	fmt.Println("Аккаунт активирован!")
	return nil
}

func (u *User) SendOTP() {
	otp := make([]rune, 5)

	for i := 0; i < 5; i++ {
		otp[i] = rand.Int31n(10) + 48
	}

	fmt.Printf("Ваш одноразовый пароль - %v\n\n", string(otp))
}

func main() {

	user := User{
		Username:            "Паша",
		Password:            "12345",
		Role:                "пользователь",
		FailedLoginAttempts: 0,
		StatusAccount:       "Activated",
	}

	////fmt.Println(user.HasAccess("консоль админа"))
	//user.SendOTP()
	user.Authenticate("123")
	user.Authenticate("12")
	user.Authenticate("1234")
	user.Authenticate("1234")
	user.Authenticate("12345")
	user.Authenticate("aaa")
	fmt.Println(user)
}
