package bdd

import (
	"context"
	"fase-4-hf-order/external/db/rds/postgres"
	l "fase-4-hf-order/external/logger"
	ps "fase-4-hf-order/external/strings"
	repositories "fase-4-hf-order/internal/adapters/driven/repositories/sql"
	"fase-4-hf-order/internal/core/application"
	"fase-4-hf-order/internal/core/domain/entity"
	"fase-4-hf-order/internal/core/domain/entity/dto"
	"fase-4-hf-order/internal/core/useCase"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/marcos-dev88/genv"
)

// go test -v -count=1 -failfast -run ^Test_GetOrderByID$
func Test_GetOrderByID(t *testing.T) {
	if err := genv.New("../../.env"); err != nil {
		l.Errorf("error set envs %v", " | ", err)
	}

	l.Info("====> TEST GetOrderByID <====")

	type Input struct {
		ID int64 `json:"id"`
	}

	type Output struct {
		Output *dto.OutputOrder `json:"output"`
	}

	tests := []struct {
		scenario          string
		name              string
		input             Input
		shouldReturnError bool
		shouldBeNull      bool
		expectedOutput    Output
	}{
		{
			scenario: "Sending a valid and existing ID",
			name:     "success_valid_id",
			input: Input{
				ID: 1,
			},
			shouldReturnError: false,
			shouldBeNull:      false,
			expectedOutput: Output{
				Output: &dto.OutputOrder{
					ID:          1,
					ClientUUID:  "1",
					VoucherUUID: "1",
					Items: []entity.OrderItems{
						{
							ProductUUID: "1",
							Quantity:    1,
							TotalPrice:  36,
							Discount:    0.0,
						},
					},
					Status:           "RECEIVED",
					VerificationCode: "abc001",
				},
			},
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(*testing.T) {

			// config
			ctx := context.Background()
			dbDuration, err := time.ParseDuration(os.Getenv("DB_DURATION"))

			if err != nil {
				log.Fatalf("error: %v", err)

			}

			db := postgres.NewPostgresDB(
				ctx,
				os.Getenv("DB_REGION"),
				os.Getenv("DB_HOST"),
				os.Getenv("DB_PORT"),
				os.Getenv("DB_NAME"),
				os.Getenv("DB_USER"),
				os.Getenv("DB_PASSWORD"),
				dbDuration,
			)

			repoOrder := repositories.NewOrderDB(ctx, db)

			repoOrderItem := repositories.NewOrderItemDB(ctx, db)

			uc := useCase.NewOrderUseCase()

			app := application.NewApplication(repoOrder, repoOrderItem, uc)
			// final config

			l.Info("----------------")
			l.Info(fmt.Sprintf("-> Scenario: %s", tc.scenario))
			l.Info(fmt.Sprintf("-> Input: %s", ps.MarshalString(tc.input)))
			l.Info(fmt.Sprintf("-> ExpectedOutput: %s", ps.MarshalString(tc.expectedOutput)))
			l.Info(fmt.Sprintf("-> Should return error: %v", tc.shouldReturnError))
			l.Info("----------------")

			order, err := app.GetOrderByID(tc.input.ID)

			if (!tc.shouldReturnError) && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !tc.shouldBeNull {
				if order.ClientUUID != tc.expectedOutput.Output.ClientUUID {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.ClientUUID, order.ClientUUID)
				}

				if order.VoucherUUID != tc.expectedOutput.Output.VoucherUUID {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.VoucherUUID, order.VoucherUUID)
				}

				if order.Status != tc.expectedOutput.Output.Status {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.Status, order.Status)
				}

				if order.VerificationCode != tc.expectedOutput.Output.VerificationCode {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.VerificationCode, order.VerificationCode)
				}

			}
			l.Info(fmt.Sprintf("====> Success running scenario: [%s] <====", tc.scenario))
		})
		l.Info("====> Finish BDD Test GetOrderByID <====")
	}
}

