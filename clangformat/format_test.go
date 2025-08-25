package clangformat_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/go-xlan/clang-format/clangformat"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/must"
	"github.com/yyle88/osexec"
	"github.com/yyle88/rese"
)

func TestClangFormatDryRun(t *testing.T) {
	// 创建临时目录用于测试
	tempDIR := rese.V1(os.MkdirTemp("", "clang-format-test-*"))
	defer func() { must.Done(os.RemoveAll(tempDIR)) }()

	// 创建一个格式不规范的 C++ 文件（clang-format 主要用于 C/C++ 格式化）
	cppFile := filepath.Join(tempDIR, "test.cpp")
	const originalContent = `#include<iostream>
using namespace std;

int main(){
int   x=10;
string name  =  "test";
  cout<<name<<endl;
    return 0;
}`

	must.Done(os.WriteFile(cppFile, []byte(originalContent), 0644))

	// 读取原始内容并打印
	t.Log("=== 原始文件内容 ===")
	t.Log(originalContent)

	// 执行 DryRun 获取格式化后的内容
	execConfig := osexec.NewExecConfig().WithDebug()
	style := clangformat.NewStyle()

	output, err := clangformat.DryRun(execConfig, cppFile, style)
	require.NoError(t, err)
	require.NotEmpty(t, output)

	// 打印格式化后的内容
	t.Log("=== 格式化后的内容 ===")
	t.Log(string(output))

	// 打印预期结果用于对比
	t.Log("=== 预期的格式化结果 ===")
	const expectedResult = `#include <iostream>
using namespace std;

int main() {
  int x = 10;
  string name = "test";
  cout << name << endl;
  return 0;
}
`
	t.Log(expectedResult)

	// 验证格式化确实发生了变化
	require.NotEqual(t, strings.TrimSpace(originalContent), strings.TrimSpace(string(output)))

	// 验证格式化结果符合预期（去除首尾空白符进行比较）
	require.Equal(t, strings.TrimSpace(expectedResult), strings.TrimSpace(string(output)))

	// 验证原文件内容未被修改（DryRun 不应该修改文件）
	actualContent := rese.V1(os.ReadFile(cppFile))
	require.Equal(t, originalContent, string(actualContent))
}

func TestClangFormatInPlace(t *testing.T) {
	// 创建临时目录用于测试
	tempDIR := rese.V1(os.MkdirTemp("", "clang-format-inplace-test-*"))
	defer func() { must.Done(os.RemoveAll(tempDIR)) }()

	// 创建另一个格式不规范的 C++ 文件
	cppFile := filepath.Join(tempDIR, "service.cpp")
	const originalContent = `#include<vector>
#include<string>

class DataService{
public:
void processData(vector<string>&data){
for(auto&item:data){
cout<<item<<endl;
}
}
};`

	must.Done(os.WriteFile(cppFile, []byte(originalContent), 0644))

	// 读取原始内容并打印
	t.Log("=== 格式化前的文件内容 ===")
	t.Log(originalContent)

	// 执行格式化操作
	execConfig := osexec.NewExecConfig().WithDebug()
	style := clangformat.NewStyle()

	output, err := clangformat.Format(execConfig, cppFile, style)
	require.NoError(t, err)

	// Format 操作通常不返回内容（因为直接修改文件），所以 output 可能为空
	t.Log("=== Format 操作输出 ===")
	if len(output) > 0 {
		t.Log(string(output))
	} else {
		t.Log("(Format 操作无输出，已直接修改文件)")
	}

	// 读取格式化后的文件内容
	formattedContent := rese.V1(os.ReadFile(cppFile))
	t.Log("=== 格式化后的文件内容 ===")
	t.Log(string(formattedContent))

	// 打印预期结果用于对比
	t.Log("=== 预期的格式化结果 ===")
	const expectedResult = `#include <string>
#include <vector>

class DataService {
 public:
  void processData(vector<string>& data) {
    for (auto& item : data) {
      cout << item << endl;
    }
  }
};
`
	t.Log(expectedResult)

	// 验证文件内容确实被修改了
	require.NotEqual(t, strings.TrimSpace(originalContent), strings.TrimSpace(string(formattedContent)))

	// 验证格式化结果符合预期（去除首尾空白符进行比较）
	require.Equal(t, strings.TrimSpace(expectedResult), strings.TrimSpace(string(formattedContent)))
}

func TestStyleConfiguration(t *testing.T) {
	// 测试默认样式创建
	style := clangformat.NewStyle()
	require.NotNil(t, style)
	require.Equal(t, "Google", style.BasedOnStyle)
	require.Equal(t, 2, style.IndentWidth)
	require.Equal(t, 0, style.ColumnLimit)
	require.False(t, style.AlignConsecutiveAssignments)
}

func TestCustomStyleConfiguration(t *testing.T) {
	// 测试自定义样式配置
	customStyle := &clangformat.Style{
		BasedOnStyle:                "LLVM",
		IndentWidth:                 4,
		ColumnLimit:                 80,
		AlignConsecutiveAssignments: true,
	}

	// 创建临时目录和文件
	tempDIR := rese.V1(os.MkdirTemp("", "clang-format-custom-test-*"))
	defer func() { must.Done(os.RemoveAll(tempDIR)) }()

	cppFile := filepath.Join(tempDIR, "custom.cpp")
	const originalContent = `int main(){
int x=1;
int y=2;
return x+y;
}`

	must.Done(os.WriteFile(cppFile, []byte(originalContent), 0644))

	// 使用自定义样式执行 DryRun
	execConfig := osexec.NewExecConfig().WithDebug()
	output, err := clangformat.DryRun(execConfig, cppFile, customStyle)
	require.NoError(t, err)
	require.NotEmpty(t, output)

	t.Log("=== 自定义样式格式化结果 ===")
	t.Log(string(output))

	// 验证格式化发生了变化
	require.NotEqual(t, strings.TrimSpace(originalContent), strings.TrimSpace(string(output)))
}
