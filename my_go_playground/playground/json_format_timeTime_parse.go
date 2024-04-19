package playground

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type VamoGremio struct {
	PaidDate time.Time `json:"paid_date,omitempty"`
}

func MakeVamoGremioJson() {
	gremioVazio := VamoGremio{}
	parsed := VamoGremio{}

	val, err := json.Marshal(&gremioVazio)

	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(val, &parsed)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(val)
	fmt.Println(parsed)
}

func Zap() {
	getMonolitoGaveta := "2023-10-26T00:00:00Z"
	listV2Inst := "2023-10-25T21:00:00-03:00"

	listV2Parsed, err := time.Parse(time.RFC3339, listV2Inst)
	if err != nil {
		log.Fatal(err)
	}
	gaveta, err := time.Parse(time.RFC3339, getMonolitoGaveta)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("List v2 -> ", listV2Parsed.UTC())
	fmt.Println("Gaveta v1 -> ", gaveta.UTC())

}
