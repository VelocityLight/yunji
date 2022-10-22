package common

type Billing struct {
	Product string  `json:"product" db:"line_item_product_code"`
	Cost    float64 `json:"cost" db:"line_item_unblended_cost"`
	UsedBy  string  `json:"used_by" db:"resource_tags_user_usedby"`
	Tag     string  `json:"tag" db:"resource_tags_user_component"`
}

type QueryBillingOpts struct {
	Tag     string `json:"tag"`
	UsedBy  string `json:"used_by"`
	Product string `json:"product"`

	Offset, Limit uint
}

type GetCostByTeamResponse struct {
	Team string  `json:"team" db:"team"`
	Cost float64 `json:"cost" db:"cost"`
}
