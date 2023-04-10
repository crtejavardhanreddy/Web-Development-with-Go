package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

type waiter struct {
	WaiterId       string `json:"WaiterId"`
	WaiterName     string `json:"WaiterName"`
	WaiterPhone    int    `json:"Waiter_Phone_Number"`
	RestaurantId   int    `json:"RestaurantId"`
	RestaurantName string `json:"RestaurantName"`
}

type Restaurant struct {
	Rid     int    `json:"Rid"`
	RName   string `json:"RName"`
	R_Ph_No int    `json:"R_Ph_No"`
}

type User struct {
	UserId   int    `json:"UserId"`
	UserName string `json:"UserName"`
	Ph_No    int    `json:"Ph_No"`
	Age      int    `json:"Age"`
	Gender   string `json:"Gender"`
}

func main() {
	fmt.Println("Connecting the database...")
	fmt.Println("Starting the server on 9080...")
	db, err = sql.Open("mysql", "root:Teja@7483@tcp(127.0.0.1:3306)/splitbill")
	if err != nil {
		fmt.Println("Error in connecting the database", err)
	}
	defer db.Close()
	http.HandleFunc("/waiter/", getDetails)
	http.HandleFunc("/waiters", getAllWaiters)
	http.HandleFunc("/restaurant", addRestaurant)
	http.HandleFunc("/delete/", deleteById)
	http.HandleFunc("/update/", updateById)
	http.HandleFunc("/restaurants", GetAllRestaurants)
	http.HandleFunc("/users", GetAllUsers)
	http.ListenAndServe(":9080", nil)
}

func getDetails(w http.ResponseWriter, r *http.Request) {
	idstr := r.URL.Path[len("/waiter/"):]
	// id, err := strconv.Atoi(idstr)
	if err != nil {
		fmt.Println("Error in the string convertion", err)

	}
	var W waiter
	rows := db.QueryRow("select waiter.wid,waiter.wname,waiter.w_phno, restaurant.rid ,restaurant.rname from waiter join restaurant on waiter.rname=restaurant.rname where wid=?", idstr)
	err = rows.Scan(&W.WaiterId, &W.WaiterName, &W.WaiterPhone, &W.RestaurantId, &W.RestaurantName)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(W)
	fmt.Printf("%+v", W)
}

func getAllWaiters(w http.ResponseWriter, r *http.Request) {
	Waiters := []waiter{}
	rows, err := db.Query("select waiter.wid,waiter.wname,waiter.w_phno, restaurant.rid ,restaurant.rname from waiter join restaurant where waiter.rname=restaurant.rname")
	if err != nil {
		fmt.Println("Query error", err)
	}
	for rows.Next() {
		var wai waiter
		err := rows.Scan(&wai.WaiterId, &wai.WaiterName, &wai.WaiterPhone, &wai.RestaurantId, &wai.RestaurantName)
		if err != nil {
			fmt.Println("Unable to read the data form database...")
		}
		Waiters = append(Waiters, wai)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Waiters)
	fmt.Printf("%+v\n", Waiters)
}

func addRestaurant(w http.ResponseWriter, r *http.Request) {
	var res Restaurant
	err := json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		fmt.Println("Unable to decode the input", err)
	}
	insert, err := db.Prepare("Insert Into restaurant(Rid,RName,R_Ph_No)values(?,?,?)")
	if err != nil {
		fmt.Println("Unable to insert the data", err)
	}
	_, err = insert.Exec(res.Rid, res.RName, res.R_Ph_No)
	if err != nil {
		fmt.Println("Unable to scan the data", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Data added successfully"))
}

func deleteById(w http.ResponseWriter, r *http.Request) {
	idstr := r.URL.Path[len("/delete/"):]
	id, err := strconv.Atoi(idstr)
	if err != nil {
		fmt.Println("Unable to convert into int", err)
	}
	db.Exec("DELETE FROM Restaurant WHERE Rid=?", id)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Data deleted successfully..."))
}

func updateById(w http.ResponseWriter, r *http.Request) {
	idstr := r.URL.Path[len("/update/"):]
	id, err := strconv.Atoi(idstr)
	if err != nil {
		fmt.Println("Unable to get the id...")
	}
	db.Query("UPDATE restaurant SET RName='Paradise' WHERE Rid=?", id)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Updated successfully..."))
}

func GetAllRestaurants(w http.ResponseWriter, r *http.Request) {
	Restaurants := []Restaurant{}
	rows, err := db.Query("Select * from restaurant")
	if err != nil {
		fmt.Println("Unable to fetch the data", err)
	}
	for rows.Next() {
		var res Restaurant
		err = rows.Scan(&res.Rid, &res.RName, &res.R_Ph_No)
		if err != nil {
			fmt.Println("Unable to scan the data", err)
		}

		Restaurants = append(Restaurants, res)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Restaurants)

}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{}

	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		fmt.Println("Unable to fetch data", err)
	}
	for rows.Next() {
		var user User
		err = rows.Scan(&user.UserId, &user.UserName, &user.Age, &user.Ph_No, &user.Gender)
		if err != nil {
			fmt.Println("Unable to fetch data", err)
		}
		users = append(users, user)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
