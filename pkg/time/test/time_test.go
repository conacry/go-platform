package timeTest

import (
	"github.com/conacry/go-platform/pkg/time"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
	goTime "time"
)

type TimeToolsShould struct {
	suite.Suite
}

func TestTimeToolsShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(TimeToolsShould))
}

func (s *TimeToolsShould) TestNow() {
	nowUTCTime := time.Now()
	require.NotNil(s.T(), nowUTCTime)

	utcNsec1 := nowUTCTime.UnixNano()
	utcNsec2 := nowUTCTime.UnixNano()

	assert.Equal(s.T(), goTime.UTC, nowUTCTime.Location())
	assert.Equal(s.T(), utcNsec1, utcNsec2)
	assert.True(s.T(), utcNsec1 > 0)
}

func (s *TimeToolsShould) TestEmpty() {
	emptyTime := time.Empty()
	require.NotNil(s.T(), emptyTime)

	assert.Equal(s.T(), int64(0), emptyTime.UnixNano())
}

func (s *TimeToolsShould) TestFromTime_ReturnTime() {
	expectedGoTime := goTime.Now()

	actualTime := time.FromTime(expectedGoTime)
	require.NotNil(s.T(), actualTime)

	assert.Equal(s.T(), goTime.UTC, actualTime.Location())
	assert.True(s.T(), actualTime.Time.Equal(expectedGoTime))

	require.False(s.T(), actualTime.Time.IsZero())
	assert.Equal(s.T(), expectedGoTime.UnixNano(), actualTime.Time.UnixNano())
}

func (s *TimeToolsShould) TestFromUnixNano_ValidSeconds_ReturnTime() {
	expectedTime := time.Now()
	require.NotNil(s.T(), expectedTime)

	actualTime := time.FromUnixNano(expectedTime.UnixNano())
	require.NotNil(s.T(), actualTime)

	assert.Equal(s.T(), goTime.UTC, actualTime.Location())
	assert.Equal(s.T(), expectedTime, actualTime)
}

func (s *TimeToolsShould) TestFromUnixNano_ZeroSeconds_ReturnEmptyTime() {
	emptyTime := time.Empty()
	require.NotNil(s.T(), emptyTime)

	timeFromZero := time.FromUnixNano(0)
	require.NotNil(s.T(), timeFromZero)

	assert.Equal(s.T(), emptyTime, timeFromZero)
}

func (s *TimeToolsShould) TestParse_ReturnTime() {
	expectedTime := time.Now()
	require.NotNil(s.T(), expectedTime)

	layout := goTime.RFC3339Nano
	expectedTimeStr := expectedTime.Format(layout)

	parsedTime, err := time.Parse(layout, expectedTimeStr)
	assert.Equal(s.T(), expectedTime, parsedTime)
	assert.NoError(s.T(), err)
}

func (s *TimeToolsShould) TestLocal_ReturnTime() {
	nowTime := time.Now()
	require.NotNil(s.T(), nowTime)

	nowNsec := nowTime.UnixNano()

	localNowTime := nowTime.Local()
	require.NotNil(s.T(), localNowTime)

	localNowNsec := localNowTime.UnixNano()

	assert.Equal(s.T(), goTime.UTC, nowTime.Location())
	assert.Equal(s.T(), goTime.Local, localNowTime.Location())
	assert.Equal(s.T(), nowNsec, localNowNsec)
}

func (s *TimeToolsShould) TestAdd_ReturnTime() {
	nowTime := time.Now()
	require.NotNil(s.T(), nowTime)

	expectedTimeNsec := nowTime.UnixNano() + int64(goTime.Second)

	actualTime := nowTime.Add(goTime.Second)
	require.NotNil(s.T(), actualTime)

	actualTimeNsec := actualTime.UnixNano()

	assert.Equal(s.T(), expectedTimeNsec, actualTimeNsec)
	assert.NotEqual(s.T(), nowTime.UnixNano(), actualTimeNsec)
}

func (s *TimeToolsShould) TestSub_ReturnTime() {
	nowTime := time.Now()
	require.NotNil(s.T(), nowTime)

	expectedTimeNsec := nowTime.UnixNano() - int64(goTime.Second)

	actualTime := nowTime.Sub(goTime.Second)
	require.NotNil(s.T(), actualTime)

	actualTimeNsec := actualTime.UnixNano()

	assert.Equal(s.T(), expectedTimeNsec, actualTimeNsec)
	assert.NotEqual(s.T(), nowTime.UnixNano(), actualTimeNsec)
}

