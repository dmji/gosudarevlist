package model

type Entity struct {
	Name string
}

type OptionsGetItems struct {
	Count  int
	Offset int

	SearchQuery string
}
