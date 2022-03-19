package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/bnert/mfr"
	mfru "github.com/bnert/mfr/utils"
)

var (
	csvData string = `mable pines,twin,12
dipper pines,twin,12
stanford pines,great uncle,60
stanely pines,great uncle,60
stanely pines,great uncle,60
wendy corduroy,friend,15`
)

type FamilyMember struct {
	Name     string `json:"name"`
	Relation string `json:"relation"`
	Age      int    `json:"age"`
}

type Family struct {
	Surname string         `json:"surname"`
	Members []FamilyMember `json:"members"`
}

func main() {
	data, err := csv.NewReader(strings.NewReader(csvData)).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	validData := mfr.Filter[[]string](data,
		func(ctx mfr.Ctx[[]string]) bool {
			return len(ctx.Item) == 3 && len(ctx.Item[1]) != 0 && ctx.Item[1] != "friend"
		},
	)

	distinctMembers := mfru.DistinctBy[[]string, string](validData,
		func(parts []string) string {
			return strings.Join(parts, "")
		},
	)

	members := mfr.Map[[]string, FamilyMember](distinctMembers,
		func(ctx mfr.Ctx[[]string]) FamilyMember {
			age, _ := strconv.Atoi(ctx.Item[2])
			return FamilyMember{
				Name:     ctx.Item[0],
				Relation: ctx.Item[1],
				Age:      age,
			}
		},
	)

	family := mfr.Reduce[FamilyMember, Family](members, Family{Surname: "Pines"},
		func(ctx mfr.Ctx[FamilyMember], acc Family) Family {
			acc.Members = append(acc.Members, ctx.Item)
			return acc
		},
	)

	j, err := json.MarshalIndent(family, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(j))
}
