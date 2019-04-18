package mdata_payload

import (
	"github.com/hyperledger/sawtooth-sdk-go/processor"
	"reflect"
	"testing"
)

var testPayloads = []struct {
	in         []byte
	outPayload *MdPayload
	outError   error
}{
	{nil, nil, &processor.InvalidTransactionError{Msg: "Must contain payload"}},
	{[]byte("create"), nil, &processor.InvalidTransactionError{Msg: "Payload is malformed"}}, //len<2
	{[]byte("create,00012345600012,000000001400245446"), &MdPayload{Action: "create", Gtin: "00012345600012", Mtrl: "000000001400245446"}, nil},
	//{[]byte("create,00012345600012"), nil, &processor.InvalidTransactionError{Msg: "Mtrl is required for create and update"}},
	//{[]byte("update,00012345600012,000000001400245446"), MdPayload{Action: "update", Gtin: "00012345600012", Mtrl: "000000001400245446"}, nil},
	//{[]byte("update,00012345600012"), nil, &processor.InvalidTransactionError{Msg: "Mtrl is required for create and update"}},
}

func compareExpectedActualError(expectedErr error, actualError error) bool {
	var areEqual bool
	if expectedErr != nil {
		areEqual = expectedErr.Error() == actualError.Error()
	} else {
		areEqual = reflect.TypeOf(expectedErr) == reflect.TypeOf(actualError)
	}
	return areEqual
}

func compareStructs(expected, actual *MdPayload) bool {
	expected_fields := reflect.TypeOf(expected)
	expected_values := reflect.ValueOf(expected)
	num_expected_fields := expected_fields.NumField()

	actual_fields := reflect.TypeOf(actual)
	actual_values := reflect.ValueOf(actual)
	num_actual_fields := actual_fields.NumField()

	if num_expected_fields != num_actual_fields {
		return false
	}

	var areEqual bool
	for i := 0; i < num; i++ {
		expected_field := expected_fields.Field(i)
		expected_value := expected_values.Field(i)
		actual_field := actual_fields.Field(i)
		actual_value := actual_values.Field(i)

		if expected_field != actual_field || expected_value != actual_value {
			areEqual = false
		}
		if areEqual == false {
			return areEqual
		}
	}
	areEqual = true
	return areEqual
}

func compareExpectedActualPayload(expectedPayload *MdPayload, actualPayload *MdPayload) bool {
	var areEqual bool
	if expectedPayload != nil {
		areEqual = compareStructs(expectedPayload, actualPayload)
	} else {
		areEqual = reflect.TypeOf(expectedPayload) == reflect.TypeOf(actualPayload)
	}
	return areEqual
}

func TestFromBytes(t *testing.T) {
	for _, tt := range testPayloads {
		payload, err := FromBytes(tt.in)
		if compareExpectedActualPayload(tt.outPayload, payload) != true || compareExpectedActualError(tt.outError, err) != true {
			t.Errorf("FromBytes(%v) => GOT %v, %v, WANT %v, %v", tt.in, payload, err, tt.outPayload, tt.outError)
		}
	}
}
