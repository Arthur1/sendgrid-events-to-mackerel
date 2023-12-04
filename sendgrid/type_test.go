package sendgrid

import (
	"encoding/json"
	"testing"
)

func TestEventTypeMarshalJSON(t *testing.T) {
	cases := []struct {
		in   EventType
		want string
	}{
		{in: DeliveryProcessedEventType, want: `"processed"`},
		{in: DeliveryDroppedEventType, want: `"dropped"`},
		{in: DeliveryDeliveredEventType, want: `"delivered"`},
		{in: DeliveryDeferredEventType, want: `"deferred"`},
		{in: DeliveryBounceEventType, want: `"bounce"`},
		{in: EngagementOpenEventType, want: `"open"`},
		{in: EngagementClickEventType, want: `"click"`},
		{in: EngagementSpamReportEventType, want: `"spamreport"`},
		{in: EngagementUnsubscribeEventType, want: `"unsubscribe"`},
		{in: EngagementGroupUnsubscribeEventType, want: `"group_unsubscribe"`},
		{in: EngagementGroupResubscribeEventType, want: `"group_resubscribe"`},
	}
	for _, tt := range cases {
		t.Run(tt.want, func(t *testing.T) {
			b, err := json.Marshal(tt.in)
			got := string(b)
			if err != nil {
				t.Errorf("want no error, got error: %s", err)
			}
			if tt.want != got {
				t.Errorf("want=%s, got=%s", tt.want, got)
			}
		})
	}
}

func TestEventTypeUnmarshalJSON(t *testing.T) {
	cases := []struct {
		in      string
		want    EventType
		wantErr bool
	}{
		{in: `"processed"`, want: DeliveryProcessedEventType},
		{in: `"dropped"`, want: DeliveryDroppedEventType},
		{in: `"delivered"`, want: DeliveryDeliveredEventType},
		{in: `"deferred"`, want: DeliveryDeferredEventType},
		{in: `"bounce"`, want: DeliveryBounceEventType},
		{in: `"open"`, want: EngagementOpenEventType},
		{in: `"click"`, want: EngagementClickEventType},
		{in: `"spamreport"`, want: EngagementSpamReportEventType},
		{in: `"unsubscribe"`, want: EngagementUnsubscribeEventType},
		{in: `"group_unsubscribe"`, want: EngagementGroupUnsubscribeEventType},
		{in: `"group_resubscribe"`, want: EngagementGroupResubscribeEventType},
		{in: `"hoge"`, wantErr: true},
	}
	for _, tt := range cases {
		t.Run(tt.in, func(t *testing.T) {
			var got EventType
			b := []byte(tt.in)
			err := json.Unmarshal(b, &got)
			if tt.wantErr && err == nil {
				t.Errorf("want error, got nil")
			} else if !tt.wantErr && err != nil {
				t.Errorf("want no error, got error: %s", err)
			}
			if tt.want != "" && tt.want != got {
				t.Errorf("want=%s, got=%s", tt.want, got)
			}
		})
	}
}
