package main 
 import "fmt" 
 func main() {     
	var n int 
    fmt.Print("bilangan: ")     
	fmt.Scanln(&n) 
    fmt.Print("Faktor: ")     
	for i := 1; i <= n; i++ {         
		if n%i == 0 {             
			fmt.Print(i, " ") 
        }     
	}     
	fmt.Println() 
    var jumlahFaktor int = 0    
	for i := 1; i <= n; i++ {         
		if n%i == 0 {             
			jumlahFaktor++ 
        } 
    }      
	if jumlahFaktor == 2 { 
        fmt.Println("Prima: true") 
    } else {         
		fmt.Println("Prima: false") 
    } 
} 
 
