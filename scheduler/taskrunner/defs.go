package taskrunner

const (
	READY_TO_DISPATCH = "d" //发送任务下达
	READY_TO_EXECUTE = "e" //接收者收到这个指令后就会从dataChan里面读取数据
	CLOSE = "c"
	VIDEO_PATH = "./videos/"  //本地video存储路径
)
//别名类型
//控制通道，用来发送指令
type controlChan chan string

//数据通道，交换数据，interface{}是一种泛型实现方式，但还不能算真正的泛型
type dataChan chan interface{}

//fn定义一个处理数据通道的函数类型
type fn func(dc dataChan) error