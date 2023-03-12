//go:build ignore

package main

import (
	"log"

	"entgo.io/ent/entc/gen"

	"github.com/frisbm/enthistory"

	"entgo.io/ent/entc"
)

func main() {
	if err := entc.Generate("./schema",
		&gen.Config{},
		entc.Extensions(enthistory.NewHistoryExtension("userId")),
	); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}