package mdata_state

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testGtin string = "01234567891234"
var testGtinAddress string = makeAddress(testGtin)

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
			outProduct: &Product{Gtin: testGtin, Mtrl: "", State: ""},
			err:        nil,
		},
	}

	for name, test := range tests {
		t.Logf("Running test case: %s", name)

		testContext := &mockContext{}

		if name == "existingProduct" {
			returnAddress := make(map[string][]byte)
			returnAddress[testGtinAddress] = []byte(testGtin)
			testContext.On("GetState", []string{testGtin}).Return(
				returnAddress,
				nil,
			)
		}
		if name == "emptyProduct" {
			testContext.On("GetState", []string{testGtin}).Return(
				nil,
				nil,
			)
		}
		if name == "error" {
			testContext.On("GetState", []string{testGtin}).Return(
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
