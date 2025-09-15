// clang-format-batch: Protocol Buffers and C/C++ batch file formatter
// Provides batch formatting workflow for multiple file types in projects
// Supports project-wide batch formatting with customizable extension lists
//
// clang-format-batch: Protocol Buffers 和 C/C++ 批量文件格式化工具
// 为项目中的多种文件类型提供批量格式化工作流程
// 支持项目范围的批量格式化，可自定义扩展名列表
package main

import (
	"os"
	"strings"

	"github.com/go-xlan/clang-format/clangformat"
	"github.com/go-xlan/clang-format/protoformat"
	"github.com/spf13/cobra"
	"github.com/yyle88/must"
	"github.com/yyle88/osexec"
	"github.com/yyle88/rese"
)

func main() {
	// Get current working DIR as default project root
	// 获取当前工作 DIR 作为默认项目根目录
	projectPath := rese.C1(os.Getwd())

	// Command line flags
	// 命令行标志
	var extensionsFlag string

	// Create and configure root command
	// 创建并配置根命令
	rootCmd := &cobra.Command{
		Use:   "clang-format-batch",
		Short: "Batch file formatter using clang-format",
		Long:  "clang-format-batch formats multiple file types with specified extensions using clang-format",
		Run: func(cmd *cobra.Command, args []string) {
			// Parse extensions from flag
			// 从标志解析扩展名
			var extensions []string
			for _, extension := range strings.Split(extensionsFlag, ",") {
				extension = strings.TrimSpace(extension)
				if extension == "" {
					continue
				}
				if !strings.HasPrefix(extension, ".") {
					extension = "." + extension
				}
				extensions = append(extensions, extension)
			}
			if len(extensions) == 0 {
				cmd.PrintErrln("ERROR: no valid extensions provided. Use --extensions to set file extensions.")
				return
			}

			// Create execution config
			// 创建执行配置
			execConfig := osexec.NewExecConfig().WithPath(projectPath)

			// Format files for each extension
			// 为每个扩展名格式化文件
			for _, extension := range extensions {
				switch extension {
				case ".proto":
					must.Done(protoformat.FormatProject(execConfig, projectPath, protoformat.NewStyle()))
				case ".c", ".cpp", ".cxx", ".cc", ".h", ".hpp", ".hxx":
					must.Done(clangformat.FormatProject(execConfig, projectPath, extension, clangformat.NewStyle()))
				default:
					cmd.PrintErrln("Warning: unsupported extension '" + extension + "', skipping")
				}
			}
		},
	}

	// Add flags
	// 添加标志
	rootCmd.Flags().StringVarP(&extensionsFlag, "extensions", "e", "", "comma-separated file extensions (e.g., .proto,.c,.cpp,.h)")

	// Execute the CLI application
	// 执行 CLI 应用程序
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
