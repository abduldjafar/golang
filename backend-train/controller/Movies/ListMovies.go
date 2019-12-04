package Movies

import (
	"backend-code/model"
	"backend-code/util"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func ListMovies(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	user := model.Accounts{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		util.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	client2 := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       1,  // use default DB
	})

	var countx int
	vars := mux.Vars(r)
	countz := vars["count"]
	counts, _ := strconv.Atoi(countz)
	movies := []model.Movies{}
	var response interface{}
	currentTime := time.Now()
	date := currentTime.Format("01-02-2006")

	val2, err := client.Get(user.Email + "_" + date).Result()
	value, err2 := client2.Get(user.Email).Result()
	ref, _ := strconv.Atoi(value)
	if err == redis.Nil || (ref == 1 && err == redis.Nil) || (err != redis.Nil && err2 == redis.Nil){
		db.Find(&movies).Count(&countx)

		min := 1
		max := countx
		value := 0
		rands := (rand.Intn(max-min) + min)

		if rands+4 > max {
			value = rands - 4
		} else {
			value = rands
		}

		if err := db.Offset(value).Limit(counts).Find(&movies).Error; err != nil {
			var resp = map[string]interface{}{"status": false, "message": "Something Wrong"}
			fmt.Println(resp)
		}

		movie, _ := json.Marshal(movies)
		result := string(movie)
		err := client.Set(user.Email+"_"+date, result, 0).Err()

		if err != nil {
			panic(err)
		}
		response = movies
	} else if err != nil {
		response = map[string]interface{}{"error": nil}
	} else {
		//fmt.Println("key2", val2)
		data := []model.Movies{}
		_ = json.Unmarshal([]byte(val2), &data)
		response = data
	}

	if err == redis.Nil  {
		err := client2.Set(user.Email, 1, 0).Err()
		if err != nil {
			panic(err)
		}
	} else  if  (err2 == redis.Nil && err != redis.Nil) {
		err := client2.Set(user.Email, 0, 1*time.Minute).Err()
		if err != nil {
			panic(err)
		}
	} else if err2 != nil {
		panic(err)
	} else {
		err := client2.Set(user.Email, 0, 1*time.Minute).Err()
		if err != nil {
			panic(err)
		}

	}

	util.RespondJSON(w, http.StatusOK, response)

}
