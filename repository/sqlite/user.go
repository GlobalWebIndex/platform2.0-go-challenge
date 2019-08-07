package sqlite

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"gwi/model"
	"gwi/repository"
	"gwi/utils"

	_ "github.com/mattn/go-sqlite3"
)

type UserStore struct {
	db *sql.DB
}

func UserDBConnect(dbName string) (*UserStore, error) {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, err
	}
	return &UserStore{db: db}, nil
}

func (s UserStore) UserDBInitialize(create bool) error {
	if create {
		err := s.userDBCreate()
		if err != nil {
			return err
		}
	}
	return nil
}

func (s UserStore) userDBCreate() error {
	statement, err := s.db.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY)")
	if err != nil {
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		return err
	}
	return nil
}

func (s UserStore) Create(us *model.User) (*model.User, error) {
	statement, err := s.db.Prepare("INSERT INTO users DEFAULT VALUES")
	if err != nil {
		return nil, err
	}
	res, err := statement.Exec()
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	us.ID = uint(id)
	return us, nil
}

func (s UserStore) Retrieve(us *model.User) (*model.User, error) {
	row := s.db.QueryRow("SELECT id FROM users WHERE id = ?", us.ID)
	if row == nil {
		return nil, utils.ErrNotFound
	}
	if err := row.Scan(&us.ID); err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.ErrNotFound
		}
		return nil, err
	}
	return us, nil
}

func (s UserStore) Delete(us *model.User) (*model.User, error) {
	statement, err := s.db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return nil, err
	}
	_, err = statement.Exec(us.ID)
	if err != nil {
		return nil, err
	}
	return us, nil
}

func (s UserStore) GetAllUsers() ([]*model.User, error) {
	users := make([]*model.User, 0)
	rows, err := s.db.Query("SELECT id FROM users")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var us model.User
		if err := rows.Scan(&us.ID); err != nil {
			return nil, err
		}
		users = append(users, &us)
	}
	return users, nil
}

func (s UserStore) GetUserFavourites(us *model.User, startIndex *uint, pageSize uint, ar repository.AssetRepo) (*model.User, error) {
	rows, err := s.db.Query("SELECT a.id, description, atype, serialized_fields FROM assets a inner join favorites f on a.id = f.aid inner join users u on u.id = f.uid WHERE u.id = ? and a.id > ? ORDER BY a.id LIMIT ?", us.ID, startIndex, pageSize)
	if err != nil {
		return nil, err
	}
	us.Assets = make(map[string]model.Asset)
	for rows.Next() {
		var id uint
		var desc string
		var atype string
		var serializedFields string
		if err := rows.Scan(&id, &desc, &atype, &serializedFields); err != nil {
			return nil, err
		}
		as := &model.Asset{ID: id, Description: desc}
		switch atype {
		case "chart":
			as.Chart = new(model.Chart)
			if err := json.Unmarshal([]byte(serializedFields), as.Chart); err != nil {
				fmt.Println(err.Error())
				return nil, err
			}
		case "insight":
			as.Insight = new(model.Insight)
			if err := json.Unmarshal([]byte(serializedFields), as.Insight); err != nil {
				return nil, err
			}
		case "audience":
			as.Audience = new(model.Audience)
			if err := json.Unmarshal([]byte(serializedFields), as.Audience); err != nil {
				return nil, err
			}
		}
		us.Assets[fmt.Sprint(as.ID)] = *as
		*startIndex = as.ID
	}
	return us, nil
}

func (s UserStore) AddUserFavourite(us *model.User, ar repository.AssetRepo, as *model.Asset) (*model.User, error) {
	statement, err := s.db.Prepare("INSERT INTO favorites (uid, aid) VALUES (?, ?)")
	if err != nil {
		return nil, err
	}
	res, err := statement.Exec(us.ID, as.ID)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	us.ID = uint(id)
	return us, nil
}

func (s UserStore) RemoveUserFavourite(us *model.User, ar repository.AssetRepo, as *model.Asset) (*model.User, error) {
	statement, err := s.db.Prepare("DELETE FROM favorites WHERE uid = ? and aid = ?")
	if err != nil {
		return nil, err
	}
	res, err := statement.Exec(us.ID, as.ID)
	if err != nil {
		return nil, err
	}
	affect, err := res.RowsAffected()
	if err != nil || affect == 0 {
		return nil, utils.ErrNotFound
	}
	return us, nil
}

func (s UserStore) Close() error {
	return s.db.Close()
}
