package model

// ArthasRequest 请求Arthas API的结构体
type ArthasRequest struct {
	Url         string
	Action      string
	Command     string
	ExecTimeout uint64
}

// ArthasRespondBody Arthas API 返回的Body结构体
type ArthasRespondBody struct {
	Command     string                   `json:"command"`
	JobId       uint64                   `json:"jobId"`
	JobStatus   string                   `json:"jobStatus"`
	Results     []map[string]interface{} `json:"results"`
	TimeExpired bool                     `json:"timeExpired"`
}

// ArthasRespond Arthas API 返回的Json结构体 嵌套上面的Body结构体
type ArthasRespond struct {
	Body      ArthasRespondBody `json:"body"`
	SessionId string            `json:"sessionId"`
	State     string            `json:"state"`
}
