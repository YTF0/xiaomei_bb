package deploy

import (
	"github.com/YTF0/xiaomei_bb/release"
	"github.com/lovego/config/config"
)

func GetCommonArgs(svcName, env, tag string) []string {
	args := []string{`-e`, config.EnvVar + `=` + env}

	service := release.GetService(env, svcName)
	args = append(args, service.Options...)
	args = append(args, service.ImageName(tag))
	args = append(args, service.Command...)
	return args
}
