package main

import(
	"golang.org/x/crypto/bcrypt"
)

func password_Crypt(password []byte) ([]byte, error){
	return bcrypt.GenerateFromPassword(password, 10);
}

func compare_Crypt(password []byte, salt []byte, hash []byte) error{
	salted_password := make([]byte, len(password) + len(salt));
	copy(salted_password[:len(password)], password[:]);
	copy(salted_password[len(password):], salt[:]);	
	return bcrypt.CompareHashAndPassword(hash, salted_password);
}
