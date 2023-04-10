// package main

// import "fmt"

// func printBytes(s string){
// 	fmt.Printf("Bytes: ")
// 	for i:= 0; i < len(s); i++{
// 		fmt.Printf("%x ", s[i])
// 	}
// }

// func printCharacter(s string){
// 	fmt.Printf(("Character: "))
// 	runes := []rune(s)
// 	for i:= 0; i < len(runes); i++{
// 		fmt.Printf("%c ", runes[i])
// 	}
// }

// func main(){
// 	name := "Hello world"
// 	name2 := "S34enot"
// 	fmt.Printf("String: %s\n", name)
// 	printBytes(name)
// 	fmt.Println()
// 	printCharacter(name)
// 	fmt.Println()
// 	printBytes(name2)
// 	fmt.Println()
// 	printCharacter(name2)
// }
// package main

// import (  
//     "fmt"
// )

// func main() {  
//     strings := "Helloworld"
// 	for _, value := range strings{
// 		if value >= 'a' && value <= 'z'{
// 			value -= 32
// 		}
// 		// fmt.Println(value)
// 		new_string := []byte{}
// 		new_string = append(new_string, byte(value))
// 		fmt.Print(string(new_string))
// 	} 
// }



