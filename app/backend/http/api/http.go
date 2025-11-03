package server

type IHTTPServer interface {
	Run()
}

type HTTPServer struct{}

func (h HTTPServer) Run() {}
