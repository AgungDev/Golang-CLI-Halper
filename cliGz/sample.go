package cligz

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func IntegerResult(msg, errmsg string) int{
	fmt.Print(msg+" ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	result, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(strings.Repeat("=", 50))
		fmt.Println(errmsg)
		result = IntegerResult(msg, errmsg)
	}
	return result
}

func StringResult(msg string) string{
	fmt.Print(msg+" ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	result := scanner.Text()
	return result
}

func Form(forms []string) []string {
	var result_slice []string
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < len(forms); i++ {
		fmt.Print(forms[i]+" : ")
		scanner.Scan()
		result_slice = append(result_slice, scanner.Text())
	}
	return result_slice
}

func Menu(text string,menu_item []string) int {
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < len(menu_item); i++ {
		fmt.Printf("%d. %s\n", i+1, menu_item[i])
	}
	fmt.Println(strings.Repeat("=", 50))
	fmt.Print(text+" : ")
	scanner.Scan()
	result,_ := strconv.Atoi(scanner.Text())
	if result > len(menu_item) || result < 1 {
		fmt.Println(strings.Repeat("=", 50))
		fmt.Println("Sorry, your select items not detect")
		fmt.Println(strings.Repeat("=", 50))
		return Menu(text, menu_item)
	}
	fmt.Println(strings.Repeat("=", 50))
	return result
}