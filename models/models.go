package models

type SocialMediaPostData struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    int    `json:"category"`
	UserID      int    `json:"userid"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type SocialMediaAddPostApiRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    int    `json:"category"`
	UserID      int    `json:"userid"`
}

type SocialMediaAddPostApiResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

type SocialMediaGetPostListApiRequest struct {
	UserID   int `json:"userid"`
	Category int `json:"category"`
}

type SocialMediaGetPostListApiResponse struct {
	Message    string                `json:"message"`
	StatusCode int                   `json:"status_code"`
	Data       []SocialMediaPostData `json:"data"`
}
