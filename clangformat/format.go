// Package clangformat: Clang-Format execution engine for Go applications
// Provides smart wrapper for clang-format CLI tool with structured configuration
// Supports both dry-run preview and direct file formatting operations
// Features customizable formatting styles with JSON-based configuration
//
// clangformat: Go 应用的 Clang-Format 执行引擎
// 为 clang-format CLI 工具提供智能包装，支持结构化配置
// 支持预览模式和直接文件格式化操作
// 提供基于 JSON 配置的可自定义格式化样式
package clangformat

import (
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/osexec"
)

// Style represents the configuration structure for clang-format styling options
// Contains formatting parameters that control code appearance and alignment
// JSON struct tags must match exact clang-format CLI parameter names
// Example: -style="{BasedOnStyle: Google, IndentWidth: 4, ColumnLimit: 0, AlignConsecutiveAssignments: true}"
//
// Style 代表 clang-format 样式选项的配置结构
// 包含控制代码外观和对齐的格式化参数
// JSON 标签必须与 clang-format CLI 参数名称完全匹配
// 示例: -style="{BasedOnStyle: Google, IndentWidth: 4, ColumnLimit: 0, AlignConsecutiveAssignments: true}"
type Style struct {
	BasedOnStyle                string `json:"BasedOnStyle"`                // Base style template (Google, LLVM, etc.) // 基础样式模板 (Google, LLVM 等)
	IndentWidth                 int    `json:"IndentWidth"`                 // Number of spaces for indentation (usually 2 or 4) // 缩进空格数（通常为 2 或 4）
	ColumnLimit                 int    `json:"ColumnLimit"`                 // Maximum line length limit // 最大行长度限制
	AlignConsecutiveAssignments bool   `json:"AlignConsecutiveAssignments"` // Whether to align assignments at equal signs // 是否在等号处对齐赋值
}

// NewStyle creates a default Style configuration with Google-based formatting
// Returns a style configured with 2-space indentation and no column limit
// Uses Google style as base template with conservative alignment settings
//
// NewStyle 创建默认的 Style 配置，基于 Google 格式化样式
// 返回配置了 2 空格缩进和无列限制的样式
// 使用 Google 样式作为基础模板，采用保守的对齐设置
func NewStyle() *Style {
	return &Style{
		BasedOnStyle:                "Google",
		IndentWidth:                 2,
		ColumnLimit:                 0,
		AlignConsecutiveAssignments: false,
	}
}

// DryRun executes clang-format in preview mode without modifying the target file
// Returns the formatted content as output bytes for inspection
// Useful for validating formatting changes before applying them
//
// DryRun 在预览模式下执行 clang-format，不修改目标文件
// 返回格式化内容作为输出字节供检查
// 适用于在应用更改之前验证格式化效果
func DryRun(config *osexec.ExecConfig, protoPath string, style *Style) (output []byte, err error) {
	return run(config, []string{protoPath, "-style", neatjsons.Sjson(style)})
}

// Format executes clang-format with in-place modification flag (-i)
// Applies formatting changes to the target file
// Use clang-format --help to see all available options and flags
//
// Format 使用就地修改标志 (-i) 执行 clang-format
// 直接对目标文件应用格式化更改
// 使用 clang-format --help 查看所有可用选项和标志
func Format(config *osexec.ExecConfig, protoPath string, style *Style) (output []byte, err error) {
	return run(config, []string{"-i", protoPath, "-style", neatjsons.Sjson(style)})
}

// run executes the clang-format command with specified arguments
// Requires clang-format to be installed and accessible in PATH
// Installation: brew install clang-format (macOS) or equivalent package manager
// Verification: clang-format --version
//
// run 使用指定参数执行 clang-format 命令
// 需要安装 clang-format 并在 PATH 中可访问
// 安装: brew install clang-format (macOS) 或等效的包管理器
// 验证: clang-format --version
func run(config *osexec.ExecConfig, args []string) (output []byte, err error) {
	return config.Exec("clang-format", args...)
}
