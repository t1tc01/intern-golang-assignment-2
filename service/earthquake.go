package service

import (
	"context"
	"encoding/json"
	"golang-intern/assignment-2/db"
	"golang-intern/assignment-2/ent"
	"golang-intern/assignment-2/ent/earthquake"
	"golang-intern/assignment-2/model"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Get
func GetAllEarthquake(c *fiber.Ctx) error {
	eqs, err := querryAllEathquake(ctx)

	if err != nil {
		return err
	}

	eqs_model, err := convertEarthquakeToBytes(eqs)

	if err != nil {
		return err
	}

	result, err := json.Marshal(eqs_model)

	if err != nil {
		return err
	}

	return c.Send(result)
}

func GetEarthquakeByLimitOffset(c *fiber.Ctx, limit int, offset int) error {
	eqs, err := querryLimitOffsetEarthquake(ctx, limit, offset)

	if err != nil {
		return err
	}

	eqs_model, err := convertEarthquakeToBytes(eqs)

	if err != nil {
		return err
	}

	result, err := json.Marshal(eqs_model)

	if err != nil {
		return err
	}

	return c.Send(result)
}

func GetEarthquakeFilteredByID(c *fiber.Ctx, id int) error {
	eqs, err := querryEarthquakeFilteredByID(ctx, id)

	if err != nil {
		return err
	}

	eqs_model, err := convertEarthquakeToBytes(eqs)

	if err != nil {
		return err
	}

	result, err := json.Marshal(eqs_model)

	if err != nil {
		return err
	}

	return c.Send(result)
}

// querry
func querryAllEathquake(c context.Context) ([]*ent.Earthquake, error) {
	eqs, err := db.Client.Earthquake.Query().All(c)

	if err != nil {
		return nil, err
	}

	return eqs, nil
}

func querryLimitOffsetEarthquake(c context.Context, limit int, offset int) ([]*ent.Earthquake, error) {
	eqs, err := db.Client.Earthquake.Query().
		Limit(limit).
		Offset(offset).
		All(c)

	if err != nil {
		return nil, err
	}

	return eqs, nil
}

func querryEarthquakeFilteredByID(c context.Context, id int) ([]*ent.Earthquake, error) {
	eqs, err := db.Client.Earthquake.Query().
		Where(earthquake.ID(id)).
		All(c)

	if err != nil {
		return nil, err
	}

	return eqs, nil
}

// convert
func convertEarthquakeToBytes(eqs []*ent.Earthquake) ([]model.EarthquakeModel, error) {
	eqs_model := make([]model.EarthquakeModel, 1)

	for _, eq := range eqs {

		tstamp := time.Date(eq.Time.Year(),
			eq.Time.Month(),
			eq.Time.Day(), 0, 0, 0, 0, time.UTC).
			Unix()
		updated_tstamp := time.Date(eq.UpdatedTime.Year(),
			eq.UpdatedTime.Month(),
			eq.UpdatedTime.Day(), 0, 0, 0, 0, time.UTC).
			Unix()
		created_tstamp := time.Date(eq.CreatedAt.Year(),
			eq.CreatedAt.Month(),
			eq.CreatedAt.Day(), 0, 0, 0, 0, time.UTC).
			Unix()

		updatedRow_tstamp := time.Date(eq.UpdatedAt.Year(),
			eq.UpdatedAt.Month(),
			eq.UpdatedAt.Day(), 0, 0, 0, 0, time.UTC).
			Unix()

		eq_model := model.EarthquakeModel{
			ID:          eq.ID,
			Mag:         eq.Mag,
			Time:        tstamp,
			Updated:     updated_tstamp,
			Tz:          eq.Tz,
			URL:         eq.URL,
			Detail:      eq.Detail,
			Status:      eq.Status,
			Tsunami:     eq.Tsunami,
			Sig:         eq.Sig,
			Net:         eq.Net,
			Code:        eq.Code,
			NST:         eq.Nst,
			Dmin:        eq.Dmin,
			RMS:         eq.Rms,
			Gap:         eq.Gap,
			MagType:     eq.MagType,
			Type:        eq.EqType,
			CreatedTime: created_tstamp,
			UpdatedTime: updatedRow_tstamp,
		}

		eqs_model = append(eqs_model, eq_model)
	}
	return eqs_model, nil
}
