package users

import (
	"fmt"
	"memes/code/db"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(username, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error while hashing password (error: %v)", err)
	}

	err = db.AddUser(username, hash)
	if err != nil {
		return err
	}

	return nil
}

func LoginUser(username, password string) (bool, error) {
	hash, err := db.GetHash(username)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil {
		return false, nil
		//return fmt.Errorf("ERROR: validating hash failed (username:%s, password:%s, error: %v)", username, password, err)
	}

	return true, nil
}
