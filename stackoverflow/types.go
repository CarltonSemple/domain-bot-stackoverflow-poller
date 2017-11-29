package stackoverflow

type answersContainer struct {
	Items          []Answer `json:"items"`
	HasMore        bool     `json:"has_more"`
	QuotaMax       int      `json:"quota_max"`
	QuotaRemaining int      `json:"quota_remaining"`
}

type questionsContainer struct {
	Items          []QuestionInfo `json:"items"`
	HasMore        bool           `json:"has_more"`
	QuotaMax       int            `json:"quota_max"`
	QuotaRemaining int            `json:"quota_remaining"`
}

type Answer struct {
	Owner struct {
		Reputation   int    `json:"reputation"`
		UserID       int    `json:"user_id"`
		UserType     string `json:"user_type"`
		ProfileImage string `json:"profile_image"`
		DisplayName  string `json:"display_name"`
		Link         string `json:"link"`
	} `json:"owner"`
	IsAccepted       bool   `json:"is_accepted"`
	Score            int    `json:"score"`
	LastActivityDate int    `json:"last_activity_date"`
	LastEditDate     int    `json:"last_edit_date,omitempty"`
	CreationDate     int    `json:"creation_date"`
	AnswerID         int    `json:"answer_id"`
	QuestionID       int    `json:"question_id"`
	Body             string `json:"body"`
	BodyMarkdown     string `json:"body_markdown"`
}

type QuestionInfo struct {
	Tags  []string `json:"tags"`
	Owner struct {
		Reputation   int    `json:"reputation"`
		UserID       int    `json:"user_id"`
		UserType     string `json:"user_type"`
		ProfileImage string `json:"profile_image"`
		DisplayName  string `json:"display_name"`
		Link         string `json:"link"`
	} `json:"owner"`
	IsAnswered       bool   `json:"is_answered"`
	ViewCount        int    `json:"view_count"`
	AnswerCount      int    `json:"answer_count"`
	Score            int    `json:"score"`
	LastActivityDate int    `json:"last_activity_date"`
	CreationDate     int    `json:"creation_date"`
	QuestionID       int    `json:"question_id"`
	Link             string `json:"link"`
	Title            string `json:"title"`
}
