package server

import (
	"context"
	"github.com/qubic/go-qubic/clients/core"
	"github.com/qubic/go-qubic/common"
	qubicpb "github.com/qubic/go-qubic/proto/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ qubicpb.QuotteryServiceServer = &QuotteryService{}

type QuotteryService struct {
	qubicpb.UnimplementedQuotteryServiceServer
	coreClient *core.Client
}

// NewQuotteryService creates the service and registers it to the grpc server
func NewQuotteryService(coreClient *core.Client) *QuotteryService {
	service := QuotteryService{coreClient: coreClient}

	return &service
}

func (s *QuotteryService) GetBasicInfo(ctx context.Context, _ *emptypb.Empty) (*qubicpb.BasicInfo, error) {
	bi, err := s.coreClient.QuotteryClient().GetBasicInfo(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return bi, nil
}

func (s *QuotteryService) GetBetInfo(ctx context.Context, req *qubicpb.GetBetInfoRequest) (*qubicpb.BetInfo, error) {
	bi, err := s.coreClient.QuotteryClient().GetBetInfo(ctx, req.BetId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return bi, nil
}

func (s *QuotteryService) GetActiveBets(ctx context.Context, _ *emptypb.Empty) (*qubicpb.ActiveBets, error) {
	ab, err := s.coreClient.QuotteryClient().GetActiveBets(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return ab, nil
}

func (s *QuotteryService) GetActiveBetsByCreator(ctx context.Context, req *qubicpb.GetActiveBetsByCreatorRequest) (*qubicpb.ActiveBets, error) {
	ab, err := s.coreClient.QuotteryClient().GetActiveBetsByCreator(ctx, common.Identity(req.CreatorId))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return ab, nil
}

func (s *QuotteryService) GetBettorsByBetOption(ctx context.Context, req *qubicpb.GetBettorsByBetOptionRequest) (*qubicpb.BetOptionBettors, error) {
	bettors, err := s.coreClient.QuotteryClient().GetBettorsByBetOption(ctx, req.BetId, req.BetOption)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return bettors, nil
}
