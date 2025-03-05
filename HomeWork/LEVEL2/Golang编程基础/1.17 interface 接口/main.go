package main

import "fmt"

type PayMethod interface {
	Pay(int)
}

type CreditCard struct {
	balance int
	limit   int
}

func (c *CreditCard) Pay(amout int) {
	if c.balance < amout {
		fmt.Println("余额不足")
		return
	}
	c.balance -= amout
}

func anyParam(param interface{}) {
	fmt.Println("param: ", param)
}

func main() {
	c := CreditCard{balance: 100, limit: 1000}
	c.Pay(200)
	var a PayMethod = &c
	fmt.Println("a.Pay: ", a)

	var b interface{} = &c
	fmt.Println("b: ", b)

	anyParam(c)
	anyParam(1)
	anyParam("123")
	anyParam(a)

	/* 	余额不足
	   	a.Pay:  &{100 1000}
	   	b:  &{100 1000}
	   	param:  {100 1000}
	   	param:  1
	   	param:  123
	   	param:  &{100 1000} */
}
