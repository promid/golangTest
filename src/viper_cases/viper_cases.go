package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	// 定义并解析命令行参数
	var configFile string
	flag.StringVar(&configFile, "config", "", "path to config file")
	flag.Int("port", 8080, "server port")
	flag.String("hostname", "example.com", "server hostname")

	flag.Parse()

	// 将 flag 包中的命令行参数绑定到 Viper
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		panic(err)
	}

	// 如果提供了配置文件，则读取它
	if configFile != "" {
		viper.SetConfigFile(configFile)
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file, %s", err)
		}
	}

	// 设置环境变量前缀，用于读取相关环境变量
	viper.SetEnvPrefix("MYAPP")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// 使用 Viper 获取配置信息

	fmt.Println("Port:", viper.GetInt("port"))
	fmt.Println("Hostname:", viper.GetString("hostname"))
	fmt.Println("Feature 1 Enabled:", viper.GetBool("features.feature1"))
	fmt.Println("Feature 2 Enabled:", viper.GetBool("features.feature2"))

	// go run src/viper_cases/viper_cases.go --port 9090 --hostname "myhost.com"

	/*
		Port: 9090
		Hostname: myhost.com
		Feature 1 Enabled: false
		Feature 2 Enabled: false
	*/

	// export MYAPP_port=9091
	// go run src/viper_cases/viper_cases.go --hostname "myhost.com"

	/*
		Port: 9091
		Hostname: myhost.com
		Feature 1 Enabled: false
		Feature 2 Enabled: false
	*/
}
