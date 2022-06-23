package presenter

import (
	"net/http"

	"github.com/8mamo10/gcl/internal/presenter/schema"
	"github.com/labstack/echo/v4"
)

func (p *Presenter) GetUsers(c echo.Context, params schema.GetUsersParams) error {
	ctx := c.Request().Context()
	limit := defaultLimit
	if params.Limit != nil {
		limit = int(*params.Limit)
	}
	offset := defaultOffset
	if params.Offset != nil {
		offset = int(*params.Offset)
	}

	results, err := p.userUC.GetAll(ctx, limit, offset)
	if err != nil {
		return err
	}

	res := make(schema.UsersResponse, len(results))

	for i, result := range results {
		u := &schema.User{
			Id:        result.ID,
			AccountId: result.AccountID,
			Name:      result.Name,
			Email:     result.Email,
		}

		res[i] = *u
	}

	return c.JSON(http.StatusOK, res)
}
