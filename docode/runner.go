package docode

import (
	"github.com/tscolari/docode/config"
	"github.com/tscolari/docode/runtime"
)

type Runner struct {
	config config.Configuration
	docker runtime.Wrapper
}

func NewWithWrapper(config config.Configuration, docker runtime.Wrapper) *Runner {
	return &Runner{
		config: config,
		docker: docker,
	}
}

func New(configuration config.Configuration) *Runner {
	return &Runner{
		config: configuration,
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
