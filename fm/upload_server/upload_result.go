package upload_server

type UploadResult struct {
	Error string `json:"error"`
	Ret   int    `json:"ret"`
	Url string `json:"url"`
}
