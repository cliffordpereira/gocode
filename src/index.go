package main

import ("fmt"
        "net/http"
        "io/ioutil"
        "encoding/json")

func main(){
    http.HandleFunc("/",handlerequest);
    http.HandleFunc("/createList",createList);
    http.ListenAndServe(":8080",nil)
};

var collection []taskList;

type taskList struct {
    task []string
}

type Message struct {
	Id   int `json:"id"`
	Name string `json:"name"`
}

//type Task struct {
//	Id   int `json:"id"`
//	Name string `json:"name"`
//}

var arr = make([]interface{}, 0);

func createList(w http.ResponseWriter , r *http.Request){
    b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
    
}

func handlerequest(w http.ResponseWriter , r *http.Request){
    b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
    
	// Unmarshal
	var msg Message
	err = json.Unmarshal(b,&msg)
    
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
    List=append(List,msg);
    fmt.Printf("%+v\n",List);
    

	output, err := json.Marshal(&msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
