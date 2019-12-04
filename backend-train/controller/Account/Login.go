package Account

import (
	"backend-code/model"
	"backend-code/util"
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func findone(db *gorm.DB, email, password string) map[string]interface{} {
	account := &model.Accounts{}

	if err := db.Where("email = ?", email).First(account).Error; err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Email address not found"}
		return resp
	}

	errf := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials. Please try again"}
		return resp
	}

	tokenString := GetToken(account.Email)

	var resp = map[string]interface{}{"status": false, "message": "logged in"}
	resp["token"] = tokenString
	resp["user"] = account
	return resp
}

func Login(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	user := model.Accounts{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		util.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	email := user.Email
	passwd := user.Password

	accountLogin := GetAccountOr404(db, email, w, r)

	if accountLogin == nil {
		return
	}
	resp := findone(db, email, passwd)
	util.RespondJSON(w, http.StatusAccepted, resp)
}
