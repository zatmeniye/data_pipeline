package keeper

type Dataset struct {
	Columns []Column         `json:"columns"`
	Rows    []map[string]any `json:"rows"`
}

type Column struct {
	Typ  string `json:"typ"`
	Name string `json:"name"`
}
