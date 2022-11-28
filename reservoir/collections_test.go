package reservoir

import "testing"

func TestGetCollections(t *testing.T) {
	type args struct {
		slug string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "happy path",
			args: args{
				slug: "azuki",
			},
			want: "azuki",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := GenerateReservoirClient()
			got, err := c.GetCollections(tt.args.slug)
			if err != nil {
				t.Errorf("GetCollections() error = %v", err)
				return
			}
			if got.Collections[0].Slug != tt.want {
				t.Errorf("GetCollections() = %v, want %v", got.Collections[0].Slug, tt.want)
			}
		})
	}
}
