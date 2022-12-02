package account

import "fmt"

type FirstAccount struct {
	Name string
	Age  int
}

func (c *FirstAccount) ChangeName(newName string) string {
	c.Name = newName
	return "Name has successfully changed"
}

func (c *FirstAccount) ChangeAge(newAge int) string {
	c.Age = newAge
	return "Age has successfully changed"
}

func (c *FirstAccount) GetAccount() {
	if len(c.Name) == 0 {
		c.Name = "N/A"
	}
	fmt.Printf("First Account: [Name: %s, Age: %v]", c.Name, c.Age)
}
