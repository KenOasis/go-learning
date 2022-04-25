package main

import (
	"errors"
	"fmt"
	"net/http"
)

// Money example: basic struct, method practice
type Money struct {
	Name string
	RateToUSDollar float64
}

func (m Money) String() string {
	return fmt.Sprintf("%s is %.4f to 1 U.S. dollar)", m.Name, m.RateToUSDollar);
}

func (m Money) transferTo(target Money, amount float64) float64 {
	return amount * target.RateToUSDollar / m.RateToUSDollar;
}

func runMoneyExample() {
	rmb := Money{
		"Chinese Yuan(RMB)",
		6.54,
	}
	euro := Money{
		"Euro",
		0.93,
	}
	fmt.Println(rmb)
	fmt.Println(euro)
	transferf1 := rmb.transferTo;
	transferf2 := Money.transferTo;
	fmt.Println(transferf1(euro, 100))
	fmt.Println(transferf2(rmb, euro, 100))
}

// Int tree example
type IntTree struct {
	val int
	left, right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}
	return it
}

func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.Contains(val)
	case val > it.val:
		return it.right.Contains(val)
	default:
		return true
	}
}

func runIntTreeExample() {
	var it *IntTree
	it = it.Insert(5)
	it = it.Insert(3)
	it = it.Insert(10)
	it = it.Insert(2)
	fmt.Println(it.Contains(2))
	fmt.Println(it.Contains(11))
}

// Basic Interface example



// declare the interface
// type Logic interface {
// 	Process(s string) string
// }

// // Client code which declare an interface
// type Client struct {
// 	L Logic
// }

// // Client code call the interface
// func (c Client) runProcess() {
// 	fmt.Println(c.L.Process("Hello World!"));
// }

// type LogicProvider struct{}

// // defined a concrete method/function which "implemented" the interface below
// func (lp LogicProvider) Process(s string) string{
// 	return fmt.Sprintf("Process with: %s", s)
// }

// func runSimpleInterface() {
// 	// Bind the concrete menthod with the inteface of the client code
// 	c := Client{ L: LogicProvider{}}

// 	// invoke the method of the interface
// 	c.runProcess()
// }

// Interface Injection

func LogOutPut(message string) {
	fmt.Println(message)
}

type SimpleDaataStore struct {
	userData map[string]string
}

func (sds SimpleDaataStore) UserNameForID(userID string) (string, bool) {
	name, ok := sds.userData[userID]
	return name, ok
}

func NewSimpleDataStore() SimpleDaataStore {
	return SimpleDaataStore{
		userData: map[string]string{
			"1": "Jimmy",
			"2": "David",
			"3": "Jackson",
		},
	}
}

type DataStore interface {
	UserNameForID(userID string) (string, bool)
}

type Logger interface {
	Log(message string)
}
// adapter to inject Logger function 
type LoggerAdapter func(message string)

func (lg LoggerAdapter) Log(message string) {
	lg(message)
}

type SimpleLogic struct {
	l Logger
	ds DataStore
}

func(sl SimpleLogic) SayHello(userID string) (string, error) {
	sl.l.Log("in Say Hello for " + userID)
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Hello, " + name, nil
}

func(sl SimpleLogic) SayGoodbye(userID string) (string, error) {
	sl.l.Log("in Say good bye for " + userID)
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "GoodBye, " + name, nil
}

func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic{
		l: l,
		ds: ds,
	}
}

type Logic interface {
	SayHello(userID string) (string, error)
}

type Controller struct {
	l Logger
	logic Logic
}

func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
	c.l.Log("In SayHello")
	userID := r.URL.Query().Get("user_id")
	message, err := c.logic.SayHello(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(message))
}

func NewController(l Logger, logic Logic) Controller {
	return Controller{
		l: l,
		logic: logic,
	}
}

func runInjectionExample() {
	l := LoggerAdapter(LogOutPut)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)
	c := NewController(l, logic)
	// this one user the 
	http.HandleFunc("/hello", c.SayHello)
	http.ListenAndServe(":8080", nil)
}
func main() {

	// runMoneyExample()
	
	// runIntTreeExample()

	// runSimpleInterface()
	runInjectionExample()
}