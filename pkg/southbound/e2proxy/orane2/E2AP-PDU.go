// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package orane2

// #cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
// #cgo LDFLAGS: -lm
// #include <stdio.h>
// #include <stdlib.h>
// #include <assert.h>
// #include "E2AP-PDU.h"
//#include "GNB-ID-Choice.h"
//#include "GlobalE2node-gNB-ID.h"
//#include "GlobalE2node-ID.h"
//#include "ProtocolIE-Field.h"
//#include "InitiatingMessage.h"
//#include "SuccessfulOutcome.h"
import "C"
import (
	"encoding/binary"
	"fmt"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/e2ctypes"
	"unsafe"
)

// XerDecodeE2apPdu - the main entry to decode E2 message in XER format
func XerDecodeE2apPdu(bytes []byte) (*e2ctypes.E2AP_PDUT, error) {

	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_E2AP_PDU)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2apPdu((*C.E2AP_PDU_t)(unsafePtr))
}

// PerDecodeE2apPdu - the main entry to decode E2 message in PER format
func PerDecodeE2apPdu(bytes []byte) (*e2ctypes.E2AP_PDUT, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_E2AP_PDU)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeE2apPdu((*C.E2AP_PDU_t)(unsafePtr))
}

// XerEncodeE2apPdu - the main function for testing from higher level api to C api
func XerEncodeE2apPdu(e2apPdu *e2ctypes.E2AP_PDUT) ([]byte, error) {
	cE2apPdu, err := newE2apPdu(e2apPdu)
	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_E2AP_PDU, unsafe.Pointer(cE2apPdu))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// PerEncodeE2apPdu - the main function for encoding from higher level api to C api
func PerEncodeE2apPdu(e2apPdu *e2ctypes.E2AP_PDUT) ([]byte, error) {
	cE2apPdu, err := newE2apPdu(e2apPdu)
	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_E2AP_PDU, unsafe.Pointer(cE2apPdu))
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// newE2apPdu - create a structure in C
func newE2apPdu(e2apPdu *e2ctypes.E2AP_PDUT) (*C.E2AP_PDU_t, error) {
	var present C.E2AP_PDU_PR
	choiceC := [8]byte{}
	switch choice := e2apPdu.Choice.(type) {
	case *e2ctypes.E2AP_PDUT_InitiatingMessage:
		present = C.E2AP_PDU_PR_initiatingMessage

		im, err := newInitiatingMessage(choice.InitiatingMessage)
		if err != nil {
			return nil, err
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(im))))
	case *e2ctypes.E2AP_PDUT_SuccessfulOutcome:
		present = C.E2AP_PDU_PR_successfulOutcome

		so, err := newSuccessfulOutcome(choice.SuccessfulOutcome)
		if err != nil {
			return nil, err
		}
		binary.LittleEndian.PutUint64(choiceC[0:], uint64(uintptr(unsafe.Pointer(so))))
	default:
		return nil, fmt.Errorf("newE2apPdu() %T not yet implemented", choice)
	}

	e2apPduC := C.E2AP_PDU_t{
		present: present,
		choice:  choiceC,
	}

	return &e2apPduC, nil
}

func decodeE2apPdu(e2apPduC *C.E2AP_PDU_t) (*e2ctypes.E2AP_PDUT, error) {
	e2apPdu := new(e2ctypes.E2AP_PDUT)
	switch e2apPduC.present {
	case C.E2AP_PDU_PR_initiatingMessage:
		// https://sunzenshen.github.io/tutorials/2015/05/09/cgotchas-intro.html
		initMsgC := *(**C.InitiatingMessage_t)(unsafe.Pointer(&e2apPduC.choice[0]))

		initMsg, err := decodeInitiatingMessage(initMsgC)
		if err != nil {
			return nil, err
		}
		e2apPdu.Choice = &e2ctypes.E2AP_PDUT_InitiatingMessage{
			InitiatingMessage: initMsg,
		}

	case C.E2AP_PDU_PR_successfulOutcome:
		// https://sunzenshen.github.io/tutorials/2015/05/09/cgotchas-intro.html
		initMsgC := *(**C.SuccessfulOutcome_t)(unsafe.Pointer(&e2apPduC.choice[0]))

		initMsg, err := decodeSuccessfulOutcome(initMsgC)
		if err != nil {
			return nil, err
		}
		e2apPdu.Choice = &e2ctypes.E2AP_PDUT_SuccessfulOutcome{
			SuccessfulOutcome: initMsg,
		}
	default:
		return nil, fmt.Errorf("PerDecodeE2apPdu decoding %v not yet implemented", e2apPduC.present)
	}

	return e2apPdu, nil
}