func (s *TimeToolsShould) TestEqual_TimesIsEqual_ReturnTrue() {
	nowTime := time.Now()
	require.NotNil(s.T(), nowTime)

	t1 := time.FromUnixNano(nowTime.UnixNano())
	require.NotNil(s.T(), t1)

	t2 := time.FromUnixNano(nowTime.UnixNano())
	require.NotNil(s.T(), t2)

	assert.True(s.T(), t1.Equal(t2))
}

func (s *TimeToolsShould) TestEqual_TimesIsNotEqual_ReturnFalse() {
	nowTime := time.Now()
	require.NotNil(s.T(), nowTime)

	t1 := time.FromUnixNano(nowTime.UnixNano())
	require.NotNil(s.T(), t1)

	t2 := time.FromUnixNano(nowTime.UnixNano() + int64(goTime.Second))
	require.NotNil(s.T(), t2)

	assert.False(s.T(), t1.Equal(t2))
}

func (s *TimeToolsShould) TestBefore_TimeIsBefore_ReturnTrue() {
	nowTime := time.Now()
	require.NotNil(s.T(), nowTime)

	pastTime := nowTime.Sub(goTime.Second)
	require.NotNil(s.T(), pastTime)

	assert.True(s.T(), pastTime.Before(nowTime))
}

func (s *TimeToolsShould) TestBefore_TimeIsAfter_ReturnFalse() {
	nowTime := time.Now()
	require.NotNil(s.T(), nowTime)

	futureTime := nowTime.Add(goTime.Second)
	require.NotNil(s.T(), futureTime)

	assert.False(s.T(), futureTime.Before(nowTime))
}

func (s *TimeToolsShould) TestAfter_TimeIsAfter_ReturnTrue() {
	nowTime := time.Now()
	require.NotNil(s.T(), nowTime)

	futureTime := nowTime.Add(goTime.Second)
	require.NotNil(s.T(), futureTime)

	assert.True(s.T(), futureTime.After(nowTime))
}

func (s *TimeToolsShould) TestAfter_TimeIsBefore_ReturnFalse() {
	nowTime := time.Now()
	require.NotNil(s.T(), nowTime)

	pastTime := nowTime.Sub(goTime.Second)
	require.NotNil(s.T(), pastTime)

	assert.False(s.T(), pastTime.After(nowTime))
}

func (s *TimeToolsShould) TestUnix_TimeIsZero_ReturnZero() {
	emptyTime := time.Empty()
	require.NotNil(s.T(), emptyTime)

	nsec := emptyTime.Unix()

	assert.Equal(s.T(), int64(0), nsec)
}

func (s *TimeToolsShould) TestUnix_TimeIsNotZero_ReturnNotZero() {
	expectedTime := time.Now()
	require.NotNil(s.T(), expectedTime)

	expectedNsec := expectedTime.UnixNano()

	actualTime := time.FromUnixNano(expectedNsec)
	require.NotNil(s.T(), actualTime)

	actualSec := actualTime.Unix()

	assert.Equal(s.T(), expectedNsec/1e9, actualSec)
}

func (s *TimeToolsShould) TestUnixNano_TimeIsZero_ReturnZero() {
	emptyTime := time.Empty()
	require.NotNil(s.T(), emptyTime)

	nsec := emptyTime.UnixNano()

	assert.Equal(s.T(), int64(0), nsec)
}

func (s *TimeToolsShould) TestUnixNano_TimeIsNotZero_ReturnNotZero() {
	expectedTime := time.Now()
	require.NotNil(s.T(), expectedTime)

	expectedNsec := expectedTime.UnixNano()

	actualTime := time.FromUnixNano(expectedNsec)
	require.NotNil(s.T(), actualTime)

	actualNsec := actualTime.UnixNano()

	assert.Equal(s.T(), expectedNsec, actualNsec)
}

func (s *TimeToolsShould) TestUnixMilli_TimeIsZero_ReturnZero() {
	emptyTime := time.Empty()
	require.NotNil(s.T(), emptyTime)

	msec := emptyTime.UnixMilli()

	assert.Equal(s.T(), int64(0), msec)
}

func (s *TimeToolsShould) TestUnixMilli_TimeIsNotZero_ReturnNotZero() {
	expectedTime := time.Now()
	require.NotNil(s.T(), expectedTime)

	expectedMsec := expectedTime.UnixMilli()

	actualTime := time.FromUnixMillis(expectedMsec)
	require.NotNil(s.T(), actualTime)

	actualMsec := actualTime.UnixMilli()

	assert.Equal(s.T(), expectedMsec, actualMsec)
}
