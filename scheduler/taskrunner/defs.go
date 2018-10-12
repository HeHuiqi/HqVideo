package taskrunner


const (
	READY_TO_DISPATCH = "d"
	READY_TO_EXECUTE = "e"
	CLOSE = "c"

	VIDEO_PATH = "./videos/"
)
//定义几种类型
type controlChan chan string
type dataChan chan interface{}
type fn func(dc dataChan) error


