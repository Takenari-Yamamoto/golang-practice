package main

import (
	"log"
	"strings"

	"github.com/Takenari-Yamamoto/golang-grpc-server/database/emulator"
	"github.com/Takenari-Yamamoto/golang-grpc-server/database/schema"
)

// エミュレータ起動用のコマンド
func main() {
	td := schema.TableDefine(schema.UserTable).DDL()

	var ddl []string
	for _, v := range strings.Split(td, "\n") {
		if v == "" {
			continue
		}
		ddl = append(ddl, v)
	}

	client := emulator.SetupDatabase(ddl)

	log.Printf("Spanner client is created successfully: %v", client)

}
