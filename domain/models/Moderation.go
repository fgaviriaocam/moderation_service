package models

import "time"

type Moderation struct {
	Input         string          `json:"input"`
	Ouput         string          `json:"ouput,omitempty"`
	ApplicationID string          `json:"application_id"`
	CreatedDate   time.Time       `json:"created_date,omitempty"`
	LangProccesed []LangProccesed `json:"lang_proccesed,omitempty"`
}
