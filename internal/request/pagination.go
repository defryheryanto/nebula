package request

import (
	"net/http"
	"net/url"
	"strconv"

	handlederror "github.com/defryheryanto/nebula/internal/errors"
)

func GetPagination(r *http.Request, defaultPage, defaultPageSize int64) (int64, int64, error) {
	pageString := r.URL.Query().Get("page")
	pageSizeString := r.URL.Query().Get("pageSize")

	page, err := strconv.ParseInt(pageString, 10, 64)
	if err != nil {
		return 0, 0, handlederror.ValidationError(err.Error()).WithMessage("Page number not valid")
	}

	pageSize, err := strconv.ParseInt(pageSizeString, 10, 64)
	if err != nil {
		return 0, 0, handlederror.ValidationError(err.Error()).WithMessage("Page Size not valid")
	}

	return page, pageSize, nil
}

func GetPreviousPageLink(originalURL url.URL) (string, error) {
	queryParams := originalURL.Query()

	pageStr := queryParams.Get("page")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return "", err
	}
	if page <= 1 {
		return "", nil
	}
	page--

	queryParams.Set("page", strconv.Itoa(page))

	originalURL.RawQuery = queryParams.Encode()

	return originalURL.String(), nil
}

func GetNextPageLink(originalURL url.URL) (string, error) {
	queryParams := originalURL.Query()

	pageStr := queryParams.Get("page")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return "", err
	}
	page++

	queryParams.Set("page", strconv.Itoa(page))

	originalURL.RawQuery = queryParams.Encode()

	return originalURL.String(), nil
}
