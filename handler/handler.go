package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	// . "github.com/connection"
	"github.com/connection"
	_ "github.com/connection"
	_ "github.com/encryption"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/model"
	_ "github.com/model"
	"github.com/utility"
	_ "github.com/utility"
	// _ "gopkg.in/connection"
)

//model
type Users model.Users
type Address model.Address

// Registartion
func Registration(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if request.Method == "POST" {

		var (
			UserMap map[string]string
		)

		db := connection.DbConn()

		UserMap = make(map[string]string)
		var u Users

		decoder := json.NewDecoder(request.Body)
		err := decoder.Decode(&u)

		if err != nil {
			panic(err)
		}

		email := u.Email

		password := utility.CreateHash(u.Password)
		external_id := ""

		insForm, err := db.Prepare("INSERT INTO users (email, password,external_id) VALUES(?,?,?)")
		log.Println("ghfrdes")
		if err != nil {
			UserMap["msg"] = "Email Id Already exist."
			json.NewEncoder(w).Encode(UserMap)
		}
		res, errs := insForm.Exec(email, password, external_id)
		if errs != nil {
			UserMap["msg"] = "Email Id Already exist."
			json.NewEncoder(w).Encode(UserMap)
		} else {

			id, _ := res.LastInsertId()

			UserMap["UserID"] = strconv.FormatInt(id, 10)
			UserMap["Email"] = email
			json.NewEncoder(w).Encode(UserMap)
		}
	}
}

// Login
func Login(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if request.Method == "POST" {
		var (
			// users   users
			UserMap map[string]string
		)
		db := connection.DbConn()
		UserMap = make(map[string]string)
		var u Users
		decoder := json.NewDecoder(request.Body)
		err := decoder.Decode(&u)

		if err != nil {
			panic(err)
		}

		email := u.Email
		password := utility.CreateHash(u.Password)

		rows, err := db.Query("select user_id, email, password from users where email = ? and password = ?", email, password)

		if err != nil {
			UserMap["msg"] = "Something wrong"
			json.NewEncoder(w).Encode(UserMap)
		} else {

			var flag = 0
			for rows.Next() {

				flag = 1
				err = rows.Scan(&u.UserID, &u.Email, &u.Password)
				if err != nil {
					UserMap["msg"] = "Something wrong"
					json.NewEncoder(w).Encode(UserMap)
				} else {
					UserMap["UserID"] = strconv.Itoa(u.UserID)
					UserMap["Email"] = u.Email

				}

			}
			defer rows.Close()
			if flag == 0 {
				UserMap["msg"] = "invalid email or password"
				json.NewEncoder(w).Encode(UserMap)
			} else {
				json.NewEncoder(w).Encode(UserMap)
			}

		}

	}
}

// Show Addresses
func ShowAddress(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if request.Method == "GET" {
		var (
			UserMap map[string]string
		)
		db := connection.DbConn()
		UserMap = make(map[string]string)
		var a Address

		params := mux.Vars(request)
		user_id := params["userid"]
		rows, err := db.Query("select user_id, email, address_id, contact_number, is_active from addresses where is_deleted = ? and user_id=?", 0, user_id)
		if err != nil {
			panic(err.Error())

		}

		var arr = []map[string]string{}

		for rows.Next() {
			err = rows.Scan(&a.UserID, &a.Email, &a.AddressID, &a.ContactNumber, &a.IsActive)

			if err != nil {
				UserMap["msg"] = "something wrong "
				json.NewEncoder(w).Encode(UserMap)
			}
			UserMap = make(map[string]string)
			UserMap["user_id"] = strconv.Itoa(a.UserID)
			UserMap["email"] = a.Email
			UserMap["address_id"] = strconv.Itoa(a.AddressID)
			UserMap["contact_number"] = a.ContactNumber
			UserMap["is_active"] = strconv.Itoa(a.IsActive)

			arr = append(arr, UserMap)

		}

		defer rows.Close()
		json.NewEncoder(w).Encode(arr)

	}
}

// Add Address
func AddAddress(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if request.Method == "POST" {
		var (
			UserMap map[string]string
		)
		db := connection.DbConn()
		UserMap = make(map[string]string)
		var a Address
		decoder := json.NewDecoder(request.Body)
		err := decoder.Decode(&a)

		if err != nil {
			UserMap["msg"] = "Something Wrong"
			json.NewEncoder(w).Encode(UserMap)
		} else {
			name := a.Name
			email := a.Email
			contact_number := a.ContactNumber
			is_active := a.IsActive
			user_id := a.UserID

			insForm, err := db.Prepare("INSERT INTO addresses (name,email, contact_number,is_active,user_id) VALUES(?,?,?,?,?)")
			if err != nil {
				panic(err.Error())
			}
			res, err := insForm.Exec(name, email, contact_number, is_active, user_id)
			id, _ := res.LastInsertId()

			UserMap["address_id"] = strconv.FormatInt(id, 10)
			UserMap["user_id"] = strconv.Itoa(user_id)
			UserMap["is_active"] = strconv.Itoa(is_active)
			UserMap["email"] = email
			UserMap["name"] = name
			UserMap["contact_number"] = contact_number

			json.NewEncoder(w).Encode(UserMap)
		}
	}
}

// Update Address
func UpdateAddress(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if request.Method == "POST" {
		var (
			UserMap map[string]string
		)
		db := connection.DbConn()
		UserMap = make(map[string]string)
		var a Address
		decoder := json.NewDecoder(request.Body)
		err := decoder.Decode(&a)

		if err != nil {
			UserMap["msg"] = "Something Wrong"
			json.NewEncoder(w).Encode(UserMap)
		} else {
			params := mux.Vars(request)
			name := a.Name
			email := a.Email
			contact_number := a.ContactNumber
			is_active := a.IsActive
			user_id := a.UserID
			address_id := params["addressid"]
			insForm, err := db.Prepare("UPDATE addresses set name=?, email=?, contact_number=?, is_active=?,user_id=? where address_id=?")
			if err != nil {
				panic(err.Error())
			}
			_, err = insForm.Exec(name, email, contact_number, is_active, user_id, address_id)

			UserMap["address_id"] = address_id //strconv.FormatInt(address_id, 10)
			UserMap["user_id"] = strconv.Itoa(user_id)
			UserMap["is_active"] = strconv.Itoa(is_active)
			UserMap["email"] = email
			UserMap["name"] = name
			UserMap["contact_number"] = contact_number

			json.NewEncoder(w).Encode(UserMap)
		}
	}
}

// Delete Address
func DeleteAddress(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if request.Method == "GET" {
		var (
			UserMap map[string]string
		)
		db := connection.DbConn()
		UserMap = make(map[string]string)

		params := mux.Vars(request)
		address_id := params["addressid"]

		insForm, err := db.Prepare("UPDATE addresses set is_deleted=1 where address_id=?")
		if err != nil {
			UserMap["success"] = "something wrong"
			json.NewEncoder(w).Encode(UserMap)
		} else {
			_, err = insForm.Exec(address_id)
			UserMap["success"] = "Delete Successfully."
			json.NewEncoder(w).Encode(UserMap)
		}

	}
}
