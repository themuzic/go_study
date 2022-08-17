package main

import (
	"fmt"

	"github.com/themuzic/learngo/mydict"
)

func main() {
	// account example
	/*
		account := accounts.NewAccount("Nico")
		account.Deposit(10)

			err := account.Withdraw(20)
			if err != nil {
				// log.Fatalln : Println 하고 프로그램 종료
				// log.Fatalln(err)
				fmt.Println(err)
			}
			fmt.Println(account.String())
	*/

	// dictionary example

	/*
		// Search
		dictionary := mydict.Dictionary{"first": "First baseWord"}
		definition, err := dictionary.Search("first")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(definition)
		}
	*/

	/*

		// Add
		dictionary := mydict.Dictionary{}
		baseWord := "hello"
		definition := "Greeting"
		err := dictionary.Add(baseWord, definition)
		if err != nil {
			fmt.Println(err)
		}
		hello, _ := dictionary.Search(baseWord)
		fmt.Println("found", hello, "/ difinetion :", hello)
		err2 := dictionary.Add(baseWord, definition)
		if err2 != nil {
			fmt.Println(err2)
		}

	*/

	/*

		// Update
		dictionary := mydict.Dictionary{}
		baseWord := "hello"
		dictionary.Add(baseWord, "First")
		err := dictionary.Update(baseWord, "Second")
		if err != nil {
			fmt.Println(err)
		}
		word, _ := dictionary.Search(baseWord)
		fmt.Println(word)

	*/

	// Delete
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(word)
	fmt.Println(dictionary)

}
