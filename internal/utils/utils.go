// Package utils: Internal utility functions for file system operations and navigation
// Provides specialized file walking capabilities with extension filtering
// Designed for internal use within the clang-format project ecosystem
// Supports recursive path navigation with custom processing callbacks
//
// utils: 文件系统操作和导航的内部工具函数
// 提供带扩展名过滤的专用文件遍历功能
// 专为 clang-format 项目生态系统内部使用而设计
// 支持带自定义处理回调的递归路径导航
package utils

import (
	"os"
	"path/filepath"
)

// WalkFilesWithExt traverses a file structure and processes files with matching extensions
// Executes the provided run function on each file that matches the specified extension
// Skips paths, handles errors with care, and validates file info before processing
// Returns any error encountered during navigation or callback execution
//
// WalkFilesWithExt 遍历文件结构并处理匹配扩展名的文件
// 对每个匹配指定扩展名的文件执行提供的 run 函数
// 跳过路径，细心处理错误，处理前验证文件信息
// 返回导航或回调执行期间遇到的任何错误
func WalkFilesWithExt(root string, extension string, run func(path string, info os.FileInfo) error) (err error) {
	err = filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info == nil {
				return nil
			}
			if info.IsDir() {
				return nil
			}
			if filepath.Ext(path) == extension {
				return run(path, info)
			}
			return nil
		},
	)
	return err
}
