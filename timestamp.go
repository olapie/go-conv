package conv

import "time"

func ToTimeP(sec *int64) *time.Time {
	if sec == nil {
		return nil
	}
	t := time.Unix(*sec, 0)
	return &t
}

func FromTimeP(t *time.Time) *int64 {
	if t == nil {
		return nil
	}
	i := t.Unix()
	return &i
}
