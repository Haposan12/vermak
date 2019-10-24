package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strings"
	"vermak2/constants"
	u "vermak2/utils"
)


type Token struct {
	UserID uint
	jwt.StandardClaims
}

type Account struct {
	gorm.Model
	Email		string	`json:"email"`
	Password 	string	`json:"password"`
	Token		string	`json:"token";sql:"-"`
}

//Validasi
func (account *Account)Validate() (map[string]interface{}, bool){
	if !strings.Contains(account.Email, "@"){
		return u.Message(constants.FalseStatus, "Email address is requires"), false
	}

	if len(account.Password) < 6{
		return u.Message(constants.FalseStatus, "Password is required"), false
	}

	//emall must unique
	temp := &Account{}

	err := GetDB().Table("accounts").Where("email = ?", account.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return u.Message(constants.FalseStatus, "Connection error. Please retry"), false
	}
	if temp.Email != ""{
		return u.Message(constants.FalseStatus, "Email address already in use by another account"), false
	}
	return u.Message(constants.TrueStatus, "Requirement passed"), true
}

func (account *Account) Create() (map[string]interface{})  {
	if resp, ok := account.Validate(); !ok{
		return resp
	}

	hashPass, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashPass)

	GetDB().Create(account)

	if account.ID <= 0{
		return u.Message(constants.FalseStatus, "Failed to create account, connection error.")
	}

	//create new jwt token
	tk := &Token{UserID:account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	account.Password = ""

	response := u.Message(constants.TrueStatus, "Account has been created")
	//response["account"] = account
	return response
}

func Login(email , password string)(map[string]interface{})  {
	account := &Account{}

	err := GetDB().Table("accounts").Where("email = ?", email).First(account).Error
	if err != nil{
		if err == gorm.ErrRecordNotFound{
			return u.Message(constants.FalseStatus, "Email address not found")
		}
		return u.Message(constants.FalseStatus, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err ==  bcrypt.ErrMismatchedHashAndPassword{
		return u.Message(constants.FalseStatus, "Invalid login credentials. Please try again.")
	}

	account.Password = ""
	tk := &Token{UserID:account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	resp := u.Message(constants.TrueStatus, "Logged In")
	//resp["account"] = account

	return resp
}

func GetUser(u uint) *Account  {
	acc := &Account{}
	GetDB().Table("accounts").Where("id = ?", u).First(acc)
	if acc.Email == ""{
		return nil
	}
	acc.Password = ""
	return acc
}