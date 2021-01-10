package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/mingrammer/go-todo-rest-api-example/app/model"
)

//Get all Items
func GetAllOptionValues(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	data := []model.OptionValue{}
	db.Find(&data)
	respondJSON(w, http.StatusOK, data)
}

//Get item as per user
func GetOptionValuesByUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	data := []model.Product{}
	vars := mux.Vars(r)
	uid := vars["uid"]
	if err := db.Where("uid=?", uid).Find(&data).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, data)
}

func CreateOptionValue(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	data := model.OptionValue{}

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

func GetOptionValue(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	uid := vars["optionValueId"]
	data := getOptionValueOr404(db, uid, w, r)
	if data == nil {
		return
	}
	respondJSON(w, http.StatusOK, data)
}

func UpdateOptionValue(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	uid := vars["pId"]
	data := getOptionValueOr404(db, uid, w, r)
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

func DeleteOptionValue(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	uid := vars["optionValueId"]
	data := getOptionValueOr404(db, uid, w, r)
	if data == nil {
		return
	}
	if err := db.Delete(&data).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

// getOptionValueOr404 gets a project instance if exists, or respond the 404 error otherwise
func getOptionValueOr404(db *gorm.DB, uid string, w http.ResponseWriter, r *http.Request) *model.OptionValue {
	data := model.OptionValue{}
	if err := db.Where("option_value_id=?", uid).First(&data).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &data
}
