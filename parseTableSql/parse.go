package parseTableSql

import (
	"fmt"
	"regexp"
	"strings"
)

//需要将字段的``更换成""
func Parse(tableSql string, structName string) {

	Regexp, _ := regexp.Compile(`"(.*)"(.*)\([\d+]+.*[\d+]?\).*`)

	strArray := Regexp.FindAllStringSubmatch(tableSql, -1)
	fmt.Println(len(strArray))

	var fields []string
	var cases []string
	var comment string
	for _, str := range strArray {
		var fieldName string
		for k, str := range strings.Split(str[1], `_`) {
			if k == 0 {
				fieldName += strings.ToUpper(str[:1]) + str[1:]
				continue
			}
			fieldName += strings.ToUpper(str[:1]) + str[1:]
		}

		if strings.Index(str[0], `COMMENT`) > 0 {

			comment = "//" + str[0][strings.Index(str[0], `COMMENT`)+7:]
		}

		_type := GetType(str[2], strings.Index(str[0], `UNSIGNED`) > 0)
		fields = append(fields, fmt.Sprintf(`%s	%s  %s`, fieldName, _type, comment))
		cases = append(cases, fmt.Sprintf(tempcase, str[1], fieldName, GetRowType(_type)))
	}

	str := fmt.Sprintf(template, structName, strings.Join(fields, `
	`), structName, structName, structName, structName, structName, structName, strings.Join(cases, `
`))
	fmt.Println(str)
}

func GetType(atype string, IsUsigned bool) string {

	switch strings.TrimSpace(atype) {

	case "tinyint":
		if IsUsigned {
			return "uint8"
		} else {
			return "int8"
		}
	case "smallint":
		if IsUsigned {
			return "uint16"
		} else {
			return "int16"
		}
	case "mediumint":
		if IsUsigned {
			return "uint32"
		} else {
			return "int32"
		}
	case "integer":
		fallthrough
	case "int":
		if IsUsigned {
			return "uint32"
		} else {
			return "int32"
		}
	case "bigint":
		if IsUsigned {
			return "uint64"
		} else {
			return "int64"
		}
	case "float":
		return "float32"
	case "double":
		return "float64"
	case "decimal":
		return "float64"
	case "date":
		return "string"
	case "time":
		return "string"
	case "datetime":
		return "string"
	case "timestamp":
		return "string"
	case "char":
		return "string"
	case "varchar":
		return "string"
	case "tinyblob":
		return "string"
	case "tinytext":
		return "string"
	case "blob":
		return "string"
	case "text":
		return "string"
	case "mediumblob":
		return "string"
	case "mediumtext":
		return "string"
	case "longblob":
		return "string"
	case "longtext":
		return "string"
	default:
		return "string"
	}

}

func GetRowType(_type string) string {

	switch _type {
	case "int8":
		return "int8(row.Int64(res.Map(v)))"
	case "uint8":
		return "uint8(row.Uint64(res.Map(v)))"
	case "uint16":
		return "uint16(row.Uint64(res.Map(v)))"
	case "int16":
		return "int16(row.Int64(res.Map(v)))"
	case "uint32":
		return "uint32(row.Uint64(res.Map(v)))"
	case "int32":
		return "int32(row.Int64(res.Map(v)))"
	case "uint64":
		return "row.Uint64(res.Map(v))"
	case "int64":
		return "row.Int64(res.Map(v))"
	case "float32":
		return "float32(row.Float(res.Map(v)))"
	case "float64":
		return "row.Float(res.Map(v))"
	case "string":
		return "row.Str(res.Map(v))"
	default:
		return ""
	}
}

const template = `

type %s struct {
	%s
}


func (u %s) RowMapping(row mysql.Row, res mysql.Result) *%s {
	if row == nil {
		return nil
	}
	records := u.RowsMapping([]mysql.Row{row}, res)
	if len(records) == 0 {
		return nil
	}
	return records[0]
}

func (u %s) RowsMapping(rows []mysql.Row, res mysql.Result) []*%s {
	if len(rows) == 0 {
		return nil
	}

	fields := GetResFields(res)
	var records []*%s
	for _, row := range rows {
		record := &%s{}
		for _, v := range fields {
			switch v {
			%s
			}
		}
		records = append(records, record)
	}

	return records
}


func GetResFields(res mysql.Result) []string {
	var ret []string
	for _, v := range res.Fields() {
		ret = append(ret, strings.ToLower(v.Name))
	}
	return ret
}


`

const tempcase = `	case "%s": record.%s = %s`
