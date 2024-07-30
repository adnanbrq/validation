package rules

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T) {
	rule := TimeRule{}
	noTime := []string{errNoTime}
	notPast := []string{errTimeNotInPast}
	notFuture := []string{errTimeNotInFuture}
	notToday := []string{errTimeNotToday}
	notYesterday := []string{errTimeNotYesterday}

	// Name
	assert.NotEmpty(t, rule.Name())

	// Input not time
	assert.Equal(t, noTime, rule.Validate(nil, nil))
	assert.Equal(t, noTime, rule.Validate(false, nil))
	assert.Equal(t, noTime, rule.Validate(uintptr(5), nil))
	assert.Equal(t, noTime, rule.Validate(rand.Int31(), nil))
	assert.Equal(t, noTime, rule.Validate(rand.Int63(), nil))
	assert.Equal(t, noTime, rule.Validate(rand.Float32(), nil))
	assert.Equal(t, noTime, rule.Validate(rand.Float64(), nil))
	assert.Equal(t, noTime, rule.Validate(map[string]string{}, nil))
	assert.Equal(t, noTime, rule.Validate(make(chan int), nil))

	// No Option
	assert.Equal(t, noErrs, rule.Validate(time.Now(), nil))
	assert.Equal(t, noErrs, rule.Validate(time.Now(), ""))
	assert.Equal(t, noErrs, rule.Validate(time.Now(), true))
	assert.Equal(t, noErrs, rule.Validate(time.Now(), map[string]string{}))
	assert.Equal(t, noErrs, rule.Validate(time.Now(), make(chan int)))

	// INFO: Be aware that adding or subtracting time can lead to a fail depending on the time the test is run.

	// Today
	assert.Equal(t, noErrs, rule.Validate(time.Now(), "today"))
	assert.Equal(t, notToday, rule.Validate(time.Now().Add(-time.Hour*24), "today"))
	assert.Equal(t, notToday, rule.Validate(time.Now().Add(time.Hour*24), "today"))

	// Yesterday
	assert.Equal(t, notYesterday, rule.Validate(time.Now(), "yesterday"))
	assert.Equal(t, noErrs, rule.Validate(time.Now().Add(-time.Hour*24), "yesterday"))
	assert.Equal(t, notYesterday, rule.Validate(time.Now().Add(time.Hour*24), "yesterday"))

	// Future
	assert.Equal(t, notFuture, rule.Validate(time.Now(), "future")) // err because time.Now() is requested before the comparing timestamp inside Validate
	assert.Equal(t, noErrs, rule.Validate(time.Now().Add((time.Hour*24)), "future"))
	assert.Equal(t, noErrs, rule.Validate(time.Now().Add(time.Hour), "future"))
	assert.Equal(t, notFuture, rule.Validate(time.Now().Add(-(time.Hour*24)), "future"))
	assert.Equal(t, notFuture, rule.Validate(time.Now().Add(-(time.Hour)), "future"))

	// Past
	assert.Equal(t, noErrs, rule.Validate(time.Now(), "past")) // valid as the timestamp inside the validate is requestd later than the given input time.Now()
	assert.Equal(t, noErrs, rule.Validate(time.Now().Add(-(time.Hour*24)), "past"))
	assert.Equal(t, notPast, rule.Validate(time.Now().Add(time.Hour), "past"))
}
