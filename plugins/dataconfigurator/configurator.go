package dataconfigurator

import (
	"github.com/gogo/status"
	"github.com/ligato/cn-infra/datasync"
	"github.com/ligato/cn-infra/logging"
	"github.com/ligato/vpp-agent/pkg/util"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"

	rpc "github.com/ligato/vpp-agent/api/dataconfigurator"
	"github.com/ligato/vpp-agent/pkg/models"
	kvs "github.com/ligato/vpp-agent/plugins/kvscheduler/api"
	"github.com/ligato/vpp-agent/plugins/orchestrator"
)

// configuratorService implements DataSyncer service.
type configuratorService struct {
	log      logging.Logger
	dispatch *orchestrator.Plugin
	dumpService
}

func (svc *configuratorService) Get(context.Context, *rpc.GetRequest) (*rpc.GetResponse, error) {
	config := newData()

	util.PlaceProtos(svc.dispatch.ListData(), config.LinuxData, config.VppData)

	return &rpc.GetResponse{Config: config}, nil
}

// Update adds configuration data present in data request to the VPP/Linux
func (svc *configuratorService) Update(ctx context.Context, req *rpc.UpdateRequest) (*rpc.UpdateResponse, error) {
	protos := util.ExtractProtos(req.Update.VppData, req.Update.LinuxData)

	var kvPairs []datasync.ProtoWatchResp
	for _, p := range protos {
		key, err := models.GetKey(p)
		if err != nil {
			svc.log.Debug("models.GetKey failed: %s", err)
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		kvPairs = append(kvPairs, &orchestrator.ProtoWatchResp{
			Key: key,
			Val: p,
		})
	}

	if req.FullResync {
		ctx = kvs.WithResync(ctx, kvs.FullResync, true)
	}

	if _, err := svc.dispatch.PushData(ctx, kvPairs); err != nil {
		st := status.New(codes.FailedPrecondition, err.Error())
		return nil, st.Err()
	}

	return &rpc.UpdateResponse{}, nil
}

// Delete removes configuration data present in data request from the VPP/linux
func (svc *configuratorService) Delete(ctx context.Context, req *rpc.DeleteRequest) (*rpc.DeleteResponse, error) {
	protos := util.ExtractProtos(req.Delete.VppData, req.Delete.LinuxData)

	var kvPairs []datasync.ProtoWatchResp
	for _, p := range protos {
		key, err := models.GetKey(p)
		if err != nil {
			svc.log.Debug("models.GetKey failed: %s", err)
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		kvPairs = append(kvPairs, &orchestrator.ProtoWatchResp{
			Key: key,
			Val: nil,
		})
	}

	if _, err := svc.dispatch.PushData(ctx, kvPairs); err != nil {
		st := status.New(codes.FailedPrecondition, err.Error())
		return nil, st.Err()
	}

	return &rpc.DeleteResponse{}, nil
}
