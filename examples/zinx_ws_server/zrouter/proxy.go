/**
 * Package zrouter
 * @Author: tbb
 * @Date: 2024/5/11 16:11
 */
package zrouter

import (
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/zlog"
	"github.com/aceld/zinx/znet"
)

//ProxyRouter ping test 自定义路由
type ProxyRouter struct {
	znet.BaseRouter
	server ziface.IConnection
	client ziface.IConnection
}

//Handle Ping Handle
func (p *ProxyRouter) Handle(request ziface.IRequest) {
	//先读取客户端的数据，再回写ping...ping...ping
	zlog.Debug("recv from client : path=", request.GetConnection().GetConnPath(), ", data=", string(request.GetData()))

	path := request.GetConnection().GetConnPath()
	if p.IsClient(path) {
		zlog.Debug("client init, path:", path)
		p.client = request.GetConnection()

		if p.server != nil {
			zlog.Debug("client to server:", string(request.GetData()))
			err := p.server.SendWsMsg(request.GetData())
			if err != nil {
				zlog.Error(err)
			}
		}
	} else {
		zlog.Debug("server init, path:", path)
		p.server = request.GetConnection()

		if p.client != nil {
			zlog.Debug("server to client:", string(request.GetData()))
			err := p.client.SendWsMsg(request.GetData())
			if err != nil {
				zlog.Error(err)
			}
		}
	}

}

func (p *ProxyRouter) IsClient(path string) bool {
	if path == "/ws/ssh_sdk_login" {
		return false
	}
	return true
}
