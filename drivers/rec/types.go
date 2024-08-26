package _rec

// py程序的响应体
type RespFromPy struct {
	Token string `json:"token"`
}

// Entity
type Entity map[string]string

// list接口响应体
type ListResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Entity     FilesEntity `json:"entity"`
}

// download接口响应体
type downloadResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Entity     Entity `json:"entity"`
}

// 文件结构
type Files struct {
	CreaterUserNumber   string `json:"creater_user_number"`
	CreaterUserRealName string `json:"creater_user_real_name"`
	CreaterUserAvatar   string `json:"creater_user_avatar"`
	Number              string `json:"number"`
	ParentNumber        string `json:"parent_number"`
	DiskType            string `json:"disk_type"`
	IsHistory           bool   `json:"is_history"`
	Name                string `json:"name"`
	Type                string `json:"type"`
	FileExt             string `json:"file_ext"`
	FileType            string `json:"file_type"`
	Bytes               int64  `json:"bytes"`
	Hash                string `json:"hash"`
	TranscodeStatus     string `json:"transcode_status"`
	IsStar              bool   `json:"is_star"`
	IsLock              bool   `json:"is_lock"`
	LockReason          string `json:"lock_reason"`
	ShareCount          int    `json:"share_count"`
	LastUpdateDate      string `json:"last_update_date"`
	ParentPathNumber    string `json:"parent_path_number"`
	ReviewStatus        string `json:"review_status"`
	Version             int    `json:"version"`
}

// 文件列表data
type FilesEntity struct {
	Total                 int     `json:"total"`
	Datas                 []Files `json:"datas"`
	OperationSeriesNumber int     `json:"operation_series_number"`
}

// 下载请求body
type DownLoadBody struct {
	FilesList   []string `json:"files_list"`
	GroupNumber string   `json:"group_number"`
}
