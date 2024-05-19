package helpers

import (
	"strconv"
)

func GetSortParams(sort, orderBy string) string {
	if sort != "ASC" && sort != "DESC" {
		sort = "ASC"
	}

	if orderBy == "" {
		orderBy = "name"
	}

	return orderBy + " " + sort
}

func GetPaginationParams(oldLimit, oldPage string) (int, int, int) {
	page, _ := strconv.Atoi(oldPage)
	if page == 0 {
		page = 1
	}

	limit, _ := strconv.Atoi(oldLimit)
	if limit == 0 {
		limit = 5
	}

	offset := (page - 1) * limit

	return page, limit, offset
}
