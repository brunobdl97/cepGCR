package internal

import "testing"

func TestNewCep(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"valid 8 digits", "12345678", false},
		{"too short", "123", true},
		{"too long", "123456789", true},
		{"non digit", "12345a78", true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cep, err := NewCep(tc.input)
			if tc.wantErr && err == nil {
				t.Fatalf("expected error for %s", tc.input)
			}
			if !tc.wantErr {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if cep.Get() != tc.input {
					t.Fatalf("expected %s, got %s", tc.input, cep.Get())
				}
			}
		})
	}
}
