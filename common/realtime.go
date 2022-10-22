package common

type RealtimeEvent struct {
	EventID string `json:"event_id" db:"event_id"`
	// Who
	AccountID string `json:"account_id" db:"account_id"`

	// What
	ProductCode   string `json:"product_code" db:"product_code"`
	ProductName   string `json:"product_name" db:"product_name"`
	ProductRegion string `json:"product_region" db:"product_region"`
	ResourceID    string `json:"resource_id" db:"resource_id"`

	// When
	CreatedTime string `json:"created_time" db:"created_time"`
	// How
	UsageType string `json:"usage_type" db:"usage_type"`
	Operation string `json:"operation" db:"operation"`

	// Other meta
	UsedByTag string `json:"user_component_tag" db:"used_by"`
}

type QueryRealtimeEventOpts struct {
	ProductCode string `json:"product"`

	Offset, Limit uint
}
