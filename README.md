# go-generic 泛型定义

## 介绍

go 1.18+ 的泛型类型定义，提供常见go机器类型在数学下的定义。

以便于通过清晰的类型分类实现逻辑清晰的泛型定义

## 分层定义结构

### 1. 数字包`number`定义

数字包中提供了标准的数学数字的定义，其中包括以下类型

| 数字类型 | 全称             | 简称 |
| -------- | ---------------- | ---- |
| 自然数   | `NaturalNumber`  | `N`  |
| 整数     | `Integer`        | `Z`  |
| 有理数   | `RationalNumber` | `Q`  |
| 实数     | `RealNumber`     | `R`  |
| 复数     | `ComplexNumber`  | `C`  |

你可以使用以上常见泛型类型的全称或简称来声明泛型类型，例如下面的使用说明中对有理数类的引用

## 使用说明

### 1. 在 import 块中导入本包

像下面这样在对应的包中引用本仓库下的对应包，例如数字包 `number`。下面的例子使用了改包中定义的**有理数类型**`RationalNumber`：

```go
package main

import (
    "fmt"
	"gitee.com/ryjer/go-generic/number"
)

func add[T number.RationalNumber](a, b T) (sum T) {
	return a + b
}

func main() {
	fmt.Printf("3 + 5 = %v\n", add[int](3, 5))
}
```

对于常用的定义包，通常建议另外起一个短一些的缩写或别名，例如将 `number` 包重命名为 `num`，数字包也定义了有理数集的缩写`Q`来指代可尚的有理数

```go
package main

import (
    "fmt"
	num "gitee.com/ryjer/go-generic/number"
)

func add[T num.Q](a, b T) (sum T) {
	return a + b
}

func main() {
	fmt.Printf("3 + 5 = %v\n", add[int](3, 5))
}
```

### 2. 使用 go mod tidy 记录版本

完成上面的导入后，在你的项目仓库目录下执行以下命令，将依赖项写入 `go.mod`文件中

```bash
go mod tidy
```

## 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request