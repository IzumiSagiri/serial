package serial

import "time"
import "log"

type myStruct struct {
	Tag, Duration string
}

type Status struct {
	TagCounter,
	PlayerCounter uint
	TimeStarted time.Time
	DurationSum,
	DurationNow,
	DurationSession,
	DurationPassed,
	DurationLeft time.Duration
	Data   []myStruct
	Ticked bool
}

func (status *Status) Tick() {
	status.Ticked = true
	status.TagCounter++
	status.DurationSum += status.DurationSession
	var err error
	status.DurationSession, err = time.ParseDuration(status.Data[status.TagCounter].Duration)
	if err != nil {
		log.Fatal(err)
	}
}

func (status *Status) Set() {
	status.DurationNow = time.Since(status.TimeStarted)
	status.DurationPassed = status.DurationNow - status.DurationSum
	if status.DurationPassed >= status.DurationSession {
		status.Tick()
		status.DurationPassed = status.DurationNow - status.DurationSum
	}
	status.DurationLeft = status.DurationSession - status.DurationPassed
}

func (status *Status) Start() {
	status.Data = []myStruct{
		myStruct{"announce_start", "1m10s"},
		myStruct{"relax_1min", "1m"},
		myStruct{"relax_2min", "1m"},
		myStruct{"relax_3min", "1m"},
		myStruct{"relax_4min", "1m"},
		myStruct{"relax_5min", "1m"},
		myStruct{"relax_6min", "1m"},
		myStruct{"relax_7min", "1m"},
		myStruct{"pre_60stimu", "30s"},
		myStruct{"60stimu1", "1s"},
		myStruct{"post_60stimu1_4sec", "4s"},
		myStruct{"post_60stimu1_rest", "50s"},
		myStruct{"60stimu2", "1s"},
		myStruct{"post_60stimu2_4sec", "4s"},
		myStruct{"post_60stimu2_rest", "55s"},
		myStruct{"60stimu3", "1s"},
		myStruct{"post_60stimu3_4sec", "4s"},
		myStruct{"post_60stimu3_rest", "35s"},
		myStruct{"60_1000Hz", "1s"},
		myStruct{"post_60_1000Hz_4sec", "4s"},
		myStruct{"post_60_1000Hz_rest", "25s"},
		myStruct{"90stimu1", "1s"},
		myStruct{"post_90stimu1_4sec", "4s"},
		myStruct{"post_90stimu1_rest", "35s"},
		myStruct{"90stimu2", "1s"},
		myStruct{"post_90stimu2_4sec", "4s"},
		myStruct{"post_90stimu2_rest", "25s"},
		myStruct{"90stimu3", "1s"},
		myStruct{"post_90stimu3_4sec", "4s"},
		myStruct{"post_90stimu3_rest", "50s"},
		myStruct{"announce_final", "10s"},
		myStruct{"for_further_use", "60s"},
	}
	status.TimeStarted = time.Now()
	var err error
	status.DurationSession, err = time.ParseDuration(status.Data[status.TagCounter].Duration)
	if err != nil {
		log.Fatal(err)
	}
}
