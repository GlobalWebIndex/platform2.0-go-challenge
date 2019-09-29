package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	jsonpatch "github.com/evanphx/json-patch"

	"gwiapi/app/auth"
	"gwiapi/app/helpers"
	"gwiapi/app/models"
	"gwiapi/app/responses"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// GetAssets will bring asset of the user
func GetAssets(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	page, _ := helpers.ExportParam(r, "page", "1")
	pg, err := strconv.ParseUint(page, 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	if pg < 1 {
		responses.ERROR(w, http.StatusNotAcceptable, errors.New("Page number is not valid"))
		return
	}

	limit, _ := helpers.ExportParam(r, "limit", "10")
	lmt, err := strconv.ParseUint(limit, 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	if lmt < 1 && lmt > 100 {
		responses.ERROR(w, http.StatusNotAcceptable, errors.New("Number of limit in records result is not supported"))
		return
	}

	// Select only User Assets from Token
	uid, err := auth.ExtractTokenID(r)
	a := &models.Asset{}

	assets, err := a.GetAssets(db, uid, uint32(pg), uint32(lmt))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	// ids := []uint32{}
	// for _, o := range *assets {
	// 	ids = append(ids, o.ID)
	// }

	assetsResponse := []models.Asset{}
	assetsResponse, err = a.GetAssetsWithDetails(db, assets, uid)

	responses.JSON(w, http.StatusOK, assetsResponse)
}

func GetAsset(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	id, found := mux.Vars(r)["id"]
	if !found {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Asset's Id was not passed to Endpoint"))
		return
	}
	assetId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Provider Asset's Id could not be parsed"))
		return
	}

	// Select only User Asset from Token
	uid, err := auth.ExtractTokenID(r)

	a := models.Asset{}
	a, err = models.GetAsset(db, uint32(assetId), uid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	if a.ID == 0 {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	responses.JSON(w, http.StatusOK, a)
}

func DeleteAsset(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	id, found := mux.Vars(r)["id"]
	if !found {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("No asset Id provided"))
		return
	}
	assetId, err := strconv.ParseUint(id, 10, 32)

	// Select only User Asset from Token
	uid, err := auth.ExtractTokenID(r)

	a := &models.Asset{}
	rowsDeleted, err := a.DeleteAsset(db, uint32(assetId), uid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", assetId))
	w.Header().Set("Rows Affected", fmt.Sprintf("%d", rowsDeleted))
	responses.JSON(w, http.StatusNoContent, "")
}

func UpdateAsset(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	uid, err := auth.ExtractTokenID(r)

	id, found := mux.Vars(r)["id"]

	// *Uneccesary based on Route definition
	if !found {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("No asset Id provided"))
		return
	}
	assetId, err := strconv.ParseUint(id, 10, 32)

	// Get Original
	//asset := models.Asset{}
	originalAsset, err := models.GetAsset(db, uint32(assetId), uid)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	originalJSON, err := json.Marshal(originalAsset)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	// Get Patch
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	patchOperation, err := jsonpatch.DecodePatch(body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Apply Patch
	targetJSON, err := patchOperation.Apply(originalJSON)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	targetAsset := models.Asset{}
	err = json.Unmarshal(targetJSON, &targetAsset)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	updatedAsset, err := targetAsset.UpdateAsset(db, uint32(assetId), uid)
	responses.JSON(w, http.StatusOK, updatedAsset)
}

func CreateAssetChart(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	uid, err := auth.ExtractTokenID(r)

	a := &models.Asset{}
	c := &models.Chart{}

	// Parse Body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = json.Unmarshal(body, &c)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	a, c, err = c.CreateAssetChart(db, uid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s/%s/chart/%d", r.Host, r.RequestURI, a.ID))
	responses.JSON(w, http.StatusOK, a)
}

func CreateAssetInsight(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	uid, err := auth.ExtractTokenID(r)

	a := &models.Asset{}
	i := &models.Insight{}

	// Parse Body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = json.Unmarshal(body, &i)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	a, i, err = i.CreateAssetInsight(db, uid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s/%s/insight/%d", r.Host, r.RequestURI, a.ID))
	responses.JSON(w, http.StatusOK, a)
}

func CreateAssetAudience(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	uid, err := auth.ExtractTokenID(r)

	ass := &models.Asset{}
	aud := &models.Audience{}

	// Parse Body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = json.Unmarshal(body, &aud)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	ass, aud, err = aud.CreateAssetAudience(db, uid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s/%s/audience/%d", r.Host, r.RequestURI, ass.ID))
	responses.JSON(w, http.StatusOK, ass)
}
