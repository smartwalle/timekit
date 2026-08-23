# timekit

Go 泛型时区安全时间处理库，内部统一存储 UTC，通过泛型时区参数在编译期强制区分时区，避免 UTC / 本地时间混用问题。

## 安装

```bash
go get github.com/smartwalle/timekit
```

## 内置时区包

| 包名 | 时区（英文） | 大洲 | IANA 时区 | 偏移 |
| --- | --- | --- | --- | --- |
| `utc` | 协调世界时（UTC） | 通用 | `UTC` | +00:00 |
| `local` | 本地时区（Local） | 通用 | `time.Local` | 随系统配置 |
| `cst` | 中国标准时间（China Standard Time） | 亚洲 | `Asia/Shanghai` | +08:00 |
| `shanghai` | 中国大陆（Shanghai） | 亚洲 | `Asia/Shanghai` | +08:00 |
| `hongkong` | 中国香港（Hong Kong） | 亚洲 | `Asia/Hong_Kong` | +08:00 |
| `taipei` | 中国台湾（Taipei） | 亚洲 | `Asia/Taipei` | +08:00 |
| `tokyo` | 日本（Tokyo） | 亚洲 | `Asia/Tokyo` | +09:00 |
| `seoul` | 韩国（Seoul） | 亚洲 | `Asia/Seoul` | +09:00 |
| `singapore` | 新加坡（Singapore） | 亚洲 | `Asia/Singapore` | +08:00 |
| `bangkok` | 泰国（Bangkok） | 亚洲 | `Asia/Bangkok` | +07:00 |
| `hochiminh` | 越南（Ho Chi Minh） | 亚洲 | `Asia/Ho_Chi_Minh` | +07:00 |
| `jakarta` | 印度尼西亚（Jakarta） | 亚洲 | `Asia/Jakarta` | +07:00 |
| `kualalumpur` | 马来西亚（Kuala Lumpur） | 亚洲 | `Asia/Kuala_Lumpur` | +08:00 |
| `manila` | 菲律宾（Manila） | 亚洲 | `Asia/Manila` | +08:00 |
| `dubai` | 阿联酋（Dubai） | 亚洲 | `Asia/Dubai` | +04:00 |
| `riyadh` | 沙特阿拉伯（Riyadh） | 亚洲 | `Asia/Riyadh` | +03:00 |
| `kolkata` | 印度（Kolkata） | 亚洲 | `Asia/Kolkata` | +05:30 |
| `kathmandu` | 尼泊尔（Kathmandu） | 亚洲 | `Asia/Kathmandu` | +05:45 |
| `kabul` | 阿富汗（Kabul） | 亚洲 | `Asia/Kabul` | +04:30 |
| `tehran` | 伊朗（Tehran） | 亚洲 | `Asia/Tehran` | +03:30 |
| `yangon` | 缅甸（Yangon） | 亚洲 | `Asia/Yangon` | +06:30 |
| `london` | 英国（London） | 欧洲 | `Europe/London` | +00:00 |
| `paris` | 法国（Paris） | 欧洲 | `Europe/Paris` | +01:00 |
| `berlin` | 德国（Berlin） | 欧洲 | `Europe/Berlin` | +01:00 |
| `amsterdam` | 荷兰（Amsterdam） | 欧洲 | `Europe/Amsterdam` | +01:00 |
| `zurich` | 瑞士（Zurich） | 欧洲 | `Europe/Zurich` | +01:00 |
| `stockholm` | 瑞典（Stockholm） | 欧洲 | `Europe/Stockholm` | +01:00 |
| `warsaw` | 波兰（Warsaw） | 欧洲 | `Europe/Warsaw` | +01:00 |
| `madrid` | 西班牙（Madrid） | 欧洲 | `Europe/Madrid` | +01:00 |
| `rome` | 意大利（Rome） | 欧洲 | `Europe/Rome` | +01:00 |
| `lisbon` | 葡萄牙（Lisbon） | 欧洲 | `Europe/Lisbon` | +00:00 |
| `reykjavik` | 冰岛（Reykjavik） | 欧洲 | `Atlantic/Reykjavik` | +00:00 |
| `athens` | 希腊（Athens） | 欧洲 | `Europe/Athens` | +02:00 |
| `helsinki` | 芬兰（Helsinki） | 欧洲 | `Europe/Helsinki` | +02:00 |
| `istanbul` | 土耳其（Istanbul） | 欧洲 | `Europe/Istanbul` | +03:00 |
| `moscow` | 俄罗斯莫斯科（Moscow） | 欧洲 | `Europe/Moscow` | +03:00 |
| `casablanca` | 摩洛哥（Casablanca） | 非洲 | `Africa/Casablanca` | +01:00 |
| `cairo` | 埃及（Cairo） | 非洲 | `Africa/Cairo` | +02:00 |
| `johannesburg` | 南非（Johannesburg） | 非洲 | `Africa/Johannesburg` | +02:00 |
| `nairobi` | 肯尼亚（Nairobi） | 非洲 | `Africa/Nairobi` | +03:00 |
| `newyork` | 美国东部（New York） | 美洲 | `America/New_York` | -05:00 |
| `chicago` | 美国中部（Chicago） | 美洲 | `America/Chicago` | -06:00 |
| `denver` | 美国山地（Denver） | 美洲 | `America/Denver` | -07:00 |
| `losangeles` | 美国西部（Los Angeles） | 美洲 | `America/Los_Angeles` | -08:00 |
| `honolulu` | 美国夏威夷（Honolulu） | 美洲 | `Pacific/Honolulu` | -10:00 |
| `anchorage` | 美国阿拉斯加（Anchorage） | 美洲 | `America/Anchorage` | -09:00 |
| `toronto` | 加拿大东部（Toronto） | 美洲 | `America/Toronto` | -05:00 |
| `vancouver` | 加拿大西部（Vancouver） | 美洲 | `America/Vancouver` | -08:00 |
| `mexicocity` | 墨西哥（Mexico City） | 美洲 | `America/Mexico_City` | -06:00 |
| `saopaulo` | 巴西（Sao Paulo） | 美洲 | `America/Sao_Paulo` | -03:00 |
| `buenosaires` | 阿根廷（Buenos Aires） | 美洲 | `America/Argentina/Buenos_Aires` | -03:00 |
| `santiago` | 智利（Santiago） | 美洲 | `America/Santiago` | -04:00 |
| `sydney` | 澳大利亚东部（Sydney） | 大洋洲 | `Australia/Sydney` | +10:00 |
| `melbourne` | 澳大利亚墨尔本（Melbourne） | 大洋洲 | `Australia/Melbourne` | +10:00 |
| `perth` | 澳大利亚西部（Perth） | 大洋洲 | `Australia/Perth` | +08:00 |
| `adelaide` | 澳大利亚阿德莱德（Adelaide） | 大洋洲 | `Australia/Adelaide` | +09:30 |
| `darwin` | 澳大利亚达尔文（Darwin） | 大洋洲 | `Australia/Darwin` | +09:30 |
| `brisbane` | 澳大利亚布里斯班（Brisbane） | 大洋洲 | `Australia/Brisbane` | +10:00 |
| `auckland` | 新西兰（Auckland） | 大洋洲 | `Pacific/Auckland` | +12:00 |

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
