// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/unclesp1d3r/cipherswarm-agent-sdk/internal/utils"
	"time"
)

type TaskStatus struct {
	OriginalLine    string         `json:"original_line"`
	Time            time.Time      `json:"time"`
	Session         string         `json:"session"`
	HashcatGuess    HashcatGuess   `json:"hashcat_guess"`
	Status          int64          `json:"status"`
	Target          string         `json:"target"`
	Progress        []int64        `json:"progress"`
	RestorePoint    int64          `json:"restore_point"`
	RecoveredHashes []int64        `json:"recovered_hashes"`
	RecoveredSalts  []int64        `json:"recovered_salts"`
	Rejected        int64          `json:"rejected"`
	DeviceStatuses  []DeviceStatus `json:"device_statuses"`
	TimeStart       time.Time      `json:"time_start"`
	EstimatedStop   time.Time      `json:"estimated_stop"`
}

func (t TaskStatus) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TaskStatus) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TaskStatus) GetOriginalLine() string {
	if o == nil {
		return ""
	}
	return o.OriginalLine
}

func (o *TaskStatus) GetTime() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Time
}

func (o *TaskStatus) GetSession() string {
	if o == nil {
		return ""
	}
	return o.Session
}

func (o *TaskStatus) GetHashcatGuess() HashcatGuess {
	if o == nil {
		return HashcatGuess{}
	}
	return o.HashcatGuess
}

func (o *TaskStatus) GetStatus() int64 {
	if o == nil {
		return 0
	}
	return o.Status
}

func (o *TaskStatus) GetTarget() string {
	if o == nil {
		return ""
	}
	return o.Target
}

func (o *TaskStatus) GetProgress() []int64 {
	if o == nil {
		return []int64{}
	}
	return o.Progress
}

func (o *TaskStatus) GetRestorePoint() int64 {
	if o == nil {
		return 0
	}
	return o.RestorePoint
}

func (o *TaskStatus) GetRecoveredHashes() []int64 {
	if o == nil {
		return []int64{}
	}
	return o.RecoveredHashes
}

func (o *TaskStatus) GetRecoveredSalts() []int64 {
	if o == nil {
		return []int64{}
	}
	return o.RecoveredSalts
}

func (o *TaskStatus) GetRejected() int64 {
	if o == nil {
		return 0
	}
	return o.Rejected
}

func (o *TaskStatus) GetDeviceStatuses() []DeviceStatus {
	if o == nil {
		return []DeviceStatus{}
	}
	return o.DeviceStatuses
}

func (o *TaskStatus) GetTimeStart() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.TimeStart
}

func (o *TaskStatus) GetEstimatedStop() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.EstimatedStop
}
