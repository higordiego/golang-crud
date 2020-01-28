package model

import (
	_ "github.com/go-sql-driver/mysql"
	Config "github.com/golang-crud/config"
	"golang.org/x/crypto/bcrypt"
	"log"
)

// User tipo de estrutura de dados
type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) setPassword() {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}
	u.Password = string(hash)
}

// Authenticator - autenticação do usuario
func Authenticator(u *User) (err error) {
	u.setPassword()
	if err = Config.DB.Model(u).Where("email ? and password = ?", u.Email, u.Password).Error; err != nil {
		return err
	}
	return nil
}

// FindOne - busca de usuario por seu id
func FindOne(u *User, id string) (err error) {
	if err = Config.DB.Begin().First(u, id).Error; err != nil {
		return err
	}
	return nil
}

// FindAll retornando todos os usuários
func FindAll(u *[]User) (err error) {
	if err = Config.DB.Find(u).Error; err != nil {
		return err
	}
	return nil
}

// CreateUser criando usuarios e setando password
func CreateUser(u *User) (err error) {
	u.setPassword()
	if err = Config.DB.Create(u).Error; err != nil {
		return err
	}
	return nil
}

// UpdateUser - update de user
func UpdateUser(u *User, id string) (err error) {
	if err = Config.DB.Model(u).Update(u).Where("id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
