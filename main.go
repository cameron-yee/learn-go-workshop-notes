package main

import (
    "fmt"
    "net/http"
    "html/template"
    "log"
    "strings"
    "io/ioutil"
    "encoding/json"

//    //"reflect"
)

// STRUCT DEFS
// For html template variables
type Todo struct {
    Id string
    Title string
    Content string
}

// For html template variables
type PageVariables struct {
    PageTitle string
    PageTodos []Todo
}

type Planet struct {
    Name string
}

// For Star Wars API
type Person struct {
    Name string `json:"name"`
    HomeworldUrl string `json:"homeworld"`
    Homeworld *Planet
}
// END STRUCT DEFS


// Global var because we're not using a db
var todos []Todo

func main() {
    // GET/POST request routes
    http.HandleFunc("/", getHomepage)
    http.HandleFunc("/todos", getTodos)
    http.HandleFunc("/add-todo", addTodo)

    // Star Wars API endpoints
    http.HandleFunc("/people/", getPerson)

    fmt.Println("Server is up and running on localhost:8000")
    http.ListenAndServe(":8000", nil)
}

func getHomepage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func getTodos(w http.ResponseWriter, r *http.Request) {
    pageVariables := PageVariables{
        PageTitle: "Intro to Go",
        PageTodos: todos,
    }

    t, err := template.ParseFiles("todos.html")

    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        log.Print("Template parse error :( ", err)
        return
    }

    //Reassign err var b/c this won't be reached if err is hit before
    err = t.Execute(w, pageVariables)

    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        log.Print("Template parse error :( ", err)
        return
    }
}

func addTodo(w http.ResponseWriter, r *http.Request) {
    // Parse our incoming form (contains todo data)
    err := r.ParseForm()

    // Handle any errors
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        log.Print("Form request parse error :( ", err)
        return
    }

    // Create new todo (using our struct)
    todo := Todo{ // Type is inferred
        Title: r.FormValue("title"), // form name field name="title"
        Content: r.FormValue("content"),
    }

    // Append the new todo to our existing array
    todos = append(todos, todo)
    fmt.Println(todos) 
    // Watch it act weird and then fix it to redirect

    http.Redirect(w, r, "/todos", http.StatusSeeOther)
}

//Method is broken
// Struct method for Person struct
func (p *Person) getHomePlanet() {
    http.Get(p.HomeworldUrl)

    // Turn it into something we can use with Go
    json.Unmarshal(bytes, &p)
    fmt.Println(p)

    // Go make a separate request for the homeworld
    res, err := p.getHomePlanet()

    if err != nil {
        fmt.Println(err)
        return
    }

    // _ means we don't want to use the second return value err
    bytes, _ := ioutil.ReadAll(res.Body)
    err = json.Unmarshal(bytes, &p.Homeworld)
}

func getPerson(w http.ResponseWriter, r *http.Request) {
    // Create a person struct with the fields I care about
    var p Person

    // Grab the id from the url
    id := strings.TrimPrefix(r.URL.Path, "/people/") // Get rid of "people" and be left with id
    url := fmt.Sprintf("https://swapi.co/api/people/%s", id)

    res, err := http.Get(url)

    // Handle errors
    if err != nil {
        fmt.Fprintf(w, err.Error())
        http.Error(w, err.Error(), http.StatusBadRequest)
        log.Print("Error in API request", err)
        return
    }

    // Parse body of API response
    bytes, err := ioutil.ReadAll(res.Body)

    if err != nil {
        fmt.Fprintf(w, err.Error())
        http.Error(w, err.Error(), http.StatusBadRequest)
        log.Print("Error parsing API response", err)
        return
    }

    fmt.Println(bytes)
}

