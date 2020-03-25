package newpackage
import "fmt"

var sum int = 0
var str2 string = "Exit"
var str1 string = "Please select an option:"
var str3 string =	"1 – Print Covid19 cases in Pakistan" 
var str4 string =	"2 – Print Covid19 cases in SouthKorea"
var str5 string =	"3 – Print Covid19 cases in France"
var str6 string =	"4 – Print my personalized message to Coronavirus"
var str7 string =	"0 – Exit"
	
func F(){
	fmt.Println(str1)
	fmt.Println(str3)
	fmt.Println(str4)
	fmt.Println(str5)
	fmt.Println(str6)
	fmt.Println(str7)
	var num int
	fmt.Scanf("%d", &num)
	
	for sum <= 4 {
		switch num {
		case 1:
			fmt.Println("Covid19 cases in Pakistan are: ")
			fmt.Println("1000")
			
		case 2:
			fmt.Println("Covid19 cases in SouthKorea are: ")
			fmt.Println("9137")

		case 3:
			fmt.Println("Covid19 cases in France are: ")
			fmt.Println("22,304")
		
		case 4:
			fmt.Println("Covid19 is a big problem for us and we need to battle it as a team and stay at home to stay safe")
		
		default:
			fmt.Printf(str2)
		}
		sum++
	}

}
