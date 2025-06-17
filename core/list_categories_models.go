package core

// ************ ListCategory ************

type ListCategoryRoot struct {
	Data ListCategory `json:"data,omitempty"`
}

type ListCategory struct {
	Type          string                     `json:"type"`
	ID            string                     `json:"id"`
	Attributes    ListCategoryAttributes     `json:"attributes"`
	Relationships *ListCategoryRelationships `json:"relationships,omitempty"`
}

type ListCategoryAttributes struct {
	Name           string `json:"name"`
	CreatedAt      string `json:"created_at,omitempty"`
	UpdatedAt      string `json:"updated_at,omitempty"`
	OrganizationId int    `json:"organization_id,omitempty"`
}

type ListCategoryRelationships struct {
	Organization ListCategoryOrganization `json:"organization"`
}

type ListCategoryOrganization struct {
	Data ListCategoryRelationshipOrganizationData `json:"data,omitempty"`
}

type ListCategoryRelationshipOrganizationData struct {
	Type string `json:"type,omitempty"`
	ID   string `json:"id,omitempty"`
}
