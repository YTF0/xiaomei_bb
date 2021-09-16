package deploy

import (
	"fmt"

	"github.com/YTF0/xiaomei_bb/release"
	"github.com/lovego/cmd"
)

func rmDeploy(svcName, env, feature string) error {
	script := fmt.Sprintf(`
for name in $(docker ps -af name='%s' --format '{{.Names}}'); do
	docker stop $name >/dev/null 2>&1 && docker rm $name
done
`, release.ServiceNameRegexp(svcName, env))
	return eachNodeRun(env, script, feature)
}

func eachNodeRun(env, script, feature string) error {
	for _, node := range release.GetCluster(env).GetNodes(feature) {
		if _, err := node.Run(cmd.O{}, script); err != nil {
			return err
		}
	}
	return nil
}
