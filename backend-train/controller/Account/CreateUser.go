package Account

import (
	"backend-code/mail"
	"backend-code/model"
	"backend-code/util"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/bcrypt"
)

func RandomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(65 + rand.Intn(25)) //A=65 and Z = 65+25
	}
	return string(bytes)
}

func CreateNewUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	user := model.Accounts{}
	result := model.Accounts{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		util.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	res, err := password.Generate(18, 10, 0, false, false)
	user.Password = res
	passwdForEmail := user.Password
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println(err)
		util.RespondError(w, 500, "Password Encryption  failed")
	}
	user.Password = string(pass)
	user.CreatedAt = time.Now()

	db.Table("accounts").Select("email").Where("email = ?", user.Email).Scan(&result)
	if result.Email == "" {
		if err := db.Save(&user).Error; err != nil {
			util.RespondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		go func() {
			mail.SendEmailOk(user.Email, passwdForEmail)
		}()
		var resp = map[string]interface{}{"message": "account created"}
		tokenString := GetToken(user.Email)
		resp["token"] = tokenString
		util.RespondJSON(w, http.StatusAccepted, resp)

	} else {
		go func() {
			mail.SendEmailError(user.Email)
		}()
		util.RespondError(w, 500, "email exist")
	}

}
