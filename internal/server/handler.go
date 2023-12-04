package server

import (
	"io"
	"log/slog"
	"net/http"
	"os"

	"github.com/Arthur1/sendgrid-events-to-mackerel/sendgrid"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func NewHTTPHandler() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	r.Post("/sendgrid/events", postSendgridEvents)
	return r
}

func postSendgridEvents(w http.ResponseWriter, r *http.Request) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// TODO: Sendgridからのリクエストか検証する
	// cf.) https://github.com/sendgrid/sendgrid-go/tree/main/helpers/eventwebhook

	bytes, err := io.ReadAll(r.Body)
	logger.Info("request", "body", string(bytes))
	if err != nil {
		panic(err.Error())
	}
	events, err := sendgrid.ParseEventsJSON(bytes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// MEMO: 実際には各イベントに含まれるタイムスタンプを尊重したい。
	// Webhookの送信から処理完了までの遅延がメトリックの時刻に影響してしまうからだ。
	// すなわち、ここでは分ごとaggregateしてログを出す、もしくは集計せずにログを出すのが望ましい実装だろう。
	// しかし、cloudwatch-logs-aggregatorがtimestampの上書きに対応していないので、ひとまずこれで。
	eventsCountByType := sendgrid.CalcDeliveryEventsCountByType(events)
	logger.Info(
		"sendgrid delivery events count",
		"processedCount", eventsCountByType.DeliveryProcessed,
		"droppedCount", eventsCountByType.DeliveryDropped,
		"deliveredCount", eventsCountByType.DeliveryDelivered,
		"deferredCount", eventsCountByType.DeliveryDeferred,
		"bounceCount", eventsCountByType.DeliveryBounce,
	)
	w.Write([]byte("success."))
}
