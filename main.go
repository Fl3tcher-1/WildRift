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

// func topLane(input string) {

// 	akali := Champion{s
// 		"Akali",
// 		[]string{"Twisted Fate", "Irelia", "Yasuo"},
// 		[]string{"Fiora", "Jax", "Nasus", "Galio", "Garen", "Diana", "Renekton"},
// 		//  "Roaming & executing enemies",
// 		//  "Mid Game",
// 	}
// 	camille := Champion{
// 		"Camille",
// 		[]string{"Yasuo", "Jarvan", "Jax"},
// 		[]string{"Fiora", "Jax", "Nasus", "Teemo", "Garen", "Renekton", "Gragas"},
// 		//  "Roaming, ganking and isolating enemies",
// 		//  "Mid Game",
// 	}
// 	darius := Champion{
// 		"Darius",
// 		[]string{"Nasus", "Garen", "Yasuo"},
// 		[]string{"Kennen", "Vayne", "Jax", "Teemo", "Zed", "Renekton"},
// 		//  "Roaming & executing enemies",
// 		//  "Mid Game",
// 	}
// 	diana := Champion{
// 		"Diana",
// 		[]string{"Zed", "Fizz", "Katarina"},
// 		[]string{"Renekton", "Ziggs", "Irelia", "Galio"},
// 		//  "Roaming & executing enemies",
// 		//  "Mid Game",
// 	}
// 	mundo := Champion{
// 		"Dr Mundo",
// 		[]string{"Lee Sin", "Singed", "Jax"},
// 		[]string{"Darius", "Shyvana", "Olaf", "Fiora", "Vayne", "Mundo"},
// 		//  "Roaming & executing enemies",
// 		//  "Mid Game",
// 	}
// 	fiora := Champion{
// 		"Fiora",
// 		[]string{"Garen", "Akali", "Jax"},
// 		[]string{"Garen", "Malphite", "Pantheon", "Darius", "Singed", "Teemo", "Renekton", "Kayle"},
// 		//  "Roaming & executing enemies",
// 		//  "Mid Game",
// 	}

// 	garen := Champion{
// 		"Garen",
// 		[]string{"Nasus", "Jax", "Yasuo"},
// 		[]string{"Fiora", "Teemo", "Darius", "Pantheon", "Kayle"},
// 		//  "Roaming & executing enemies",
// 		//  "Mid Game",
// 	}
// 	gragas := Champion{
// 		"Gragas",
// 		[]string{"Camille", "Evelynn", "Xin Zhao"},
// 		[]string{"Fiora", "Yasuo", "Ahri", "Fizz", "Evelynn", "Akali", "Diana", "Ziggs", "Lee Sin", "Jax"},
// 		//  "Roaming & executing enemies",
// 		//  "Mid Game",
// 	}
// 	graves := Champion{
// 		"Graves",
// 		[]string{"Olaf", "Fiora", "Rengar", "Jarvan"},
// 		[]string{"Evelynn", "Lee Sin", "Rammus"},
// 		//  "Roaming & executing enemies",
// 		//  "Mid Game",
// 	}
// 	irelia := Champion{
// 		"Irelia",
// 		[]string{"Mundo", "Malphite", "Diana", "Jayce"},
// 		[]string{"Fizz", "Zed", "Galio", "Akali", "Yasuo", "Teemo", "Ziggs", "Seraphine", "Jaxx", "Akali"},
// 		//  "Roaming & executing enemies",
// 		//  "Mid Game",
// 	}
// 	jarvan := Champion{
// 		"Jarvan IV",
// 		[]string{"Olaf", "Evelynn", "Wukong"},
// 		[]string{"Graves", "Vi", "Lee Sin", "Xin Zhao", "Khazix", "Rammus"},
// 		//  "Roaming & executing enemies",
// 		//  "Mid Game",
// 	}
// 	jax := Champion{
// 		"Jax",
// 		[]string{"Tryndamere", "Xin Zhao", "Yasuo"},
// 		[]string{"Fiora", "Darius", "Garen", "Mundo", "Malphite", "Pantheon", "Renekton", "Singed"},
// 		//  "Roaming & executing enemies",
// 		//  "Mid Game",
// 	}
// 	jayce := Champion{
// 		"Jayce",
// 		[]string{"Fiora", "Kennen", "Darius"},
// 		[]string{"Irelia", "Nasus", "Wukong", "Galio", "Garen", "Diana", "Renekton", "Sett"},
// 		//  "Roaming & executing enemies",
// 		//  "Mid Game",
// 	}
// 	kayle := Champion{
// 		"Kayle",
// 		[]string{"Garen", "Fiora", "Mundo"},
// 		[]string{"Tryndamere", "Irelia", "Riven", "Renekton", "Sett"},
// 		//  "Roaming & executing enemies",
// 		//  "Mid Game",
// 	}
// 	kennen := Champion{
// 		"Kennen",
// 		[]string{"Jax", "Camille", "Teemo"},
// 		[]string{"Fiora", "Tryndamere", "Jarvan", "Rengar", "Nasus", "Mundo", "Garen"},
// 		//  "Roaming & executing enemies",
// 		//  "Mid Game",
// 	}
// 	malphite := Champion{
// 		"Malphite",
// 		[]string{"Tryndamere", "Fiora", "Jax"},
// 		[]string{"Nasus", "Akali", "Olaf", "Singed", "Garen", "Darius"},
// 		//  "Roaming & executing enemies",
// 		//  "Mid Game",
// 	}
// 	nasus := Champion{
// 		"Nasus",
// 		[]string{"Fiora", "Malphite", "Akali"},
// 		[]string{"Darius", "Singed", "Wukong", "Garen", "Fiora", "Vayne", "Kennen", "Olaf"},
// 		//  "Roaming & executing enemies",
// 		//  "Mid Game",
// 	}
// 	olaf := Champion{
// 		"Olaf",
// 		[]string{"Nasus", "Shyvana", "Malphite"},
// 		[]string{"Evelynn", "Yi", "Khazix", "Gragas", "Lee Sin", "Jax", "Fiora", "Jarvan"},
// 		//  "Roaming & executing enemies",
// 		//  "Mid Game",
// 	}
// 	pantheon := Champion{
// 		"Pantheon",
// 		[]string{"Yasuo", "Camille", "Fiora"},
// 		[]string{"Wukong", "Teemo", "Malphite", "Singed", "Rengar", "Akali", "Renekton"},
// 		//  "Roaming & executing enemies",
// 		//  "Mid Game",
// 	}
// 	riven := Champion{
// 		"Riven",
// 		[]string{"Kayle", "Irelia", "Yasuo", "Pantheon", "Wukong", "Jax"},
// 		[]string{"Renekton", "Teemo", "Fiora", "Nasus", "Garen", "Gragas", "Wukong"},
// 		//  "Roaming & executing enemies",
// 		//  "Mid Game",
// 	}
// 	renekton := Champion{
// 		"Renekton",
// 		[]string{"Riven", "Camille", "Yasuo"},
// 		[]string{"Garen", "Teemo", "Vayne", "Kennen", "Gragas", "Rengar", "Malphite"},
// 		//  "Roaming & executing enemies",
// 		//  "Mid Game",
// 	}

