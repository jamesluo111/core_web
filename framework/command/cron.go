package command

import (
	"fmt"
	"github.com/erikdubbelboer/gspt"
	"github.com/jamesluo111/core_web/framework/cobra"
	"github.com/jamesluo111/core_web/framework/contract"
	"github.com/sevlyar/go-daemon"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

var cronDaemon = false

func initCronCommand() *cobra.Command {
	cronStartCommand.Flags().BoolVarP(&cronDaemon, "deamon", "d", false, "start serve deamon")
	cronCommand.AddCommand(cronStartCommand)
	return cronCommand
}

var cronCommand = &cobra.Command{
	Use:   "cron",
	Short: "定时任务相关命令",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			cmd.Help()
		}
		return nil
	},
}

var cronStartCommand = &cobra.Command{
	Use:   "start",
	Short: "启动cron常驻进程",
	RunE: func(cmd *cobra.Command, args []string) error {
		//获取容器
		container := cmd.GetContainer()
		//获取容器中的app服务
		appService := container.MustMake(contract.AppKey).(contract.App)

		//设置cron的日志地址和进程id地址
		pidFolder := appService.RuntimeFolder()
		serverPidFile := filepath.Join(pidFolder, "cron.pid")
		logFolder := appService.LogFolder()
		serverLogFile := filepath.Join(logFolder, "cron.log")
		currrentFolder := appService.BaseFolder()

		// deamon模式
		if cronDaemon {
			//创建一个Context
			cntext := &daemon.Context{
				//设置pid文件
				PidFileName: serverPidFile,
				PidFilePerm: 0664,
				//设置日志文件
				LogFileName: serverLogFile,
				LogFilePerm: 0640,
				//设置工作路径
				WorkDir: currrentFolder,
				//设置所有文件的mask，默认为750
				Umask: 027,
				//子进程的参数，按照这个参数设置，子进程的命令为 ./main cron start --deamon=true
				Args: []string{"", "cron", "start", "--deamon=true"},
			}
			//启动子进程，d不为空表示当前是父进程，d为空表示当前为子进程
			d, err := cntext.Reborn()
			if err != nil {
				return err
			}
			if d != nil {
				//父进程直接打印启动成功消息，不做任何操作
				fmt.Println("cron serve started, pid:", d.Pid)
				fmt.Println("log file:", serverLogFile)
				return nil
			}
			defer cntext.Release()
			fmt.Println("deamon started")
			gspt.SetProcTitle("hade cron")
			cmd.Root().Cron.Run()
		}

		//not deamon mode
		fmt.Println("start cron job")
		content := strconv.Itoa(os.Getpid())
		fmt.Println("[PID]", content)
		err := ioutil.WriteFile(serverPidFile, []byte(content), 0664)
		if err != nil {
			return err
		}
		gspt.SetProcTitle("hade cron")
		cmd.Root().Cron.Run()
		return nil
	},
}
