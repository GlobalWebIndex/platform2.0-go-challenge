package sqlite

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"gwi/model"
	"gwi/utils"

	_ "github.com/mattn/go-sqlite3"
)

type AssetStore struct {
	db *sql.DB
}

func AssetDBConnect(dbName string) (*AssetStore, error) {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, err
	}
	return &AssetStore{db: db}, nil
}

func (s AssetStore) AssetDBInitialize(create bool) error {
	if create {
		err := s.assetDBCreate()
		if err != nil {
			return err
		}
	}
	return nil
}

func (s AssetStore) assetDBCreate() error {
	statement, err := s.db.Prepare("CREATE TABLE IF NOT EXISTS assets (id INTEGER PRIMARY KEY, description TEXT, atype TEXT, serialized_fields TEXT)")
	if err != nil {
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		return err
	}
	statement, err = s.db.Prepare("CREATE TABLE IF NOT EXISTS favorites (id INTEGER PRIMARY KEY, uid INTEGER, aid INTEGER, " +
		" CONSTRAINT fk_user " +
		"FOREIGN KEY (uid) " +
		"REFERENCES users(id) " +
		"ON DELETE CASCADE, " +
		" CONSTRAINT fk_asset " +
		"FOREIGN KEY (aid) " +
		"REFERENCES assets(id) " +
		"ON DELETE CASCADE, UNIQUE(uid,aid) " +
		")")
	if err != nil {
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		return err
	}
	return nil
}

func (s AssetStore) Create(as *model.Asset) (*model.Asset, error) {
	statement, err := s.db.Prepare("INSERT INTO assets (description, atype, serialized_fields) VALUES (?, ?, ?)")
	if err != nil {
		return nil, err
	}
	sf, atype, err := serializeValidObject(as)
	if err != nil {
		return nil, err
	}
	res, err := statement.Exec(as.Description, atype, sf)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	as.ID = uint(id)
	return as, nil
}

func (s AssetStore) Retrieve(as *model.Asset) (*model.Asset, error) {
	row := s.db.QueryRow("SELECT description, atype, serialized_fields FROM assets WHERE id = ?", as.ID)
	if row == nil {
		return nil, utils.ErrNotFound
	}
	var desc string
	var atype string
	var serializedFields string
	if err := row.Scan(&desc, &atype, &serializedFields); err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.ErrNotFound
		}
		return nil, err
	}
	as = &model.Asset{ID: as.ID}
	as.Description = desc
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
	return as, nil
}
func (s AssetStore) Update(as *model.Asset) (*model.Asset, error) {
	statement, err := s.db.Prepare("UPDATE assets SET description = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}
	_, err = statement.Exec(as.Description, as.ID)
	if err != nil {
		return nil, err
	}
	return as, nil
}
func (s AssetStore) Delete(as *model.Asset) (*model.Asset, error) {
	statement, err := s.db.Prepare("DELETE FROM assets WHERE id = ?")
	if err != nil {
		return nil, err
	}
	_, err = statement.Exec(as.ID)
	if err != nil {
		return nil, err
	}
	return as, nil
}

func (s AssetStore) Close() error {
	return s.db.Close()
}

func serializeValidObject(a *model.Asset) (string, string, error) {
	if err := a.IsValid(); err != nil {
		return "", "", err
	}
	if a.Chart != nil {
		b, err := json.Marshal(a.Chart)
		if err != nil {
			return "", "", err
		}
		return string(b), "chart", nil
	} else if a.Insight != nil {
		b, err := json.Marshal(a.Insight)
		if err != nil {
			return "", "", err
		}
		return string(b), "insight", nil
	}
	b, err := json.Marshal(a.Audience)
	if err != nil {
		return "", "", err
	}
	return string(b), "audience", nil
}
