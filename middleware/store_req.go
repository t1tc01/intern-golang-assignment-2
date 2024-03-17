package middleware

import (
	"context"
	"golang-intern/assignment-2/db"
	"golang-intern/assignment-2/ent"
	"time"
)

func StoreAPIRequestRespond(ctx context.Context,
	requestTime time.Time,
	requestParameter string, requestHeaders string, requestBody string,
	responseTime time.Time,
	responseStatus int, responseHeaders string, responseBody string) (*ent.Apireq, error) {

	currentTime := time.Now()

	apireq := db.Client.Apireq.Create().
		SetReqTime(requestTime).
		SetReqParam(requestParameter).
		SetReqHeaders(requestHeaders).
		SetReqBody(requestBody).
		SetRespTime(responseTime).
		SetRespStatus(responseStatus).
		SetRespHeaders(responseHeaders).
		SetRespBody(responseBody).
		SetCreatedAt(currentTime).
		SetUpdatedAt(currentTime).
		SaveX(ctx)

	return apireq, nil
}
