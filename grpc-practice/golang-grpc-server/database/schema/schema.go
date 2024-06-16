package schema

import "github.com/Takenari-Yamamoto/golang-grpc-server/database/spanner"

// table name
const (
	UserTable = "Users"
)

var (
	tables = []*spanner.TableDefine{
		{
			Name: UserTable,
			Columns: []*spanner.ColumnDefine{
				{
					Name:      "ID",
					Attribute: "STRING(MAX)",
				},
				{
					Name:      "Name",
					Attribute: "STRING(MAX)",
				},
				{
					Name:      "Age",
					Attribute: "INT64",
				},
			},
		},
	}

	tableMap = map[string]*spanner.TableDefine{}
)

func init() {
	tableMap = make(map[string]*spanner.TableDefine, len(tables))
	for _, v := range tables {
		tableMap[v.Name] = v
	}
}

func TableDefine(tableName string) *spanner.TableDefine {
	return tableMap[tableName]
}
