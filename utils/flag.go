package utils

import (
	"fmt"

	"github.com/yuneejang/webserver/config"
	cli "gopkg.in/urfave/cli.v1"
)

//cmd
var (
	//TODO: To setup database, make init cmd
	InitCommand = cli.Command{
		Action: MigrateFlags(initServer),
		Name:   "init",
		Usage:  "Initialize a new Server",
		Flags: []cli.Flag{
			ConfigFilePath,
		},
		Category: "Initailizing Server.",
		Description: `
The init command initializes a new Server.
It will do something in the future`,
	}
)

//global flags
var (
	ConfigFilePath = cli.StringFlag{
		Name:  "config",
		Usage: "실행에 필요한 config 파일을 지정합니다. 자세한 옵션 설명은 README.md 파일을 참조해주세요.",
	}
	LogFlag = cli.BoolFlag{
		Name:  "log",
		Usage: "`Log 를 파일로 남깁니다.",
	}
	LogPathFlag = cli.StringFlag{
		Name:  "logpath",
		Usage: "`Log 를 기록할 경로입니다.",
	}
	// DevelopFlag = cli.BoolFlag{
	// 	Name:  "develop",
	// 	Usage: "`develop` 모드로 실행합니다. table 자동생성됩니다. (해당 flag를 사용하지 않으면 기본적으로 production 모드로 실행됩니다.) ",
	// }
	EnableAPI = cli.BoolFlag{
		Name:  "enableapi",
		Usage: "Enable Vote Server APIs",
	}
)

func initServer(ctx *cli.Context) error {
	fmt.Println("TODO : ? ")
	return nil
}

// MigrateFlags sets the global flag from a local flag when it's set.
// This is a temporary function used for migrating old command/flags to the
// new format.
//
// e.g. gsym account new --keystore /tmp/mykeystore --lightkdf
//
// is equivalent after calling this method with:
//
// gsym --keystore /tmp/mykeystore --lightkdf account new
//
// This allows the use of the existing configuration functionality.
// When all flags are migrated this function can be removed and the existing
// configuration functionality must be changed that is uses local flags
func MigrateFlags(action func(ctx *cli.Context) error) func(*cli.Context) error {
	return func(ctx *cli.Context) error {
		for _, name := range ctx.FlagNames() {
			if ctx.IsSet(name) {
				ctx.GlobalSet(name, ctx.String(name))
			}
		}
		return action(ctx)
	}
}

func MakeConfig(ctx *cli.Context) (cfg *config.Config, err error) {

	if ctx.GlobalIsSet(ConfigFilePath.Name) {
		//yaml config를 load 합니다.
		path := ctx.GlobalString(ConfigFilePath.Name)
		//logger.Info("custom config file path", ":", path)
		cfg, err = config.LoadConfigFile(path)
		if err != nil {
			//	logger.Info("ERROR", ":", err)
			return nil, err
		}
	} else {
		cfg = &config.DefaultConfig
	}
	return cfg, nil
}
