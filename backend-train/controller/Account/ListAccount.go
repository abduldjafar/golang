package Account

import (
	"backend-code/model"
	"backend-code/util"
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
)

func ListAccounts(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	user := model.Accounts{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		util.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	accountFull := GetAccountOr404(db, user.Email, w, r)
	util.RespondJSON(w, http.StatusAccepted, accountFull)
}
