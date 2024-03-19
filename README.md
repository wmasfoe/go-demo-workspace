> 学习中的记录，不是笔记!
> 防止忘记，记录操作的过程

## 模块

每个模块都有 go.mod

现有模块 B

模块A 中使用 模块B 代码，应该：

```sh
go mod init example.com/module-a
# 本地引用模块
go mod edit -replace example.com/module-b=../module-b
go mod tidy
```

引用第三方模块，比如 `gin`：

```sh
go mod init example.com/module-a
go get github.com/gin-gonic/gin
go mod tidy
```

`src/main.go` 中使用 `gin` 启动了一个服务器例子。启动：

```sh
cd ./src
go run .
```
