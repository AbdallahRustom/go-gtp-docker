// Copyright 2019-2023 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/wmnsk/go-gtp/gtpv1"
	"github.com/wmnsk/go-gtp/gtpv1/ie"
	"github.com/wmnsk/go-gtp/gtpv1/message"
)

func TestIEs(t *testing.T) {
	cases := []struct {
		description string
		structured  *ie.IE
		serialized  []byte
	}{
		{
			"IMSI",
			ie.NewIMSI("123451234567890"),
			[]byte{0x02, 0x21, 0x43, 0x15, 0x32, 0x54, 0x76, 0x98, 0xf0},
		}, {
			"PacketTMSI",
			ie.NewPacketTMSI(0xbeebee),
			[]byte{0x05, 0x00, 0xbe, 0xeb, 0xee},
		}, {
			"AuthenticationTriplet",
			ie.NewAuthenticationTriplet(
				[]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
				[]byte{0xde, 0xad, 0xbe, 0xef},
				[]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77},
			),
			[]byte{
				0x09,
				0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff,
				0xde, 0xad, 0xbe, 0xef,
				0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77,
			},
		}, {
			"MAPCause",
			ie.NewMAPCause(gtpv1.MAPCauseSystemFailure),
			[]byte{0x0b, 0x22},
		}, {
			"PTMSISignature",
			ie.NewPTMSISignature(0xbeebee),
			[]byte{0x0c, 0xbe, 0xeb, 0xee},
		}, {
			"MSValidated",
			ie.NewMSValidated(true),
			[]byte{0x0d, 0xff},
		}, {
			"Recovery",
			ie.NewRecovery(1),
			[]byte{0x0e, 0x01},
		}, {
			"SelectionMode",
			ie.NewSelectionMode(gtpv1.SelectionModeMSorNetworkProvidedAPNSubscribedVerified),
			[]byte{0x0f, 0xf0},
		}, {
			"TEIDDataI",
			ie.NewTEIDDataI(0xdeadbeef),
			[]byte{0x10, 0xde, 0xad, 0xbe, 0xef},
		}, {
			"TEIDCPlane",
			ie.NewTEIDCPlane(0xdeadbeef),
			[]byte{0x11, 0xde, 0xad, 0xbe, 0xef},
		}, {
			"TEIDDataII",
			ie.NewTEIDDataII(0xdeadbeef),
			[]byte{0x12, 0xde, 0xad, 0xbe, 0xef},
		}, {
			"TeardownInd",
			ie.NewTeardownInd(true),
			[]byte{0x13, 0xff},
		}, {
			"NSAPI",
			ie.NewNSAPI(0x05),
			[]byte{0x14, 0x05},
		}, {
			"RANAPCause",
			ie.NewRANAPCause(gtpv1.MAPCauseUnknownSubscriber),
			[]byte{0x15, 0x01},
		}, {
			"EndUserAddress/v4",
			ie.NewEndUserAddress("1.1.1.1"),
			[]byte{0x80, 0x00, 0x06, 0xf1, 0x21, 0x01, 0x01, 0x01, 0x01},
		}, {
			"EndUserAddress/v6",
			ie.NewEndUserAddress("2001::1"),
			[]byte{
				0x80, 0x00, 0x12, 0x00,
				0x57, 0x20, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
			},
		}, {
			"AccessPointName",
			ie.NewAccessPointName("some.apn.example"),
			[]byte{
				0x83, 0x00, 0x11,
				0x04, 0x73, 0x6f, 0x6d, 0x65, 0x03, 0x61, 0x70, 0x6e, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
			},
		}, {
			"GSNAddressV4",
			ie.NewGSNAddress("1.1.1.1"),
			[]byte{0x85, 0x00, 0x04, 0x01, 0x01, 0x01, 0x01},
		}, {
			"GSNAddressV6",
			ie.NewGSNAddress("2001::1"),
			[]byte{
				0x85, 0x00, 0x10,
				0x20, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
			},
		}, {
			"MSISDN",
			ie.NewMSISDN("818012345678"),
			[]byte{0x86, 0x00, 0x07, 0x91, 0x18, 0x08, 0x21, 0x43, 0x65, 0x87},
		}, {
			"AuthenticationQuintuplet",
			ie.NewAuthenticationQuintuplet(
				[]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
				[]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
				[]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
				[]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
				[]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
			),
			[]byte{
				0x88, 0x00, 0x52,
				0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff,
				0x10,
				0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff,
				0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff,
				0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff,
				0x10,
				0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff,
			},
		}, {
			"ExtensionHeaderTypeList",
			ie.NewExtensionHeaderTypeList(
				message.ExtHeaderTypePDUSessionContainer,
				message.ExtHeaderTypeUDPPort,
			),
			[]byte{0x8d, 0x02, 0x85, 0x40},
		}, {
			"CommonFlags",
			ie.NewCommonFlags(0, 1, 0, 0, 0, 0, 0, 0),
			[]byte{0x94, 0x00, 0x01, 0x40},
		}, {
			"APNRestriction",
			ie.NewAPNRestriction(gtpv1.APNRestrictionPrivate1),
			[]byte{0x95, 0x00, 0x01, 0x03},
		}, {
			"RATType",
			ie.NewRATType(gtpv1.RatTypeEUTRAN),
			[]byte{0x97, 0x00, 0x01, 0x06},
		}, {
			"UserLocationInformationWithCGI",
			ie.NewUserLocationInformationWithCGI("123", "45", 0xff, 0),
			[]byte{0x98, 0x00, 0x08, 0x00, 0x21, 0xf3, 0x54, 0x00, 0xff, 0x00, 0x00},
		}, {
			"UserLocationInformationWithSAI",
			ie.NewUserLocationInformationWithSAI("123", "45", 0xff, 0),
			[]byte{0x98, 0x00, 0x08, 0x01, 0x21, 0xf3, 0x54, 0x00, 0xff, 0x00, 0x00},
		}, {
			"UserLocationInformationWithRAI",
			ie.NewUserLocationInformationWithRAI("123", "45", 0xff, 0),
			[]byte{0x98, 0x00, 0x07, 0x02, 0x21, 0xf3, 0x54, 0x00, 0xff, 0x00},
		}, {
			"MSTimeZone",
			ie.NewMSTimeZone(9*time.Hour, 0), // XXX - should be updated with more realistic value
			[]byte{0x99, 0x00, 0x02, 0x63, 0x00},
		}, {
			"MSTimeZone",
			ie.NewMSTimeZone(2*time.Hour, 0),
			[]byte{0x99, 0x00, 0x02, 0x80, 0x00},
		}, {
			"IMEISV",
			ie.NewIMEISV("123450123456789"),
			[]byte{0x9a, 0x00, 0x08, 0x21, 0x43, 0x05, 0x21, 0x43, 0x65, 0x87, 0xf9},
		}, {
			"ULITimestamp",
			ie.NewULITimestamp(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			[]byte{0xd6, 0x00, 0x04, 0xdf, 0xd5, 0x2c, 0x00},
		}, {
			"ChargingID",
			ie.NewChargingID(0xffffffff),
			[]byte{0x7f, 0xff, 0xff, 0xff, 0xff},
		},
		{
			"PrivateExtension",
			ie.NewPrivateExtension(0x0080, []byte{0xde, 0xad, 0xbe, 0xef}),
			[]byte{
				// Type, Length
				0xff, 0x00, 0x06,
				// Value
				0x00, 0x80, 0xde, 0xad, 0xbe, 0xef,
			},
		},
	}

	for _, c := range cases {
		t.Run("Marshal/"+c.description, func(t *testing.T) {
			got, err := c.structured.Marshal()
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(got, c.serialized); diff != "" {
				t.Error(diff)
			}
		})

		t.Run("Parse/"+c.description, func(t *testing.T) {
			got, err := ie.Parse(c.serialized)
			if err != nil {
				t.Fatal(err)
			}

			opt := cmp.AllowUnexported(*got, *c.structured)
			if diff := cmp.Diff(got, c.structured, opt); diff != "" {
				t.Error(diff)
			}
		})
	}
}
