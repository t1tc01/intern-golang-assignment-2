package service

import (
	"context"
	"encoding/json"
	"golang-intern/assignment-2/db"
	"golang-intern/assignment-2/ent"
	"golang-intern/assignment-2/ent/apireq"
	"golang-intern/assignment-2/model"
	"golang-intern/assignment-2/util"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/gofiber/fiber/v2"
)

func GetRequestByHour(c *fiber.Ctx, hourNumber int) error {
	apireqs, err := queryRequestByHour(util.Ctx, hourNumber)

	if err != nil {
		return err
	}

	apireq_model, err := convertRequestToBytes(apireqs)

	if err != nil {
		return err
	}

	result, err := json.Marshal(apireq_model)

	if err != nil {
		return err
	}

	return c.Send(result)
}

func GetRequestByWeek(c *fiber.Ctx, weekNumber int) error {
	apireqs, err := queryRequestByWeek(util.Ctx, weekNumber)

	if err != nil {
		return err
	}

	apireq_model, err := convertRequestToBytes(apireqs)

	if err != nil {
		return err
	}

	result, err := json.Marshal(apireq_model)

	if err != nil {
		return err
	}

	return c.Send(result)
}

func queryRequestByHour(c context.Context, hourNumber int) ([]*ent.Apireq, error) {

	hoursAgo := time.Now().Add(-time.Duration(hourNumber) * time.Hour)

	apireqs, err := db.Client.Apireq.Query().
		Where(sql.FieldGT(apireq.FieldCreatedAt, hoursAgo)).
		All(c)

	if err != nil {
		return nil, err
	}

	return apireqs, nil

}

func queryRequestByWeek(c context.Context, weekNumber int) ([]*ent.Apireq, error) {
	weekAgo := time.Now().AddDate(0, 0, -weekNumber)

	apireqs, err := db.Client.Apireq.Query().
		Where(sql.FieldGT(apireq.FieldCreatedAt, weekAgo)).
		All(c)

	if err != nil {
		return nil, err
	}

	return apireqs, nil
}

func convertRequestToBytes(apireqs []*ent.Apireq) ([]model.ApiReq, error) {
	apireq_model := make([]model.ApiReq, 1)

	for _, req := range apireqs {

		req_tstamp := time.Date(req.ReqTime.Year(),
			req.ReqTime.Month(),
			req.ReqTime.Day(), 0, 0, 0, 0, time.UTC).
			Unix()

		resp_tstamp := time.Date(req.RespTime.Year(),
			req.RespTime.Month(),
			req.RespTime.Day(), 0, 0, 0, 0, time.UTC).
			Unix()

		created_tstamp := time.Date(req.CreatedAt.Year(),
			req.CreatedAt.Month(),
			req.CreatedAt.Day(), 0, 0, 0, 0, time.UTC).
			Unix()
		updated_tstamp := time.Date(req.UpdatedAt.Year(),
			req.UpdatedAt.Month(),
			req.UpdatedAt.Day(), 0, 0, 0, 0, time.UTC).
			Unix()

		req_model := model.ApiReq{
			ID:          req.ID,
			ReqTime:     req_tstamp,
			ReqParams:   req.ReqParam,
			ReqBody:     req.ReqBody,
			ReqHeaders:  req.ReqHeaders,
			RespTime:    resp_tstamp,
			RespStatus:  req.RespStatus,
			RespBody:    req.RespBody,
			RespHeaders: req.RespHeaders,
			CreatedAt:   created_tstamp,
			UpdatedAt:   updated_tstamp,
		}

		apireq_model = append(apireq_model, req_model)
	}

	return apireq_model, nil
}
