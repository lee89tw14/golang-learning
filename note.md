package 包名
以main為主包去呼叫其他包

初始化項目 go mod init "名稱"
用於管理項目

同文件調用命令時首字母小寫，反之大寫

以下為簡易golang格式
```go
import "fmt"

func sayHelloWorld() {
    fmt.Println("hello world")
}

func main(){
    sayHelloworld()
}
```

