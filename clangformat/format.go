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
	"os"

	"github.com/go-xlan/clang-format/internal/utils"
	"github.com/yyle88/erero"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/osexec"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
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

// FormatProject executes clang-format on files with specified extension in a project directory
// Walks through the project structure and formats all matching source files
// Takes a single extension parameter to process one file type at a time
// Returns error if any formatting operation fails during project navigation
//
// FormatProject 对项目目录中指定扩展名的文件执行 clang-format
// 遍历项目结构并格式化所有匹配的源文件
// 接受单个扩展名参数，一次处理一种文件类型
// 如果在项目导航过程中任何格式化操作失败则返回错误
func FormatProject(config *osexec.ExecConfig, projectPath string, extension string, style *Style) error {
	if err := utils.WalkFilesWithExt(projectPath, extension, func(path string, info os.FileInfo) error {
		zaplog.LOG.Debug("clang-format", zap.String("path", path))
		osmustexist.MustFile(path)

		output, err := Format(config, path, style)
		if err != nil {
			return erero.Wro(err)
		}
		if len(output) > 0 {
			zaplog.LOG.Debug("clang-format", zap.String("path", path), zap.ByteString("output", output))
		}
		return nil
	}); err != nil {
		return erero.Wro(err)
	}
	return nil
}
