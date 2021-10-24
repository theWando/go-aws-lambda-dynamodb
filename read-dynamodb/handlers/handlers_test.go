package handlers

import (
	"context"
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandleRead(t *testing.T) {
	type args struct {
		req events.APIGatewayProxyRequest
	}
	tests := []struct {
		name    string
		args    args
		want    events.APIGatewayProxyResponse
		wantErr bool
		errMsg  string
	}{
		{
			name: "Unable to get IP",
			args: args{
				req: events.APIGatewayProxyRequest{},
			},
			want:    events.APIGatewayProxyResponse{},
			wantErr: true,
		},
		{
			name: "Non 200 Response",
			args: args{
				req: events.APIGatewayProxyRequest{},
			},
			want:    events.APIGatewayProxyResponse{},
			wantErr: false,
			errMsg:  errNon200Response.Error(),
		},
		{
			name: "Unable decode IP",
			args: args{
				req: events.APIGatewayProxyRequest{},
			},
			want:    events.APIGatewayProxyResponse{},
			wantErr: false,
			errMsg:  errNon200Response.Error(),
		},
		{
			name: "Successful Request",
			args: args{
				req: events.APIGatewayProxyRequest{},
			},
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Body:       "Hello, 127.0.0.1",
			},
			wantErr: false,
			errMsg:  errNon200Response.Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewReadRestaurants(nil)(context.Background(), tt.args.req)
			if (err != nil) != tt.wantErr && tt.errMsg != err.Error() {
				t.Errorf("NewReadRestaurants() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReadRestaurants() got = %v, want %v", got, tt.want)
			}
		})
	}
}
