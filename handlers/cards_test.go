package handlers

import (
	"card-project/models"
	"strings"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/go-playground/validator/v10"
)

func init() {
	validate = validator.New()
}

func TestCardPostValidation(t *testing.T) {
	cases := []struct {
		name    string
		payload models.NewCard
		wantErr bool
	}{
		{
			name: "valid card",
			payload: models.NewCard{
				UserID: 1,
				BankID: 1,
				Number: 1234,
				CreateDate: strfmt.NewDateTime(),
			},
			wantErr: false,
		},
		{
			name: "missing user id",
			payload: models.NewCard{
				UserID: 0,
				BankID: 1,
				Number: 1234,
				CreateDate: strfmt.NewDateTime(),
			},
			wantErr: true,
		},
		{
			name: "missing bank id",
			payload: models.NewCard{
				UserID: 1,
				BankID: 0,
				Number: 1234,
				CreateDate: strfmt.NewDateTime(),
			},
			wantErr: true,
		},
		{
			name: "missing number",
			payload: models.NewCard{
				UserID: 1,
				BankID: 1,
				CreateDate: strfmt.NewDateTime(),
			},
			wantErr: true,
		},
		{
			name: "missing create date",
			payload: models.NewCard{
				UserID: 1,
				BankID: 1,
				Number: 1234,
			},
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := validate.Struct(c.payload)

			if c.wantErr {
				if err == nil {
					t.Fatalf("expected error but got nil")
				}

				validationErrors := err.(validator.ValidationErrors)
				errors := strings.Builder{}
				for _, e := range validationErrors {
					errors.WriteString(e.Error())
				}

				t.Logf("Generated error message:\n %v", errors.String())

			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
			}
		})
	}
}
