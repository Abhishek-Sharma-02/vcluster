package cmd

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/loft-sh/log"
	"github.com/loft-sh/utils/pkg/downloader"
	"github.com/loft-sh/utils/pkg/downloader/commands"
	vclusterctllog "github.com/loft-sh/vcluster/cmd/vclusterctl/log"
)

const DefaultHomeVClusterFolder = ".vcluster"

// GetHelmBinaryPath checks for helm binary and downloads if it's not present.
func GetHelmBinaryPath(ctx context.Context, log log.BaseLogger) (string, error) {
	logger := vclusterctllog.NewBaseLogLogger(log)

	// test for helm
	helmExecutablePath, err := exec.LookPath("helm")
	if err != nil {
		_ = fmt.Errorf("seems like helm is not installed. Helm is required for the creation of a virtual cluster")
		helmExecutablePath, err = downloader.NewDownloader(commands.NewHelmV3Command(), logger, DefaultHomeVClusterFolder).EnsureCommand(ctx)
		if err != nil {
			_ = fmt.Errorf("error while installing helm")
			return "", err
		}
	}
	return helmExecutablePath, nil
}
