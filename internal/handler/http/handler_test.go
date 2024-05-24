package http

import (
	"errors"
	"fase-4-hf-order/internal/core/domain/entity/dto"
	"fase-4-hf-order/mocks"
	"io"
	"net/http/httptest"
	"strings"
	"testing"
)

// go test -v -count=1 -failfast -run ^Test_Handler$
func Test_Handler(t *testing.T) {
	type args struct {
		method string
		url    string
		body   io.Reader
	}

	tests := []struct {
		name            string
		args            args
		mockApplication mocks.MockApplication
		wantOut         string
		isWantedErr     bool
	}{
		{
			name: "success_getOrders",
			args: args{
				method: "GET",
				url:    "hermes_foods/order",
				body:   nil,
			},
			mockApplication: mocks.MockApplication{
				WantOutList: []dto.OutputOrder{
					{
						ID:               1,
						ClientUUID:       "1",
						VoucherUUID:      "2",
						Status:           "Paid",
						VerificationCode: "aab231",
					},
				},
				WantOut:     &dto.OutputOrder{},
				WantErr:     nil,
				WantOutNull: "",
			},
			wantOut:     `[{"id":1,"clientUuid":"1","voucherUuid":"2","status":"Paid","verificationCode":"aab231"}]`,
			isWantedErr: false,
		},
		{
			name: "nil_getOrders",
			args: args{
				method: "GET",
				url:    "hermes_foods/order",
				body:   nil,
			},
			mockApplication: mocks.MockApplication{
				WantOutList: nil,
				WantOut:     &dto.OutputOrder{},
				WantErr:     nil,
				WantOutNull: "",
			},
			wantOut:     `{"error": "order not found"}`,
			isWantedErr: false,
		},
		{
			name: "error_getOrders",
			args: args{
				method: "GET",
				url:    "hermes_foods/order",
				body:   nil,
			},
			mockApplication: mocks.MockApplication{
				WantOut:     nil,
				WantErr:     errors.New("errGetOrders"),
				WantOutNull: "",
			},
			wantOut:     `{"error": "error to get order by ID: errGetOrders"}`,
			isWantedErr: false,
		},
		{
			name: "success_getByID",
			args: args{
				method: "GET",
				url:    "hermes_foods/order/100000",
				body:   nil,
			},
			mockApplication: mocks.MockApplication{
				WantOut:     &dto.OutputOrder{},
				WantErr:     nil,
				WantOutNull: "",
			},
			wantOut:     "{}",
			isWantedErr: false,
		},
		{
			name: "order_null_getByID",
			args: args{
				method: "GET",
				url:    "hermes_foods/order/100000",
				body:   nil,
			},
			mockApplication: mocks.MockApplication{
				WantOut:     nil,
				WantErr:     nil,
				WantOutNull: "",
			},
			wantOut:     `{"error": "order not found"}`,
			isWantedErr: false,
		},
		{
			name: "error_getByID",
			args: args{
				method: "GET",
				url:    "hermes_foods/order/100000",
				body:   nil,
			},
			mockApplication: mocks.MockApplication{
				WantOut:     nil,
				WantErr:     errors.New("errGetOrderByID"),
				WantOutNull: "",
			},
			wantOut:     `{"error": "error to get order by ID: errGetOrderByID"}`,
			isWantedErr: false,
		},
		{
			name: "success_save",
			args: args{
				method: "POST",
				url:    "hermes_foods/order",
				body:   strings.NewReader(`{}`),
			},
			mockApplication: mocks.MockApplication{
				WantOut: &dto.OutputOrder{
					CreatedAt: "",
				},
				WantErr:     nil,
				WantOutNull: "",
			},
			wantOut:     `{}`,
			isWantedErr: false,
		},
		{
			name: "error_save_unmarshal",
			args: args{
				method: "POST",
				url:    "hermes_foods/order/",
				body:   strings.NewReader(`<=>`),
			},
			mockApplication: mocks.MockApplication{
				WantOut: &dto.OutputOrder{
					ID:               0,
					ClientUUID:       "",
					VoucherUUID:      "",
					Items:            nil,
					Status:           "",
					VerificationCode: "",
					CreatedAt:        "",
				},
				WantErr:     nil,
				WantOutNull: "",
			},
			wantOut:     `{"error": "error to Unmarshal: invalid character '<' looking for beginning of value"}`,
			isWantedErr: true,
		},
		{
			name: "error_save",
			args: args{
				method: "POST",
				url:    "hermes_foods/order",
				body:   strings.NewReader(`{"name":"Marty", "cpf":"051119995", "email": "martybttf@bttf.com"}`),
			},
			mockApplication: mocks.MockApplication{
				WantOut: &dto.OutputOrder{
					CreatedAt: "",
				},
				WantErr:     errors.New("errSaveOrder"),
				WantOutNull: "",
			},
			wantOut:     `{"error": "error to save order: errSaveOrder"}`,
			isWantedErr: false,
		},
		{
			name: "success_update_by_id",
			args: args{
				method: "PATCH",
				url:    "hermes_foods/order/10000",
				body:   strings.NewReader(`{}`),
			},
			mockApplication: mocks.MockApplication{
				WantOut: &dto.OutputOrder{
					CreatedAt: "",
				},
				WantErr:     nil,
				WantOutNull: "",
			},
			wantOut:     `{}`,
			isWantedErr: false,
		},
		{
			name: "nil_update_by_id",
			args: args{
				method: "PATCH",
				url:    "hermes_foods/order/10000",
				body:   strings.NewReader(`{}`),
			},
			mockApplication: mocks.MockApplication{
				WantOut:     nil,
				WantErr:     nil,
				WantOutNull: "",
			},
			wantOut:     `{"error": "order not found"}`,
			isWantedErr: false,
		},
		{
			name: "error_update_by_id",
			args: args{
				method: "PATCH",
				url:    "hermes_foods/order/10000",
				body:   strings.NewReader(`{"name":"Marty", "cpf":"051119995", "email": "martybttf@bttf.com"}`),
			},
			mockApplication: mocks.MockApplication{
				WantOut: &dto.OutputOrder{
					CreatedAt: "",
				},
				WantErr:     errors.New("errUpdateOrderByID"),
				WantOutNull: "",
			},
			wantOut:     `{"error": "error to get order by ID: errUpdateOrderByID"}`,
			isWantedErr: false,
		},
		{
			name: "error_route_not_found",
			args: args{
				method: "PATCH",
				url:    "/hermes_foods/order",
				body:   strings.NewReader(`{"name":"Marty", "cpf":"051119995", "email": "martybttf@bttf.com"}`),
			},
			mockApplication: mocks.MockApplication{
				WantOut: &dto.OutputOrder{
					CreatedAt: "",
				},
				WantErr:     errors.New("errSaveOrder"),
				WantOutNull: "",
			},
			wantOut:     `{"error": "route PATCH /hermes_foods/order not found"}`,
			isWantedErr: false,
		},
	}

	for _, tc := range tests {
		h := NewHandler(tc.mockApplication)
		t.Run(tc.name, func(*testing.T) {

			req := httptest.NewRequest(tc.args.method, "/", tc.args.body)
			req.URL.Path = tc.args.url
			rw := httptest.NewRecorder()

			h.Handler(rw, req)

			response := rw.Result()
			defer response.Body.Close()

			b, err := io.ReadAll(response.Body)

			if (!tc.isWantedErr) && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if strings.TrimSpace(string(b)) != strings.TrimSpace(tc.wantOut) {
				t.Errorf("expected: %s\ngot: %s", tc.wantOut, string(b))

			}

		})
	}
}

