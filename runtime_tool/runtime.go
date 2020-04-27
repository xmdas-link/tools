package runtime_tool

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"time"
)

// 返回当前函数命
func PrintFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

// 返回调用函数名
func PrintCallerName() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}

// panic 日志
func PanicDump() {
	if err := recover(); err != nil {
		exeName := os.Args[0] //获取程序名称

		now := time.Now()  //获取当前时间
		pid := os.Getpid() //获取进程ID

		timeStr := now.Format("20060102150405")                          //设定时间格式
		fName := fmt.Sprintf("%s-%d-%s-dump.log", exeName, pid, timeStr) //保存错误信息文件名:程序名-进程ID-当前时间（年月日时分秒）
		fmt.Println("dump to file ", fName)

		f, fileErr := os.Create(fName)
		if fileErr != nil {
			fmt.Printf("create dump file fail. %v, %v", fName, fileErr)
			return
		}
		defer f.Close()
		f.WriteString(fmt.Sprintf("%v\r\n", err)) //输出panic信息
		f.WriteString("========\r\n")
		f.WriteString(string(debug.Stack())) //输出堆栈信息
	}
}
