package logs

type Filter struct {
	Page        int64  `json:"page"`
	PageSize    int64  `json:"page_size"`
	ServiceName string `json:"service_name"`
	Search      string `json:"search"`
}

func (f *Filter) GetPagination() (bool, int64, int64) {
	if f.Page == 0 {
		return false, 0, 0
	}

	if f.PageSize == 0 {
		return true, f.Page, 10
	}

	return true, f.Page, f.PageSize
}
