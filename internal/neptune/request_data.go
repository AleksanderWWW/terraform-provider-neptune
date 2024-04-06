package neptune

type requestData struct {
	headers  map[string]string
	params   map[string]string
	bodyJson []byte
}

func newRequestData(headers, params map[string]string, bodyJson []byte) *requestData {
	return &requestData{
		headers:  headers,
		params:   params,
		bodyJson: bodyJson,
	}
}
