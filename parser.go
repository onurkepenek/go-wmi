package gowmi

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	CIM_EMPTY         = 0
	CIM_SINT8         = 16
	CIM_UINT8         = 17
	CIM_SINT16        = 2
	CIM_UINT16        = 18
	CIM_SINT32        = 3
	CIM_UINT32        = 19
	CIM_SINT64        = 20
	CIM_UINT64        = 21
	CIM_REAL32        = 4
	CIM_REAL64        = 5
	CIM_BOOLEAN       = 11
	CIM_STRING        = 8
	CIM_DATETIME      = 101
	CIM_REFERENCE     = 102
	CIM_CHAR16        = 103
	CIM_OBJECT        = 13
	CIM_FLAG_ARRAY    = 0x2000
	CIM_ARR_SINT8     = CIM_FLAG_ARRAY | CIM_SINT8
	CIM_ARR_UINT8     = CIM_FLAG_ARRAY | CIM_UINT8
	CIM_ARR_SINT16    = CIM_FLAG_ARRAY | CIM_SINT16
	CIM_ARR_UINT16    = CIM_FLAG_ARRAY | CIM_UINT16
	CIM_ARR_SINT32    = CIM_FLAG_ARRAY | CIM_SINT32
	CIM_ARR_UINT32    = CIM_FLAG_ARRAY | CIM_UINT32
	CIM_ARR_SINT64    = CIM_FLAG_ARRAY | CIM_SINT64
	CIM_ARR_UINT64    = CIM_FLAG_ARRAY | CIM_UINT64
	CIM_ARR_REAL32    = CIM_FLAG_ARRAY | CIM_REAL32
	CIM_ARR_REAL64    = CIM_FLAG_ARRAY | CIM_REAL64
	CIM_ARR_BOOLEAN   = CIM_FLAG_ARRAY | CIM_BOOLEAN
	CIM_ARR_STRING    = CIM_FLAG_ARRAY | CIM_STRING
	CIM_ARR_DATETIME  = CIM_FLAG_ARRAY | CIM_DATETIME
	CIM_ARR_REFERENCE = CIM_FLAG_ARRAY | CIM_REFERENCE
	CIM_ARR_CHAR16    = CIM_FLAG_ARRAY | CIM_CHAR16
	CIM_ARR_OBJECT    = CIM_FLAG_ARRAY | CIM_OBJECT
	CIM_ILLEGAL       = 0xfff
	CIM_TYPEMASK      = 0x2FFF
)

func parse(item string) any {

	ind := strings.Index(item, ":")

	if ind < 0 {
		return nil
	}

	str := item[:ind]
	d_type, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return nil
	}
	str = item[ind+1:]

	if str == "(null)" {
		return nil
	}

	switch d_type {
	case CIM_EMPTY:
		return nil
	case CIM_SINT8:
		res, _ := strconv.ParseInt(str, 10, 8)
		return int8(res)
	case CIM_UINT8:
		res, _ := strconv.ParseUint(str, 10, 8)
		return uint8(res)
	case CIM_SINT16:
		res, _ := strconv.ParseInt(str, 10, 16)
		return int16(res)
	case CIM_UINT16:
		res, _ := strconv.ParseUint(str, 10, 16)
		return uint16(res)
	case CIM_SINT32:
		res, _ := strconv.ParseInt(str, 10, 32)
		return int32(res)
	case CIM_UINT32:
		res, _ := strconv.ParseUint(str, 10, 32)
		return uint32(res)
	case CIM_SINT64:
		res, _ := strconv.ParseInt(str, 10, 64)
		return int64(res)
	case CIM_UINT64:
		res, _ := strconv.ParseUint(str, 10, 64)
		return uint64(res)
	case CIM_REAL32:
		return nil
	case CIM_REAL64:
		return nil
	case CIM_BOOLEAN:
		res, _ := strconv.ParseBool(str)
		return res
	case CIM_STRING:
		return str
	case CIM_DATETIME:
		year := str[0:4]
		month := str[4:6]
		day := str[6:8]
		hour := str[8:10]
		minute := str[10:12]
		second := str[12:14]

		tm, _ := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%s-%s-%s %s:%s:%s", year, month, day, hour, minute, second))
		return tm
	case CIM_ARR_STRING:
		return strings.Split(str, ";")
	}
	return str
}
