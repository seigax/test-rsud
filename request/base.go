package request

type BasePaginateRequest struct {
	Sort     []string `json:"sort,omitempty;query:sort"`
	Search   string   `json:"search,omitempty;query:search"`
	Page     uint     `json:"page,omitempty;query:page"`
	Limit    uint     `json:"limit,omitempty;query:limit"`
	Preloads []string `json:"-"`
}

func (query *BasePaginateRequest) GetOffset() uint {
	return (query.Page - 1) * query.Limit
}
