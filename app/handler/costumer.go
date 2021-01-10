package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/mingrammer/go-todo-rest-api-example/app/model"
	_ "github.com/mingrammer/go-todo-rest-api-example/app/model"
)

func GetAllCustomers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	data := []model.Customer{}
	db.Find(&data)
	respondJSON(w, http.StatusOK, data)
}

func GetCustomersByUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	data := []model.Customer{}
	vars := mux.Vars(r)
	uid := vars["uid"]
	if err := db.Where("uid=?", uid).Find(&data).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, data)
}

func CreateCustomer(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	data := model.Customer{}

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

func GetCustomer(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	uid := vars["customer_id"]
	data := getCustomerOr404(db, uid, w, r)
	if data == nil {
		return
	}
	respondJSON(w, http.StatusOK, data)
}

func UpdateCustomer(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	uid := vars["customer_id"]
	data := getCustomerOr404(db, uid, w, r)
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

func DeleteCustomer(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	uid := vars["customer_id"]
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
func getCustomerOr404(db *gorm.DB, uid string, w http.ResponseWriter, r *http.Request) *model.Customer {
	data := model.Customer{}
	if err := db.Where("customer_id=?", uid).First(&data).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &data
}
