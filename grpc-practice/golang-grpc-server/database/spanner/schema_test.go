package spanner_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/Takenari-Yamamoto/golang-grpc-server/database/spanner"
)

func TestDDL(t *testing.T) {
	cases := map[string]struct {
		tableName string
		define    spanner.TableDefine
		want      string
	}{
		"Check: DDL": {
			tableName: "Tests",
			define: spanner.TableDefine{
				Name: "Tests",
				Columns: []*spanner.ColumnDefine{
					{
						Name:      "ID",
						Attribute: "STRING(MAX) NOT NULL",
					},
					{
						Name:      "CreatedAt",
						Attribute: "TIMESTAMP NOT NULL OPTIONS (allow_commit_timestamp=true)",
					},
					{
						Name:      "Status",
						Attribute: "STRING(MAX) NOT NULL",
					},
					{
						Name:      "Memo",
						Attribute: "STRING(MAX) NOT NULL",
					},
				},
				PrimaryKeys: []string{"ID", "CreatedAt"},
				Indexes: []*spanner.IndexDefine{
					{
						Name: "test_idx",
						Columns: []*spanner.IndexColumn{
							{
								Name:    "ID",
								OrderBy: "ASC",
							},
							{
								Name: "CreatedAt",
							},
						},
						Unique:   false,
						Storings: []string{"Status", "Memo"},
					},
				},
			},
			want: `
			CREATE TABLE Tests (
				ID STRING(MAX) NOT NULL,
				CreatedAt TIMESTAMP NOT NULL OPTIONS (allow_commit_timestamp=true),
				Status STRING(MAX) NOT NULL,
				Memo STRING(MAX) NOT NULL,
			) PRIMARY KEY (ID, CreatedAt);
			CREATE INDEX test_idx ON Tests (
				ID ASC, CreatedAt
			) STORING (
				Status, Memo
			);`,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := tc.define.DDL()
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("result didn't match (-want / +got)\n%s\nfull ddl: %s", diff, got)
			}
		})
	}
}
