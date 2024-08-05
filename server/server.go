package server

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/qubic/go-qubic/clients/core"
	"github.com/qubic/go-qubic/common"
	qubicpb "github.com/qubic/go-qubic/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
	"net/http"
)

var _ qubicpb.QubicServiceServer = &Server{}

type TransactionInfo struct {
	timestamp uint64
	moneyFlew bool
}

type Server struct {
	qubicpb.UnimplementedQubicServiceServer
	listenAddrGRPC string
	listenAddrHTTP string
	coreClient     *core.Client
}

func NewServer(listenAddrGRPC, listenAddrHTTP string, coreClient *core.Client) *Server {
	return &Server{
		listenAddrGRPC: listenAddrGRPC,
		listenAddrHTTP: listenAddrHTTP,
		coreClient:     coreClient,
	}
}

func (s *Server) GetBasicInfo(ctx context.Context, _ *emptypb.Empty) (*qubicpb.BasicInfo, error) {
	bi, err := s.coreClient.QuotteryClient().GetBasicInfo(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return bi, nil
}

func (s *Server) GetBetInfo(ctx context.Context, req *qubicpb.GetBetInfoRequest) (*qubicpb.BetInfo, error) {
	bi, err := s.coreClient.QuotteryClient().GetBetInfo(ctx, req.BetId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return bi, nil
}

func (s *Server) GetActiveBets(ctx context.Context, _ *emptypb.Empty) (*qubicpb.ActiveBets, error) {
	ab, err := s.coreClient.QuotteryClient().GetActiveBets(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return ab, nil
}

func (s *Server) GetActiveBetsByCreator(ctx context.Context, req *qubicpb.GetActiveBetsByCreatorRequest) (*qubicpb.ActiveBets, error) {
	ab, err := s.coreClient.QuotteryClient().GetActiveBetsByCreator(ctx, common.Identity(req.CreatorId))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return ab, nil
}

func (s *Server) GetBettorsByBetOption(ctx context.Context, req *qubicpb.GetBettorsByBetOptionRequest) (*qubicpb.BetOptionBettors, error) {
	bettors, err := s.coreClient.QuotteryClient().GetBettorsByBetOption(ctx, req.BetId, req.BetOption)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return bettors, nil
}

func (s *Server) Start() error {
	srv := grpc.NewServer(
		grpc.MaxRecvMsgSize(600*1024*1024),
		grpc.MaxSendMsgSize(600*1024*1024),
	)
	qubicpb.RegisterQubicServiceServer(srv, s)
	reflection.Register(srv)

	lis, err := net.Listen("tcp", s.listenAddrGRPC)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func() {
		if err := srv.Serve(lis); err != nil {
			panic(err)
		}
	}()

	if s.listenAddrHTTP != "" {
		go func() {
			mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{EmitDefaultValues: true, EmitUnpopulated: false},
			}))
			opts := []grpc.DialOption{
				grpc.WithTransportCredentials(insecure.NewCredentials()),
				grpc.WithDefaultCallOptions(
					grpc.MaxCallRecvMsgSize(600*1024*1024),
					grpc.MaxCallSendMsgSize(600*1024*1024),
				),
			}

			if err := qubicpb.RegisterQubicServiceHandlerFromEndpoint(
				context.Background(),
				mux,
				s.listenAddrGRPC,
				opts,
			); err != nil {
				panic(err)
			}

			if err := http.ListenAndServe(s.listenAddrHTTP, mux); err != nil {
				panic(err)
			}
		}()
	}

	return nil
}
