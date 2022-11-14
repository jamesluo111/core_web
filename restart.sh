#!/bin/bash
now=$(date +%s)
NOWPATH=$(
  cd $(dirname $0)
  pwd
)
function log() {
  echo "[$(($(date +%s) - now))]  $@"
}
SERVICE_NAME=$1
if [ ${#SERVICE_NAME} -eq 0 ]; then
	log "输入要重启的服务......"
	log "app-1"
	exit
fi

if [ ${#SERVICE_NAME} -eq 1 ]; then
	log "进入编译main服务....."
 	cd $NOWPATH/
  	rm -rf main
	GOOS=linux GOARCH=amd64 go build main.go
	log "服务编译中>>>>"
	log "重启服务......"
	kill -s 9 `ps -ef | grep ./main | grep -v grep | awk '{print $2}'`
	cd $NOWPATH && nohup ./main app start &
	log "重启完成"
fi
exit
