package core

// ************ FormFields ************

type FormFieldsRoot struct {
	Data FormFields `json:"data,omitempty"`
}

type FormFields struct {
	Type          string                   `json:"type"`
	ID            string                   `json:"id"`
	Attributes    FormFieldsAttributes     `json:"attributes"`
	Relationships *FormFieldsRelationships `json:"relationships,omitempty"`
}

type FormFieldsAttributes struct {
	Label       string `json:"label,omitempty"`
	Description string `json:"description,omitempty"`
	Required    bool   `json:"required,omitempty"`
	Settings    string `json:"settings,omitempty"`
	Field_type  string `json:"field_type,omitempty"`
	Sequence    int    `json:"sequence,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

type FormFieldsRelationships struct {
	Form                FormFieldsRelationshipForm                `json:"form"`
	Fieldable           FormFieldsRelationshipFieldable           `json:"fieldable"`
	Options             FormFieldsRelationshipOptions             `json:"options"`
	FormFieldConditions FormFieldsRelationshipFormFieldConditions `json:"form_field_conditions"`
}

type FormFieldsRelationshipForm struct {
	Data FormFieldsRelationshipFormData `json:"data,omitempty"`
}

type FormFieldsRelationshipFieldable struct {
	Data FormFieldsRelationshipFieldableData `json:"data,omitempty"`
}

type FormFieldsRelationshipOptions struct {
	Data FormFieldsRelationshipOptionsData `json:"data,omitempty"`
}

type FormFieldsRelationshipFormFieldConditions struct {
	Data FormFieldsRelationshipFormFieldConditionsData `json:"data,omitempty"`
}

type FormFieldsRelationshipFormData struct {
	Type string `json:"type,omitempty"`
	ID   string `json:"id,omitempty"`
}

type FormFieldsRelationshipFieldableData struct {
	Type string `json:"type,omitempty"`
	ID   string `json:"id,omitempty"`
}

type FormFieldsRelationshipOptionsData struct {
	Type string `json:"type,omitempty"`
	ID   string `json:"id,omitempty"`
}

type FormFieldsRelationshipFormFieldConditionsData struct {
	Type string `json:"type,omitempty"`
	ID   string `json:"id,omitempty"`
}
