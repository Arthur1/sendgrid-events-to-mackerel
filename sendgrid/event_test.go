package sendgrid

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// cf.) https://docs.sendgrid.com/for-developers/tracking-events/event
const deliveryProcessedEvents = `
[
  {
      "email":"example@test.com",
      "timestamp":1513299569,
      "pool": {
            "name": "new_MY_test",
            "id": 210
        },
      "smtp-id":"<14c5d75ce93.dfd.64b469@ismtpd-555>",
      "event":"processed",
      "category":["cat facts"],
      "sg_event_id":"rbtnWrG1DVDGGGFHFyun0A==",
      "sg_message_id":"14c5d75ce93.dfd.64b469.filter0001.16648.5515E0B88.000000000000000000000"
  }
]
`

const deliveryDroppedEvents = `
[
  {
      "email":"example@test.com",
      "timestamp":1513299569,
      "smtp-id":"<14c5d75ce93.dfd.64b469@ismtpd-555>",
      "event":"dropped",
      "category":["cat facts"],
      "sg_event_id":"zmzJhfJgAfUSOW80yEbPyw==",
      "sg_message_id":"14c5d75ce93.dfd.64b469.filter0001.16648.5515E0B88.0",
      "reason":"Bounced Address",
      "status":"5.0.0"
  }
]
`

const deliveryDeliveredEvents = `
[
   {
      "email":"example@test.com",
      "timestamp":1513299569,
      "smtp-id":"<14c5d75ce93.dfd.64b469@ismtpd-555>",
      "event":"delivered",
      "category":["cat facts"],
      "sg_event_id":"rWVYmVk90MjZJ9iohOBa3w==",
      "sg_message_id":"14c5d75ce93.dfd.64b469.filter0001.16648.5515E0B88.0",
      "response":"250 OK"
   }
]
`

const deliveryDeferredEvents = `
[
   {
      "email":"example@test.com",
      "timestamp":1513299569,
      "smtp-id":"<14c5d75ce93.dfd.64b469@ismtpd-555>",
      "event":"deferred",
      "category":["cat facts"],
      "sg_event_id":"t7LEShmowp86DTdUW8M-GQ==",
      "sg_message_id":"14c5d75ce93.dfd.64b469.filter0001.16648.5515E0B88.0",
      "response":"400 try again later",
      "attempt":"5"
   }
]
`

const deliveryBounceEvents = `
[
   {
      "email":"example@test.com",
      "timestamp":1513299569,
      "smtp-id":"<14c5d75ce93.dfd.64b469@ismtpd-555>",
      "bounce_classification":"invalid",
      "event":"bounce",
      "category":["cat facts"],
      "sg_event_id":"6g4ZI7SA-xmRDv57GoPIPw==",
      "sg_message_id":"14c5d75ce93.dfd.64b469.filter0001.16648.5515E0B88.0",
      "reason":"500 unknown recipient",
      "status":"5.0.0",
      "type":"bounce"
   }
]
`

