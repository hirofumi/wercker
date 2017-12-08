package cmd

import (
	"github.com/wercker/wercker/core"
	stepscmd "github.com/wercker/wercker/steps/cmd"
)

func cmdStepPublish(opts *core.WerckerStepOptions) error {
	publishOpts := &stepscmd.PublishStepOptions{
		Endpoint:  opts.StepRegistryURL,
		AuthToken: opts.AuthToken,
		Owner:     opts.Owner,
		StepDir:   ".",
		TempDir:   "",
	}
	return stepscmd.PublishStep(publishOpts)
}
