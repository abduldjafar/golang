package Account

import (
	"backend-code/model"
	"backend-code/util"
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
)

func DeleteAccount(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	user := model.Accounts{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		util.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	name := user.Email
	account := GetAccountOr404(db, name, w, r)

	data := db.Where("password = ? ", account.Password).Delete(&account)

	if err := data.Error; err != nil {
		util.RespondError(w, http.StatusInternalServerError, "Internal Service Error")
		return
	}

	util.RespondOk(w, http.StatusOK, "Account Deleted")
}
