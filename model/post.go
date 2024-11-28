package model

type Post struct {
	ID        int        `json:"id"`
	UserID    string     `json:"user_id"`
	Text      string     `json:"text"`
	ImagePath string     `json:"image_path"`
	Location  [2]float64 `json:"location"` // [x, y] 좌표
	Category  *string    `json:"category"` // NULL 가능
	Accuracy  *float64   `json:"accuracy"` // NULL 가능
	CreatedAt string     `json:"created_at"`
}

type FilteredPost struct {
	ID        int        `json:"id"`
	UserID    string     `json:"user_id"`
	Text      string     `json:"text"`
	ImagePath string     `json:"image_path"`
	Location  [2]float64 `json:"location"` // [x, y] 좌표
	Category  *string    `json:"category"` // NULL 가능
	Accuracy  *float64   `json:"accuracy"` // NULL 가능
}
