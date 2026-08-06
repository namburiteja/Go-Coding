package report

import (
	"context"
	"encoding/json"
	"sort"
	"strconv"
)

// MerchantFeeCollected matches the previous public report JSON shape.
type MerchantFeeCollected struct {
	ID                int32  `json:"id"`
	Name              string `json:"name"`
	TotalFeeCollected string `json:"total_fee_collected"`
}

type Service struct {
	customers *CustomersAPI
	merchants *MerchantsAPI
	ledger    *LedgerAPI
}

func NewService(customers *CustomersAPI, merchants *MerchantsAPI, ledger *LedgerAPI) *Service {
	return &Service{
		customers: customers,
		merchants: merchants,
		ledger:    ledger,
	}
}

func (s *Service) GetUsersAtCreditLimit(ctx context.Context) (json.RawMessage, error) {
	return s.customers.GetUsersAtCreditLimit(ctx)
}

func (s *Service) GetCustomersWithDue(ctx context.Context) (json.RawMessage, error) {
	return s.customers.GetCustomersWithDue(ctx)
}

func (s *Service) GetCustomerDueByName(ctx context.Context, name string) (json.RawMessage, error) {
	return s.customers.GetCustomerDueByName(ctx, name)
}

// GetAllMerchantsFeeCollected joins merchant names with ledger fee totals in Go
// (replaces the former SQL LEFT JOIN across merchants + transactions).
func (s *Service) GetAllMerchantsFeeCollected(ctx context.Context) ([]MerchantFeeCollected, error) {
	names, err := s.merchants.GetNames(ctx)
	if err != nil {
		return nil, err
	}

	fees, err := s.ledger.GetMerchantFeeTotals(ctx)
	if err != nil {
		return nil, err
	}

	feeByMerchant := make(map[int32]string, len(fees))
	for _, fee := range fees {
		feeByMerchant[fee.MerchantID] = fee.TotalFeeCollected
	}

	out := make([]MerchantFeeCollected, 0, len(names))
	for _, m := range names {
		total := "0"
		if v, ok := feeByMerchant[m.ID]; ok && v != "" {
			total = v
		}
		out = append(out, MerchantFeeCollected{
			ID:                m.ID,
			Name:              m.Name,
			TotalFeeCollected: total,
		})
	}

	sort.Slice(out, func(i, j int) bool {
		ai, _ := strconv.ParseFloat(out[i].TotalFeeCollected, 64)
		aj, _ := strconv.ParseFloat(out[j].TotalFeeCollected, 64)
		if ai == aj {
			return out[i].ID < out[j].ID
		}
		return ai > aj
	})

	return out, nil
}
