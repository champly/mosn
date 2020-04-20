package handler

import (
	"context"
	"fmt"

	"mosn.io/api"
	"mosn.io/mosn/pkg/protocol/xprotocol/dubbo"
	"mosn.io/mosn/pkg/types"
)

func init() {
	XHandler[dubbo.ProtocolName] = GetDubboRouteHandler
	CustomerPort = append(CustomerPort, dubboPort, dubboMosnPort)
}

const (
	dubboClusterPre = "outbound|%d|%s|dubbo-%s-%s"

	// consumer  -> 127.0.0.1:20880(mosn-consumer)->provider:20881(mosn-provider)->127.0.0.1:20880
	dubboPort     = 20880
	dubboMosnPort = 20881
)

type dubboHandler struct {
	route  api.Route
	header api.HeaderMap
}

func GetDubboRouteHandler(route api.Route, header api.HeaderMap) types.RouteHandler {
	return &dubboHandler{route: route, header: header}
}

func (d *dubboHandler) IsAvailable(ctx context.Context, manager types.ClusterManager) (types.ClusterSnapshot, types.HandlerStatus) {
	if d.route == nil {
		return nil, types.HandlerNotAvailable
	}

	service, ok := d.header.Get(dubbo.ServiceNameHeader)
	if !ok {
		return nil, types.HandlerNotAvailable
	}
	method, ok := d.header.Get(dubbo.MethodNameHeader)
	if !ok {
		return nil, types.HandlerNotAvailable
	}
	// TODO: support through subset build cluster
	var subset string

	listenerPort, ok := ctx.Value(types.ContextKeyListenerPort).(int)
	if !ok {
		return nil, types.HandlerNotAvailable
	}

	var clusterName string
	if listenerPort == dubboMosnPort {
		clusterName = d.Route().RouteRule().ClusterName()
	} else {
		clusterName = fmt.Sprintf(dubboClusterPre, dubboMosnPort, subset, service, method)
	}
	snapshot := manager.GetClusterSnapshot(context.Background(), clusterName)

	return snapshot, types.HandlerAvailable
}

func (d *dubboHandler) Route() api.Route {
	return d.route
}