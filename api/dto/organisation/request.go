package dto

type CreateOrgRequestModel struct {
	Name        string      `json:"name"`
	DisplayName string      `json:"display_name"`
	// OrgMetadata OrgMetadata `json:"metadata"`
}