package peopleMethods

type People struct {
	Name string
	Age  int
}

func (p People) AgeLimit() string {
	if p.Age < 60 {
		return "You're accepted into the program"
	}
	return "I'm sorry, but you don't meet the age requirement outlined in our guidelines"
}
