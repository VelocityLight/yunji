package store

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"

	"yunji/common"
	"yunji/utils/log"
	"yunji/utils/sql"
)

type BillingService struct {
	db *sql.Database
}

func NewBillingService(db *sql.Database) *BillingService {
	return &BillingService{db}
}

func (s *BillingService) Update() {}

func (s *BillingService) Get() {}

func (s *BillingService) Delete() {}

func (s *BillingService) List(ctx context.Context, opts common.QueryBillingOpts) ([]common.Billing, error) {
	var billings []common.Billing
	return billings, nil
}

func (s *BillingService) GetCostByTeam(ctx context.Context) ([]common.GetCostByTeamResponse, error) {
	var res []common.GetCostByTeamResponse
	err := s.db.SelectContext(ctx, &res, `
		select team, sum(cost) as cost from (
			select usedby, CASE
				WHEN usedby in ('admin','gardener', 'dbaas-control-plane-seed',
				'dbaas-control-plane-shoot', 'gardener-shootmaster', 'gardener-shoot', 'Infra') THEN 'Infra'
				WHEN lower(usedby) = 'customer-tidb' THEN 'TiDB User'
				WHEN lower(usedby) in ('dev-us-east-1-f02', 'dev-us-east-1-f01', 'staging-us-east-1-f01')  THEN 'Dev tier/Serverless'
				WHEN lower(usedby) = 'dbaas-central' THEN 'Cloud Platform'
				ELSE 'Others(security,no tag resources,common services)'
			END as team , cost from (
				select resource_tags_user_usedby as usedby, round(sum(line_item_unblended_cost),2)  as cost from dev_billing
				where line_item_unblended_cost > 0 and line_item_product_code !=  'ComputeSavingsPlans'
				group by  resource_tags_user_usedby  order by cost
			) t ) t2 group by team`)
	return res, err
}

func (s *BillingService) GetUsedByTags(ctx context.Context) ([]common.UsedByTag, error) {
	var res []common.UsedByTag
	err := s.db.SelectContext(ctx, &res, `
		select resource_tags_user_usedby from dev_billing group by resource_tags_user_usedby`)
	return res, err
}

func (s *BillingService) GetTags(ctx context.Context) ([]common.Tag, error) {
	var res []common.Tag
	err := s.db.SelectContext(ctx, &res, `
		select resource_tags_user_component from dev_billing group by resource_tags_user_component`)
	return res, err
}

func (s *BillingService) GetTrends(ctx context.Context, opts common.GetTrendOpts) (common.GetTrendResponse, error) {
	var res common.GetTrendResponse
	where := []exp.Expression{}
	if len(opts.Tags) > 0 {
		where = append(where, goqu.I("dev_billing.resource_tags_user_component").In(opts.Tags))
	}

	if len(opts.Service) > 0 {
		where = append(where, goqu.I("dev_billing.resource_tags_user_usedby").In(opts.Service))
	}

	b := sql.Builder.From(goqu.T("dev_billing")).Where(where...).Prepared(true)

	query, args, _ := b.Select(
		goqu.L("DATE_FORMAT(line_item_usage_start_date, '%Y%m%d') AS time"),
		goqu.L("SUM(line_item_unblended_cost) AS cost"),
		goqu.L("line_item_product_code AS service"),
	).GroupBy(goqu.L("time, service")).ToSQL()

	log.Log.Infof("where: %v, query: %s, args: %v", where, query, args)
	err := s.db.SelectContext(ctx, &res.Body, query, args...)
	return res, err
}

func (s *BillingService) Create() {}
