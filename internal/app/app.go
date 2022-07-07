package app

import (
	"database/sql"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type App struct {
	Router   *mux.Router
	Database *sql.DB
}

func (app *App) SetupRouter() {
	app.Router.
		Methods("GET").
		Path("/endpoint_get").
		HandlerFunc(app.getFunction)

	app.Router.
		Methods("POST").
		Path("/endpoint").
		HandlerFunc(app.postFunction)
}

func (app *App) getFunction(w http.ResponseWriter, r *http.Request) {
	
	var data DbData
	rows,err := app.Database.Query("SELECT * FROM `test` ")
	if err != nil {
		log.Fatal(err)
	}
	
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&data.ID, &data.UUID); 
		json.NewEncoder(w).Encode(data)
	
	}
}	

func (app *App) postFunction(w http.ResponseWriter, r *http.Request) {
	_, err := app.Database.Exec("INSERT INTO `test` (uuid) VALUES ('uuid')")
	if err != nil {
		log.Fatal("Database INSERT failed")
	}

	log.Println("You called a thing!")
	w.WriteHeader(http.StatusOK)
}