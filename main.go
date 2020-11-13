package main

import (
	"fmt"
	"strings"

	"github.com/legacyofmu/learngo/accounts"
	"github.com/legacyofmu/learngo/mydict"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// naked return
// defer : function이 끝났을때 실행되는 함수 지정
func lenAndUpper2(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	//return len(name), strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers { // _ : index
		//fmt.Println(index, number)
		total += number
	}
	return total
}

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func low() {
	a := 2
	b := a
	a = 10
	fmt.Println(a, b)
	fmt.Println(&a, &b) // & : address

	a = 2
	c := &a // c : pointer to a
	a = 10
	fmt.Println(a, *c)
	fmt.Println(&a, c)
}

func testArray() {
	names := []string{"dhpark", "lynn", "dal", "marl"}
	names = append(names, "flynn")
	fmt.Println(names)
}

func testMap() {
	dhpark := map[string]string{"name": "dhpark", "age": "12"}
	fmt.Println(dhpark)
	for key, value := range dhpark {
		fmt.Println(key, value)
	}
}

type person struct {
	name    string
	age     int
	favFood []string
}

func testStruct() {
	favFood := []string{"kimchi", "ramen"}
	dhpark := person{"dhpark", 18, favFood}
	fmt.Println(dhpark.name)
}

func testAccount() {
	account := accounts.NewAccount("dhpark")
	fmt.Println(account)
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance())
}

func testDictionary() {
	dictionary := mydict.Dictionary{"first": "First word"}
	//dictionary["hello"] = "hello"
	fmt.Println(dictionary)
	definition, err := dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
	err1 := dictionary.Add("hello", "Greeting")
	if err1 != nil {
		fmt.Println(err1)
	}
	definition2, err2 := dictionary.Search("hello")
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(definition2)
	}
	err3 := dictionary.Update("hello", "Bye")
	if err3 != nil {
		fmt.Println(err3)
	}
	word, _ := dictionary.Search("hello")
	fmt.Println(word)
	dictionary.Delete("hello")
	word4, err4 := dictionary.Search("hello")
	if err4 != nil {
		fmt.Println(err4)
	} else {
		fmt.Println(word4)
	}
}

func main() {
	totalLength, upperName := lenAndUpper2("dhpark")
	fmt.Println(totalLength, upperName)
	repeatMe("dhpark", "lynn", "dal", "marl", "flynn")
	total := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(total)
	fmt.Println(canIDrink(16))
	low()
	testArray()
	testMap()
	testStruct()
	testAccount()
	testDictionary()
}
