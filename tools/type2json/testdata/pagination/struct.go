package pagination

type Page struct {
	Count    int64 `json:"count"`
	PageSize int64 `json:"page_size"`
	PageNum  int64 `json:"page_num"`
}
