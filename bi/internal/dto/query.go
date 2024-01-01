package dto

type QueryDto struct {
	SourceId uint32 `json:"sourceId"`
	Query    string `json:"query"`
}
