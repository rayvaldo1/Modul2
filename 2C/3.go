package main 
 import "fmt" 
 func main() {     
	var azka int 
    fmt.Print("bilangan: ")     
	fmt.Scanln(&azka) 
    fmt.Print("Faktor: ")     
	for i := 1; i <= azka; i++ {         
		if azka%i == 0 {             
			fmt.Print(i, " ") 
        }     
	}     
	fmt.Println() 
    var jumlahFaktor int = 0    
	for i := 1; i <= azka; i++ {         
		if azka%i == 0 {             
			jumlahFaktor++ 
        } 
    }      
	if jumlahFaktor == 2 { 
        fmt.Println("Prima: true") 
    } else {         
		fmt.Println("Prima: false") 
    } 
} 
 
