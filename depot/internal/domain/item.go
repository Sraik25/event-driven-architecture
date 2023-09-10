package domain

type Items map[string]*Item

type Item struct {
	ProductNAme string
	Quantity    int
}
