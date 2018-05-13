package createList

import ("fmt"
        "net/http"
        "encoding/json"
        "./server"
       )

var collection=main.collection;

func createList(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
    var data=map[string]interface{}{
        
    };
	err := decoder.Decode(&data)
    if err != nil {
		panic(err)
	}
    var output=map[string]interface{}{
    };
    
    if val, ok := data["Name"]; ok {
        fmt.Println(val);
        fmt.Println(ok);
        fmt.Println(data["Name"])
        key :=data["Name"].(string);
        fmt.Println(key);
        if obj,exists := collection[key]; !exists {
            fmt.Println(obj);
            collection[string(key)]=map[string]interface{}{
            }
            output=collection;
        } else {
            output["data"]="List Existing";
        }
    } else {
        output["data"]="Please specify name of the list";
    }
    final , err := json.Marshal(&output)
    if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
    w.Header().Set("content-type", "application/json")
	w.Write(final)
}