package sendgrid

import (
	"encoding/json"
	"fmt"
)

func ParseEventsJSON(bytes []byte) ([]Event, error) {
	var items []map[string]any
	if err := json.Unmarshal(bytes, &items); err != nil {
		return nil, err
	}
	events := make([]Event, 0, len(items))
	for _, item := range items {
		eventType, ok := item["event"]
		if !ok {
			return nil, fmt.Errorf("event key is required")
		}
		eventTypeString, ok := eventType.(string)
		if !ok {
			return nil, fmt.Errorf("event should be string")
		}
		b, err := json.Marshal(item)
		if err != nil {
			return nil, err
		}
		switch EventType(eventTypeString) {
		case DeliveryProcessedEventType:
			var event DeliveryProcessedEvent
			if err := json.Unmarshal(b, &event); err != nil {
				return nil, err
			}
			events = append(events, event)
		case DeliveryDroppedEventType:
			var event DeliveryDroppedEvent
			if err := json.Unmarshal(b, &event); err != nil {
				return nil, err
			}
			events = append(events, event)
		case DeliveryDeliveredEventType:
			var event DeliveryDeliveredEvent
			if err := json.Unmarshal(b, &event); err != nil {
				return nil, err
			}
			events = append(events, event)
		case DeliveryDeferredEventType:
			var event DeliveryDeferredEvent
			if err := json.Unmarshal(b, &event); err != nil {
				return nil, err
			}
			events = append(events, event)
		case DeliveryBounceEventType:
			var event DeliveryBounceEvent
			if err := json.Unmarshal(b, &event); err != nil {
				return nil, err
			}
			events = append(events, event)
		case EngagementOpenEventType, EngagementClickEventType, EngagementSpamReportEventType, EngagementUnsubscribeEventType, EngagementGroupUnsubscribeEventType, EngagementGroupResubscribeEventType:
			// TODO: support engagement events
			continue
		default:
			return nil, fmt.Errorf("unexpected event type: %s", eventTypeString)
		}
	}
	return events, nil
}

type DeliveryEventsCountByType struct {
	DeliveryProcessed int64
	DeliveryDropped   int64
	DeliveryDelivered int64
	DeliveryDeferred  int64
	DeliveryBounce    int64
}

func CalcDeliveryEventsCountByType(events []Event) DeliveryEventsCountByType {
	result := DeliveryEventsCountByType{}
	for _, event := range events {
		switch event.GetEventType() {
		case DeliveryProcessedEventType:
			result.DeliveryProcessed++
		case DeliveryDroppedEventType:
			result.DeliveryDropped++
		case DeliveryDeliveredEventType:
			result.DeliveryDelivered++
		case DeliveryDeferredEventType:
			result.DeliveryDeferred++
		case DeliveryBounceEventType:
			result.DeliveryBounce++
		}
	}
	return result
}
