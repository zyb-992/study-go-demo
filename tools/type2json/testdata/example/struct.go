package example

import "github.com/zyb-992/demo/tools/type2json/testdata/pagination"

type ExampleStruct struct {
	PresetInt64    int64                       `json:"preset_int_64"`
	PresetStr      string                      `json:"preset_str"`
	PresetFloat64  float64                     `json:"preset_float_64"`
	PresetSlice    []int64                     `json:"preset_slice"`
	PresetMap      map[string]string           `json:"preset_map"`
	PresetMapSlice map[string][]int64          `json:"preset_map_slice"`
	PresetMapMap   map[string]map[string]int64 `json:"preset_map_map"`

	NormalStruct Normal `json:"normal_struct"`
	// this field should be pass, because it is unexported field
	normalStruct Normal `json:"normal_struct_pass"`

	// consider embedded field, should be marshal
	Normal

	// third party package
	pagination.Page
	PageInfo pagination.Page `json:"page_info"`
	// this field should be pass, because it is unexported field
	pageInfo pagination.Page `json:"unexported_page_info"`
}

type Normal struct {
	Doc      string  `json:"doc"`
	Name     string  `json:"name"`
	Age      int64   `json:"age"`
	StaffIds []int64 `json:"staff_ids"`
}
