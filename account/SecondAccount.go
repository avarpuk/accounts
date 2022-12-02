package account

import "fmt"

type SecondAccount struct {
	Name string
	Age  int
}

func (c *SecondAccount) ChangeName(newName string) string {
	c.Name = newName
	return "Name has successfully changed"
}

func (c *SecondAccount) ChangeAge(newAge int) string {
	c.Age = newAge
	return "Age has successfully changed"
}

func (c *SecondAccount) GetAccount() {
	fmt.Printf("Second Account: [Name: %s, Age: %v]", c.Name, c.Age)
}
