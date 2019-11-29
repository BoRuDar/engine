package wsapi

import (
	"fmt"
	"github.com/webitel/engine/app"
	"github.com/webitel/engine/model"
	"time"
)

func (api *API) InitCall() {
	api.Router.Handle("subscribe_call", api.ApiWebSocketHandler(api.subscribeSelfCalls))
	api.Router.Handle("un_subscribe_call", api.ApiWebSocketHandler(api.unSubscribeSelfCalls))
	api.Router.Handle("call_invite", api.ApiAsyncWebSocketHandler(api.callInvite))
	api.Router.Handle("call_hangup", api.ApiWebSocketHandler(api.callHangup))
	api.Router.Handle("call_hold", api.ApiWebSocketHandler(api.callHold))
	api.Router.Handle("call_unhold", api.ApiWebSocketHandler(api.callUnHold))
	api.Router.Handle("call_dtmf", api.ApiWebSocketHandler(api.callDTMF))
	api.Router.Handle("call_mute", api.ApiWebSocketHandler(api.callMute))
	api.Router.Handle("call_blind_transfer", api.ApiWebSocketHandler(api.callBlindTransfer))
	api.Router.Handle("call_bridge", api.ApiWebSocketHandler(api.callBridge))

	api.Router.Handle("test", api.ApiAsyncWebSocketHandler(api.test))

	api.Router.Handle("sip_proxy", api.ApiWebSocketHandler(api.sipProxy))
}

func (api *API) test(conn *app.WebConn, req *model.WebSocketRequest) (map[string]interface{}, *model.AppError) {
	time.Sleep(time.Second * 10)
	return nil, nil
}

func (api *API) sipProxy(conn *app.WebConn, req *model.WebSocketRequest) (map[string]interface{}, *model.AppError) {
	var ok bool
	var data string
	if data, ok = req.Data["data"].(string); !ok {
		return nil, NewInvalidWebSocketParamError(req.Action, "data")
	}
	//conn.Sip.Send([]byte(data))
	if data != "" {
	}
	return nil, nil
}

func (api *API) subscribeSelfCalls(conn *app.WebConn, req *model.WebSocketRequest) (map[string]interface{}, *model.AppError) {
	h, e := api.App.GetHubById(req.Session.Domain(0))
	if e != nil {
		return nil, e
	}

	return nil, h.SubscribeSessionCalls(conn)
}

func (api *API) unSubscribeSelfCalls(conn *app.WebConn, req *model.WebSocketRequest) (map[string]interface{}, *model.AppError) {
	h, e := api.App.GetHubById(req.Session.Domain(0))
	if e != nil {
		return nil, e
	}

	return nil, h.UnSubscribeCalls(conn)
}

