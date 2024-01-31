package servers

import (
	"qnc/biz/servers/utils"
	"qnc/pkg/constants"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/websocket"
)

// channel通道
var ToClientChan chan clientInfo

// channel通道结构体
type clientInfo struct {
	ClientId   string
	SendUserId string
	MessageId  string
	MsgType    string
	Code       int
	Msg        string
	Data       *string
}

type RetData struct {
	MessageId  string      `json:"messageId"`
	SendUserId string      `json:"sendUserId"`
	MsgType    string      `json:"msg_type"`
	Code       int         `json:"code"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
}

// 心跳间隔
var heartbeatInterval = 25 * time.Second

func init() {
	ToClientChan = make(chan clientInfo, 1000)
}

var Manager = NewClientManager() // 管理者

// 发送信息到指定客户端
func SendMessage2Client(clientId string, sendUserId string, msgType string, code int, msg string, data *string) (messageId string) {
	messageId = utils.GenUUID()
	if utils.IsCluster() {
		_, _, _, isLocal, err := utils.GetAddrInfoAndIsLocal(clientId)
		if err != nil {
			hlog.Errorf("%s", err)
			return
		}

		//如果是本机则发送到本机
		if isLocal {
			SendMessage2LocalClient(messageId, clientId, sendUserId, msgType, code, msg, data)
		} else {
			//TODO
			// SendRpc2Client(addr, messageId, sendUserId, clientId, code, msg, data)
		}
	} else {
		//如果是单机服务，则只发送到本机
		SendMessage2LocalClient(messageId, clientId, sendUserId, msgType, code, msg, data)
	}

	return
}

// 关闭客户端
func CloseClient(clientId, systemId string) {
	if utils.IsCluster() {
		_, _, _, isLocal, err := utils.GetAddrInfoAndIsLocal(clientId)
		if err != nil {
			hlog.Errorf("%s", err)
			return
		}

		if isLocal {
			CloseLocalClient(clientId, systemId)
		} else {
			//TODO
			// CloseRpcClient(addr, clientId, systemId)
		}
	} else {
		//如果是单机服务，则只发送到本机
		CloseLocalClient(clientId, systemId)
	}

	return
}

func AddClient2Group(systemId string, groupName string, clientId string, userId string, extend string) {
	if utils.IsCluster() {
		_, _, _, isLocal, err := utils.GetAddrInfoAndIsLocal(clientId)
		if err != nil {
			hlog.Errorf("%s", err)
			return
		}

		if isLocal {
			if client, err := Manager.GetByClientId(clientId); err == nil {
				Manager.AddClient2LocalGroup(groupName, client, userId, extend)
			} else {
				hlog.Error(err)
			}
		} else {
			//TODO
			// SendRpcBindGroup(addr, systemId, groupName, clientId, userId, extend)
		}
	} else {
		if client, err := Manager.GetByClientId(clientId); err == nil {
			//如果是单机，就直接添加到本地group了
			Manager.AddClient2LocalGroup(groupName, client, userId, extend)
		}
	}
}

// 发送信息到指定分组
func SendMessage2Group(systemId, sendUserId, groupName string, msgType string, code int, msg string, data *string) (messageId string) {
	messageId = utils.GenUUID()
	if utils.IsCluster() {
		//TODO
		// go SendGroupBroadcast(systemId, messageId, sendUserId, groupName, code, msg, data)
	} else {
		//如果是单机服务，则只发送到本机
		Manager.SendMessage2LocalGroup(systemId, messageId, sendUserId, groupName, msgType, code, msg, data)
	}
	return
}

// 发送信息到指定系统
func SendMessage2System(systemId, sendUserId string, msgType string, code int, msg string, data string) {
	messageId := utils.GenUUID()
	if utils.IsCluster() {
		//TODO
		// SendSystemBroadcast(systemId, messageId, sendUserId, code, msg, &data)
	} else {
		Manager.SendMessage2LocalSystem(systemId, messageId, sendUserId, msgType, code, msg, &data)
	}
}

// 获取分组列表
func GetOnlineList(systemId *string, groupName *string) map[string]interface{} {
	var clientList []string
	if utils.IsCluster() {
		//TODO
		// clientList = GetOnlineListBroadcast(systemId, groupName)
	} else {
		//如果是单机服务，则只发送到本机
		retList := Manager.GetGroupClientList(utils.GenGroupKey(*systemId, *groupName))
		clientList = append(clientList, retList...)
	}

	return map[string]interface{}{
		"count": len(clientList),
		"list":  clientList,
	}
}

// 通过本服务器发送信息
func SendMessage2LocalClient(messageId, clientId string, sendUserId string, msgType string, code int, msg string, data *string) {
	hlog.Debug(map[string]interface{}{
		"host":     constants.GlobalSetting.LocalHost,
		"port":     constants.CommonSetting.HttpPort,
		"clientId": clientId,
	})
	hlog.Info("send to channel")
	ToClientChan <- clientInfo{ClientId: clientId, MessageId: messageId, SendUserId: sendUserId, MsgType: msgType, Code: code, Msg: msg, Data: data}
	return
}

// 发送关闭信号
func CloseLocalClient(clientId, systemId string) {
	if conn, err := Manager.GetByClientId(clientId); err == nil && conn != nil {
		if conn.SystemId != systemId {
			return
		}
		Manager.DisConnect <- conn
		hlog.Debug(map[string]interface{}{
			"host":     constants.GlobalSetting.LocalHost,
			"port":     constants.CommonSetting.HttpPort,
			"clientId": clientId,
		})
		hlog.Info("kickoff client")
	}
	return
}

// 监听并发送给客户端信息
func WriteMessage() {
	for {
		clientInfo := <-ToClientChan
		hlog.Debug(map[string]interface{}{
			"host":       constants.GlobalSetting.LocalHost,
			"port":       constants.CommonSetting.HttpPort,
			"clientId":   clientInfo.ClientId,
			"messageId":  clientInfo.MessageId,
			"sendUserId": clientInfo.SendUserId,
			"msgType":    clientInfo.MsgType,
			"code":       clientInfo.Code,
			"msg":        clientInfo.Msg,
			"data":       clientInfo.Data,
		})
		hlog.Info("send to local")
		if conn, err := Manager.GetByClientId(clientInfo.ClientId); err == nil && conn != nil {
			if err := Render(conn.Socket, clientInfo.MessageId, clientInfo.SendUserId, clientInfo.Code, clientInfo.MsgType, clientInfo.Msg, clientInfo.Data); err != nil {
				Manager.DisConnect <- conn
				hlog.Errorf("client offline error: ", map[string]interface{}{
					"host":       constants.GlobalSetting.LocalHost,
					"port":       constants.CommonSetting.HttpPort,
					"clientId":   clientInfo.ClientId,
					"messageId":  clientInfo.MessageId,
					"sendUserId": clientInfo.SendUserId,
					"msgType":    clientInfo.MsgType,
					"code":       clientInfo.Code,
					"msg":        clientInfo.Msg,
					"data":       clientInfo.Data,
				})
			}
		}
	}
}

func Render(conn *websocket.Conn, messageId string, sendUserId string, code int, msgType string, message string, data interface{}) error {
	return conn.WriteJSON(RetData{
		Code:       code,
		MessageId:  messageId,
		SendUserId: sendUserId,
		MsgType:    msgType,
		Msg:        message,
		Data:       data,
	})
}

// 启动定时器进行心跳检测
func PingTimer() {
	go func() {
		ticker := time.NewTicker(heartbeatInterval)
		defer ticker.Stop()
		for {
			<-ticker.C
			//发送心跳
			for clientId, conn := range Manager.AllClient() {
				if err := conn.Socket.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(time.Second)); err != nil {
					Manager.DisConnect <- conn
					hlog.Errorf("send heart beat failed: %s total connect：%d", clientId, Manager.Count())
				}
			}
		}

	}()
}