func TestParseEventsJSON(t *testing.T) {
	cases := map[string]struct {
		in      string
		want    []Event
		wantErr bool
	}{
		"DeliveryProcessedEvents": {
			in: deliveryProcessedEvents,
			want: []Event{
				DeliveryProcessedEvent{
					EventType:   "processed",
					Email:       "example@test.com",
					Timestamp:   1513299569,
					Pool:        Pool{Name: "new_MY_test", ID: 210},
					SMTPID:      "<14c5d75ce93.dfd.64b469@ismtpd-555>",
					Category:    []string{"cat facts"},
					SGEventID:   "rbtnWrG1DVDGGGFHFyun0A==",
					SGMessageID: "14c5d75ce93.dfd.64b469.filter0001.16648.5515E0B88.000000000000000000000",
				},
			},
		},
		"DeliveryDroppedEvents": {
			in: deliveryDroppedEvents,
			want: []Event{
				DeliveryDroppedEvent{
					EventType:   "dropped",
					Email:       "example@test.com",
					Timestamp:   1513299569,
					SMTPID:      "<14c5d75ce93.dfd.64b469@ismtpd-555>",
					Category:    []string{"cat facts"},
					SGEventID:   "zmzJhfJgAfUSOW80yEbPyw==",
					SGMessageID: "14c5d75ce93.dfd.64b469.filter0001.16648.5515E0B88.0",
					Reason:      "Bounced Address",
					Status:      "5.0.0",
				},
			},
		},
		"DeliveryDeliveredEvents": {
			in: deliveryDeliveredEvents,
			want: []Event{
				DeliveryDeliveredEvent{
					EventType:   "delivered",
					Email:       "example@test.com",
					Timestamp:   1513299569,
					SMTPID:      "<14c5d75ce93.dfd.64b469@ismtpd-555>",
					Category:    []string{"cat facts"},
					SGEventID:   "rWVYmVk90MjZJ9iohOBa3w==",
					SGMessageID: "14c5d75ce93.dfd.64b469.filter0001.16648.5515E0B88.0",
					Response:    "250 OK",
				},
			},
		},
		"DeliveryDeferredEvents": {
			in: deliveryDeferredEvents,
			want: []Event{
				DeliveryDeferredEvent{
					EventType:   "deferred",
					Email:       "example@test.com",
					Timestamp:   1513299569,
					SMTPID:      "<14c5d75ce93.dfd.64b469@ismtpd-555>",
					Category:    []string{"cat facts"},
					SGEventID:   "t7LEShmowp86DTdUW8M-GQ==",
					SGMessageID: "14c5d75ce93.dfd.64b469.filter0001.16648.5515E0B88.0",
					Response:    "400 try again later",
					Attempt:     "5",
				},
			},
		},
		"DeliveryBounceEvents": {
			in: deliveryBounceEvents,
			want: []Event{
				DeliveryBounceEvent{
					EventType:            "bounce",
					Email:                "example@test.com",
					Timestamp:            1513299569,
					SMTPID:               "<14c5d75ce93.dfd.64b469@ismtpd-555>",
					BounceClassification: "invalid",
					Category:             []string{"cat facts"},
					SGEventID:            "6g4ZI7SA-xmRDv57GoPIPw==",
					SGMessageID:          "14c5d75ce93.dfd.64b469.filter0001.16648.5515E0B88.0",
					Reason:               "500 unknown recipient",
					Status:               "5.0.0",
					BounceType:           "bounce",
				},
			},
		},
	}
	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			got, err := ParseEventsJSON([]byte(tt.in))
			if err != nil {
				t.Errorf("want no error, got error: %s", err)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestCalcDeliveryEventsCountByType(t *testing.T) {
	cases := map[string]struct {
		in   []Event
		want DeliveryEventsCountByType
	}{
		"eventsが空なら全て0": {
			in:   []Event{},
			want: DeliveryEventsCountByType{},
		},
		"eventsを種別ごと数える": {
			in: []Event{
				DeliveryProcessedEvent{},
				DeliveryDroppedEvent{},
				DeliveryDroppedEvent{},
				DeliveryDeliveredEvent{},
				DeliveryDeliveredEvent{},
				DeliveryDeliveredEvent{},
				DeliveryDeferredEvent{},
				DeliveryDeferredEvent{},
				DeliveryDeferredEvent{},
				DeliveryDeferredEvent{},
				DeliveryBounceEvent{},
				DeliveryBounceEvent{},
				DeliveryBounceEvent{},
				DeliveryBounceEvent{},
				DeliveryBounceEvent{},
			},
			want: DeliveryEventsCountByType{
				DeliveryProcessed: 1,
				DeliveryDropped:   2,
				DeliveryDelivered: 3,
				DeliveryDeferred:  4,
				DeliveryBounce:    5,
			},
		},
	}
	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			got := CalcDeliveryEventsCountByType(tt.in)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Error(diff)
			}
		})
	}
}
