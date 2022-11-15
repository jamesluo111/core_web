package contract

import "time"

// DistributedKey 定义字符串凭证
const DistributedKey = "hade:distributed"

// Distributed 分布式服务
type Distributed interface {
	// Select 分布式服务器,所有节点对某个服务进行抢占，只选择其中一个节点
	// ServiceName 服务名称
	// appID 当前appID
	// holdTime 分布式选择器hold住的时间
	// 返回值
	// SelectAppID 分布式选择器最终选择的App
	// err 异常才返回，如果没有被选择，不返回err
	Select(ServiceName string, appID string, holdTime time.Duration) (selectAppID string, err error)
}