func (api *API) callHangup(conn *app.WebConn, req *model.WebSocketRequest) (map[string]interface{}, *model.AppError) {
	var ok bool
	var id, nodeId string

	if id, ok = req.Data["id"].(string); !ok {
		return nil, NewInvalidWebSocketParamError(req.Action, "id")
	}
	if nodeId, ok = req.Data["node_id"].(string); !ok {
		return nil, NewInvalidWebSocketParamError(req.Action, "node_id")
	}

	var cause = req.GetFieldString("cause")

	if cli, err := api.App.CallManager().CallClient(nodeId); err != nil {
		return nil, err
	} else {
		err = cli.HangupCall(id, cause)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func (api *API) callBlindTransfer(conn *app.WebConn, req *model.WebSocketRequest) (map[string]interface{}, *model.AppError) {
	var ok bool
	var id, nodeId, destination string

	if id, ok = req.Data["id"].(string); !ok {
		return nil, NewInvalidWebSocketParamError(req.Action, "id")
	}
	if nodeId, ok = req.Data["node_id"].(string); !ok {
		return nil, NewInvalidWebSocketParamError(req.Action, "node_id")
	}
	if destination, ok = req.Data["destination"].(string); !ok || len(destination) < 1 {
		return nil, NewInvalidWebSocketParamError(req.Action, "destination")
	}

	if cli, err := api.App.CallManager().CallClient(nodeId); err != nil {
		return nil, err
	} else {
		err = cli.BlindTransfer(id, destination)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func (api *API) callHold(conn *app.WebConn, req *model.WebSocketRequest) (map[string]interface{}, *model.AppError) {
	var ok bool
	var id, nodeId string

	if id, ok = req.Data["id"].(string); !ok {
		return nil, NewInvalidWebSocketParamError(req.Action, "id")
	}
	if nodeId, ok = req.Data["node_id"].(string); !ok {
		return nil, NewInvalidWebSocketParamError(req.Action, "node_id")
	}

	if cli, err := api.App.CallManager().CallClient(nodeId); err != nil {
		return nil, err
	} else {
		err = cli.Hold(id)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func (api *API) callDTMF(conn *app.WebConn, req *model.WebSocketRequest) (map[string]interface{}, *model.AppError) {
	var ok bool
	var id, nodeId string
	var key string

	if id, ok = req.Data["id"].(string); !ok {
		return nil, NewInvalidWebSocketParamError(req.Action, "id")
	}
	if nodeId, ok = req.Data["node_id"].(string); !ok {
		return nil, NewInvalidWebSocketParamError(req.Action, "node_id")
	}
	if key, ok = req.Data["dtmf"].(string); !ok || len(key) < 1 {
		return nil, NewInvalidWebSocketParamError(req.Action, "dtmf")
	}

	if cli, err := api.App.CallManager().CallClient(nodeId); err != nil {
		return nil, err
	} else {
		err = cli.DTMF(id, []rune(key)[0])
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (api *API) callUnHold(conn *app.WebConn, req *model.WebSocketRequest) (map[string]interface{}, *model.AppError) {
	var ok bool
	var id, nodeId string

	if id, ok = req.Data["id"].(string); !ok {
		return nil, NewInvalidWebSocketParamError(req.Action, "id")
	}
	if nodeId, ok = req.Data["node_id"].(string); !ok {
		return nil, NewInvalidWebSocketParamError(req.Action, "node_id")
	}

	if cli, err := api.App.CallManager().CallClient(nodeId); err != nil {
		return nil, err
	} else {
		err = cli.UnHold(id)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func (api *API) callInvite(conn *app.WebConn, req *model.WebSocketRequest) (map[string]interface{}, *model.AppError) {
	var ok bool
	var callId string
	var destinationNumber, destinationName string
	var variables map[string]interface{}

	callId = model.NewUuid()

	if destinationNumber, ok = req.Data["toNumber"].(string); !ok {
		return nil, NewInvalidWebSocketParamError(req.Action, "toNumber")
	}

	if destinationName, ok = req.Data["toName"].(string); !ok {
		destinationName = destinationNumber
	}

	info, err := api.App.GetUserCallInfo(conn.UserId, conn.DomainId)
	if err != nil {
		return nil, err
	}

	invite := &model.CallRequest{
		Endpoints:   info.GetCallEndpoints(),
		Destination: destinationNumber,
		Variables: map[string]string{
			model.CALL_VARIABLE_ID:                callId,
			model.CALL_VARIABLE_DIRECTION:         model.CALL_DIRECTION_INTERNAL,
			model.CALL_VARIABLE_DISPLAY_DIRECTION: model.CALL_DIRECTION_OUTBOUND,
			model.CALL_VARIABLE_USER_ID:           fmt.Sprintf("%v", conn.UserId),
			model.CALL_VARIABLE_DOMAIN_ID:         fmt.Sprintf("%v", conn.DomainId),
			model.CALL_VARIABLE_SOCK_ID:           conn.Id(),

			"sip_h_X-Webitel-Destination": destinationNumber,

			"origination_uuid":      callId,
			"absolute_codec_string": "pcma",

			"hangup_after_bridge":        "true",
			"effective_caller_id_number": info.Extension,
			"effective_caller_id_name":   info.Name,
			"effective_callee_id_name":   destinationName,
			"effective_callee_id_number": destinationNumber,

			"origination_caller_id_name":   destinationName,
			"origination_caller_id_number": destinationNumber,
			"origination_callee_id_name":   info.Name,
			"origination_callee_id_number": info.Extension,
		},
		Timeout:      0,
		CallerName:   destinationName,
		CallerNumber: destinationNumber,
		//Applications: []*model.CallRequestApplication{
		//	{
		//		AppName: "transfer",
		//		Args:    "9999",
		//	},
		//},
	}

	if variables, ok = req.Data["variables"].(map[string]interface{}); ok {
		for k, v := range variables {
			switch v.(type) {
			case string:
				invite.AddUserVariable(k, v.(string))
			case interface{}:
				invite.AddUserVariable(k, fmt.Sprintf("%v", v))
			}
		}
	}

	_, err = api.App.CallManager().MakeOutboundCall(invite)

	if err != nil {
		return nil, err
	}

	data := map[string]interface{}{}
	data["call_id"] = callId
	return data, nil
}

func (api *API) callMute(conn *app.WebConn, req *model.WebSocketRequest) (map[string]interface{}, *model.AppError) {
	var ok, mute bool
	var id, nodeId string

	if id, ok = req.Data["id"].(string); !ok {
		return nil, NewInvalidWebSocketParamError(req.Action, "id")
	}
	if nodeId, ok = req.Data["node_id"].(string); !ok {
		return nil, NewInvalidWebSocketParamError(req.Action, "node_id")
	}
	if mute, ok = req.Data["mute"].(bool); !ok {
		return nil, NewInvalidWebSocketParamError(req.Action, "mute")
	}

	if cli, err := api.App.CallManager().CallClient(nodeId); err != nil {
		return nil, err
	} else {
		err = cli.Mute(id, mute)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func (api *API) callBridge(conn *app.WebConn, req *model.WebSocketRequest) (map[string]interface{}, *model.AppError) {
	var ok bool
	var id, nodeId, id2, nodeId2 string

	if id, ok = req.Data["id"].(string); !ok {
		return nil, NewInvalidWebSocketParamError(req.Action, "id")
	}
	if nodeId, ok = req.Data["node_id"].(string); !ok {
		return nil, NewInvalidWebSocketParamError(req.Action, "node_id")
	}
	if id2, ok = req.Data["parent_id"].(string); !ok {
		return nil, NewInvalidWebSocketParamError(req.Action, "parent_id")
	}
	if nodeId2, ok = req.Data["parent_node_id"].(string); !ok {
		return nil, NewInvalidWebSocketParamError(req.Action, "parent_node_id")
	}

	api.App.CallManager().Bridge(id, nodeId, id2, nodeId2)

	return nil, nil
}
