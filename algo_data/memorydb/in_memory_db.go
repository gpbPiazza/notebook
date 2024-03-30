package memorydb

import (
	"fmt"
	"sort"
)

const (
	setCommand            = "SET"
	getCommand            = "GET"
	deleteCommand         = "DELETE"
	scanCommand           = "SCAN"
	scanByPrefixCommand   = "SCAN_BY_PREFIX"
	setAtCommand          = "SET_AT"
	getAtCommand          = "GET_AT"
	deleteAtCommand       = "DELETE_AT"
	scanAtCommand         = "SCAN_AT"
	scanByPrefixAtCommand = "SCAN_BY_PREFIX_AT"
)

type record map[string]string

type database map[string]record

func newDatabase() database {
	return make(database)
}

func (db database) Set(key, field, val string) string {
	_, ok := db[key]
	if !ok {
		db[key] = make(record)
	}

	db[key][field] = val

	return ""
}

func (db database) Get(key, field string) string {
	record, ok := db[key]
	if ok {
		if val, exists := record[field]; exists {
			return val
		}
	}

	return ""
}

func (db database) Delete(key, field string) string {
	if record, exists := db[key]; exists {
		if _, exists := record[field]; exists {
			delete(record, field)
			return "true"
		}
	}

	return "false"
}

func (db database) Scan(key string) string {
	record, exists := db[key]
	if !exists {
		return ""
	}

	var fields []string
	for field := range record {
		fields = append(fields, field)
	}

	return db.scanResponse(fields, record)
}

func (db database) ScanByPrefix(key, prefix string) string {
	record, exists := db[key]
	if !exists {
		return ""
	}

	var fields []string
	for field := range record {
		if stringHasPrefix(field, prefix) {
			fields = append(fields, field)
		}
	}
	return db.scanResponse(fields, record)
}

func stringHasPrefix(target, prefix string) bool {
	if len(target) < len(prefix) {
		return false
	}

	for i := 0; i < len(prefix); i++ {
		if target[i] != prefix[i] {
			return false
		}
	}

	return true
}

func (db database) scanResponse(fields []string, rec record) string {
	var result string
	sort.Strings(fields)

	for _, field := range fields {
		val, exists := rec[field]
		if !exists {
			continue
		}
		if result == "" {
			result = db.scanFormat(field, val)
			continue
		}
		result = fmt.Sprintf("%s, %s", result, db.scanFormat(field, val))
	}
	return result
}

func (db database) scanFormat(field, val string) string {
	return fmt.Sprintf("%s(%s)", field, val)
}

func (db database) SetAt(key, field, val, timestamp string) string {

	return ""
}

func solution(queries [][]string) []string {
	db := newDatabase()
	result := make([]string, 0)

	for _, query := range queries {
		switch query[0] {
		case setCommand:
			result = append(result, db.Set(query[1], query[2], query[3]))
		case getCommand:
			result = append(result, db.Get(query[1], query[2]))
		case deleteCommand:
			result = append(result, db.Delete(query[1], query[2]))
		case scanCommand:
			result = append(result, db.Scan(query[1]))
		case scanByPrefixCommand:
			result = append(result, db.ScanByPrefix(query[1], query[2]))
		case setAtCommand:
			result = append(result, db.SetAt(query[1], query[2], query[3], query[4]))
		case getAtCommand:
			// result = append(result, db.Get(query[1],query[2]))
		case deleteAtCommand:
			// result = append(result, db.Delete(query[1],query[2]))
		case scanAtCommand:
			// result = append(result, db.Scan(query[1]))
		case scanByPrefixAtCommand:
			// result = append(result, db.ScanByPrefix(query[1], query[2]))
		}

	}

	return result
}
