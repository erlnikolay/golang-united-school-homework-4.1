package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here
	sourceRuneString := []rune(input)
	//fmt.Println(sourceRuneString)
	resRuneString := make([]rune, len(sourceRuneString), cap(sourceRuneString))
	for iter := len(sourceRuneString) - 1; iter >= 0; iter-- {
		//fmt.Println(iter)
		resRuneString[(len(sourceRuneString)-1)-iter] = sourceRuneString[iter]
	}
	output = string(resRuneString)
	return output
}

//func main() {
//	fmt.Println(ReverseString("А роза упала на лапу азорА"))
//}
