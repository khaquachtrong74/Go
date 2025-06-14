package operations
import "fmt" 

func ReadDB(){
	_, err := OpenDB()
	if err != nil{
		fmt.Println("Error Open DB")
	}else{
		fmt.Println("Success Open DB")
	}
}
