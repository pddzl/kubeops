package request

type DynamicResource struct {
	Content string `json:"content" validated:"required"`
}
