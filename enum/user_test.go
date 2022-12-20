package enum

import "testing"

func TestRole_String(t *testing.T) {
	tests := []struct {
		name string
		r    Role
		want string
	}{
		{
			name: "success admin role",
			r:    Admin,
			want: "admin",
		},
		{
			name: "success admin role",
			r:    User,
			want: "user",
		},
		{
			name: "success admin role",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.String(); got != tt.want {
				t.Errorf("Role.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
