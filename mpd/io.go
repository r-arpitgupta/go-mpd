package mpd

import (
	"io"
	"os"
	"time"

	"github.com/Eyevinn/dash-mpd/xml"
)

// ReadFromFile reads and unmarshals an MPD from a file.
func ReadFromFile(path string) (*MPD, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	mpd := MPD{}
	err = xml.Unmarshal(data, &mpd)
	if err != nil {
		return nil, err
	}
	return &mpd, nil
}

// ReadFromString reads and unmarshals an MPD from a string
func ReadFromString(str string) (*MPD, error) {
	mpd := MPD{}
	err := xml.Unmarshal([]byte(str), &mpd)
	if err != nil {
		return nil, err
	}
	return &mpd, nil
}

// Write marshals and writes an MPD.
func (m *MPD) Write(w io.Writer) (int, error) {
	data, err := xml.MarshalIndent(m, "", "  ")
	if err != nil {
		return 0, err
	}
	return w.Write(data)
}

const RFC3339MS string = "2006-01-02T15:04:05.999Z07:00"

// ConvertToDateTime converts a number of seconds to a UTC DateTime by cropping to ms precision.
func ConvertToDateTime(seconds float64) DateTime {
	s := int64(seconds)
	ns := int64((seconds - float64(s)) * 1_000_000_000)
	t := time.Unix(s, ns).UTC()

	return DateTime(t.Format(RFC3339MS))
}

// ConvertToDateTime converts an integral number of seconds to a UTC DateTime.
func ConvertToDateTimeS(seconds int64) DateTime {
	t := time.Unix(seconds, 0).UTC()
	return DateTime(t.Format(RFC3339MS))
}

// ConvertToDateTime converts an integral number of milliseconds to a UTC DateTime.
func ConvertToDateTimeMS(ms int64) DateTime {
	seconds := ms / 1000
	ns := (ms - 1000*seconds) * 1_000_000
	t := time.Unix(seconds, ns).UTC()
	return DateTime(t.Format(RFC3339MS))
}

// ConvertToSeconds converts a DateTime to a number of seconds.
func (dt DateTime) ConvertToSeconds() (float64, error) {
	t, err := time.Parse(RFC3339MS, string(dt))
	if err != nil {
		return 0, err
	}
	return float64(t.UnixNano()) / 1_000_000_000, nil
}
