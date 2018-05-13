# gocode
Create List :

Link :  <domain link>/createList

Request Body : 
{
	"Name":"List1"
}

Delete List :

Link :  <domain link>/deleteList

Request Body : 
{
	"Name":"List1"
}

Add Task :

Link : <domain link>/addTask

 Request Body :  
{
	"listName":"List1",
	"task":{
		"name":"task1",
		"data":"task 1 data"
	}
} 

Delete Task :

Link :  <domain link>/deleteTask

Request Body :  
{
	"listName":"List1",
	"taskName":"task1"
}

Update Task :

Link : <domain link>/updateTask

Request Body :  
{
	"listName":"List1",
	"taskName":"task1",
	"newData":"task 1 updated data"
}
