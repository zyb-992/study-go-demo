package reflect

type ListReq struct {
	Id       int64                  `json:"id" column:"id"`
	Name     string                 `json:"name" column:"name"`
	PageSize int64                  `json:"page_size" column:"" parse:"no"`
	PageNum  int64                  `json:"page_num" column:""   parse:"no"`
	Embed    EmbedReq               `json:"embed" column:"embed"`
	IntList  []int64                `json:"int_list" column:"int_list"`
	StrList  []string               `json:"str_list" column:"str_list"`
	Map      map[string]interface{} `json:"map" column:"map"`
}

type EmbedReq struct {
	EmbedId   int64  `json:"embed_id" column:"embed_id"`
	EmbedName string `json:"embed_name" column:"embed_name"`
}
