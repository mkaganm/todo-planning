package data

// GetDevelopers is a struct that represents a developer from database
func GetDevelopers() ([]Developer, error) {

	db := InitDB()
	defer Close(db)

	var developers []Developer
	if err := db.Find(&developers).Error; err != nil {
		return nil, err
	}

	return developers, nil
}
