package sales_usecase

import (
	"github.com/beto-ouverney/go-affiliates/backend/internal/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_removeDuplicate(t *testing.T) {
	assertions := assert.New(t)

	tests := []struct {
		describe    string
		argProduct  []entities.Product
		argProducer []entities.Producer
		want        []entities.Product
		want1       []entities.Producer
		msg         string
		msg1        string
	}{
		{
			describe: "Remove duplicate Product entities",
			argProduct: []entities.Product{
				{

					Name:       "Product 1",
					ProducerId: 1,
				},
				{

					Name:       "Product 2",
					ProducerId: 2,
				},
				{
					Name:       "Product 1",
					ProducerId: 1,
				},
			},
			argProducer: []entities.Producer{
				{
					Name: "Producer 1",
				},
				{
					Name: "Producer 2",
				},
				{
					Name: "Producer 1",
				},
			},
			want: []entities.Product{
				{
					Name:       "Product 1",
					ProducerId: 1,
				},
				{
					Name:       "Product 2",
					ProducerId: 2,
				},
			},
			want1: []entities.Producer{
				{
					Name: "Producer 1",
				},
				{
					Name: "Producer 2",
				},
			},
			msg:  "Should remove duplicate Product entities",
			msg1: "Should remove duplicate Producer entities",
		},
	}
	for _, tt := range tests {
		t.Run(tt.describe, func(t *testing.T) {
			got := removeDuplicate(tt.argProduct)
			assertions.Equal(tt.want, got, tt.msg)
			got1 := removeDuplicate(tt.argProducer)
			assertions.Equal(tt.want1, got1, tt.msg1)
		})
	}
}
