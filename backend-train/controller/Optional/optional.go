package Optional

import (
	"backend-code/util"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"net/http"
)

func Optional(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]
	var response interface{}

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       1,  // use default DB
	})

	val2, err := client.Get(email).Result()
	if err == redis.Nil {
		response = map[string]interface{}{"refresh": "1"}
		util.RespondJSON(w, http.StatusOK, response)

	} else if err != nil {
		response = map[string]interface{}{"error": err.Error()}
		util.RespondJSON(w, http.StatusInternalServerError, response)
	} else {
		response = map[string]interface{}{"refresh": val2}
		util.RespondJSON(w, http.StatusOK, response)
	}

}
