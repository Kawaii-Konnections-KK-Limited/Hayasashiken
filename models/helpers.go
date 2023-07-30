package models

func GetAllRecords() ([]Links, error) {

	var records []Links
	result := database_.Find(&records)

	return records, result.Error
}
