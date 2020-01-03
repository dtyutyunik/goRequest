package main

import(
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
)

type Article struct {
	Id      string `json:"Id"`
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func returnAllArticles(w http.ResponseWriter, r *http.Request){

	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

// func postArticle(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprintf(w, "Article Posted")
// }

func createNewArticle(w http.ResponseWriter, r *http.Request) {
 
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
	
}

func deleteArticle(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range Articles {

		if article.Id == id{
		// updates our Articles array to remove the 
            Articles = append(Articles[:index], Articles[index+1:]...)
        }
	}

}


func returnSingleArticle(w http.ResponseWriter, r*http.Request){

		vars := mux.Vars(r)
		key := vars["id"]

		found:= 0

		for _, article := range Articles{
			if article.Id == key{
				json.NewEncoder(w).Encode(article)
				found++
			}
			
		}
		if(found==0){
			json.NewEncoder(w).Encode("Not Found")
		}
    
}

func handleRequests(){

	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true);

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", 
	returnAllArticles).Methods("GET")

	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")

	myRouter.HandleFunc("/article/{id}", returnSingleArticle).Methods("GET")

	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")

	
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main(){

	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
}
		
handleRequests()
}