package credit

import (
	"carbX/app/services/credit"
	"carbX/app/system/log"
	context "context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type server struct {
	UnimplementedCreditServer
}

func (s *server) Credit(ctx context.Context, req *CreditRequest) (*CreditResponse, error) {
	credit, err := credit.GetCreditService().GetCreditBySSRN(ctx, int32(req.GetSsrn()))
	if err != nil {
		return &CreditResponse{}, status.Error(404, "Credit not found")
	}
	return &CreditResponse{
		ProjectId:     credit.ProjectID,
		Standard:      credit.Std,
		Methodology:   credit.Methodology,
		Region:        credit.Region,
		StorageMethod: credit.Storagemethod,
		Method:        credit.Method,
		EmissionType:  credit.Emissiontype,
		Category:      credit.Catergory,
		Uri:           credit.Uri,
		Beneficiary:   credit.Beneficiary,
	}, nil
}

func Start() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Error("failed to start listener: %v", err)
		panic(err)
	}
	s := grpc.NewServer()
	RegisterCreditServer(s, &server{})
	log.Info("listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Error("failed to start server: %v", err)
		panic(err)
	}
}