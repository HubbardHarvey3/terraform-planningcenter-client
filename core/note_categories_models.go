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
	NoteCategory           string `json:"note"`
	CreatedAt              string `json:"created_at"`
	UpdatedAt              string `json:"updated_at"`
	DisplayDate            string `json:"display_date"`
	NoteCategoryCategoryId int    `json:"note_category_id"`
}

type NoteCategoryRelationships struct {
	NoteCategoryCategory NoteCategory             `json:"note_category"`
	Organization         NoteCategoryOrganization `json:"organization"`
	Person               NoteCategoryPerson       `json:"person,omitempty"`
	CreatedBy            NoteCategoryCreatedBy    `json:"created_by"`
}

type NoteCategoryOrganization struct {
	Data NoteCategoryRelationshipData `json:"data,omitempty"`
}

type NoteCategoryPerson struct {
	Data NoteCategoryRelationshipData `json:"data,omitempty"`
}

type NoteCategoryCreatedBy struct {
	Data NoteCategoryRelationshipData `json:"data,omitempty"`
}

type NoteCategoryRelationshipData struct {
	Type string `json:"type,omitempty"`
	ID   string `json:"id,omitempty"`
}
