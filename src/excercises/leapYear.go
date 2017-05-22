package excercises

func IsLeapYear(year int) bool {
	if year % 4 == 0 {
		if year % 100 == 0 {
			if year % 400 == 0 {
				return true
			} else {
				return false
			}
		} else {
			return true
		}
	}
	return false 
}

//func main() {
//	var year int
//	fmt.Print("Enter year: ")
//	fmt.Scanf("%d", &year)
//	if excercises.IsLeapYear(year) {
//		fmt.Println(year , " is leap")
//	} else {
//		fmt.Println(year , " is not leap")
//	}
//}

