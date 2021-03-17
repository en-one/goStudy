package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Print(i)
		}(i)
	}
	time.Sleep(time.Millisecond * 50)

}

/*go基本调度机制GPM

M系统线程
P processor go实现的携程处理器
G goroutine

每一个P都有一个g的队列（协程处理队列），以及一个m真正的内核线程（工作）

如果创建一个goroutine运行，这个goroutine会被放到调度器的全局运行队列中，
之后，调度器江浙写队列中的goroutine分配给P（逻辑处理器），并放到这个P对应的本地运行队列中
队列中的goroutine等着P执行


如果有一个g处理时间很长，会不会导致队列里的其他在等待

	1、go创建的时候，会有一个守护线程，做一个计数，记录每一个p完成的协程数量，一段时间内，数量没有变化，会向这个协程任务栈里面插入一个特别的标记，当协程运行，
遇到非内联函数就会读到这个标记，中断自己，插到等候协程队列的队尾，切换成别的协程继续运行
	2、当某一个协程被系统中断（io），需要等待的时候，M和g就会分离出去执行，调度器给该P创建一个新M，绑定到P

*/
