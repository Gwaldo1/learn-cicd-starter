package auth

import (
    "net/http"
    "testing"
)

func TestGetAPIKey(t *testing.T) {
    tests := []struct {
        name    string
        input   string
        want    string
        wantErr bool
    }{
        {name: "valid", input: "ApiKey abc123", want: "abc123", wantErr: false},
        {name: "empty", input: "", want: "", wantErr: true},
        {name: "invalid", input: "Basic xyz789", want: "", wantErr: true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            headers := make(http.Header)
            if tt.input != "" {
                headers.Set("Authorization", tt.input)
            }
            got, err := GetAPIKey(headers)
            if (err != nil) != tt.wantErr {
                t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
            }
        })
    }
}