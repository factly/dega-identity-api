package organization

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/factly/identity/model"
	"github.com/factly/identity/util/render"
	"github.com/go-chi/chi"
)

func delete(w http.ResponseWriter, r *http.Request) {
	organizationID := chi.URLParam(r, "organization_id")
	orgID, err := strconv.Atoi(organizationID)

	organization := &model.Organization{}
	organization.ID = uint(orgID)

	// check record exists or not
	err = model.DB.First(&organization).Error
	if err != nil {
		return
	}

	// check the permission of host
	hostID, _ := strconv.Atoi(r.Header.Get("X-User"))
	host := &model.OrganizationUser{}

	err = model.DB.Model(&model.OrganizationUser{}).Where(&model.OrganizationUser{
		OrganizationID: uint(orgID),
		UserID:         uint(hostID),
		Role:           "owner",
	}).First(host).Error

	if err != nil {
		fmt.Println(err)
		return
	}

	// delete
	model.DB.Delete(&organization)

	render.JSON(w, http.StatusOK, nil)
}