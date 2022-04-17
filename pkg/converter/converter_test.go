package converter

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestConvert_Simple(t *testing.T) {
	type from struct {
		ID   int
		Slug string
	}

	type to struct {
		ID   int
		Slug string
	}

	fromStruct := from{ID: 34, Slug: "test"}

	toStruct, err := ConvertStruct[from, to](fromStruct)

	assert.Nil(t, err)

	assert.Equal(t, fromStruct.ID, toStruct.ID)
	assert.Equal(t, fromStruct.Slug, toStruct.Slug)
}

func TestConvert_Similar(t *testing.T) {
	type from struct {
		ID   int64
		Slug string
	}

	type to struct {
		ID   int
		Slug string
	}

	fromStruct := from{ID: 34, Slug: "test"}

	toStruct, err := ConvertStruct[from, to](fromStruct)

	assert.Nil(t, err)

	assert.Equal(t, 34, toStruct.ID)
	assert.Equal(t, fromStruct.Slug, toStruct.Slug)
}

func TestConvert_Tag(t *testing.T) {
	type from struct {
		ID   int
		Slug string `convert:"slug"`
	}

	type to struct {
		ID          int
		AnotherName string `convert:"slug"`
	}

	fromStruct := from{ID: 34, Slug: "test"}

	toStruct, err := ConvertStruct[from, to](fromStruct)

	assert.Nil(t, err)

	assert.Equal(t, fromStruct.ID, toStruct.ID)
	assert.Equal(t, fromStruct.Slug, toStruct.AnotherName)
}

func TestConvert_StringFromTime(t *testing.T) {
	type from struct {
		ID       int
		DateTime time.Time
	}

	type to struct {
		ID       int
		DateTime string
	}

	now := time.Now()
	fromStruct := from{ID: 34, DateTime: now}

	toStruct, err := ConvertStruct[from, to](fromStruct)

	assert.Nil(t, err)

	assert.Equal(t, fromStruct.ID, toStruct.ID)
	assert.Equal(t, now.Format(time.RFC3339), toStruct.DateTime)
}

func TestConvert_TimeFromString(t *testing.T) {
	type from struct {
		ID       int
		DateTime string
	}

	type to struct {
		ID       int
		DateTime time.Time
	}

	now := time.Now()
	fromStruct := from{ID: 34, DateTime: now.Format(time.RFC3339)}

	toStruct, err := ConvertStruct[from, to](fromStruct)

	assert.Nil(t, err)

	assert.Equal(t, fromStruct.ID, toStruct.ID)
	assert.Equal(t, now.Format(time.RFC3339), toStruct.DateTime.Format(time.RFC3339))
}

func TestConvert_FromNullable(t *testing.T) {
	type from struct {
		ID     int
		UserID *int
	}

	type to struct {
		ID     int
		UserID int
	}

	userID := 8
	fromStruct := from{ID: 34, UserID: &userID}

	toStruct, err := ConvertStruct[from, to](fromStruct)

	assert.Nil(t, err)

	assert.Equal(t, fromStruct.ID, toStruct.ID)
	assert.Equal(t, *fromStruct.UserID, toStruct.UserID)
}

func TestConvert_ToNullable(t *testing.T) {
	type from struct {
		ID     int
		UserID int
	}

	type to struct {
		ID     int
		UserID *int
	}

	fromStruct := from{ID: 34, UserID: 8}

	toStruct, err := ConvertStruct[from, to](fromStruct)

	assert.Nil(t, err)

	assert.Equal(t, fromStruct.ID, toStruct.ID)
	assert.NotNil(t, toStruct.UserID)
	assert.Equal(t, fromStruct.UserID, *toStruct.UserID)
}

