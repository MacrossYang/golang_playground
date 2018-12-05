package main
import "fmt"
import "net/http"
import "encoding/json"
import "log"
import "github.com/gorilla/mux"
//import "github.com/go-martini/martini"

type Person struct {
        ID              string  `json:"id, omitempty"`
        Firstname       string  `json:"firstname, omitempty"`
        Lastname        string  `json:"lastname, omitempty"`
        Address         *Address`json:"address, omitempty"`
}

type Address struct {
        City    string  `json:"city, omitempty"`
        State   string  `json:"state, omitempty"`
}

var people []Person



func main(){
	fmt.Printf("hello, world.\n")
	//sub1()
	//sub2()
	sample()
	//fileserver()
}

func sample(){
        router := mux.NewRouter()
        people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Marshon", Address: &Address{ City: "City X", State: "State Y"}})
        people = append(people, Person{ID: "2", Firstname: "Jack", Lastname: "Marshon", Address: &Address{ City: "City A", State: "State B"}})

        router.HandleFunc("/people", GetPeople).Methods("GET")
        router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
        //router.HandleFunc("/people/{id}", NewPerson).Methods("POST")
        //router.HandleFunc("/people/{id}", DelPerson).Methods("DELETE")
        log.Fatal(http.ListenAndServe("192.168.11.49:8080", router))
}

func GetPeople(w http.ResponseWriter, r *http.Request){
        json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request){
        params := mux.Vars(r)
        for _ , value := range people{
                if value.ID == params["id"]{
                        json.NewEncoder(w).Encode(value)
                        return
                }
        }
        json.NewEncoder(w).Encode(&Person{})
}

/*
func sub1(){
	m := martini.Classic()
	m.Get("/", func() string {
		return "hello world"
	})
	fmt.Println("sub1 run")
	m.Run()
}

func sub2(){
	m2 := martini.Classic()
	m2.Get("/get/:key", func(r render.Render, pool *redis.Pool, params martini.Params) {
		key := params["keys"]
		c := pool.Get()
		defer c.Close()

		value, err :=redis.String(c.Do("Get", key))

		if err != nil {
			message := fmt.Sprintf("Could not Get %s", key)

			r.JSON(400, map[string]interface{}{
				"status" : "ERR",
				"message" : message})
		} else {
			r.JSON(200, map[string]interface{}{
				"status" : "OK",
				"value" : value})
		}
	})
	m2.Run()
}*/

func fileserver(){
	fmt.Println("file server run")
	http.ListenAndServe("192.168.11.49:8080",http.FileServer(http.Dir(".")))
}
