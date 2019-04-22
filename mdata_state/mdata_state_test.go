package mdata_state

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testGtin string = "01234567891234"
var testMtrl string = "12345-67890"
var testState string = "ACTIVE"
var testGtinAddress string = makeAddress(testGtin)
var testProduct Product = Product{
	Gtin:  testGtin,
	Mtrl:  testMtrl,
	State: testState,
}
var testProductSlice []*Product = make([]*Product,1)
testProductSlice[0] = testProduct

func TestGetProduct(t *testing.T) {
	sampleError := errors.New("sample")

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
			returnAddress := make(map[string][]byte)
			returnAddress[testGtinAddress] = serialize(testProductSlice)
			testContext.On("GetState", []string{testGtinAddress}).Return(
				returnAddress,
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

		product, err := testState.GetProduct(testGtin)
		assert.Equal(t, test.outProduct, product)
		assert.Equal(t, test.err, err)

		testContext.AssertExpectations(t)

	}
}

func TestDeleteProducts(t *testing.T) {

}

func TestLoadProducts(t *testing.T) {

}
