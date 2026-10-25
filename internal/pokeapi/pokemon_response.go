package pokeapi

type Pokemon struct {
	Name string `json:"name"`
	Exp  int    `json:"base_experience"`
}
