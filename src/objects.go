package main
import ("fmt")

func main(){
    var collection=map[string]interface{}{
        "list1":map[string]interface{}{
        },
    };
    
    collection["list1"]="data";
//    delete (collection,"list1");
    fmt.Println(collection);
};