func TestConvert_ToNullableTime(t *testing.T) {
	type from struct {
		ID       int
		DateTime string
	}

	type to struct {
		ID       int
		DateTime *time.Time
	}

	now := time.Now()
	fromStruct := from{ID: 34, DateTime: now.Format(time.RFC3339)}

	toStruct, err := ConvertStruct[from, to](fromStruct)

	assert.Nil(t, err)

	assert.Equal(t, fromStruct.ID, toStruct.ID)
	assert.NotNil(t, toStruct.DateTime)
	assert.Equal(t, fromStruct.DateTime, toStruct.DateTime.Format(time.RFC3339))
}

func TestConvert_FromNullableTime(t *testing.T) {
	type from struct {
		ID       int
		DateTime *time.Time
	}

	type to struct {
		ID       int
		DateTime string
	}

	now := time.Now()
	fromStruct := from{ID: 34, DateTime: &now}

	toStruct, err := ConvertStruct[from, to](fromStruct)

	assert.Nil(t, err)

	assert.Equal(t, fromStruct.ID, toStruct.ID)
	assert.Equal(t, fromStruct.DateTime.Format(time.RFC3339), toStruct.DateTime)
}

func TestConvert_FromNullableTime_Nil(t *testing.T) {
	type from struct {
		ID       int
		DateTime *time.Time
	}

	type to struct {
		ID       int
		DateTime string
	}

	fromStruct := from{ID: 34}

	toStruct, err := ConvertStruct[from, to](fromStruct)

	assert.Nil(t, err)

	assert.Equal(t, fromStruct.ID, toStruct.ID)
	assert.Equal(t, "", toStruct.DateTime)
}

func TestConvertSlice(t *testing.T) {
	type from struct {
		ID   int
		Slag string
	}

	type to struct {
		ID   int
		Slag string
	}

	fromSlice := []from{
		{ID: 1, Slag: "test1"},
		{ID: 2, Slag: "test2"},
		{ID: 3, Slag: "test3"},
		{ID: 4, Slag: "test4"},
	}

	toSlice, err := ConvertSlice[from, to](fromSlice)

	assert.Nil(t, err)

	assert.Equal(t, fromSlice[0].ID, toSlice[0].ID)
	assert.Equal(t, fromSlice[0].Slag, toSlice[0].Slag)

	assert.Equal(t, fromSlice[3].ID, toSlice[3].ID)
	assert.Equal(t, fromSlice[3].Slag, toSlice[3].Slag)
}

type benchFrom struct {
	ID       int
	Slag     string
	UserID   *int
	Value    float32
	DateTime time.Time
	EndTime  *time.Time
}

type benchTo struct {
	ID       int
	Slag     string
	UserID   int
	Value    *float32
	DateTime string
	EndTime  string
}

func BenchmarkConvertStruct(b *testing.B) {
	userID := 8
	now := time.Now()
	sample := benchFrom{
		ID:       34,
		Slag:     "asdfoiwfeewf",
		UserID:   &userID,
		Value:    0.45,
		DateTime: time.Now(),
		EndTime:  &now,
	}
	for i := 0; i < b.N; i++ {
		ConvertStruct[benchFrom, benchTo](sample)
	}
}

func manualConvert(from benchFrom) benchTo {
	var userID int
	if from.UserID != nil {
		userID = *from.UserID
	}

	var endDate string
	if from.EndTime != nil {
		endDate = from.DateTime.Format(time.RFC3339)
	}

	return benchTo{
		ID:       from.ID,
		Slag:     from.Slag,
		UserID:   userID,
		Value:    &from.Value,
		DateTime: from.DateTime.Format(time.RFC3339),
		EndTime:  endDate,
	}
}

func BenchmarkManualConvert(b *testing.B) {
	userID := 8
	now := time.Now()
	sample := benchFrom{
		ID:       34,
		Slag:     "asdfoiwfeewf",
		UserID:   &userID,
		Value:    0.45,
		DateTime: time.Now(),
		EndTime:  &now,
	}
	for i := 0; i < b.N; i++ {
		manualConvert(sample)
	}
}