// go test -v -count=1 -failfast -run ^Test_HealthCheck$
func Test_HealthCheck(t *testing.T) {
	type args struct {
		method string
		url    string
		body   io.Reader
	}
	tests := []struct {
		name            string
		args            args
		wantOut         string
		mockApplication mocks.MockApplication
		isWantedErr     bool
	}{
		{
			name: "success",
			args: args{
				method: "GET",
				url:    "/",
				body:   nil,
			},
			wantOut:     `{"status": "OK"}`,
			isWantedErr: false,
		},
		{
			name: "error_method_not_allowed",
			args: args{
				method: "POST",
				url:    "/",
				body:   nil,
			},
			wantOut:     `{"error": "method not allowed"}`,
			isWantedErr: true,
		},
	}

	for _, tc := range tests {
		h := NewHandler(tc.mockApplication)
		t.Run(tc.name, func(*testing.T) {

			req := httptest.NewRequest(tc.args.method, tc.args.url, tc.args.body)
			rw := httptest.NewRecorder()

			h.HealthCheck(rw, req)

			response := rw.Result()
			defer response.Body.Close()

			b, err := io.ReadAll(response.Body)

			if (!tc.isWantedErr) && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if strings.TrimSpace(string(b)) != strings.TrimSpace(tc.wantOut) {
				t.Errorf("expected: %s\ngot: %s", tc.wantOut, string(b))

			}
		})
	}
}
