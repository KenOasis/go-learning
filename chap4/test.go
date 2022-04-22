package main

import (
	"errors"
	"fmt"
)

type factContent struct {
	amountPerServ float64
	anountUnit string
	dailyPercent int
}

type supplementFacts map[string]factContent

func printSupplementFacts(facts supplementFacts) {
	if len(facts) == 0 {
		fmt.Println("There is no supplement in the facts list.")
		return 
	} else {
		defer func(){
			fmt.Println("Printing end!");
		}()
	}
	fmt.Println("The supplement facts is/are:")
	for k, v := range facts {
		fmt.Printf("%s\t%f %s\t%d%%\n", k, v.amountPerServ, v.anountUnit, v.dailyPercent)
	}
}

func addSupplement(facts supplementFacts, sup string, amount float64, unit string, daily int) (bool, error){
	if _, ok := facts[sup]; ok {
		error := errors.New(sup + " already exists in the facts.")
		return false, error
	}
	newFact := factContent{amount, unit, daily}
	facts[sup] = newFact
	return true, nil	
}

func removeSupplement(facts supplementFacts, sup string) factContent{
	if v, ok := facts[sup]; ok {
		delete(facts, sup)
		return v 
	} else {
		return factContent{}
	}
}

func main() {
	facts := supplementFacts{}
	printSupplementFacts(facts)
	addSupplement(facts, "Vitamin A", 900, "mcg", 100)
	printSupplementFacts(facts)
	addSupplement(facts, "Vitamin C", 99, "mg", 110)
	addSupplement(facts, "Vitamin D", 25, "mcg", 120)
	printSupplementFacts(facts)
	removeFacts := []string{"Vitamin B", "Vitamin C"}
	emptyfact := factContent{}
	for _, v := range removeFacts{
		fact := removeSupplement(facts, v)
		if fact == emptyfact {
			fmt.Println(v, "is not in the facts")
		} else {
			fmt.Println(v, "is successfully removed from the facts")
		}
	}
	printSupplementFacts(facts)

	var c1 rune = 35885
	var c2 rune = 38182
	var c3 rune = 20581
	fmt.Printf("%c%c%c\n", c1, c2, c3)
	name := "谭锦健"
	for _, v := range name {
		fmt.Printf("%c", v)
	}
	fmt.Println()
}