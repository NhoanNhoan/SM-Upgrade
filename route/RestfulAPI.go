package route

type RestfulAPI interface {
	DoGet()
	DoPost()
	DoPut()
	DoDelete()
}
