package dto

type Response struct {
	Errors []Error     `json:"errors,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

type Error struct {
	Id     string `json:"id,omitempty"`     // a unique identifier for this particular occurrence of the problem.
	Title  string `json:"title,omitempty"`  // a short, human-readable summary of the problem that SHOULD NOT change from occurrence to occurrence of the problem, except for purposes of localization.
	Detail string `json:"detail,omitempty"` // a human-readable explanation specific to this occurrence of the problem. Like title, this field’s value can be localized.
	Status string `json:"status,omitempty"` // the HTTP status code applicable to this problem, expressed as a string value.
	Code   string `json:"code,omitempty"`   // an application-specific error code, expressed as a string value.
}

type Request struct {
}
