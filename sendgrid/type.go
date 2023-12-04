package sendgrid

import (
	"encoding/json"
	"fmt"
)

type EventType string

const (
	DeliveryProcessedEventType          EventType = "processed"
	DeliveryDroppedEventType            EventType = "dropped"
	DeliveryDeliveredEventType          EventType = "delivered"
	DeliveryDeferredEventType           EventType = "deferred"
	DeliveryBounceEventType             EventType = "bounce"
	EngagementOpenEventType             EventType = "open"
	EngagementClickEventType            EventType = "click"
	EngagementSpamReportEventType       EventType = "spamreport"
	EngagementUnsubscribeEventType      EventType = "unsubscribe"
	EngagementGroupUnsubscribeEventType EventType = "group_unsubscribe"
	EngagementGroupResubscribeEventType EventType = "group_resubscribe"
)

var (
	p                  = DeliveryProcessedEventType
	_ json.Marshaler   = p
	_ json.Unmarshaler = &p
)

func (t EventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(t))
}

func (t *EventType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("data should be a string, got %s", data)
	}
	switch s {
	case "processed":
		*t = DeliveryProcessedEventType
	case "dropped":
		*t = DeliveryDroppedEventType
	case "delivered":
		*t = DeliveryDeliveredEventType
	case "deferred":
		*t = DeliveryDeferredEventType
	case "bounce":
		*t = DeliveryBounceEventType
	case "open":
		*t = EngagementOpenEventType
	case "click":
		*t = EngagementClickEventType
	case "spamreport":
		*t = EngagementSpamReportEventType
	case "unsubscribe":
		*t = EngagementUnsubscribeEventType
	case "group_unsubscribe":
		*t = EngagementGroupUnsubscribeEventType
	case "group_resubscribe":
		*t = EngagementGroupResubscribeEventType
	default:
		return fmt.Errorf("invalid EventType %s", s)
	}
	return nil
}

type Pool struct {
	Name string `json:"name"`
	ID   int64  `json:"id"`
}

type Event interface {
	GetEventType() EventType
}

type DeliveryProcessedEvent struct {
	EventType   EventType `json:"event"`
	Email       string    `json:"email"`
	Timestamp   int64     `json:"timestamp"`
	Pool        Pool      `json:"pool"`
	SMTPID      string    `json:"smtp-id"`
	Category    []string  `json:"category"`
	SGEventID   string    `json:"sg_event_id"`
	SGMessageID string    `json:"sg_message_id"`
}

func (DeliveryProcessedEvent) GetEventType() EventType {
	return DeliveryProcessedEventType
}

type DeliveryDroppedEvent struct {
	EventType   EventType `json:"event"`
	Email       string    `json:"email"`
	Timestamp   int64     `json:"timestamp"`
	SMTPID      string    `json:"smtp-id"`
	Category    []string  `json:"category"`
	SGEventID   string    `json:"sg_event_id"`
	SGMessageID string    `json:"sg_message_id"`
	Reason      string    `json:"reason"`
	Status      string    `json:"status"`
}

func (DeliveryDroppedEvent) GetEventType() EventType {
	return DeliveryDroppedEventType
}

type DeliveryDeliveredEvent struct {
	EventType   EventType `json:"event"`
	Email       string    `json:"email"`
	Timestamp   int64     `json:"timestamp"`
	SMTPID      string    `json:"smtp-id"`
	Category    []string  `json:"category"`
	SGEventID   string    `json:"sg_event_id"`
	SGMessageID string    `json:"sg_message_id"`
	Response    string    `json:"response"`
}

func (DeliveryDeliveredEvent) GetEventType() EventType {
	return DeliveryDeliveredEventType
}

type DeliveryDeferredEvent struct {
	EventType   EventType `json:"event"`
	Email       string    `json:"email"`
	Timestamp   int64     `json:"timestamp"`
	SMTPID      string    `json:"smtp-id"`
	Category    []string  `json:"category"`
	SGEventID   string    `json:"sg_event_id"`
	SGMessageID string    `json:"sg_message_id"`
	Response    string    `json:"response"`
	Attempt     string    `json:"attempt"`
}

func (DeliveryDeferredEvent) GetEventType() EventType {
	return DeliveryDeferredEventType
}

type DeliveryBounceEvent struct {
	EventType            EventType `json:"event"`
	Email                string    `json:"email"`
	Timestamp            int64     `json:"timestamp"`
	SMTPID               string    `json:"smtp-id"`
	BounceClassification string    `json:"bounce_classification"`
	Category             []string  `json:"category"`
	SGEventID            string    `json:"sg_event_id"`
	SGMessageID          string    `json:"sg_message_id"`
	Response             string    `json:"response"`
	Reason               string    `json:"reason"`
	Status               string    `json:"status"`
	BounceType           string    `json:"type"`
}

func (DeliveryBounceEvent) GetEventType() EventType {
	return DeliveryBounceEventType
}

var (
	_ Event = DeliveryProcessedEvent{}
	_ Event = DeliveryDroppedEvent{}
	_ Event = DeliveryDeliveredEvent{}
	_ Event = DeliveryDeferredEvent{}
	_ Event = DeliveryBounceEvent{}
)