// 	sett := Champion{

// 		"Sett",
// 		[]string{"Tanks", "Irelia", "Yasuo", "Kayle"},
// 		[]string{"Jayce", "Darius", "Kennen", "Riven 50/50"},
// 		//  "Roaming & executing enemies",
// 		//  "Mid Game",
// 	}
// 	singed := Champion{
// 		"Singed",
// 		[]string{"Nasus", "Malphite", "Jax"},
// 		[]string{"Rengar", "Darius", "Vayne", "Camille", "Kennen", "Garen", "Teemo", "Mundo"},
// 		//  "Roaming & executing enemies",
// 		//  "Mid Game",
// 	}
// 	teemo := Champion{
// 		"Teemo",
// 		[]string{"Yi", "Garen", "Tryndamere"},
// 		[]string{"Pantheon", "Vayne", "Wukong", "Malphite", "Mundo", "Kennen", "Yasuo", "Akali"},
// 		//  "Roaming & executing enemies",
// 		//  "Mid Game",
// 	}
// 	tryndamere := Champion{
// 		"Tryndamere",
// 		[]string{"Darius", "Garen", "Nasus"},
// 		[]string{"Singed", "Malphite", "Wukong", "Jax", "Renekton", "Fiora"},
// 		//  "Roaming & executing enemies",
// 		//  "Mid Game",
// 	}
// 	vayne := Champion{
// 		"Vayne",
// 		[]string{"Tanks", "Ezreal", "Nasus", "Jhin"},
// 		[]string{"Draven", "Varus", "Caitlyn", "Miss Fortune", "Lee Sin", "Jinx"},
// 		//  "Roaming & executing enemies",
// 		//  "Mid Game",
// 	}
// 	wukong := Champion{
// 		"Wukong",
// 		[]string{"Jax", "Xin Zhao", "Yi", "Fiora"},
// 		[]string{"Garen", "Darius", "Lee Sin", "Singed", "Mundo", "Shyvana", "Lee Sin"},
// 		//  "Roaming & executing enemies",
// 		//  "Mid Game",
// 	}
// 	yasuo := Champion{
// 		"Yasuo",
// 		[]string{"Ahri", "Twisted Fate", "Orianna"},
// 		[]string{"Renekton", "Fizz", "Darius", "Annie", "Garen", "Malphite", "Pantheon"},
// 		//  "Roaming & executing enemies",
// 		//  "Mid Game",
// 	}

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

	file, err := os.OpenFile("champions.txt", os.O_RDWR, 0644)
	if err != nil {
		CheckErr(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	type finalChamps struct {
		champ database.Champion
	}

	var Data []string
	counter := 0
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, ",")
		Data = append(Data, champ(items))
		counter++
	}

	var name = ""
	var role = ""
	var strong = ""
	var weak = ""

	for champNumber, _ := range Data {
		for i, v := range strings.Split(Data[champNumber], "\n") {
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

func champ(s []string) string {
	return strings.Join(s, "\n")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
		db, err := sql.Open("sqlite3", "./database/champion.db")
	if err != nil {
		CheckErr(err)
	}

	data := database.Connect(db)
	database.Connect(db)

	tpl := template.Must(template.ParseGlob("templates/*"))
	tpl.ExecuteTemplate(w, "home.html", data.GetChamp())
}

func Handler(w http.ResponseWriter, r *http.Request) {
	//pathLink refers to the location where a file is stored e.g. "./images/ahri.jpg"
	pathLink := "./images"
	pathLink += r.URL.Path
	//name is used in order not to manually have to write the endpoint--- instead parsing it directly from url path as it will refer to an image to be served
	name:= r.URL.Path
	pathLink += ".jpg"

	switch r.URL.Path {
	case "/style":
		http.ServeFile(w, r, "./templates/style.css")
	case "/home":
		homeHandler(w, r)
	case name:
		http.ServeFile(w,r, pathLink)
	}
}
