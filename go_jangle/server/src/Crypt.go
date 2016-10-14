package main

import(
	"golang.org/x/crypto/bcrypt"
)

func Password_Crypt(password []byte) ([]byte, error){
	return bcrypt.GenerateFromPassword(password, 10);
}

func Compair_Crypt(password []byte, hash []byte) error{
	return bcrypt.CompareHashAndPassword(hash, password);
}
