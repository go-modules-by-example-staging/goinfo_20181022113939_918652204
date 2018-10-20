package contributors

type Person struct {
	Name string
}

var all = [...]Person{
	Person{Name: "Robert Griesemer"},
	Person{"Rob Pike"},
	Person{"Ken Thompson"},
	Person{"Russ Cox"},
	Person{"Ian Lance Taylor"},
}

func Details() []Person {
	res := all
	return res[:]
}
