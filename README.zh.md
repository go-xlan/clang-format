# clang-format

clang-format 的 Go 封装工具，专为 Protocol Buffers 格式化而设计。

## 英文文档

[ENGLISH README](README.md)

## 核心特性

🎯 **智能 Proto 格式化**: 智能的 clang-format 包装器，默认使用 Google 样式  
⚡ **双重操作模式**: 支持预览（DryRun）和就地格式化两种模式  
🔄 **批量处理**: 递归的项目级 .proto 文件格式化  
🌍 **可配置样式**: 支持 JSON 配置的自定义格式化样式  
📋 **全面日志**: 详细的操作日志和结构化输出

## 安装

```bash
go install github.com/go-xlan/clang-format@latest
```

## 前置要求

在系统中安装 clang-format：

```bash
# macOS
brew install clang-format

# Ubuntu/Debian
sudo apt-get install clang-format

# 验证安装
clang-format --version
```

## 使用方法

### Protocol Buffers 格式化（主要功能）

```go
package main

import (
    "fmt"
    "github.com/go-xlan/clang-format/protoformat"
    "github.com/yyle88/must"
    "github.com/yyle88/osexec"
    "github.com/yyle88/rese"
)

func main() {
    execConfig := osexec.NewExecConfig()
    style := protoformat.NewStyle()
    
    // 预览 .proto 文件格式化 (DryRun)
    output := rese.V1(protoformat.DryRun(execConfig, "example.proto", style))
    fmt.Println(string(output))
    
    // 格式化单个 .proto 文件
	rese.V1(protoformat.Format(execConfig, "example.proto", style))
    
    // 格式化整个项目（批量处理）
    must.Done(protoformat.FormatProject(execConfig, "./proto-project", style))
}
```

### 自定义样式配置

```go
customStyle := &clangformat.Style{
    BasedOnStyle:                "LLVM",
    IndentWidth:                 4,
    ColumnLimit:                 80,
    AlignConsecutiveAssignments: true,
}

output := rese.V1(protoformat.DryRun(execConfig, "example.proto", customStyle))
```

### 通用文件格式化（C/C++ 支持）

```go
import "github.com/go-xlan/clang-format/clangformat"

// 格式化 C/C++ 文件
output := rese.V1(clangformat.DryRun(execConfig, "example.cpp", style))
must.Done(clangformat.Format(execConfig, "example.cpp", style))
```

## API 参考

### clangformat 包

- `NewStyle()` - 创建默认的基于 Google 的样式配置
- `DryRun(config, path, style)` - 预览格式化而不修改文件
- `Format(config, path, style)` - 直接对文件应用格式化

### protoformat 包

- `NewStyle()` - 创建针对 Protocol Buffers 优化的样式配置
- `DryRun(config, path, style)` - 预览 .proto 文件格式化
- `Format(config, path, style)` - 格式化单个 .proto 文件
- `FormatProject(config, path, style)` - 批量格式化项目中的所有 .proto 文件

### 样式配置

```go
type Style struct {
    BasedOnStyle                string // "Google", "LLVM", "Chromium" 等
    IndentWidth                 int    // 缩进空格数
    ColumnLimit                 int    // 最大行长度 (0 = 无限制)
    AlignConsecutiveAssignments bool   // 在等号处对齐赋值
}
```

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->

## 📄 许可证类型

MIT 许可证。详见 [LICENSE](LICENSE)。

---

## 🤝 项目贡献

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **发现问题？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **功能建议？** 创建 issue 讨论您的想法
- 📖 **文档疑惑？** 报告问题，帮助我们改进文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，帮助我们优化性能
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **意见反馈？** 欢迎所有建议和宝贵意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：为面向用户的更改更新文档，并使用有意义的提交消息
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Pull Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Pull Request 和报告问题来为此项目做出贡献。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**使用这个包快乐编程！** 🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->

---

## GitHub 标星点赞

[![Stargazers](https://starchart.cc/go-xlan/clang-format.svg?variant=adaptive)](https://starchart.cc/go-xlan/clang-format)