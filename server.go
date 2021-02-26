package main
import (
    "fmt"
    "log"
    "net/http"
	"encoding/json"
	"context"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	

)
type  Article struct{
	Title string 
	Desc string  
	Content string 
}

type Articles []Article

func allArticles(w http.ResponseWriter, r*http.Request){
	articles:=Articles{
		Article{Title:"Test title" , Desc:"Test Description", Content:"Hello World"},
		Article{Title:"Test title1" , Desc:"Test Description1", Content:"Hello World1"},
	}
	fmt.Println("Endpoint Hit : All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}
 func homePage(w http.ResponseWriter, r *http.Request){
     fmt.Fprintf(w, "Homepage endpoint hit")
 }
func handleRequest(){
    http.HandleFunc("/",homePage)
	http.HandleFunc("/articles",allArticles)
    log.Fatal(http.ListenAndServe(":8090",nil))
 }

 func main(){
     handleRequest()
	 //......................
	 client,err :=mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	 if err!=nil{
		 log.Fatal(err)
	 }
	 ctx,_ :=context.WithTimeout(context.Background(),10*time.Second)
	 err=client.Connect(ctx)

	 if err!=nil{
		 log.Fatal(err)
	 }

	 dbs, err:=client.ListDatabaseNames(context.Background(),bson.D{})
	 if err!=nil{
		 log.Fatal(err)
	 }
	 for _,dbName:=range dbs{
		 fmt.Println(dbName)
	 }
 }