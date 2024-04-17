package db

import (
	"context"
	"testing"
	"tsn-ecommerce/utils"

	"github.com/stretchr/testify/require"
)

func TestQueries_GetProductWithInfo(t *testing.T) {
	// add a product
	product := addRandomProductTx(t)
	arg := GetProductWithInfoParams{
		Pk:       product.Product.Pk,
		Language: product.ProductTrans.Language,
	}

	// get product with info and check for error
	product2, err := testQueries.GetProductWithInfo(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product2)

	// check if the product is the same
	require.Equal(t, product.Product.Pk, product2.Pk)
	require.Equal(t, product.ProductTrans.Name, product2.Name)
	require.Equal(t, product.ProductTrans.Description, product2.Description)
	require.Equal(t, product.ProductTrans.CreatedAt, product2.CreatedAt)
	require.Equal(t, product.ProductTrans.UpdatedAt, product2.UpdatedAt)
}

func TestQueries_ListProductWithInfo(t *testing.T) {
	num := 10
	langCode := LanguageCode(utils.RandomLanguage())

	// create dummy products
	for i := 0; i < num; i++ {
		addProductTxWithLangCode(t, langCode)
	}

}
