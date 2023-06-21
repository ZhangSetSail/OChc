package cmd

import (
	"github.com/ZhangSetSail/OChc/pkg/k8s"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"io"
)

var cmdLong = `ochc: 这是一个测试项目，致力于实现一键导出 helm chart 包，感谢使用。`

func newLogrus() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func NewCMD(out io.Writer) *cobra.Command {
	var oCmd = &cobra.Command{
		Use:   "ochc",
		Short: "test",
		Long:  cmdLong,
	}
	oCmd.AddCommand(
		newResourceCmd(out),
	)
	return oCmd
}

func init() {
	newLogrus()
	k8s.InitK8SCli()
}
