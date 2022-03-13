package main

import (
	"embed"
	// "runtime"
	"github.com/YunFreeTech/kubefree/internal/route"
	"github.com/YunFreeTech/kubefree/internal/server"
	"github.com/YunFreeTech/kubefree/pkg/network/ip"
	"github.com/spf13/cobra"
)

var (
	configPath     string
	serverBindHost string
	serverBindPort int
)

//go:embed web/kubefree
var embedWebKubeFree embed.FS

//go:embed web/dashboard
var embedWebDashboard embed.FS

//go:embed web/terminal
var embedWebTerminal embed.FS

//go:embed script/darwin/init-kube.sh
var webkubectlEntrypointDarwin string

//go:embed script/linux/init-kube.sh
var webkubectlEntrypointLinux string

//go:embed helper/ip/qqwry.dat
var IpCommonDictionary []byte

func init() {
	RootCmd.Flags().StringVar(&serverBindHost, "server-bind-host", "", "kubefree bind address")
	RootCmd.Flags().IntVar(&serverBindPort, "server-bind-port", 0, "kubefree bind port")
	RootCmd.Flags().StringVarP(&configPath, "config-path", "c", "", "config file path")
}

var RootCmd = &cobra.Command{
	Use:   "kubefree-server",
	Short: "A dashboard for kubernetes",
	RunE: func(cmd *cobra.Command, args []string) error {
		server.EmbedWebDashboard = embedWebDashboard
		server.EmbedWebTerminal = embedWebTerminal
		server.EmbedWebKubeFree = embedWebKubeFree
		// if runtime.GOOS == "darwin" {
		// 	server.WebkubectlEntrypoint = webkubectlEntrypointDarwin
		// } else {
		// 	server.WebkubectlEntrypoint = webkubectlEntrypointLinux
		// }
		ip.IpCommonDictionary = IpCommonDictionary
		return server.Listen(route.InitRoute,
			server.WithCustomConfigFilePath(configPath),
			server.WithServerBindHost(serverBindHost),
			server.WithServerBindPort(serverBindPort))
	},
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		panic(err)
	}
}
