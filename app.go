package main

import (
    "encoding/json"
    "fmt"
    "log"
    "io/ioutil"
    "net/http"
    "github.com/gorilla/mux"
    "strconv"
)

// Article - Our struct for all articles
type Article struct {
    Id      string `json:"Id"`
    Title   string `json:"Title"`
    Desc    string `json:"desc"`
    Content string `json:"content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: homePage")
    
    fmt.Fprintf(w, "Welcome to the HomePage!")
}

func getAllArticles(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type","application/json")
    
    fmt.Println("Endpoint Hit: getAllArticles")
    
    json.NewEncoder(w).Encode(Articles)
}

func getOneArticle(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type","application/json")
    
    fmt.Println("Endpoint Hit: getOneArticle")
    
    vars := mux.Vars(r)
    key := vars["id"]

    fmt.Println(strconv.Atoi("5"))

    for _, article := range Articles {
        
        if article.Id == key {
            json.NewEncoder(w).Encode(article)
        }
    }
}

func createOneArticle(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type","application/json")
    
    fmt.Println("Endpoint Hit: createOneArticle")
    
    // get the body of our POST request
    // unmarshal this into a new Article struct
    // append this to our Articles array.    
    reqBody, _ := ioutil.ReadAll(r.Body)
    var article Article 
    json.Unmarshal(reqBody, &article)
    
    // update our global Articles array to include new article
    Articles = append(Articles, article)

    json.NewEncoder(w).Encode(article)
}

func updateOneArticle(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type","application/json")
    
    fmt.Println("Endpoint Hit: updateOneArticle")

    // json.NewEncoder(w).Encode(article)
}

func deleteOneArticle(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: deleteOneArticle")
    
    vars := mux.Vars(r)
    id := vars["id"]

    for index, article := range Articles {
        if article.Id == id {
            Articles = append(Articles[:index], Articles[index+1:]...)
        }
    }
}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/articles", getAllArticles).Methods("GET")
    myRouter.HandleFunc("/articles/{id}", getOneArticle).Methods("GET")
    myRouter.HandleFunc("/articles", createOneArticle).Methods("POST")
    myRouter.HandleFunc("/articles/{id}", updateOneArticle).Methods("PUT")
    myRouter.HandleFunc("/articles/{id}", deleteOneArticle).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
    Articles = []Article{
        Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
    }
    handleRequests()
}