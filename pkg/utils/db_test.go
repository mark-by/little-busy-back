package utils

import (
	"github.com/bmizerany/assert"
	"github.com/nleeper/goment"
	"log"
	"testing"
	"time"
)

func TestSQLSlice(t *testing.T) {
	res := SQLSlice([]int{1, 2})
	assert.Equal(t, "1,2", res)
}

func TestEndTx(t *testing.T) {
	err := goment.SetLocale("ru")
	g, err := goment.New(time.Now().AddDate(0, 0, 1))
	g2, err := goment.New(time.Now().AddDate(0, 0, 1).Add(150 * time.Minute))
	log.Println(g.Format("Do в H:mm"), "до", g2.Format("H:mm, MMMM, dddd"))
	log.Println(err)
}
