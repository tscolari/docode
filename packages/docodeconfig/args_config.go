package docodeconfig

type ArgsConfiguration struct {
	ImageName *string
	ImageTag  *string
	Ports     *map[int]int
	RunList   *[]string
	SSHKey    *string
	DontPull  *bool
	EnvSets   *map[string]string
	MountSets *map[string]string
}
