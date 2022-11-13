package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"wildrift/database"

	_ "github.com/mattn/go-sqlite3"
)

func CheckErr(e error) {
	fmt.Println(e)
}
func main() {

	db, err := sql.Open("sqlite3", "./database/champion.db")
	if err != nil {
		CheckErr(err)
	}

	data := database.Connect(db)
	database.Connect(db)

	mux := http.NewServeMux()
	mux.HandleFunc("/", Handler)
	// mux.HandleFunc("/style", css)

	file, err := os.OpenFile("champions.txt", os.O_RDWR, 0644) //open text file
	if err != nil {
		CheckErr(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file) //scan file

	type finalChamps struct { //create a struct referencing another struct that we will use as the layout
		champ database.Champion
	}
	var Data []string
	// counter := 0

	for scanner.Scan() { // scans text and splits on "," as it indicates a different column of databse to b populated
		line := scanner.Text()
		items := strings.Split(line, ",") // creates []string with newlines instead of ","
		Data = append(Data, champ(items)) //joins each entry on newline
		// counter++
	}
	//initialise variables to store date

	var name = ""
	var role = ""
	var strong = ""
	var weak = ""

	for champNumber, _ := range Data { //range over data and create a new variable in databse based on index [0]=name [1]= role [2]= strong [3]=weak, champ number refers to which champion
		for i, v := range strings.Split(Data[champNumber], "\n") { // i, v refers to the index and values inside each []champion
			switch i {
			case 0:
				name = v
			case 1:
				role = v
			case 2:
				strong = v
			case 3:
				weak = v
			}
		}
		data.CreateChampion(database.Champion{
			Name:     name,
			Role:     role,
			StrongVS: strong,
			WeakVS:   weak,
		})

	}
	fmt.Println("Starting server at port 8080:\n http://localhost:8080/home")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(500, "500 Internal server error:", err)
		fmt.Printf("main ListenAndServe error: %+v\n", err)
	}
}

//joins string
func champ(s []string) string {
	return strings.Join(s, "\n")
}

//handles home page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./database/champion.db")
	if err != nil {
		CheckErr(err)
	}

	data := database.Connect(db)
	database.Connect(db)

	tpl := template.Must(template.ParseGlob("templates/*"))
	tpl.ExecuteTemplate(w, "index.html", data.GetChamp())
}
func championHandler(w http.ResponseWriter, r *http.Request, name string) {
	db, err := sql.Open("sqlite3", "./database/champion.db")
	if err != nil {
		CheckErr(err)
	}
	path := strings.Trim(name, "/home/")

	data := database.Connect(db)
	database.Connect(db)

	for _, champion := range data.GetChamp() {
		if champion.Name == path {

			tpl := template.Must(template.ParseGlob("templates/*"))
			tpl.ExecuteTemplate(w, "champion.html", champion)
		}
	}

}

//handlers handles all the website endpoints
func Handler(w http.ResponseWriter, r *http.Request) {
	//pathLink refers to the location where a file is stored e.g. "./images/ahri.jpg"
	pathLink := "./images"
	pathLink += r.URL.Path
	//name is used in order not to manually have to write the endpoint--- instead parsing it directly from url path as it will refer to an image to be served
	name := r.URL.Path
	pathLink += ".jpg"

	// champPath:= r.URL.Path
	// fmt.Println(champPath)

	switch r.URL.Path {
	case "/style":
		http.ServeFile(w, r, "./templates/style.css")
	case "/images/favicon-32x32.png":
		http.ServeFile(w, r, "./images/favicon-32x32.png")
	case "/home":
		homeHandler(w, r)
	case name:
		if strings.Contains(name, "/home/") {
			championHandler(w, r, name)
		} else {
			http.ServeFile(w, r, pathLink)
		}
	}
}
