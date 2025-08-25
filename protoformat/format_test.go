package protoformat_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/go-xlan/clang-format/protoformat"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/must"
	"github.com/yyle88/osexec"
	"github.com/yyle88/rese"
)

func TestProtoFormatDryRun(t *testing.T) {
	// 创建临时目录用于测试
	tempDIR := rese.V1(os.MkdirTemp("", "proto-format-test-*"))
	defer func() { must.Done(os.RemoveAll(tempDIR)) }()

	// 创建一个格式不规范的 proto 文件
	protoFile := filepath.Join(tempDIR, "test.proto")
	const originalContent = `syntax = "proto3";

package test;

option go_package = "test/proto;test";

message User{
int32    id=1;
string name  =  2;
  string email=3;
}

enum Status{
UNKNOWN=0;
  ACTIVE = 1;
    INACTIVE=2;
}`

	must.Done(os.WriteFile(protoFile, []byte(originalContent), 0644))

	// 读取原始内容并打印
	t.Log("=== 原始文件内容 ===")
	t.Log(originalContent)

	// 执行 DryRun 获取格式化后的内容
	execConfig := osexec.NewExecConfig().WithDebug()
	style := protoformat.NewStyle()

	output, err := protoformat.DryRun(execConfig, protoFile, style)
	require.NoError(t, err)
	require.NotEmpty(t, output)

	// 打印格式化后的内容
	t.Log("=== 格式化后的内容 ===")
	t.Log(string(output))

	// 打印预期结果用于对比
	t.Log("=== 预期的格式化结果 ===")
	const expectedResult = `syntax = "proto3";

package test;

option go_package = "test/proto;test";

message User {
  int32 id = 1;
  string name = 2;
  string email = 3;
}

enum Status {
  UNKNOWN = 0;
  ACTIVE = 1;
  INACTIVE = 2;
}
`
	t.Log(expectedResult)

	// 验证格式化确实发生了变化
	require.NotEqual(t, strings.TrimSpace(originalContent), strings.TrimSpace(string(output)))

	// 验证格式化结果符合预期（去除首尾空白符进行比较）
	require.Equal(t, strings.TrimSpace(expectedResult), strings.TrimSpace(string(output)))

	// 验证原文件内容未被修改（DryRun 不应该修改文件）
	actualContent := rese.V1(os.ReadFile(protoFile))
	require.Equal(t, originalContent, string(actualContent))
}

func TestProtoFormatInPlace(t *testing.T) {
	// 创建临时目录用于测试
	tempDIR := rese.V1(os.MkdirTemp("", "proto-format-inplace-test-*"))
	defer func() { must.Done(os.RemoveAll(tempDIR)) }()

	// 创建另一个格式不规范的 proto 文件
	protoFile := filepath.Join(tempDIR, "service.proto")
	const originalContent = `syntax="proto3";

package service;

option go_package="service/proto;service";

message Request {
int32 page=1;
int32 size=2;
repeated string filters  =3;
}

message Response{
repeated string data=1;
  int32 total_count   = 2;
}

service TestService {
rpc GetData(Request)returns(Response);
}`

	must.Done(os.WriteFile(protoFile, []byte(originalContent), 0644))

	// 读取原始内容并打印
	t.Log("=== 格式化前的文件内容 ===")
	t.Log(originalContent)

	// 执行格式化操作
	execConfig := osexec.NewExecConfig().WithDebug()
	style := protoformat.NewStyle()

	output, err := protoformat.Format(execConfig, protoFile, style)
	require.NoError(t, err)

	// Format 操作通常不返回内容（因为直接修改文件），所以 output 可能为空
	t.Log("=== Format 操作输出 ===")
	if len(output) > 0 {
		t.Log(string(output))
	} else {
		t.Log("(Format 操作无输出，已直接修改文件)")
	}

	// 读取格式化后的文件内容
	formattedContent := rese.V1(os.ReadFile(protoFile))
	t.Log("=== 格式化后的文件内容 ===")
	t.Log(string(formattedContent))

	// 打印预期结果用于对比
	t.Log("=== 预期的格式化结果 ===")
	const expectedResult = `syntax = "proto3";

package service;

option go_package = "service/proto;service";

message Request {
  int32 page = 1;
  int32 size = 2;
  repeated string filters = 3;
}

message Response {
  repeated string data = 1;
  int32 total_count = 2;
}

service TestService {
  rpc GetData(Request) returns (Response);
}
`
	t.Log(expectedResult)

	// 验证文件内容确实被修改了
	require.NotEqual(t, strings.TrimSpace(originalContent), strings.TrimSpace(string(formattedContent)))

	// 验证格式化结果符合预期（去除首尾空白符进行比较）
	require.Equal(t, strings.TrimSpace(expectedResult), strings.TrimSpace(string(formattedContent)))
}
