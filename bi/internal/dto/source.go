package dto

type SourceDto struct {
	Id  uint32       `json:"id"`
	Dsn string       `json:"dsn"`
	Typ SourceTypDto `json:"typ"`
}

type SourceAddDto struct {
	Dsn   string `json:"dsn"`
	TypId uint32 `json:"typId"`
}
