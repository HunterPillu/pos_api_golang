package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/mingrammer/go-todo-rest-api-example/app/model"
)

//Get all Items
func GetAllHoldCartItems(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	data := []model.HoldCart{}
	db.Find(&data)
	respondJSON(w, http.StatusOK, data)
}

//Get item as per user
func GetHoldCartItemsByUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	data := []model.HoldCart{}
	vars := mux.Vars(r)
	uid := vars["uid"]
	if err := db.Where("uid=?", uid).Find(&data).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, data)
}

func CreateHoldCartItem(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	data := model.HoldCart{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&data).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, data)
}

func GetHoldCartItem(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	uid := vars["uId"]
	data := getHoldCartOr404(db, uid, w, r)
	if data == nil {
		return
	}
	respondJSON(w, http.StatusOK, data)
}

func UpdateHoldCartItem(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	uid := vars["uId"]
	data := getHoldCartOr404(db, uid, w, r)
	if data == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&data).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, data)
}

func DeleteHoldCartItem(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	uid := vars["uId"]
	data := getHoldCartOr404(db, uid, w, r)
	if data == nil {
		return
	}
	if err := db.Delete(&data).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

// getHoldCartOr404 gets a project instance if exists, or respond the 404 error otherwise
func getHoldCartOr404(db *gorm.DB, uid string, w http.ResponseWriter, r *http.Request) *model.HoldCart {
	data := model.HoldCart{}
	if err := db.Where("u_id=?", uid).First(&data).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &data
}
