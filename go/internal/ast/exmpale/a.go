package exmpale

import "encoding/json"

type A struct {
	Db json.Marshaler `json:"db"`
}
