package mdata_payload

import (
	"github.com/hyperledger/sawtooth-sdk-go/processor"
	"testing"
)

var testPayloads = []struct {
	in         []byte
	outPayload *MdPayload
	outError   error
}{
	{nil, nil, &processor.InvalidTransactionError{Msg: "Must contain payload"}},
	//{[]byte("create"), nil, &processor.InvalidTransactionError{Msg: "Payload is malformed"}}, //len<2
	//{[]byte("create,00012345600012,000000001400245446"), &MdPayload{Action: "create", Gtin: "00012345600012", Mtrl: "000000001400245446"}, nil},
	//{[]byte("create,00012345600012"), nil, &processor.InvalidTransactionError{Msg: "Mtrl is required for create and update"}},
	//{[]byte("update,00012345600012,000000001400245446"), &MdPayload{Action: "update", Gtin: "00012345600012", Mtrl: "000000001400245446"}, nil},
	//{[]byte("update,00012345600012"), nil, &processor.InvalidTransactionError{Msg: "Mtrl is required for create and update"}},
}

func TestFromBytes(t *testing.T) {
	for _, tt := range testPayloads {
		payload, err := FromBytes(tt.in)
		if &payload != tt.outPayload || &err != &tt.outError {
			t.Errorf("FromBytes(%v) => %v, %v, want %v, %v", tt.in, payload, err, tt.outPayload, tt.outError)
		}
	}
}
