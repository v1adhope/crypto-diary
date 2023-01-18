package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/v1adhope/crypto-diary/internal/entity"
	"github.com/v1adhope/crypto-diary/internal/usecase"
	"github.com/v1adhope/crypto-diary/pkg/postgres"
)

const (
	defaultEntityCap = 25
	timeModel        = "2006-01-02"
)

type PositionRepo struct {
	*postgres.Postgres
}

func (pr *PositionRepo) Create(ctx context.Context, position *entity.Position) error {
	q := `INSERT INTO positions(open_date, pair, reason, according_to_plan, percentage_risk,
                             direction, deposit, open_price, stop_loss_price,
                             take_profit_price, close_price, user_id)
        VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
        RETURNING position_id`

	err := pr.Pool.QueryRow(ctx, q,
		position.OpenDate,
		position.Pair,
		position.Reason,
		position.AccordingToPlan,
		position.Risk,
		position.Direction,
		position.Deposit,
		position.OpenPrice,
		position.StopLossPrice,
		position.TakeProfitPrice,
		position.ClosePrice,
		position.UserID).Scan(&position.ID)
	if err != nil {
		return fmt.Errorf("sql request: Create position: QueryRow: %s", err)
	}

	return nil
}

func (pr *PositionRepo) FindAll(ctx context.Context) ([]entity.Position, error) {
	q := `SELECT * FROM all_positions`

	rows, err := pr.Pool.Query(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("sql request: FinAll positions: Query: %s", err)
	}
	defer rows.Close()

	positions := make([]entity.Position, 0, defaultEntityCap)

	for rows.Next() {
		//TODO: Create DTO???
		var (
			p               entity.Position
			openDate        pgtype.Date
			accordingToPlan bool
			closePrice      pgtype.Text
		)

		err := rows.Scan(
			&p.ID,
			&openDate,
			&p.Pair,
			&p.Reason,
			&accordingToPlan,
			&p.Risk,
			&p.Direction,
			&p.Deposit,
			&p.OpenPrice,
			&p.StopLossPrice,
			&p.TakeProfitPrice,
			&closePrice,
			&p.UserID)
		if err != nil {
			return nil, fmt.Errorf("sql request: FindAll positons: Scan: %s", err)
		}

		p.OpenDate = fmt.Sprintf("%s", openDate.Time.Format(timeModel))
		p.AccordingToPlan = fmt.Sprintf("%t", accordingToPlan)
		if closePrice.Valid {
			p.ClosePrice = fmt.Sprintf("%s", closePrice.String)
		}

		positions = append(positions, p)
	}

	return positions, nil
}

// TODO: commandTag
func (pr *PositionRepo) Delete(ctx context.Context, ID *string) error {
	q := `DELETE FROM positions
        WHERE position_id = $1`

	commandTag, err := pr.Pool.Exec(ctx, q, ID)
	if err != nil {
		return fmt.Errorf("sql request: Delete positons: Scan: %s", err)
	}
	if commandTag.RowsAffected() != 1 {
		return fmt.Errorf("No row found to delete")
	}

	return nil
}

func NewPosition(pg *postgres.Postgres) usecase.PositionRepo {
	return &PositionRepo{pg}
}
