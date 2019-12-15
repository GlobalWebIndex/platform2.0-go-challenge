package repositories

import "gwi-challenge/data"

func Save(assetToCreate interface{}) error {
	db := data.GetDB()
	err := db.Save(assetToCreate).Error
	return err
}

func Delete(assetToDelete interface{}) error {
	db := data.GetDB()
	err := db.Delete(assetToDelete).Error
	return err
}
