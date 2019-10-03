//Golang visitor counter using MySQL database
//Hardik Sharma

package main
import (
	"fmt"
	"database/sql"
	_"https://github.com/go-sql-driver/mysql" //as it is not available in standard package of go
	"net"
	"net/http"
	"log"
)
	
//First collect the user ip address 
var ip1 string
ip1 = func IndexPage(w http.ResponseWriter, r *http.Request) {
  
  ip,_,_ := net.SplitHostPort(r.RemoteAddr)
  return ip 
  
}

func main() {

	var count int
	count=0
  http.HandleFunc("/", IndexPage)
  http.ListenAndServe("localhost:3062", nil)
  
  adrs:= "root:tiger@tcp(" + ip1 + ":3062)/clientcount"
  db1, err:= sql.Open("mysql", ip1)
  if err!=nil{
	fmt.Println(err)
	}
  else{
	fmt.Println("connection established succcessfuly!")
  }
	ins err :=db.Prepare("INSERT into user(ipadd) VALUES(?)")
	if err!=nil{
		log.Fatal(err)
	}
	result, err = ins.Exec(ip1) 
	if err!=nil{
		log.Fatal(err)
	}
	
	err := result.RowsAffected()
		if err != nil {
		log.Fatal(err)
		}
	
	
	var num int
	rows, err := db.QueryRow("SELECT COUNT(*) from user where ipadd = ?" , ip1)
	err:= rows.Scan(&num)
	
	if err != nil {
		log.Fatal(err)
		}
	else{
		if num == 1
			count=count+1
	}
		
	
	
	Defer rows.Close()
	
	
	defer db.Close()
	
	fmt.Println("Number of visitors are: ", count)

 }
	