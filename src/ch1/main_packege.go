package main

//不允许倒入没有使用的包
import (
	"github.com/goinaction/code/chapter2/sample/search"
	"log"
	"os"
)

func init() {
	//日志输出到标准输出
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
