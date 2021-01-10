package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/mingrammer/go-todo-rest-api-example/app/model"
)

func GetAllCategories(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	data := []model.Category{}
	db.Find(&data)
	respondJSON(w, http.StatusOK, data)
}

func GetCategories(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	data := []model.Category{}
	vars := mux.Vars(r)
	uid := vars["parent_id"]
	if err := db.Where("parent_id=?", uid).Find(&data).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, data)
}

func GetCategoriesByUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	data := []model.Category{}
	vars := mux.Vars(r)
	pid := vars["parent_id"]
	uid := vars["uid"]

	if err := db.Where("parent_id=? AND uid=?", pid, uid).Find(&data).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, data)
}

func CreateCategory(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	data := model.Category{}

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

func GetCategory(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	uid := vars["category_id"]
	data := getCategoryOr404(db, uid, w, r)
	if data == nil {
		return
	}
	respondJSON(w, http.StatusOK, data)
}

func UpdateCategory(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	uid := vars["category_id"]
	data := getCategoryOr404(db, uid, w, r)
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

func DeleteCategory(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	uid := vars["category_id"]
	data := getCustomerOr404(db, uid, w, r)
	if data == nil {
		return
	}
	if err := db.Delete(&data).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

// getCustomerOr404 gets a project instance if exists, or respond the 404 error otherwise
func getCategoryOr404(db *gorm.DB, uid string, w http.ResponseWriter, r *http.Request) *model.Category {
	data := model.Category{}
	if err := db.Where("cid=?", uid).First(&data).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &data
}