// go test -v -count=1 -failfast -run ^Test_SaveOrder$
func Test_SaveOrder(t *testing.T) {
	t.Skip("skipping this test, comment it to test it properly")

	if err := genv.New("../../.env"); err != nil {
		l.Errorf("error set envs %v", " | ", err)
	}

	l.Info("====> TEST SaveOrder <====")

	type Input struct {
		Input *dto.RequestOrder `json:"input"`
	}

	type Output struct {
		Output *dto.OutputOrder `json:"output"`
	}

	tests := []struct {
		scenario          string
		name              string
		input             Input
		shouldReturnError bool
		shouldBeNull      bool
		expectedOutput    Output
	}{
		{
			scenario: "Sending a valid input",
			name:     "success",
			input: Input{
				Input: &dto.RequestOrder{
					ClientUUID:  "1",
					VoucherUUID: "1",
					Items: []entity.OrderItems{
						{
							ProductUUID: "1",
							Quantity:    1,
							TotalPrice:  0.0,
							Discount:    0.0,
						},
					},
					Status:           "In progress",
					VerificationCode: "abc123",
				},
			},
			shouldReturnError: false,
			shouldBeNull:      false,
			expectedOutput: Output{
				Output: &dto.OutputOrder{
					ClientUUID:  "1",
					VoucherUUID: "1",
					Items: []entity.OrderItems{
						{
							ProductUUID: "1",
							Quantity:    1,
							TotalPrice:  0.0,
							Discount:    0.0,
						},
					},
					Status:           "In progress",
					VerificationCode: "abc123",
					CreatedAt:        "",
				},
			},
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(*testing.T) {

			// config
			ctx := context.Background()
			dbDuration, err := time.ParseDuration(os.Getenv("DB_DURATION"))

			if err != nil {
				log.Fatalf("error: %v", err)

			}

			db := postgres.NewPostgresDB(
				ctx,
				os.Getenv("DB_REGION"),
				os.Getenv("DB_HOST"),
				os.Getenv("DB_PORT"),
				os.Getenv("DB_NAME"),
				os.Getenv("DB_USER"),
				os.Getenv("DB_PASSWORD"),
				dbDuration,
			)

			repoOrder := repositories.NewOrderDB(ctx, db)

			repoOrderItem := repositories.NewOrderItemDB(ctx, db)

			uc := useCase.NewOrderUseCase()

			app := application.NewApplication(repoOrder, repoOrderItem, uc)
			// final config

			l.Info("----------------")
			l.Info(fmt.Sprintf("-> Scenario: %s", tc.scenario))
			l.Info(fmt.Sprintf("-> Input: %s", ps.MarshalString(tc.input)))
			l.Info(fmt.Sprintf("-> ExpectedOutput: %s", ps.MarshalString(tc.expectedOutput)))
			l.Info(fmt.Sprintf("-> Should return error: %v", tc.shouldReturnError))
			l.Info("----------------")

			order, err := app.SaveOrder(*tc.input.Input)

			if (!tc.shouldReturnError) && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !tc.shouldBeNull {
				if order.ClientUUID != tc.expectedOutput.Output.ClientUUID {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.ClientUUID, order.ClientUUID)
				}

				if order.VoucherUUID != tc.expectedOutput.Output.VoucherUUID {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.VoucherUUID, order.VoucherUUID)
				}

				if order.Status != tc.expectedOutput.Output.Status {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.Status, order.Status)
				}

				if order.VerificationCode != tc.expectedOutput.Output.VerificationCode {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.VerificationCode, order.VerificationCode)
				}

			}
			l.Info(fmt.Sprintf("====> Success running scenario: [%s] <====", tc.scenario))
		})
		l.Info("====> Finish BDD Test SaveOrder <====")
	}
}

