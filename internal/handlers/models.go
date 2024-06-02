package handlers

type ConnectionInfoModel struct {
	Address  string `json:"address"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
	Status   string `json:"status"`
}
