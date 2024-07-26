package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main(){

	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	
	for _,arg := range args {
		if arg == ""{
			z01.PrintRune('\n')
			continue
		}
		res := rev(arg)
		pr(res)
	}
}

func pr(s string) {
	for _, n := range s{
		z01.PrintRune(n)
	}
	z01.PrintRune('\n')
}

func rev (b string ) string {
	a :=split(b, " ")
	res := []string{}
for _,word := range a {
	if len(word) == 0{
		res = append(res, word)
		continue
	}
	var newWord []rune
	for i, r := range word{
		if i < len(word)-1 {
			newWord = append(newWord, toLower(r))	
		} else {
			newWord = append(newWord, toUpper(r))
		}
	}
	res = append(res, string(newWord))
	}
	str := join(res)
return str
}





// toLower converts a rune to lowercase if it is an uppercase letter
func toLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + ('a' - 'A')
	}
	return r
}

// toUpper converts a rune to uppercase if it is a lowercase letter
func toUpper(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return r - ('a' - 'A')
	}
	return r
}

func join(a[]string) string {
	res := ""
	for i, word := range a {
		if i != len(a){
			res += word + " "
		} else {
			res += word
		}
	}
	return res
}

func split(s , delim string) []string {
	result := []string{}
	start := 0

	for i:=0; i+len(delim) <=len(s); i++{
		if s[i:i+len(delim)]  == delim{
			if start < i {
				result= append(result, s[start:i])
			}
			start=i+len(delim)
		}
	}
	if start < len(s){
		result = append(result, s[start:])
	}
	return result
}


// func Split(str string )[]string{
// 	words := ""
// 	var res []string
// 	for i, v := range str{
// 		sp := "A"
// 		if string(v) != sp {
// 			words += string(str[i])
// 		}else if words != ""{
// 			res = append(res, words)
// 			words = ""
// 		}
// 	}
// 	res = append(res, words)
// 	return res
// }

// func join(slice []string) string {
// 	var result string
// 	for i, char := range slice {
// 		result += string(char)
// 		if i < len(slice)-1 {
// 			result += " "
// 		}
// 	}
	
// 	return result
// }

// func Joins(str []string)string{
// 	res := ""
// //last := str[len(str)-1]
// 	for i,r := range str{
// 		if i < len(str)-1{
// 			res += " "
// 		}
// 		res += r
// 	}
// 	//res = res - last
// 	return res 
// }

// func main() {
// 	args := os.Args[1]
// 	rony := split(args, "AD")
// 	fmt.Println(join(rony))

// }