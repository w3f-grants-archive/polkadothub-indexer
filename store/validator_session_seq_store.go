package store

import (
	"github.com/figment-networks/polkadothub-indexer/types"
	"github.com/jinzhu/gorm"
	"time"

	"github.com/figment-networks/polkadothub-indexer/model"
)

func NewValidatorSessionSeqStore(db *gorm.DB) *ValidatorSessionSeqStore {
	return &ValidatorSessionSeqStore{scoped(db, model.ValidatorSessionSeq{})}
}

// ValidatorSessionSeqStore handles operations on validators
type ValidatorSessionSeqStore struct {
	baseStore
}

// CreateIfNotExists creates the validator if it does not exist
func (s ValidatorSessionSeqStore) CreateIfNotExists(validator *model.ValidatorSessionSeq) error {
	_, err := s.FindByHeight(validator.Session)
	if isNotFound(err) {
		return s.Create(validator)
	}
	return nil
}

// FindByHeightAndStashAccount finds validator by height and stash account
func (s ValidatorSessionSeqStore) FindByHeightAndStashAccount(height int64, stash string) (*model.ValidatorSessionSeq, error) {
	q := model.ValidatorSessionSeq{
		StashAccount: stash,
	}
	var result model.ValidatorSessionSeq

	err := s.db.
		Where(&q).
		Where("start_height >= ? AND end_height <= ?", height, height).
		First(&result).
		Error

	return &result, checkErr(err)
}

// FindBySessionAndStashAccount finds validator by session and stash account
func (s ValidatorSessionSeqStore) FindBySessionAndStashAccount(session int64, stash string) (*model.ValidatorSessionSeq, error) {
	q := model.ValidatorSessionSeq{
		SessionSequence: &model.SessionSequence{
			Session: session,
		},
		StashAccount: stash,
	}
	var result model.ValidatorSessionSeq

	err := s.db.
		Where(&q).
		First(&result).
		Error

	return &result, checkErr(err)
}

// FindByHeight finds validator session sequences by height
func (s ValidatorSessionSeqStore) FindByHeight(h int64) ([]model.ValidatorSessionSeq, error) {
	var result []model.ValidatorSessionSeq

	err := s.db.
		Where("start_height >= ? AND end_height <= ?", h, h).
		Find(&result).
		Error

	return result, checkErr(err)
}

// FindByHeight finds validator session sequences by session
func (s ValidatorSessionSeqStore) FindBySession(session int64) ([]model.ValidatorSessionSeq, error) {
	q := model.ValidatorSessionSeq{
		SessionSequence: &model.SessionSequence{
			Session: session,
		},
	}
	var result []model.ValidatorSessionSeq

	err := s.db.
		Where(&q).
		Find(&result).
		Error

	return result, checkErr(err)
}

// FindMostRecent finds most recent validator session sequence
func (s *ValidatorSessionSeqStore) FindMostRecent() (*model.ValidatorSessionSeq, error) {
	validatorSeq := &model.ValidatorSessionSeq{}
	if err := findMostRecent(s.db, "time", validatorSeq); err != nil {
		return nil, err
	}
	return validatorSeq, nil
}

// DeleteOlderThan deletes validator sequence older than given threshold
func (s *ValidatorSessionSeqStore) DeleteOlderThan(purgeThreshold time.Time) (*int64, error) {
	tx := s.db.
		Unscoped().
		Where("time < ?", purgeThreshold).
		Delete(&model.ValidatorSessionSeq{})

	if tx.Error != nil {
		return nil, checkErr(tx.Error)
	}

	return &tx.RowsAffected, nil
}

type ValidatorSessionSeqSummary struct {
	Address         string         `json:"address"`
	TimeBucket      types.Time     `json:"time_bucket"`
	VotingPowerAvg  float64        `json:"voting_power_avg"`
	VotingPowerMax  float64        `json:"voting_power_max"`
	VotingPowerMin  float64        `json:"voting_power_min"`
	TotalSharesAvg  types.Quantity `json:"total_shares_avg"`
	TotalSharesMax  types.Quantity `json:"total_shares_max"`
	TotalSharesMin  types.Quantity `json:"total_shares_min"`
	ValidatedSum    int64          `json:"validated_sum"`
	NotValidatedSum int64          `json:"not_validated_sum"`
	ProposedSum     int64          `json:"proposed_sum"`
	UptimeAvg       float64        `json:"uptime_avg"`
}

// Summarize gets the summarized version of validator sequences
//func (s *ValidatorSessionSeqStore) Summarize(interval types.SummaryInterval, activityPeriods []ActivityPeriodRow) ([]ValidatorSessionSeqSummary, error) {
//	defer logQueryDuration(time.Now(), "ValidatorSessionSeqStore_Summarize")
//
//	tx := s.db.
//		Table(model.ValidatorSessionSeq{}.TableName()).
//		Select(summarizeValidatorsQuerySelect, interval).
//		Order("time_bucket").
//		Group("address, time_bucket")
//
//	if len(activityPeriods) == 1 {
//		activityPeriod := activityPeriods[0]
//		tx = tx.Or("time < ? OR time >= ?", activityPeriod.Min, activityPeriod.Max)
//	} else {
//		for i, activityPeriod := range activityPeriods {
//			isLast := i == len(activityPeriods)-1
//
//			if isLast {
//				tx = tx.Or("time >= ?", activityPeriod.Max)
//			} else {
//				duration, err := time.ParseDuration(fmt.Sprintf("1%s", interval))
//				if err != nil {
//					return nil, err
//				}
//				tx = tx.Or("time >= ? AND time < ?", activityPeriod.Max.Add(duration), activityPeriods[i+1].Min)
//			}
//		}
//	}
//
//	rows, err := tx.Rows()
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//
//	var models []ValidatorSessionSeqSummary
//	for rows.Next() {
//		var summary ValidatorSessionSeqSummary
//		if err := s.db.ScanRows(rows, &summary); err != nil {
//			return nil, err
//		}
//
//		models = append(models, summary)
//	}
//	return models, nil
//}
