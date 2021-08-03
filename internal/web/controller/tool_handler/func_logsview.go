package tool_handler

import (
	"encoding/json"

	"exams-api/configs"
	"exams-api/internal/pkg/core"
	"exams-api/pkg/file"

	"go.uber.org/zap"
)

type logsViewResponse struct {
	Logs []logData `json:"logs"`
}

type logData struct {
	Level       string  `json:"level"`
	Time        string  `json:"time"`
	Path        string  `json:"path"`
	HTTPCode    int     `json:"http_code"`
	Method      string  `json:"method"`
	Msg         string  `json:"msg"`
	TraceID     string  `json:"trace_id"`
	Content     string  `json:"content"`
	CostSeconds float64 `json:"cost_seconds"`
}

func (h *handler) LogsView() core.HandlerFunc {

	type logParseData struct {
		Level        string  `json:"level"`
		Time         string  `json:"time"`
		Caller       string  `json:"caller"`
		Msg          string  `json:"msg"`
		Domain       string  `json:"domain"`
		Method       string  `json:"method"`
		Path         string  `json:"path"`
		HTTPCode     int     `json:"http_code"`
		BusinessCode int     `json:"business_code"`
		Success      bool    `json:"success"`
		CostSeconds  float64 `json:"cost_seconds"`
		TraceID      string  `json:"trace_id"`
	}

	return func(c core.Context) {
		readLineFromEnd, err := file.NewReadLineFromEnd(configs.ProjectLogFile())
		if err != nil {
			h.logger.Error("NewReadLineFromEnd err", zap.Error(err))
		}

		logSize := 100

		obj := new(logsViewResponse)
		obj.Logs = make([]logData, logSize)

		for i := 0; i < logSize; i++ {
			content, _ := readLineFromEnd.ReadLine()

			var logParse logParseData
			_ = json.Unmarshal(content, &logParse)
			data := logData{
				Content:     string(content),
				Level:       logParse.Level,
				Time:        logParse.Time,
				Path:        logParse.Path,
				Method:      logParse.Method,
				Msg:         logParse.Msg,
				HTTPCode:    logParse.HTTPCode,
				TraceID:     logParse.TraceID,
				CostSeconds: logParse.CostSeconds,
			}

			if string(content) != "" {
				obj.Logs[i] = data
			}
		}
		c.HTML("tool_logs", obj)
	}
}
