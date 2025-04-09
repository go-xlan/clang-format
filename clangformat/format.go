package clangformat

import (
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/osexec"
)

// Style 有很多配置项，我们把它放在结构体里，这样比较明确
// -style="{BasedOnStyle: Google, IndentWidth: 4, ColumnLimit: 0, AlignConsecutiveAssignments: true, AlignConsecutiveAssignments: true}"
// 注意：结构里面的字段名和 json 标签名都不要变，确保和 clang-format 命令行需要的保持相同
type Style struct {
	BasedOnStyle                string `json:"BasedOnStyle"`
	IndentWidth                 int    `json:"IndentWidth"` //就是对齐的空格数，通常填2或者4
	ColumnLimit                 int    `json:"ColumnLimit"`
	AlignConsecutiveAssignments bool   `json:"AlignConsecutiveAssignments"` //就是赋值列是否在等号处对齐，自由选择
}

func NewStyle() *Style {
	return &Style{
		BasedOnStyle:                "Google",
		IndentWidth:                 2,
		ColumnLimit:                 0,
		AlignConsecutiveAssignments: false,
	}
}

func DryRun(config *osexec.ExecConfig, protoPath string, style *Style) (output []byte, err error) {
	return run(config, []string{protoPath, "-style", neatjsons.Sjson(style)})
}

// Format 使用命令 clang-format --help 能够看到当增加 -i 接口时，表示将结果写回到proto文件里
func Format(config *osexec.ExecConfig, protoPath string, style *Style) (output []byte, err error) {
	return run(config, []string{"-i", protoPath, "-style", neatjsons.Sjson(style)})
}

// MacOS 里安装： brew install clang-format 检查是否已安装： clang-format --version
func run(config *osexec.ExecConfig, args []string) (output []byte, err error) {
	return config.Exec("clang-format", args...)
}
