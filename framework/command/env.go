package command

import (
	"fmt"
	"github.com/jamesluo111/core_web/framework/cobra"
	"github.com/jamesluo111/core_web/framework/contract"
	"github.com/jamesluo111/core_web/framework/util"
)

func InitEnvCommand() *cobra.Command {
	envCommand.AddCommand(envListCommand)
	return envCommand
}

var envCommand = &cobra.Command{
	Use:   "env",
	Short: "环境变量",
	Run: func(cmd *cobra.Command, args []string) {
		// 获取容器
		container := cmd.GetContainer()
		envService := container.MustMake(contract.EnvKey).(contract.Env)
		// 打印环境
		fmt.Println(envService.AppEnv())
	},
}

var envListCommand = &cobra.Command{
	Use:   "list",
	Short: "获取所有环境变量",
	Run: func(cmd *cobra.Command, args []string) {
		// 获取容器
		container := cmd.GetContainer()
		envService := container.MustMake(contract.EnvKey).(contract.Env)
		envs := envService.All()
		outs := [][]string{}
		for k, v := range envs {
			outs = append(outs, []string{k, v})
		}
		util.PrettyPrint(outs)
	},
}
