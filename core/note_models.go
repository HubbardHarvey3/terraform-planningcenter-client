package core

// ************ Note ************

type NoteRoot struct {
	Data Note `json:"data,omitempty"`
}

type Note struct {
	Type          string             `json:"type"`
	ID            string             `json:"id"`
	Attributes    NoteAttributes     `json:"attributes"`
	Relationships *NoteRelationships `json:"relationships,omitempty"`
}

type NoteAttributes struct {
	Note           string `json:"note"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	DisplayDate    string `json:"display_date"`
	NoteCategoryId string `json:"note_category_id"`
}

type NoteRelationships struct {
	NoteCategory NoteCategory     `json:"note_category"`
	Organization NoteOrganization `json:"organization"`
	Person       NotePerson       `json:"person,omitempty"`
	CreatedBy    NoteCreatedBy    `json:"created_by"`
}

type NoteCategory struct {
	Data NoteRelationshipData `json:"data,omitempty"`
}

type NoteOrganization struct {
	Data NoteRelationshipData `json:"data,omitempty"`
}

type NotePerson struct {
	Data NoteRelationshipData `json:"data,omitempty"`
}

type NoteCreatedBy struct {
	Data NoteRelationshipData `json:"data,omitempty"`
}

type NoteRelationshipData struct {
	Type string `json:"type,omitempty"`
	ID   string `json:"id,omitempty"`
}
