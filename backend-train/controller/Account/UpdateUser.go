package Account

import (
	"backend-code/model"
	"backend-code/util"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// getEmployeeOr404 gets a employee instance if exists, or respond the 404 error otherwise
func GetAccountOr404(db *gorm.DB, email string, w http.ResponseWriter, r *http.Request) *model.Accounts {
	accountTemp := model.Accounts{}
	if err := db.First(&accountTemp, model.Accounts{Email: email}).Error; err != nil {
		util.RespondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &accountTemp
}

func GetAccountDelete(db *gorm.DB, email string, w http.ResponseWriter, r *http.Request) *model.Accounts {
	accountTemp := model.Accounts{}
	if err := db.First(&accountTemp, model.Accounts{Email: email}).Order("created_at desc").Error; err != nil {
		util.RespondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &accountTemp
}

func UpdateAccount(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	account := model.Accounts{}
	accountdelete := model.Accounts{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&account); err != nil {
		util.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	updatedAccount := GetAccountOr404(db, account.Email, w, r)
	accountdelete = *updatedAccount

	pass, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)

	updatedAccount.CreatedAt = time.Now()
	fmt.Println(account.Fullname)
	if account.Fullname == ""{
		updatedAccount.Fullname = accountdelete.Fullname
	} else {
		updatedAccount.Fullname = account.Fullname
	}

	if account.Birthday == ""{
		updatedAccount.Birthday = accountdelete.Birthday
	} else {
		updatedAccount.Birthday = account.Birthday
	}

	if account.Gender == ""{
		updatedAccount.Gender = accountdelete.Gender
	} else {
		updatedAccount.Gender = account.Gender
	}

	if account.Country == ""{
		updatedAccount.Country = accountdelete.Country
	} else {
		updatedAccount.Country = account.Country
	}

	if account.Plan == ""{
		updatedAccount.Plan = accountdelete.Plan
	} else {
		updatedAccount.Plan = account.Plan
	}

	if account.Password == "" {
		updatedAccount.Password = accountdelete.Password
	} else {

		updatedAccount.Password = string(pass)
	}


	if err := db.Save(&updatedAccount).Error; err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err != nil {
		fmt.Println(err)
		util.RespondError(w, 500, "Password Encryption  failed")
	}

	accountDeleted := GetAccountDelete(db, account.Email, w, r)
	db.Where("created_at = ? ", accountDeleted.CreatedAt).Delete(&accountDeleted)
	util.RespondOk(w, http.StatusOK, "User Updated")
}
