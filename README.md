# snow_flake
snow_flake算法golang实现

### 算法介绍
将64bit的二进制数字分成若干部分，每一部分存储有特定含义的数据，比如说时间戳、机器ID、序列号等等，最终生成全局唯一的有序ID。
用41位表示时间戳（大概可以支撑pow(2,41)/1000/60/60/24/365年，约等于69年），10位表示机器ID（可以继续划分为 2～3 位的 IDC 标示，可以支撑 4 个或者 8 个 IDC 机房，和7～8 位的机器 ID，支持 128-256 台机器，12位表示序列号（代表着每个节点每毫秒最多可以生成 4096 的 ID）


### 使用示例
```go
import "github.com/hhq163/snow_flake/gen_id"

node, err := gen_id.New(13)
if err != nil {
    fmt.Println(err)
    return
}
id := node.GenId()
```