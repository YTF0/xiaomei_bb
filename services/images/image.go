package images

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/YTF0/xiaomei_bb/misc"
	"github.com/YTF0/xiaomei_bb/release"
	"github.com/YTF0/xiaomei_bb/services/images/app"
	"github.com/fatih/color"
	"github.com/lovego/cmd"
	"github.com/lovego/config/config"
	"github.com/lovego/fs"
)

type Image struct {
	svcName string
}

func Get(svcName string) Image {
	switch svcName {
	case "app":
		return Image{svcName: `app`}
	case "web":
		return Image{svcName: `web`}
	case "logc":
		return Image{svcName: `logc`}
	default:
		panic(`no image for: ` + svcName)
	}
}

// 1. port env variable name
func (i Image) PortEnvVar() string {
	switch i.svcName {
	case "app", "web":
		return "ProPORT"
	default:
		return ""
	}
}

// 2. default port number
func (i Image) DefaultPort() uint16 {
	switch i.svcName {
	case "app":
		return 3000
	case "web":
		return 8000
	default:
		return 0
	}
}

// 3. flags for run
func (i Image) FlagsForRun(env string) []string {
	switch i.svcName {
	case "app":
		return []string{`-e=ProDEV=true`}
	case "web":
		return []string{
			`-e=SendfileOff=true`,
			fmt.Sprintf("-v=%s/public:/var/www/%s", release.ServiceDir(`web`), release.Name(env)),
		}
	default:
		return nil
	}
}

// 4. prepare files for build
func (i Image) prepare(env, svcDir string, flags []string) error {
	if prepare := svcDir + "/prepare"; fs.Exist(prepare) {
		environment := config.NewEnv(env)
		if _, err := cmd.Run(cmd.O{
			Dir: svcDir, Env: environment.Vars(), Print: true,
		}, prepare, flags...); err != nil {
			return err
		}
	}

	switch i.svcName {
	case "app":
		return app.Compile(true, env, flags)
	case "web":
		tmpls, err := filepath.Glob(svcDir + "/*.tmpl")
		if err != nil {
			return err
		}
		for _, tmpl := range tmpls {
			output := strings.TrimSuffix(tmpl, ".tmpl")
			log.Printf(
				"%s %s %s %s\n.",
				color.GreenString(`render`), tmpl,
				color.GreenString(`to`), output,
			)
			if err := misc.RenderFileWithEnvConfig(env, tmpl, output); err != nil {
				return err
			}
		}
	}

	return nil
}
