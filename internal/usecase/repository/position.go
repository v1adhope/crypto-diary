package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/v1adhope/crypto-diary/internal/entity"
	"github.com/v1adhope/crypto-diary/pkg/postgres"
)

const (
	defaultEntityCap = 25
)

type PositionRepo struct {
	*postgres.Postgres
}

func NewPosition(pg *postgres.Postgres) *PositionRepo {
	return &PositionRepo{pg}
}

func (pr *PositionRepo) Create(ctx context.Context, position *entity.Position) error {
	q := `INSERT INTO positions(open_date, pair, reason, strategically, percentage_risk,
                             direction, deposit, open_price, stop_loss_price,
                             take_profit_price, close_price, user_id)
        VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
        RETURNING position_id`

	err := pr.Pool.QueryRow(ctx, q,
		position.OpenDate,
		position.Pair,
		position.Reason,
		position.Strategically,
		position.Risk,
		position.Direction,
		position.Deposit,
		position.OpenPrice,
		position.StopLossPrice,
		position.TakeProfitPrice,
		nullCheck(position.ClosePrice),
		position.UserID).
		Scan(&position.ID)
	if err != nil {
		return fmt.Errorf("repository: Create position: QueryRow: %s", err)
	}

	return nil
}

func nullCheck(s string) *string {
	if strings.TrimSpace(s) == "" {
		return nil
	}

	return &s
}

func (pr *PositionRepo) FindAll(ctx context.Context, id string) ([]entity.Position, error) {
	q := `SELECT * FROM get_all_positions
        WHERE user_id = $1
        ORDER by position_id ASC`

	rows, err := pr.Pool.Query(ctx, q, id)
	if err != nil {
		return nil, fmt.Errorf("repository: FindAll position: Query: %s", err)
	}
	defer rows.Close()

	positions := make([]entity.Position, 0, defaultEntityCap)

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
	q := `DELETE FROM positions
        WHERE user_id = $1 AND position_id = $2`

	commandTag, err := pr.Pool.Exec(ctx, q, userID, positionID)
	if err != nil {
		return fmt.Errorf("repository: Delete positon: Exec: %s", err)
	}

	if commandTag.RowsAffected() != 1 {
		return entity.ErrNoFoundPosition
	}

	return nil
}

func (pr *PositionRepo) Replace(ctx context.Context, position *entity.Position) error {
	q := `UPDATE positions
        SET open_date = $1, pair = $2, reason = $3, strategically = $4,
          percentage_risk = $5, direction = $6, deposit = $7, open_price = $8,
          stop_loss_price=$9, take_profit_price=$10, close_price=$11
        WHERE position_id = $12 AND user_id = $13`

	commandTag, err := pr.Pool.Exec(ctx, q,
		position.OpenDate,
		position.Pair,
		position.Reason,
		position.Strategically,
		position.Risk,
		position.Direction,
		position.Deposit,
		position.OpenPrice,
		position.StopLossPrice,
		position.TakeProfitPrice,
		position.ClosePrice,

		position.ID,
		position.UserID,
	)
	if err != nil {
		return fmt.Errorf("repository: Replace position: Exec: %w", err)
	}

	if commandTag.RowsAffected() != 1 {
		return entity.ErrNothingToChange
	}

	return nil
}
