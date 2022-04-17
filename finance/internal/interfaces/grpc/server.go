package grpc

import (
	"context"
	"github.com/mark-by/little-busy-back/finance/internal/application"
	"github.com/mark-by/little-busy-back/finance/pkg/proto/finance"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type FinanceServer struct {
	records application.RecordI
	value   application.ValueI
}

func (f FinanceServer) Create(ctx context.Context, record *finance.Record) (*finance.Empty, error) {
	err := f.records.Create(convertProtoRecord(record))
	if err != nil {
		return nil, err
	}
	return new(finance.Empty), nil
}

func (f FinanceServer) GetRecordsForDay(ctx context.Context, request *finance.DateRequest) (*finance.Records, error) {
	records, err := f.records.GetRecordsForDay(
		time.Date(int(request.Year),
			time.Month(request.Month),
			int(request.Day), 0, 0, 0, 0, time.UTC))
	if err != nil {
		return nil, err
	}

	return convertRecords(records), nil
}

func (f FinanceServer) ProfitForMonth(ctx context.Context, request *finance.DateRequest) (*finance.Profit, error) {
	profit, err := f.value.GetProfitForMonth(int(request.Year), int(request.Month))
	if err != nil {
		return nil, err
	}

	return &finance.Profit{Value: profit}, nil
}

func (f FinanceServer) GetValuesForMonth(ctx context.Context, request *finance.DateRequest) (*finance.Values, error) {
	values, err := f.value.GetForMonth(int(request.Year), int(request.Month))
	if err != nil {
		return nil, err
	}
	return convertValues(values), nil
}

func (f FinanceServer) GetValuesForYear(ctx context.Context, request *finance.DateRequest) (*finance.Values, error) {
	values, err := f.value.GetForYear(int(request.Year))
	if err != nil {
		return nil, err
	}
	return convertValues(values), nil
}

func (f FinanceServer) Start(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("fail to licten tcp for host %s: %v", address, err)
	}

	server := grpc.NewServer()

	finance.RegisterFinanceServiceServer(server, f)

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func NewFinanceServer(records application.RecordI, value application.ValueI) *FinanceServer {
	return &FinanceServer{
		records: records,
		value:   value,
	}
}

var _ finance.FinanceServiceServer = &FinanceServer{}