// go test -v -count=1 -failfast -run ^Test_UpdateOrderByID$
func Test_UpdateOrderByID(t *testing.T) {
	t.Skip("skipping this test, comment it to test it properly")
	if err := genv.New("../../.env"); err != nil {
		l.Errorf("error set envs %v", " | ", err)
	}

	l.Info("====> TEST UpdateOrderByID <====")

	type Input struct {
		ID    int64             `json:"id"`
		Input *dto.RequestOrder `json:"input"`
	}

	type Output struct {
		Output *dto.OutputOrder `json:"output"`
	}

	tests := []struct {
		scenario          string
		name              string
		input             Input
		shouldReturnError bool
		shouldBeNull      bool
		expectedOutput    Output
	}{
		{
			scenario: "Sending a valid input",
			name:     "success",
			input: Input{
				ID: 2,
				Input: &dto.RequestOrder{
					ClientUUID:       "1",
					VoucherUUID:      "2",
					VerificationCode: "bcd002",
				},
			},
			shouldReturnError: false,
			shouldBeNull:      false,
			expectedOutput: Output{
				Output: &dto.OutputOrder{
					ClientUUID:       "1",
					VoucherUUID:      "2",
					Items:            []entity.OrderItems{},
					Status:           "IN PROGRESS",
					VerificationCode: "bcd002",
				},
			},
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(*testing.T) {

			// config
			ctx := context.Background()
			dbDuration, err := time.ParseDuration(os.Getenv("DB_DURATION"))

			if err != nil {
				log.Fatalf("error: %v", err)

			}

			db := postgres.NewPostgresDB(
				ctx,
				os.Getenv("DB_REGION"),
				os.Getenv("DB_HOST"),
				os.Getenv("DB_PORT"),
				os.Getenv("DB_NAME"),
				os.Getenv("DB_USER"),
				os.Getenv("DB_PASSWORD"),
				dbDuration,
			)

			repoOrder := repositories.NewOrderDB(ctx, db)

			repoOrderItem := repositories.NewOrderItemDB(ctx, db)

			uc := useCase.NewOrderUseCase()

			app := application.NewApplication(repoOrder, repoOrderItem, uc)
			// final config

			l.Info("----------------")
			l.Info(fmt.Sprintf("-> Scenario: %s", tc.scenario))
			l.Info(fmt.Sprintf("-> Input: %s", ps.MarshalString(tc.input)))
			l.Info(fmt.Sprintf("-> ExpectedOutput: %s", ps.MarshalString(tc.expectedOutput)))
			l.Info(fmt.Sprintf("-> Should return error: %v", tc.shouldReturnError))
			l.Info("----------------")

			order, err := app.UpdateOrderByID(tc.input.ID, *tc.input.Input)

			if (!tc.shouldReturnError) && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !tc.shouldBeNull {
				if order.ClientUUID != tc.expectedOutput.Output.ClientUUID {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.ClientUUID, order.ClientUUID)
				}

				if order.VoucherUUID != tc.expectedOutput.Output.VoucherUUID {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.VoucherUUID, order.VoucherUUID)
				}

				if order.Status != tc.expectedOutput.Output.Status {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.Status, order.Status)
				}

				if order.VerificationCode != tc.expectedOutput.Output.VerificationCode {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.VerificationCode, order.VerificationCode)
				}

			}
			l.Info(fmt.Sprintf("====> Success running scenario: [%s] <====", tc.scenario))
		})
		l.Info("====> Finish BDD Test UpdateOrderByID <====")
	}
}

