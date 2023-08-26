package middleware

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type FieldValidationError struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

type InputValidatorType int

const (
	BodyValidator InputValidatorType = iota
	ParamsValidator
	QueryValidator
)

func InputValidatorMiddleware(
	customValidator *validator.Validate,
	validatorType InputValidatorType,
	data interface{},
) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var v *validator.Validate

		if customValidator != nil {
			v = customValidator
		} else {
			v = validator.New()
		}

		var validationData interface{}

		validationData = reflect.New(reflect.TypeOf(data).Elem()).Interface()

		switch validatorType {
		case BodyValidator:
			if err := c.BodyParser(validationData); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error":   err.Error(),
					"code":    fiber.StatusBadRequest,
					"message": "BODY_PARSING_FAILED",
				})
			}
		case ParamsValidator:
			if err := c.ParamsParser(validationData); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error":   err.Error(),
					"code":    fiber.StatusBadRequest,
					"message": "PARAMS_PARSING_FAILED",
				})
			}
		case QueryValidator:
			if err := c.QueryParser(validationData); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error":   err.Error(),
					"code":    fiber.StatusBadRequest,
					"message": "QUERY_PARAMS_PARSING_FAILED",
				})
			}
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   "Invalid Validation Type",
				"code":    fiber.StatusInternalServerError,
				"message": "INVALID_VALIDATION_TYPE",
			})
		}

		if err := v.Struct(validationData); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   fieldErrorsFormatter(err.(validator.ValidationErrors)),
				"code":    fiber.StatusBadRequest,
				"message": "INPUT_VALIDATION_ERROR",
			})
		}

		return c.Next()
	}
}

func fieldErrorsFormatter(currentErrors validator.ValidationErrors) []*FieldValidationError {
	var errors []*FieldValidationError

	for _, err := range currentErrors {
		var el FieldValidationError
		el.Field = err.Field()
		el.Tag = err.Tag()
		el.Value = err.Param()
		errors = append(errors, &el)
	}

	return errors
}
