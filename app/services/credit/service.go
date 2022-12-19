package credit

import (
	"carbX/app/datastore/credit"
	"context"
	"sync"
)

type creditService struct {
	creditRepo *credit.Queries
}

var (
	creditOnce sync.Once
	credInst *creditService
)

func GetCreditService() *creditService {
	creditOnce.Do(func() {
		credInst = &creditService{
			creditRepo: credit.New(credit.RwInstancePool()),
		}
	})
	return credInst
}

func (c *creditService) GetCreditBySSRN(ctx context.Context, ssrn int32) (*credit.Credit, error) {
	return c.creditRepo.GetCreditBySSRN(ctx, ssrn)
}