// go test -v -count=1 -failfast -run ^Test_GetOrders$
func Test_GetOrders(t *testing.T) {
	if err := genv.New("../../.env"); err != nil {
		l.Errorf("error set envs %v", " | ", err)
	}

	l.Info("====> TEST GetOrders <====")

	type Input struct {
		Category string `json:"category"`
	}

	type Output struct {
		Output []dto.OutputOrder `json:"output"`
	}

	tests := []struct {
		scenario          string
		name              string
		input             Input
		shouldReturnError bool
		shouldBeNull      bool
		expectedOutput    Output
	}{
		{
			scenario: "Sending a valid and existing category",
			name:     "success_valid_id",
			input: Input{
				Category: "meal",
			},
			shouldReturnError: false,
			shouldBeNull:      false,
			expectedOutput: Output{
				Output: []dto.OutputOrder{
					{
						ClientUUID:       "1",
						VoucherUUID:      "1",
						Items:            []entity.OrderItems{},
						Status:           "RECEIVED",
						VerificationCode: "abc001",
						CreatedAt:        "",
					},
					{
						ClientUUID:       "3",
						VoucherUUID:      "3",
						Items:            []entity.OrderItems{},
						Status:           "DONE",
						VerificationCode: "cde003",
						CreatedAt:        "",
					},
					{
						ClientUUID:       "1",
						VoucherUUID:      "2",
						Items:            []entity.OrderItems{},
						Status:           "IN PROGRESS",
						VerificationCode: "bcd002",
						CreatedAt:        "",
					},
					{
						ClientUUID:       "",
						VoucherUUID:      "",
						Items:            []entity.OrderItems{},
						Status:           "FINISHED",
						VerificationCode: "cde003",
						CreatedAt:        "",
					},
					{
						ClientUUID:       "",
						VoucherUUID:      "",
						Items:            []entity.OrderItems{},
						Status:           "In progress",
						VerificationCode: "cde003",
						CreatedAt:        "",
					},
				},
			},
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(*testing.T) {

			// config
			ctx := context.Background()
			dbDuration, err := time.ParseDuration(os.Getenv("DB_DURATION"))

			if err != nil {
				log.Fatalf("error: %v", err)

			}

			db := postgres.NewPostgresDB(
				ctx,
				os.Getenv("DB_REGION"),
				os.Getenv("DB_HOST"),
				os.Getenv("DB_PORT"),
				os.Getenv("DB_NAME"),
				os.Getenv("DB_USER"),
				os.Getenv("DB_PASSWORD"),
				dbDuration,
			)

			repoOrder := repositories.NewOrderDB(ctx, db)

			repoOrderItem := repositories.NewOrderItemDB(ctx, db)

			uc := useCase.NewOrderUseCase()

			app := application.NewApplication(repoOrder, repoOrderItem, uc)
			// final config

			l.Info("----------------")
			l.Info(fmt.Sprintf("-> Scenario: %s", tc.scenario))
			l.Info(fmt.Sprintf("-> Input: %s", ps.MarshalString(tc.input)))
			l.Info(fmt.Sprintf("-> ExpectedOutput: %s", ps.MarshalString(tc.expectedOutput)))
			l.Info(fmt.Sprintf("-> Should return error: %v", tc.shouldReturnError))
			l.Info("----------------")

			orderList, err := app.GetOrders()

			if (!tc.shouldReturnError) && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !tc.shouldBeNull {
				for i := range orderList {
					l.Debugf("order: ", " | ", ps.MarshalString(orderList[i]), "item: ", i+1)
					if orderList[i].ClientUUID != tc.expectedOutput.Output[i].ClientUUID {
						t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output[i].ClientUUID, orderList[i].ClientUUID)
					}

					if orderList[i].VoucherUUID != tc.expectedOutput.Output[i].VoucherUUID {
						t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output[i].VoucherUUID, orderList[i].VoucherUUID)
					}

					if orderList[i].VerificationCode != tc.expectedOutput.Output[i].VerificationCode {
						t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output[i].VerificationCode, orderList[i].VerificationCode)
					}

					if orderList[i].Status != tc.expectedOutput.Output[i].Status {
						t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output[i].Status, orderList[i].Status)
					}

				}
			}
			l.Info(fmt.Sprintf("====> Success running scenario: [%s] <====", tc.scenario))
		})
		l.Info("====> Finish BDD Test GetOrders <====")
	}
}
