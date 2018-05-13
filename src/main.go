package main

import ("fmt"
        "net/http"
        "encoding/json"
        "./createList"
       )

func main(){
    http.HandleFunc("/createList",createList.createList());
    http.HandleFunc("/deleteList",deleteList);
    http.HandleFunc("/addTask",addTask);
    http.HandleFunc("/deleteTask",deleteTask);
    http.HandleFunc("/updateTask",updateTask);
    http.ListenAndServe(":8080",nil)
};    

var collection=map[string]interface{}{
};

func deleteList(w http.ResponseWriter, req *http.Request) {
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
            output["data"]="List Does Not Exist";
        } else {
            delete(collection,key)
            output=collection;
        }
    } else {
        output["data"]="Please specify name of the list to be deleted.";
    }
    final , err := json.Marshal(&output)
    if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
    w.Header().Set("content-type", "application/json")
	w.Write(final)  
}

func addTask(w http.ResponseWriter, req *http.Request){
	decoder := json.NewDecoder(req.Body)
    var data=map[string]interface{}{
    };
    
	err := decoder.Decode(&data)
    if err != nil {
		panic(err)
	}
    
    fmt.Println(data);
    
    var output=map[string]interface{}{
    };
    if val1 , ok := data["listName"]; ok {
        fmt.Println(val1);
        fmt.Println(ok);
        fmt.Println(data["listName"])
        key :=data["listName"].(string);
        fmt.Println(key);
        if obj,exists := collection[key]; !exists {
            fmt.Println(obj);
            output["data"]="List Does Not Exist";
        } else {
            taskName := data["task"].(map[string]interface{})["name"].(string);
            taskData := data["task"].(map[string]interface{})["data"].(string);
//            ["name"].(string);
//            taskData := data["listName"].(map[string]interface{})["task"].(map[string]interface{})["data"].(string);
            fmt.Println(taskName);
            fmt.Println(taskData);
//            fmt.Println(taskData);
            if obj,exists := collection[key].(map[string]interface{})[taskName]; !exists {
                fmt.Println(obj);
                collection[key].(map[string]interface{})[taskName]=taskData;
                output=collection;
            } else {
                output["data"]="Task already present";
            }
        }
    } else {
        output["data"]="Please specify the require details.";
    }
    final , err := json.Marshal(&output)
    if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
    w.Header().Set("content-type", "application/json")
	w.Write(final)  
}

func deleteTask(w http.ResponseWriter, req *http.Request){
	decoder := json.NewDecoder(req.Body)
    var data=map[string]interface{}{
    };
    
	err := decoder.Decode(&data)
    if err != nil {
		panic(err)
	}
    
    fmt.Println(data);
    
    var output=map[string]interface{}{
    };
    if val1 , ok := data["listName"]; ok {
        fmt.Println(val1);
        fmt.Println(ok);
        fmt.Println(data["listName"])
        key :=data["listName"].(string);
        fmt.Println(key);
        if obj,exists := collection[key]; !exists {
            fmt.Println(obj);
            output["data"]="List Does Not Exist";
        } else {
            taskName := data["taskName"].(string)
            fmt.Println(taskName);
            if obj,exists := collection[key].(map[string]interface{})[taskName]; exists {
                fmt.Println(obj);
                delete(collection[key].(map[string]interface{}),taskName);
                output=collection;
            } else {
                output["data"]="Task not present.";
            }
        }
    } else {
        output["data"]="Please specify the required details.";
    }
    final , err := json.Marshal(&output)
    if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
    w.Header().Set("content-type", "application/json")
	w.Write(final)  
}

func updateTask(w http.ResponseWriter, req *http.Request){
	decoder := json.NewDecoder(req.Body)
    var data=map[string]interface{}{
    };
    
	err := decoder.Decode(&data)
    if err != nil {
		panic(err)
	}
    
    fmt.Println(data);
    
    var output=map[string]interface{}{
    };
    if val1 , ok := data["listName"]; ok {
        fmt.Println(val1);
        fmt.Println(ok);
        fmt.Println(data["listName"])
        key :=data["listName"].(string);
        fmt.Println(key);
        if obj,exists := collection[key]; !exists {
            fmt.Println(obj);
            output["data"]="List Does Not Exist";
        } else {
            taskName := data["taskName"].(string)
            fmt.Println(taskName);
            if obj,exists := collection[key].(map[string]interface{})[taskName]; exists {
                fmt.Println(obj);
                collection[key].(map[string]interface{})[taskName]=data["newData"];
                output=collection;
            } else {
                output["data"]="Task not present.";
            }
        }
    } else {
        output["data"]="Please specify the required details.";
    }
    final , err := json.Marshal(&output)
    if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
    w.Header().Set("content-type", "application/json")
	w.Write(final)  
}
