package utils

import "codebase-app/internal/common/consts"

func FieldOrderBy(query *string, orderBy consts.OrderBy, sort string, fields ...string) {
	if orderBy != consts.ASC && orderBy != consts.DESC {
		return
	}

	for _, field := range fields {
		if field == sort {
			*query += " ORDER BY " + field
			*query += " " + string(orderBy)
			break
		}
	}
}
