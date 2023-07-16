package model

type Config struct {
	UserIDList          []string `json:"user_id_list"`
	Filter              int      `json:"filter"`
	SinceDate           string   `json:"since_date"`
	EndDate             string   `json:"end_date"`
	RandomWaitPages     []int    `json:"random_wait_pages"`
	RandomWaitSeconds   []int    `json:"random_wait_seconds"`
	GlobalWait          [][]int  `json:"global_wait"`
	WriteMode           []string `json:"write_mode"`
	PicDownload         int      `json:"pic_download"`
	VideoDownload       int      `json:"video_download"`
	FileDownloadTimeout []int    `json:"file_download_timeout"`
	ResultDirName       int      `json:"result_dir_name"`
	Cookie              string   `json:"cookie"`
	MysqlConfig         struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Charset  string `json:"charset"`
	} `json:"mysql_config"`
}

var ConfigJson = Config{}
