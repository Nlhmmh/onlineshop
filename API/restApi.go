package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// LoginUser is a struct for Login User
type LoginUser struct {
	ID             int       `json:"id"`
	FirstName      string    `json:"firstName"`
	LastName       string    `json:"lastName"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	IsAdmin        bool      `json:"isAdmin"`
	Address1       string    `json:"address1"`
	Address2       string    `json:"address2"`
	Address3       string    `json:"address3"`
	PhoneNo        string    `json:"phoneNo"`
	CreateUser     string    `json:"createUser"`
	CreateDatetime time.Time `json:"createDatetime"`
	UpdateUser     string    `json:"updateUser"`
	UpdateDatetime time.Time `json:"updateDatetime"`
}

/*
 * Main Method
 */
func main() {

	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/getAllLoginUser", getAllLoginUser).Methods("GET")
	router.HandleFunc("/getLoginUser/{id}", getLoginUser).Methods("GET")
	router.HandleFunc("/addLoginUser", addLoginUser).Methods("POST")
	router.HandleFunc("/updateLoginUser/{id}", updateLoginUser).Methods("PUT")
	router.HandleFunc("/deleteLoginUser/{id}", deleteLoginUser).Methods("DELETE")

	// Start server
	fmt.Println("# Server Starting at http://localhost:5000 !!!!")
	log.Fatal(http.ListenAndServe("127.0.0.1:5000", router))

}

/*
 * Open Database Connection
 */
func openConnection() *sql.DB {

	dbinfo := "postgresql://postgres:root@localhost/onlineshop?sslmode=disable"

	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		fmt.Println("# Failed to Open Connection to Postgres Database")
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("# Failed to Ping to Postgres Database")
		log.Fatal(err)
	}

	return db

}

/*
 * Get all LoginUser
 */
func getAllLoginUser(w http.ResponseWriter, r *http.Request) {

	db := openConnection()

	fmt.Println("# Querying all LoginUsers")
	rows, err := db.Query("SELECT * FROM loginuser")
	if err != nil {
		fmt.Println("Querying Failed!")
		log.Println(err)
		return
	}

	var loginUsers []LoginUser = []LoginUser{}
	for rows.Next() {
		var loginUser LoginUser
		err = rows.Scan(
			&loginUser.ID,
			&loginUser.FirstName,
			&loginUser.LastName,
			&loginUser.Email,
			&loginUser.Password,
			&loginUser.IsAdmin,
			&loginUser.Address1,
			&loginUser.Address2,
			&loginUser.Address3,
			&loginUser.PhoneNo,
			&loginUser.CreateUser,
			&loginUser.CreateDatetime,
			&loginUser.UpdateUser,
			&loginUser.UpdateDatetime,
		)
		if err != nil {
			fmt.Println("Querying Failed at LoginUser with ID : ", loginUser.ID)
			log.Println(err)
			return
		}
		loginUsers = append(loginUsers, loginUser)
		fmt.Println(loginUser)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(loginUsers)
	fmt.Println("# Success")

	defer rows.Close()
	defer db.Close()

}

/*
 * Get LoginUser with ID
 */
func getLoginUser(w http.ResponseWriter, r *http.Request) {

	db := openConnection()

	// get ID of LoginUser from the route parameter
	var idParam string = mux.Vars(r)["id"]
	fmt.Println("# Querying LoginUser with " + idParam)

	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	var loginUser LoginUser
	err = db.QueryRow("SELECT * FROM loginuser WHERE id = $1", &id).Scan(
		&loginUser.ID,
		&loginUser.FirstName,
		&loginUser.LastName,
		&loginUser.Email,
		&loginUser.Password,
		&loginUser.IsAdmin,
		&loginUser.Address1,
		&loginUser.Address2,
		&loginUser.Address3,
		&loginUser.PhoneNo,
		&loginUser.CreateUser,
		&loginUser.CreateDatetime,
		&loginUser.UpdateUser,
		&loginUser.UpdateDatetime,
	)
	if err != nil {
		log.Println(err)
		w.WriteHeader(404)
		w.Write([]byte("No LoginUser found with specified ID : " + idParam))
		return
	}
	fmt.Println(loginUser)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(loginUser)
	fmt.Println("# Success")

	defer db.Close()

}

/*
 * Add LoginUser with ID
 */
func addLoginUser(w http.ResponseWriter, r *http.Request) {

	db := openConnection()

	var loginUser LoginUser
	json.NewDecoder(r.Body).Decode(&loginUser)

	fmt.Println("# Adding LoginUser")
	stmt, err := db.Prepare(
		"INSERT INTO loginuser(" +
			"first_name," +
			" last_name," +
			" email," +
			" password," +
			" isAdmin," +
			" address1," +
			" address2," +
			" address3," +
			" phone_no," +
			" create_user," +
			" create_datetime," +
			" update_user," +
			" update_datetime)" +
			" VALUES ( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13 )")
	if err != nil {
		log.Println(err)
		w.WriteHeader(404)
		w.Write([]byte("Error in SQL Statement :: " + err.Error()))
		return
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		loginUser.FirstName,
		loginUser.LastName,
		loginUser.Email,
		loginUser.Password,
		loginUser.IsAdmin,
		loginUser.Address1,
		loginUser.Address2,
		loginUser.Address3,
		loginUser.PhoneNo,
		loginUser.CreateUser,
		loginUser.CreateDatetime,
		loginUser.UpdateUser,
		loginUser.UpdateDatetime,
	)
	if err != nil {
		log.Println(err)
		w.WriteHeader(404)
		w.Write([]byte("Error in SQL Statement Execution :: " + err.Error()))
		return
	}

	var status int = 0
	rowsAffected, err := res.RowsAffected()
	if rowsAffected > 0 {
		status = 1
		fmt.Println("Successfully Inserting LoginUser!")
	} else {
		status = 0
		fmt.Println("Inserting LoginUser Failed!")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(strconv.Itoa(status)))
	fmt.Println("# Success")

	defer db.Close()

}

/*
 * Update LoginUser with ID
 */
func updateLoginUser(w http.ResponseWriter, r *http.Request) {

	db := openConnection()

	// get ID of LoginUser from the route parameter
	var idParam string = mux.Vars(r)["id"]
	fmt.Println("# Querying LoginUser with " + idParam)

	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	var loginUser LoginUser
	json.NewDecoder(r.Body).Decode(&loginUser)

	fmt.Println("# Updating LoginUser")
	stmt, err := db.Prepare(
		"UPDATE loginuser SET " +
			"first_name = $1," +
			" last_name = $2," +
			" email = $3," +
			" password = $4," +
			" isAdmin = $5," +
			" address1 = $6," +
			" address2 = $7," +
			" address3 = $8," +
			" phone_no = $9," +
			" create_user = $10," +
			" create_datetime = $11," +
			" update_user = $12," +
			" update_datetime = $13" +
			" WHERE id = $14")
	if err != nil {
		log.Println(err)
		w.WriteHeader(404)
		w.Write([]byte("Error in SQL Statement :: " + err.Error()))
		return
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		loginUser.FirstName,
		loginUser.LastName,
		loginUser.Email,
		loginUser.Password,
		loginUser.IsAdmin,
		loginUser.Address1,
		loginUser.Address2,
		loginUser.Address3,
		loginUser.PhoneNo,
		loginUser.CreateUser,
		loginUser.CreateDatetime,
		loginUser.UpdateUser,
		loginUser.UpdateDatetime,
		id,
	)
	if err != nil {
		log.Println(err)
		w.WriteHeader(404)
		w.Write([]byte("Error in SQL Statement Execution :: " + err.Error()))
		return
	}

	var status int = 0
	rowsAffected, err := res.RowsAffected()
	w.Header().Set("Content-Type", "application/json")
	if rowsAffected > 0 {
		status = 1
		fmt.Println("Successfully Updating LoginUser with ID : ", id)
		w.Write([]byte(strconv.Itoa(status)))
		fmt.Println("# Success")
	} else {
		status = 0
		fmt.Println("Updating LoginUser Failed! ID : ", id)
		w.Write([]byte(strconv.Itoa(status)))
		fmt.Println("# Failed")
	}

	defer db.Close()

}

/*
 * Delete LoginUser with ID
 */
func deleteLoginUser(w http.ResponseWriter, r *http.Request) {

	db := openConnection()

	// get ID of LoginUser from the route parameter
	var idParam string = mux.Vars(r)["id"]
	fmt.Println("# Querying LoginUser with " + idParam)

	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}
	fmt.Println("# Updating LoginUser")
	stmt, err := db.Prepare("DELETE FROM loginuser WHERE id = $1")
	if err != nil {
		log.Println(err)
		w.WriteHeader(404)
		w.Write([]byte("Error in SQL Statement :: " + err.Error()))
		return
	}
	defer stmt.Close()
	res, err := stmt.Exec(id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(404)
		w.Write([]byte("Error in SQL Statement Execution :: " + err.Error()))
		return
	}

	var status int = 0
	rowsAffected, err := res.RowsAffected()
	w.Header().Set("Content-Type", "application/json")
	if rowsAffected > 0 {
		status = 1
		fmt.Println("Successfully Deleting LoginUser with ID : ", id)
		w.Write([]byte(strconv.Itoa(status)))
		fmt.Println("# Success")
	} else {
		status = 0
		fmt.Println("Deleting LoginUser Failed! ID : ", id)
		w.Write([]byte(strconv.Itoa(status)))
		fmt.Println("# Failed")
	}

	defer db.Close()

}
