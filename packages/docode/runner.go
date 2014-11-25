package docode

import (
	"github.com/tscolari/docode/packages/dockerwrapper"
	"github.com/tscolari/docode/packages/docodeconfig"
)

type Runner struct {
	config docodeconfig.Configuration
	docker dockerwrapper.DockerWrapper
}

func NewWithWrapper(config docodeconfig.Configuration, docker dockerwrapper.DockerWrapper) *Runner {
	return &Runner{
		config: config,
		docker: docker,
	}
}

func New(config docodeconfig.Configuration) *Runner {
	return &Runner{
		config: config,
		docker: dockerwrapper.New(),
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
