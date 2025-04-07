package utils

const DEFAULT_PAGE_SIZE = 20
const MAX_PAGE_SIZE = 100

func GetLimitAndOffset(pageNumber int, pageSize int) (limit int, offset int) {
	if pageNumber < 1 {
		pageNumber = 1
	}
	if pageSize < 1 {
		pageSize = DEFAULT_PAGE_SIZE
	}

	if pageSize > MAX_PAGE_SIZE {
		pageSize = MAX_PAGE_SIZE
	}

	offset = (pageNumber - 1) * pageNumber
	limit = pageSize

	return limit, offset
}
