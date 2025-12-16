package sanitize

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Address struct {
	Line1    string `sanitize:"trim_space"`
	Line2    string `sanitize:"trim_space"`
	City     string `sanitize:"trim_space"`
	Postcode string `sanitize:"trim_space,upper,strip_space"`
}

type User struct {
	FirstName       *string  `sanitize:"trim_space,capitalize"`
	LastName        string   `sanitize:"trim_space,capitalize"`
	DOB             string   `sanitize:"-"`
	Email           string   `sanitize:"lower,trim_space"`
	Phone           string   `sanitize:"trim_space,strip_space"`
	Tags            []string `sanitize:"dive,trim_space,capitalize"`
	Address         Address  `sanitize:"dive"`
	OptionalAddress *Address `sanitize:"dive"`
}

func TestApply(t *testing.T) {
	type args struct {
		v any
	}
	tests := []struct {
		name    string
		args    args
		want    any
		wantErr bool
	}{

		{
			name: "invalid case - not a pointer",
			args: args{
				v: User{},
			},
			wantErr: true,
		},
		{
			name: "invalid case - nil pointer",
			args: args{
				v: nil,
			},
			wantErr: true,
		},
		{
			name: "valid case, optional address as nil, tags empty",
			args: args{
				v: &User{
					FirstName: ptr("Codex "),
					LastName:  "monastery ",
					DOB:       "2025-11-24",
					Email:     "Codex.Monastery@email.com",
					Phone:     "+44 00000 00000",
					Address: Address{
						Line1:    "123 High Street ",
						Line2:    " ",
						City:     "London ",
						Postcode: "SW1A 1aa",
					},
				},
			},
			want: &User{
				FirstName: ptr("Codex"),
				LastName:  "Monastery",
				DOB:       "2025-11-24",
				Email:     "codex.monastery@email.com",
				Phone:     "+440000000000",
				Address: Address{
					Line1:    "123 High Street",
					Line2:    "",
					City:     "London",
					Postcode: "SW1A1AA",
				},
			},
			wantErr: false,
		},
		{
			name: "valid case, optional address as not nil, tags not empty",
			args: args{
				v: &User{
					FirstName: ptr("Codex "),
					LastName:  "monastery ",
					DOB:       "2025-11-24",
					Email:     "Codex.Monastery@email.com",
					Phone:     "+44 00000 00000",
					Address: Address{
						Line1:    "123 High Street ",
						Line2:    " ",
						City:     "London ",
						Postcode: "SW1A 1aa",
					},
					OptionalAddress: &Address{
						Line1:    "123 High Street ",
						Line2:    " ",
						City:     "London ",
						Postcode: "SW1A 1aa",
					},
					Tags: []string{"tag 1", "  tag 2 "},
				},
			},
			want: &User{
				FirstName: ptr("Codex"),
				LastName:  "Monastery",
				DOB:       "2025-11-24",
				Email:     "codex.monastery@email.com",
				Phone:     "+440000000000",
				Address: Address{
					Line1:    "123 High Street",
					Line2:    "",
					City:     "London",
					Postcode: "SW1A1AA",
				},
				OptionalAddress: &Address{
					Line1:    "123 High Street",
					Line2:    "",
					City:     "London",
					Postcode: "SW1A1AA",
				},
				Tags: []string{"Tag 1", "Tag 2"},
			},
			wantErr: false,
		},
		{
			name: "slices or array of string",
			args: args{
				v: &struct {
					Tags []string `sanitize:"dive,trim_space,capitalize"`
				}{
					Tags: []string{"tag 1", "  tag 2 "},
				},
			},
			want: &struct {
				Tags []string `sanitize:"dive,trim_space,capitalize"`
			}{Tags: []string{"Tag 1", "Tag 2"}},
			wantErr: false,
		},
		{
			name: "slices or array of structs",
			args: args{
				v: &struct {
					Addresses []Address `sanitize:"dive"`
				}{
					Addresses: []Address{
						{
							Line1:    "123 High Street ",
							Line2:    " ",
							City:     "London ",
							Postcode: "SW1A 1aa",
						},
						{
							Line1:    "124 High Street ",
							Line2:    " ",
							City:     "London ",
							Postcode: "SW1A 1aa",
						},
					},
				},
			},
			want: &struct {
				Addresses []Address `sanitize:"dive"`
			}{Addresses: []Address{
				{
					Line1:    "123 High Street",
					Line2:    "",
					City:     "London",
					Postcode: "SW1A1AA",
				},
				{
					Line1:    "124 High Street",
					Line2:    "",
					City:     "London",
					Postcode: "SW1A1AA",
				},
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Apply(tt.args.v)

			if tt.wantErr {
				assert.Error(t, err)

				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want, tt.args.v)
		})
	}
}

func ptr[T any](v T) *T {
	return &v
}
