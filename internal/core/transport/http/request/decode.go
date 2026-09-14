package core_http_request

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	core_errors "github.com/maximyunak/qzltgo/internal/core/errors"
)

var RequestValidator = validator.New()

func DecodeAndValidateRequest(r *http.Request, dest any) error {
	if err := json.NewDecoder(r.Body).Decode(dest); err != nil {
		return fmt.Errorf("decode json: %v: %w", err, core_errors.ErrInvalidArgument)
	}

	if err := RequestValidator.Struct(dest); err != nil {
		return fmt.Errorf("validate json: %v :%w", err, core_errors.ErrInvalidArgument)
	}

	return nil
}
