package gincon

type Response struct {
	Controller     *Controller
	HttpStatusCode int
	Data           map[string]any
}

func (resp *Response) JSON(httpStatusCode int, data map[string]any) {
	resp.HttpStatusCode = httpStatusCode
	resp.Data = data
}
