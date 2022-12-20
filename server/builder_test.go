package server

import (
	"reflect"
	"testing"
)

func Test_buildQueryFilter(t *testing.T) {
	type args struct {
		query map[string][]string
	}
	tests := []struct {
		name       string
		args       args
		wantResult map[string]string
	}{
		{
			name: "success",
			args: args{
				query: map[string][]string{
					"name":     {"ganteng"},
					"options":  {"1"},
					"type":     {"1,2,3"},
					"sort_by":  {"id"},
					"order_by": {"desc"},
				},
			},
			wantResult: map[string]string{
				"name":     "ganteng",
				"options":  "1",
				"order_by": "desc",
				"sort_by":  "id",
				"type":     "1,2,3",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := buildQueryFilter(tt.args.query); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("buildQueryFilter() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
