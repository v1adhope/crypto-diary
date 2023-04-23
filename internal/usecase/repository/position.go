package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/v1adhope/crypto-diary/internal/entity"
	"github.com/v1adhope/crypto-diary/pkg/postgres"
)

const (
	_defaultEntityCap = 25
)

type PositionRepo struct {
	*postgres.Postgres
}

func NewPosition(pg *postgres.Postgres) *PositionRepo {
	return &PositionRepo{pg}
}

func (pr *PositionRepo) Create(ctx context.Context, position *entity.Position) error {
	sql, args, err := pr.Builder.Insert("positions").
		Columns("open_date",
			"pair",
			"reason",
			"strategically",
			"percentage_risk",
			"direction",
			"deposit",
			"open_price",
			"stop_loss_price",
			"take_profit_price",
			"close_price",
			"user_id").
		Values(position.OpenDate,
			position.Pair,
			nullCheck(position.Reason),
			position.Strategically,
			position.Risk,
			position.Direction,
			position.Deposit,
			position.OpenPrice,
			position.StopLossPrice,
			position.TakeProfitPrice,
			nullCheck(position.ClosePrice),
			position.UserID).
		Suffix("RETURNING position_id").
		ToSql()
	if err != nil {
		return fmt.Errorf("repository: Create position: Query builder: %w", err)
	}

	err = pr.Pool.QueryRow(ctx, sql, args...).Scan(&position.ID)
	if err != nil {
		return fmt.Errorf("repository: Create position: QueryRow: %w", err)
	}

	return nil
}

func nullCheck(s string) *string {
	if strings.TrimSpace(s) == "" {
		return nil
	}

	return &s
}

func (pr *PositionRepo) FindAll(ctx context.Context, userID string, filter entity.Filter) ([]entity.Position, error) {
	q := pr.Builder.Select("*").
		From("get_all_positions").
		Where("user_id = ? AND position_id > ?", userID, filter.PaginationCursor)

	for fieldKey, field := range filter.Fields {
		if realFilterName, ok := allowedFilters.Load(fieldKey); ok {
			switch field.Operation {
			case entity.OpEq:
				q = q.Where(fmt.Sprintf("%s IN (?)", realFilterName.(string)), fmt.Sprint(strings.Join(field.Values, ",")))
			case entity.OpRange:
				q = q.Where(fmt.Sprintf("%s BETWEEN SYMMETRIC ? AND ?", realFilterName.(string)), field.Values[0], field.Values[1])
			}
		}
	}

	sql, args, err := q.OrderBy("position_id ASC").
		Limit(_defaultEntityCap).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("repository: FindAll position: Query builder: %w", err)
	}

	rows, err := pr.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("repository: FindAll position: Query: %w", err)
	}
	defer rows.Close()

	positions := make([]entity.Position, 0, _defaultEntityCap)

	for rows.Next() {
		p := &PositionDTO{}

		err := rows.Scan(
			&p.ID,
			&p.OpenDate,
			&p.Pair,
			&p.Reason,
			&p.Strategically,
			&p.Risk,
			&p.Direction,
			&p.Deposit,
			&p.OpenPrice,
			&p.StopLossPrice,
			&p.TakeProfitPrice,
			&p.ClosePrice,
			&p.UserID)
		if err != nil {
			return nil, fmt.Errorf("repository: FindAll positon: Scan: %w", err)
		}

		positions = append(positions, *p.ToEntity())
	}

	return positions, nil
}

func (pr *PositionRepo) Delete(ctx context.Context, userID, positionID string) error {
	sql, args, err := pr.Builder.Delete("positions").
		Where("user_id = ? AND position_id = ?", userID, positionID).
		ToSql()
	if err != nil {
		return fmt.Errorf("repository: Delete position: Query builder: %w", err)
	}

	commandTag, err := pr.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("repository: Delete positon: Exec: %w", err)
	}

	if commandTag.RowsAffected() != 1 {
		return entity.ErrNoFoundPosition
	}

	return nil
}

func (pr *PositionRepo) Replace(ctx context.Context, position *entity.Position) error {
	sql, args, err := pr.Builder.Update("positions").
		Set("open_date", position.OpenDate).
		Set("pair", position.Pair).
		Set("reason", position.Reason).
		Set("strategically", position.Strategically).
		Set("percentage_risk", position.Risk).
		Set("direction", position.Direction).
		Set("deposit", position.Deposit).
		Set("open_price", position.OpenPrice).
		Set("stop_loss_price", position.StopLossPrice).
		Set("take_profit_price", position.TakeProfitPrice).
		Set("close_price", position.ClosePrice).
		Where("user_id = ? AND position_id = ?", position.UserID, position.ID).
		ToSql()
	if err != nil {
		return fmt.Errorf("repository: Replace position: Query builder: %w", err)
	}

	commandTag, err := pr.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("repository: Replace position: Exec: %w", err)
	}

	if commandTag.RowsAffected() != 1 {
		return entity.ErrNothingToChange
	}

	return nil
}
