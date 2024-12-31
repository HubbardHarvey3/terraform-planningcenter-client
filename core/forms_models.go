package core

// ************ Form ************

type FormRoot struct {
	Data Form `json:"data,omitempty"`
}

type Form struct {
	Type          string             `json:"type"`
	ID            string             `json:"id"`
	Attributes    FormAttributes     `json:"attributes"`
	Relationships *FormRelationships `json:"relationships,omitempty"`
}

type FormAttributes struct {
	Active                     bool   `json:"active,omitempty"`
	Archived                   bool   `json:"archived,omitempty"`
	ArchivedAt                 string `json:"archived_at,omitempty"`
	CreatedAt                  string `json:"created_at,omitempty"`
	DeletedAt                  string `json:"deleted_at,omitempty"`
	Description                string `json:"description,omitempty"`
	LoginRequired              bool   `json:"login_required,omitempty"`
	Name                       string `json:"name,omitempty"`
	PublicUrl                  string `json:"public_url,omitempty"`
	RecentlyViewed             bool   `json:"recently_viewed,omitempty"`
	SendSumbissionNotification bool   `json:"send_submission_notification_to_submitter,omitempty"`
	SubmissionCount            int    `json:"submission_count,omitempty"`
	UpdatedAt                  string `json:"updated_at,omitempty"`
}

type FormRelationships struct {
	Campus       FormRelationshipCampus        `json:"campus"`
	FormCategory FormRelationshipsFormCategory `json:"form_category"`
}

type FormRelationshipCampus struct {
	Data FormRelationshipCampusData `json:"data,omitempty"`
}

type FormRelationshipsFormCategory struct {
	Data FormRelationshipsFormCategoryData `json:"data,omitempty"`
}

type FormRelationshipCampusData struct {
	Type string `json:"type,omitempty"`
	ID   string `json:"id,omitempty"`
}

type FormRelationshipsFormCategoryData struct {
	Type string `json:"type,omitempty"`
	ID   string `json:"id,omitempty"`
}
