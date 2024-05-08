package validators

import (
	"errors"
	"regexp"

	"github.com/somos831/somos-backend/models"
)

func (v *Validator) ValidateNewUser(newUser models.User) error {

	if err := validateUserFields(newUser); err != nil {
		return err
	}

	if newUser.Password == "" {
		return errors.New("Password is a required field")
	}

	newUsername, err := v.isUniqueUsername(newUser.Username)
	if err != nil {
		return err
	}

	if !newUsername {
		return errors.New("Username is already taken")
	}

	newAccount, err := v.isUniqueEmail(newUser.Email)
	if err != nil {
		return err
	}

	if !newAccount {
		return errors.New("An account with the given email already exists")
	}

	return nil
}

func (v *Validator) ValidateUpdatedFields(user models.User) error {

	if err := validateUserFields(user); err != nil {
		return err
	}

	// Get current values: email and username
	userRec, err := models.FindUserByID(v.DB, user.ID)
	if err != nil {
		return err
	}

	if userRec == nil {
		return errors.New("user not found")
	}

	// Email is being updated
	if userRec.Email != user.Email {
		newEmail, err := v.isUniqueEmail(user.Email)

		if err != nil {
			return err
		}

		if !newEmail {
			return errors.New("An account with the email provided already exists")
		}
	}

	// Username is being updated
	if userRec.Username != user.Username {
		uniqueUsername, err := v.isUniqueUsername(user.Username)

		if err != nil {
			return err
		}

		if !uniqueUsername {
			return errors.New("Username is already taken")
		}
	}

	err = models.UpdateUser(v.DB, &user)
	if err != nil {
		return err
	}

	return nil
}

func validateUserFields(user models.User) error {

	if user.Username == "" || user.Email == "" {
		return errors.New("Username and email are required fields")
	}

	if user.StatusID == 0 || user.RoleID == 0 {
		return errors.New("User must be assigned a role and status")
	}

	if !isValidEmail(user.Email) {
		return errors.New("Email address provided is not valid")
	}

	return nil
}

func isValidEmail(email string) bool {

	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)

	return regex.MatchString(email)
}

func (v *Validator) isUniqueUsername(username string) (bool, error) {

	exists, err := models.UserExistsByUsername(v.DB, username)
	if err != nil {
		return false, errors.New("Unable to check username: " + err.Error())
	}

	return !exists, nil
}

func (v *Validator) isUniqueEmail(email string) (bool, error) {

	exists, err := models.UserExistsByEmail(v.DB, email)
	if err != nil {
		return false, errors.New("Unable to check email: " + err.Error())
	}

	return !exists, nil
}
