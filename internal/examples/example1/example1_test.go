package example1_test

import (
	"os"
	"testing"

	"github.com/go-xlan/clang-format/internal/utils"
	"github.com/go-xlan/clang-format/protoformat"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/erero"
	"github.com/yyle88/osexec"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/runpath"
)

func TestDryRun(t *testing.T) {
	projectPath := osmustexist.ROOT(runpath.PARENT.Join("protos"))
	t.Log(projectPath)

	err := utils.WalkFilesWithExt(projectPath, ".proto", func(path string, info os.FileInfo) error {
		t.Log(path)
		output, err := protoformat.DryRun(osexec.NewExecConfig().WithDebug(), path, protoformat.NewStyle())
		if err != nil {
			return erero.Wro(err)
		}
		t.Log(string(output))
		return nil
	})
	require.NoError(t, err)
}

func TestFormatProject(t *testing.T) {
	projectPath := osmustexist.ROOT(runpath.PARENT.Join("protos"))
	t.Log(projectPath)

	style := protoformat.NewStyle()
	commandConfig := osexec.NewExecConfig().WithDebug()
	err := protoformat.FormatProject(commandConfig, projectPath, style)
	require.NoError(t, err)
}
