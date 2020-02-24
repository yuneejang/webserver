package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/gin-gonic/gin"

	"github.com/yuneejang/webserver/types"
	"gopkg.in/urfave/cli.v1"
)

func main() {

	app := cli.NewApp()
	app.Copyright = "Copyright 2020-2020 The go-symverse Authors"
	app.Version = "0.0.0"
	app.Name = "Web Server Test"
	app.Usage = "See README"
	app.Commands = []cli.Command{
		initCommand,
	}
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Flags = []cli.Flag{
		ConfigFilePath,
		//	DevelopFlag,
		LogFlag,
		LogPathFlag,
		EnableAPI,
	}
	app.Before = func(ctx *cli.Context) error {
		//
		cfg, err := makeConfig(ctx)
		if err != nil {
			fmt.Println(err)
		}
		setConfig(ctx, cfg)
		return nil
	}

	app.Action = func(ctx *cli.Context) error {
		run(ctx)
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}

}
func run(ctx *cli.Context) {

	//Step 1. Creates a gin router
	//- 1. 기본
	// Default With the Logger and Recovery middleware already attached
	router := gin.Default()
	//- 2.
	// Creates a router without any middleware by default
	//router := gin.New()

	//Step 2. Setup Router
	router = SetupRouter(router)

	//Step 3. Run Server
	router.Run(":8131")
	//세세한 http config설정이필요한 경우 아래와 같이 사용이 가능함
	// s := &http.Server{
	// 	Addr:           ":8080",
	// 	Handler:        router,
	// 	ReadTimeout:    10 * time.Second,
	// 	WriteTimeout:   10 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	// s.ListenAndServe()
}
func makeConfig(ctx *cli.Context) (cfg *types.Config, err error) {

	if ctx.GlobalIsSet(ConfigFilePath.Name) {
		//yaml config를 load 합니다.
		path := ctx.GlobalString(ConfigFilePath.Name)
		//logger.Info("custom config file path", ":", path)
		cfg, err = types.LoadConfigFile(path)
		if err != nil {
			//	logger.Info("ERROR", ":", err)
			return nil, err
		}
	} else {
		cfg = &types.DefaultConfig
	}
	return cfg, nil
}

func setConfig(ctx *cli.Context, conf *types.Config) {
	info := conf.Nodes[0]
	types.HttpAttach = "http://" + info.Host + ":" + info.HttpPort

}
