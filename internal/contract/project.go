package contract

type (
	ProjectItem struct {
		ID        int    `json:"id"`
		Title     string `json:"title"`
		HexColor  string `json:"hex_color"`
		CreatedAt string `json:"created_at"`
	}

	GetProjectsList struct {
		Data []ProjectItem `json:"data"`
	}
)
