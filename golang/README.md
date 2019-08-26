# golang代码

## 使用方法

```golang
var token = ""
var version = ""
var date = ""
var savePath = ""
e := domain.Pull(token, version, date, savePath)
if e != nil {
    panic(e)
}
```
