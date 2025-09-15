// Package protoformat: Specialized Protocol Buffers formatting engine using Clang-Format
// Provides high-level interface for formatting .proto files with intelligent project traversal
// Features batch processing capabilities with configurable styling and comprehensive logging
// Built on top of clangformat package for reliable formatting operations
//
// protoformat: 使用 Clang-Format 的专用 Protocol Buffers 格式化引擎
// 为 .proto 文件格式化提供高级接口，支持智能项目遍历
// 具有批处理功能，可配置样式和全面日志记录
// 基于 clangformat 包构建，提供可靠的格式化操作
package protoformat

import (
	"os"

	"github.com/go-xlan/clang-format/clangformat"
	"github.com/go-xlan/clang-format/internal/utils"
	"github.com/yyle88/erero"
	"github.com/yyle88/eroticgo"
	"github.com/yyle88/osexec"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

// NewStyle creates a Protocol Buffers optimized Style configuration
// Returns a Google-based style tuned for .proto file formatting
// Uses conservative settings suitable for most Protocol Buffer projects
//
// NewStyle 创建针对 Protocol Buffers 优化的样式配置
// 返回专门为 .proto 文件格式化调优的 Google 样式
// 使用适合大多数 Protocol Buffer 项目的保守设置
func NewStyle() *clangformat.Style {
	return &clangformat.Style{
		BasedOnStyle:                "Google",
		IndentWidth:                 2,
		ColumnLimit:                 0,
		AlignConsecutiveAssignments: false,
	}
}

// DryRun performs a preview formatting operation on a Protocol Buffer file
// Returns formatted content without modifying the original file
// Wraps clangformat.DryRun with proto-specific context
//
// DryRun 对 Protocol Buffer 文件执行预览格式化操作
// 返回格式化内容而不修改原始文件
// 在 proto 特定上下文中包装 clangformat.DryRun
func DryRun(config *osexec.ExecConfig, protoPath string, style *clangformat.Style) (output []byte, err error) {
	return clangformat.DryRun(config, protoPath, style)
}

// Format applies formatting changes to a Protocol Buffer file
// Modifies the target .proto file in-place with the specified style
// Wraps clangformat.Format with proto-specific context
//
// Format 直接对 Protocol Buffer 文件应用格式化更改
// 使用指定样式就地修改目标 .proto 文件
// 在 proto 特定上下文中包装 clangformat.Format
func Format(config *osexec.ExecConfig, protoPath string, style *clangformat.Style) (output []byte, err error) {
	return clangformat.Format(config, protoPath, style)
}

// FormatProject performs batch formatting operation on all .proto files in a project
// Discovers and formats all Protocol Buffer files within the specified path
// Provides detailed logging and validation with success feedback upon completion
// Uses intelligent file traversal to handle complex project structures
//
// FormatProject 对项目中的所有 .proto 文件执行批量格式化操作
// 递归发现并格式化指定路径内的所有 Protocol Buffer 文件
// 提供详细日志和验证，完成时给出成功反馈
// 使用智能文件遍历处理复杂的项目结构
func FormatProject(config *osexec.ExecConfig, projectPath string, style *clangformat.Style) error {
	if err := utils.WalkFilesWithExt(projectPath, ".proto", func(path string, info os.FileInfo) error {
		zaplog.LOG.Debug("proto-format", zap.String("path", path))
		osmustexist.MustFile(path)

		output, err := Format(config, path, style)
		if err != nil {
			return erero.Wro(err)
		}
		if len(output) > 0 {
			zaplog.LOG.Debug("proto-format", zap.String("path", path), zap.ByteString("output", output))
		}
		return nil
	}); err != nil {
		return erero.Wro(err)
	}
	eroticgo.GREEN.ShowMessage("SUCCESS")
	return nil
}
