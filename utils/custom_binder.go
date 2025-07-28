package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/gommon/log"

	"github.com/labstack/echo/v4"
	"io"
)

type CustomBinder struct{}

func (cb *CustomBinder) Bind(i interface{}, c echo.Context) error {
	var body, err = io.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(io.NopCloser(io.MultiReader(io.NewSectionReader(bytes.NewReader(body), 0, int64(len(body))))))
	decoder.DisallowUnknownFields() // 🚀 Strict JSON validation

	if err := decoder.Decode(i); err != nil {
		return fmt.Errorf("invalid JSON: check field names")
	}
	err = c.Validate(i)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
