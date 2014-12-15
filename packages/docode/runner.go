package docode

import (
	"github.com/tscolari/docode/packages/docodeconfig"
	"github.com/tscolari/docode/packages/runtime"
)

type Runner struct {
	config docodeconfig.Configuration
	docker runtime.Wrapper
}

func NewWithWrapper(config docodeconfig.Configuration, docker runtime.Wrapper) *Runner {
	return &Runner{
		config: config,
		docker: docker,
	}
}

func New(config docodeconfig.Configuration) *Runner {
	return &Runner{
		config: config,
		docker: runtime.New(),
	}
}

func (r *Runner) Run() error {
	if !r.config.DontPull {
		err := r.docker.PullImage(r.config.ImageName, r.config.ImageTag)
		if err != nil {
			return err
		}
	}

	return r.docker.Run(
		r.config.RunList,
		r.config.Ports,
		r.config.ImageName,
		r.config.ImageTag,
		r.config.SSHKey,
		r.config.EnvSets,
		r.config.MountSets,
	)
}
