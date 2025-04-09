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

func NewStyle() *clangformat.Style {
	return &clangformat.Style{
		BasedOnStyle:                "Google",
		IndentWidth:                 2,
		ColumnLimit:                 0,
		AlignConsecutiveAssignments: false,
	}
}

func DryRun(config *osexec.ExecConfig, protoPath string, style *clangformat.Style) (output []byte, err error) {
	return clangformat.DryRun(config, protoPath, style)
}

func Format(config *osexec.ExecConfig, protoPath string, style *clangformat.Style) (output []byte, err error) {
	return clangformat.Format(config, protoPath, style)
}

func FormatProject(config *osexec.ExecConfig, projectPath string, style *clangformat.Style) error {
	if err := utils.WalkFilesWithExt(projectPath, ".proto", func(path string, info os.FileInfo) error {
		zaplog.LOG.Debug("proto-format", zap.String("proto_path", path))
		osmustexist.MustFile(path)

		output, err := Format(config, path, style)
		if err != nil {
			return erero.Wro(err)
		}
		if len(output) > 0 {
			zaplog.LOG.Debug("proto-format", zap.String("proto_path", path), zap.ByteString("output", output))
		}
		return nil
	}); err != nil {
		return erero.Wro(err)
	}
	eroticgo.GREEN.ShowMessage("SUCCESS")
	return nil
}
