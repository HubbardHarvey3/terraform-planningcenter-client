package core

// ************ NoteCategory ************

type NoteCategoryRoot struct {
	Data NoteCategory `json:"data,omitempty"`
}

type NoteCategory struct {
	Type          string                     `json:"type"`
	ID            string                     `json:"id"`
	Attributes    NoteCategoryAttributes     `json:"attributes"`
	Relationships *NoteCategoryRelationships `json:"relationships,omitempty"`
}

type NoteCategoryAttributes struct {
	Name                   string `json:"name"`
	CreatedAt              string `json:"created_at,omitempty"`
	UpdatedAt              string `json:"updated_at,omitempty"`
	Locked                 bool   `json:"locked,omitempty"`
	NoteCategoryCategoryId int    `json:"note_category_id,omitempty"`
	OrganizationId         int    `json:"organization_id,omitempty"`
}

type NoteCategoryRelationships struct {
	Organization NoteCategoryOrganization `json:"organization"`
}

type NoteCategoryOrganization struct {
	Data NoteCategoryRelationshipOrganizationData `json:"data,omitempty"`
}

type NoteCategoryRelationshipOrganizationData struct {
	Type string `json:"type,omitempty"`
	ID   string `json:"id,omitempty"`
}
