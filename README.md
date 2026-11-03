# timekit

Go 泛型时区安全时间处理库，内部统一存储 UTC，通过泛型时区参数在编译期强制区分时区，避免 UTC / 本地时间混用问题。

## 安装

```bash
go get github.com/smartwalle/timekit
```

## 内置时区包

| 包名 | 中文名称 | IANA 时区 | 标准偏移 | 夏令时偏移 |
| --- | --- | --- | --- | --- |
| `utc` | 协调世界时（UTC） | `UTC` | +00:00 | — |
| `local` | 本地时区 | `time.Local` | 随系统 | 随系统 |
| `cst` | 中国标准时间（北京时间） | `Asia/Shanghai` | +08:00 | 无 |
| `shanghai` | 中国上海 | `Asia/Shanghai` | +08:00 | 无 |
| `hongkong` | 中国香港 | `Asia/Hong_Kong` | +08:00 | 无 |
| `taipei` | 中国台湾 | `Asia/Taipei` | +08:00 | 无 |
| `singapore` | 新加坡 | `Asia/Singapore` | +08:00 | 无 |
| `tokyo` | 日本东京 | `Asia/Tokyo` | +09:00 | 无 |
| `london` | 英国伦敦 | `Europe/London` | +00:00 | +01:00 |
| `newyork` | 美国纽约（美东） | `America/New_York` | -05:00 | -04:00 |

> `cst` 与 `shanghai` 指向同一个 IANA 时区，但属于两个互不兼容的类型，跨包赋值需要用 `As` 转换。
>
> 标准偏移为冬令时偏移，夏令时偏移仅在对应时区实行夏令时的期间生效（如 `Europe/London` 为 UTC+0 / UTC+1）。

## 自定义时区

内置时区未覆盖的地区，可以直接在项目中按下面的模板声明自己的时区类型，同样享受编译期的时区区分：

```go
package myzone

import (
	"time"

	"github.com/smartwalle/timekit"
)

var location = timekit.NewLocation("Asia/Kathmandu", 5*60*60+45*60)

type Timezone struct{}

func (Timezone) Location() *time.Location {
	return location
}

type Time = timekit.Time[Timezone]
```

如果不需要编译期区分，直接使用标准库即可：`time.LoadLocation("Europe/Paris")`。

## 快速开始

```go
package main

import (
    "fmt"

    "github.com/smartwalle/timekit/cst"
)

func main() {
    now := cst.Now()
    fmt.Println(now.Format("2006-01-02 15:04:05"))
}
```
