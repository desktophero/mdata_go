package mdata_state

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testGtin string = "01234567891234"
var testMtrl string = "12345-67890"
var testSetNewMtrl string = "67890-12345"
var testState string = "ACTIVE"
var testGtinAddress string = makeAddress(testGtin)
var testProduct Product = Product{
	Gtin:  testGtin,
	Mtrl:  testMtrl,
	State: testState,
}
var testSetNewProduct Product = Product{
	Gtin:  testGtin,
	Mtrl:  testSetNewMtrl,
	State: testState,
}
var sampleError = errors.New("sample")

func TestGetProduct(t *testing.T) {

	tests := map[string]struct {
		gtin       string
		outProduct *Product
		err        error
	}{
		"error": {
			gtin:       testGtin,
			outProduct: nil,
			err:        sampleError,
		},
		"emptyProduct": {
			gtin:       testGtin,
			outProduct: nil,
			err:        nil,
		},
		"existingProduct": {
			gtin:       testGtin,
			outProduct: &testProduct,
			err:        nil,
		},
	}

	for name, test := range tests {
		t.Logf("Running test case: %s", name)

		testContext := &mockContext{}

		if name == "existingProduct" {
			returnState := make(map[string][]byte)
			testProductSlice := make([]*Product, 1)
			testProductSlice[0] = &testProduct

			returnState[testGtinAddress] = serialize(testProductSlice)
			testContext.On("GetState", []string{testGtinAddress}).Return(
				returnState,
				nil,
			)
		}
		if name == "emptyProduct" {
			testContext.On("GetState", []string{testGtinAddress}).Return(
				nil,
				nil,
			)
		}
		if name == "error" {
			testContext.On("GetState", []string{testGtinAddress}).Return(
				nil,
				sampleError,
			)
		}

		testState := &MdState{
			context:      testContext,
			addressCache: make(map[string][]byte),
		}

		product, err := testState.GetProduct(test.gtin)
		assert.Equal(t, test.outProduct, product)
		assert.Equal(t, test.err, err)

		testContext.AssertExpectations(t)

	}
}

func TestSetProduct(t *testing.T) {

	tests := map[string]struct {
		gtin      string
		inProduct *Product
		err       error
	}{
		"newProduct": {
			gtin:      testGtin,
			inProduct: &testProduct,
			err:       nil,
		},
		"updateProductState": {
			gtin:      testGtin,
			inProduct: &testSetNewProduct,
			err:       nil,
		},
	}

	for name, test := range tests {
		t.Logf("Running test case: %s", name)

		testContext := &mockContext{}
		testProductSlice := make([]*Product, 1)
		testProductSlice[0] = &testProduct

		if name == "newProduct" {
			returnState := make(map[string][]byte)
			testContext.On("GetState", []string{testGtinAddress}).Return(
				returnState,
				nil,
			)

			data := serialize(testProductSlice)
			testContext.On("SetState", map[string][]byte{testGtinAddress: data}).Return(
				[]string{testGtinAddress},
				nil,
			)
		}

		if name == "updateProductState" {
			returnState := make(map[string][]byte)
			returnState[testGtinAddress] = serialize(testProductSlice)
			testContext.On("GetState", []string{testGtinAddress}).Return(
				returnState,
				nil,
			)

			data := serialize([]*Product{&testSetNewProduct})
			testContext.On("SetState", map[string][]byte{testGtinAddress: data}).Return(
				[]string{testGtinAddress},
				nil,
			)

		}

		testState := &MdState{
			context:      testContext,
			addressCache: make(map[string][]byte),
		}

		err := testState.SetProduct(test.gtin, test.inProduct)
		assert.Equal(t, test.err, err)
		testContext.AssertExpectations(t)

	}
}

func TestDeleteProduct(t *testing.T) {

	tests := map[string]struct {
		gtin string
		err  error
	}{
		"productDoesNotExist": { //Delete state at testGtin address
			gtin: testGtin,
			err:  nil,
		},
		// "storeProductsWithoutDeleted": { //If other products exist, just storeProducts at the state address of the deleted Gtin
		// 	gtin: testGtin,
		// 	err:  nil,
		// },
	}

	for name, test := range tests {
		t.Logf("Running test case: %s", name)

		testContext := &mockContext{}

		if name == "productDoesNotExist" {
			returnState := make(map[string][]byte) // Return empty map
			testContext.On("GetState", []string{testGtinAddress}).Return(
				returnState,
				nil,
			)

			data := serialize(testProductSlice)
			testContext.On("DeleteState", []string{testGtinAddress}).Return(
				nil,
				nil,
			).Once()
		}

		// if name == "storeProductsWithoutDeleted" {
		// 	returnState := make(map[string][]byte)
		// 	returnState[testGtinAddress] = serialize(testProductSlice)
		// 	testContext.On("GetState", []string{testGtinAddress}).Return(
		// 		returnState,
		// 		nil,
		// 	)

		// 	data := serialize([]*Product{&testSetNewProduct})
		// 	testContext.On("DeleteState", map[string][]byte{testGtinAddress: data}).Return(
		// 		[]string{testGtinAddress},
		// 		nil,
		// 	).Once

		// }

		testState := &MdState{
			context:      testContext,
			addressCache: make(map[string][]byte),
		}

		err := testState.DeleteProduct(test.gtin)
		assert.Equal(t, test.err, err)
		testContext.AssertExpectations(t)

	}

}
