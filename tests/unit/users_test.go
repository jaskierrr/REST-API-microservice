package handlers_test

import (
	"card-project/models"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestUserPostValidation(t *testing.T) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	cases := []struct {
		name    string
		payload models.NewUser
		wantErr bool
	}{
		{
			name: "valid user",
			payload: models.NewUser{
				FirstName: "Ivan",
				LastName:  "Shash",
			},
			wantErr: false,
		},
		{
			name: "missing first name",
			payload: models.NewUser{
				FirstName: "",
				LastName:  "Shash",
			},
			wantErr: true,
		},
		{
			name: "missing last name",
			payload: models.NewUser{
				FirstName: "Ivan",
				LastName:  "",
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

				// t.Logf("Generated error message:\n %v", errors.String())

			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
			}
		})
	}
}
