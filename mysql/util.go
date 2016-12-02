// Copyright 2015 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package mysql

import (
	"strings"
	"unicode"
)

// GetDefaultFieldLength is used for Interger Types, Flen is the display length.
// Call this when no Flen assigned in ddl.
// or column value is calculated from an expression.
// For example: "select count(*) from t;", the column type is int64 and Flen in ResultField will be 21.
// See https://dev.mysql.com/doc/refman/5.7/en/storage-requirements.html
func GetDefaultFieldLength(tp byte) int {
	switch tp {
	case TypeTiny:
		return 4
	case TypeShort:
		return 6
	case TypeInt24:
		return 9
	case TypeLong:
		return 11
	case TypeLonglong:
		return 21
	case TypeDecimal, TypeNewDecimal:
		// See https://dev.mysql.com/doc/refman/5.7/en/fixed-point-types.html
		return 10
	case TypeBit, TypeBlob:
		return -1
	default:
		//TODO: Add more types.
		return -1
	}
}

// GetDefaultDecimal returns the default decimal length for column.
func GetDefaultDecimal(tp byte) int {
	switch tp {
	case TypeDecimal, TypeNewDecimal:
		// See https://dev.mysql.com/doc/refman/5.7/en/fixed-point-types.html
		return 0
	default:
		//TODO: Add more types.
		return -1
	}
}

func isSpace(c byte) bool {
	return c == ' ' || c == '\t'
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func myMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func myMaxInt8(a, b int8) int8 {
	if a > b {
		return a
	}
	return b
}

func myMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func myMinInt8(a, b int8) int8 {
	if a < b {
		return a
	}
	return b
}

// strToInt converts a string to an integer in best effort.
// TODO: Handle overflow and add unittest.
func strToInt(str string) (int64, error) {
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return 0, nil
	}
	negative := false
	i := 0
	if str[i] == '-' {
		negative = true
		i++
	} else if str[i] == '+' {
		i++
	}
	r := int64(0)
	for ; i < len(str); i++ {
		if !unicode.IsDigit(rune(str[i])) {
			break
		}
		r = r*10 + int64(str[i]-'0')
	}
	if negative {
		r = -r
	}
	// TODO: If i < len(str), we should return an error.
	return r, nil
}
