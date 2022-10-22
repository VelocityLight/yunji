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

type DetailBilling struct {
	AccountID     string `json:"line_item_usave_account_id" db:"line_item_usave_account_id"`
	ProductCode   string `json:"product_code" db:"line_item_product_code"`
	ProductName   string `json:"product_name" db:"product_product_name"`
	ProductRegion string `json:"product_region" db:"product_region_code"`
	ResourceID    string `json:"resource_id" db:"line_item_resource_id"`
	// How
	UsageType string `json:"usage_type" db:"line_item_usage_type"`
	Operation string `json:"operation" db:"line_item_operation"`

	// Other meta
	UsedByTag string `json:"used_by" db:"resource_tags_user_usedby"`
}
