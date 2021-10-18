package main

import (
  "github.com/gorilla/mux"
  "net/http"
  "encoding/json"
  "log"
  "ioutil"
)

type Articles struct{
  ID string        `json:"id"`
  Title string     `json:"title"`
  Author string    `json:"author"`
  Link string      `json:"link"`
}

func main(){
  r := mux.NewRouter().StrictSlash(true)
  r.HandleFunc("/article", articles)
  r.HandleFunc("/article/{id}", returnSingleArticle)
  r.HandleFunc("/article", createNewArticle).Methods("POST")
  r.HandleFunc("/article/{id}", DeleteArticle).Methods("DELETE")
  log.Fatal(http.ListenAndServe(":8080", r))
}

func articles(w http.ResponseWriter, r *http.Request){
  var articles Articles
  json.NewEncoder(w).Encode(articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request){
  vars := mux.Vars(r)
  key := vars["id"]

  for _, article := range Articles{
    if article.ID == key{
      json.NewEncoder(w).Encode(article)
    }
  }
}

func createNewArticle(w http.ResponseWriter, r *http.Request){
  reqbody, _ := ioutil.ReadAll(r.Body)
  var article Articles
  json.Unmarshal(reqbody, &article)

  json.NewEncoder(w).Encode(article)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request){
  vars := mux.Vars(r)
  id := vars["id"]

  for index, article := range Articles{
    if article.ID == id{
      Articles = append(Articles[:index], Articles[index+1:]...)
    }
  }
}
