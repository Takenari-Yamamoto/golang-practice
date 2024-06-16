package spanner

import (
	"fmt"
	"strings"
)

type TableDefine struct {
	Name        string
	Columns     []*ColumnDefine
	PrimaryKeys []string
	Interleave  *InterleaveDefine
	Indexes     []*IndexDefine
}

type InterleaveOnDeleteOpt string

const (
	InterleaveOnDeleteOpt_CASCADE  = "CASCADE"
	InterleaveOnDeleteOpt_NOACTION = "NO ACTION"
)

type InterleaveDefine struct {
	ParentTable string
	OnDelete    InterleaveOnDeleteOpt
}

func (t *TableDefine) AllColumnsString() string {
	var r string
	for i, v := range t.Columns {
		if i == 0 {
			r = v.Name
			continue
		}
		r = fmt.Sprintf("%s, %s", r, v.Name)
	}
	return r
}

func (t *TableDefine) IndexKeyAndStoringColumnsString(indexName string) []string {
	var r []string
	var idx *IndexDefine
	for _, v := range t.Indexes {
		if v.Name != indexName {
			continue
		}
		idx = v
	}
	if idx == nil {
		return []string{}
	}

	for _, c := range idx.Columns {
		r = append(r, c.Name)
	}
	r = append(r, idx.Storings...)

	return r
}

func (t *TableDefine) DDL() string {
	r := fmt.Sprintf("\nCREATE TABLE %s (", t.Name)
	for _, v := range t.Columns {
		r = fmt.Sprintf("%s\n\t%s", r, v.ddl())
	}
	var pk string
	for i, v := range t.PrimaryKeys {
		if i == 0 {
			pk = v
			continue
		}
		pk = fmt.Sprintf("%s, %s", pk, v)
	}
	r = fmt.Sprintf("%s\n) PRIMARY KEY (%s)", r, pk)

	if t.Interleave != nil && len(t.Interleave.ParentTable) != 0 && len(t.Interleave.OnDelete) != 0 {
		r = fmt.Sprintf("%s,\nINTERLEAVE IN PARENT %s ON DELETE %s",
			r,
			t.Interleave.ParentTable,
			t.Interleave.OnDelete)
	}
	// End of CREATE schema
	// r = r + ";"

	if len(t.Indexes) == 0 {
		return r
	}
	for _, v := range t.Indexes {
		r = fmt.Sprintf("%s\n%s", r, v.ddl(t.Name))
	}

	return r
}

type ColumnDefine struct {
	Name      string
	Attribute string
}

func (c *ColumnDefine) ddl() string {
	return fmt.Sprintf("%s %s,", c.Name, c.Attribute)
}

type IndexColumn struct {
	Name    string
	OrderBy string
}

type IndexDefine struct {
	Name         string
	Columns      []*IndexColumn
	NullFiltered bool
	Unique       bool
	Storings     []string
}

func (i *IndexDefine) ddl(table string) string {
	r := "CREATE"
	if i.Unique {
		r = fmt.Sprintf("%s UNIQUE", r)
	}
	if i.NullFiltered {
		r = fmt.Sprintf("%s NULL_FILTERED", r)
	}
	r = fmt.Sprintf("%s INDEX %s ON %s (", r, i.Name, table)
	var bs strings.Builder
	for i, c := range i.Columns {
		if i == 0 {
			bs.WriteString("\t")
			writeColumnString(&bs, c)
			continue
		}
		bs.WriteString(", ")
		writeColumnString(&bs, c)
	}
	r = fmt.Sprintf("%s\n%s\n)", r, bs.String())
	r = fmt.Sprintf("%s%s", r, i.storing())
	r = fmt.Sprintf("%s;", r)
	return r
}

func writeColumnString(bs *strings.Builder, c *IndexColumn) {
	bs.WriteString(c.Name)
	if c.OrderBy != "" {
		bs.WriteString(fmt.Sprintf(" %s", c.OrderBy))
	}
}

func (i *IndexDefine) storing() string {
	if len(i.Storings) == 0 {
		return ""
	}
	var bs strings.Builder
	bs.WriteString(" STORING (\n")
	for i, s := range i.Storings {
		if i == 0 {
			bs.WriteString(fmt.Sprintf("\t%s", s))
			continue
		}
		bs.WriteString(fmt.Sprintf(", %s", s))
	}
	bs.WriteString("\n)")
	return bs.String()
}
