# Report service

Aggregator microservice (Phase 10). No domain database.

Public admin routes under `/reports/*` are fulfilled by calling:

- Customer `/internal/customers/reports/*`
- Merchant `/internal/merchants/reports/names`
- Ledger `/internal/transactions/reports/merchant-fees`
