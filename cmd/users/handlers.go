package main

import (
	"encoding/json"
	"net/http"

	"github.com/caseylmanus/users/dbmodels"
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type restHandler struct {
	db boil.ContextExecutor
}

//getUsers handles GET /users
func (handler *restHandler) getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.PlanetUsers(qm.Load(models.PlanetUserRels.GroupPlanetGroups)).All(r.Context(), handler.db)
	if err != nil {
		writeErr(w, err, http.StatusInternalServerError)
		return
	}
	writeJSON(w, fromDbUsers(r.Context(), users, handler.db))
}

//getUser handles GET /users/user_id
func (handler *restHandler) getUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	user, _ := models.FindPlanetUser(r.Context(), handler.db, id)
	if user != nil {
		writeJSON(w, fromDbUsers(r.Context(), models.PlanetUserSlice{user}, handler.db)[0])
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

//postUser handles POST /users
func (handler *restHandler) postUser(w http.ResponseWriter, r *http.Request) {
	var user models.PlanetUser
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		writeErr(w, err, http.StatusBadRequest)
		return
	}
	if err := validateUser(user); err != nil {
		writeErr(w, err, http.StatusBadRequest)
		return
	}
	if err := user.Insert(r.Context(), handler.db, boil.Infer()); err != nil {
		writeErr(w, err, http.StatusInternalServerError)
	}
}

//putUser handles PUT /users/user_id
func (handler *restHandler) putUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var user models.PlanetUser
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil || user.UserID != id {
		writeErr(w, err, http.StatusBadRequest)
		return
	}
	if err := validateUser(user); err != nil {
		writeErr(w, err, http.StatusBadRequest)
		return
	}
	if _, err := user.Update(r.Context(), handler.db, boil.Infer()); err != nil {
		writeErr(w, err, http.StatusInternalServerError)
	}
}

//deleteUsers handles DEL /users/user_id
func (handler *restHandler) deleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	user, _ := models.FindPlanetUser(r.Context(), handler.db, id)

	if user != nil {
		groups, err := user.GroupPlanetGroups().All(r.Context(), handler.db)
		if err != nil {
			writeErr(w, err, http.StatusInternalServerError)
		}
		user.RemoveGroupPlanetGroups(r.Context(), handler.db, groups...)
		_, err = user.Delete(r.Context(), handler.db)
		if err != nil {
			writeErr(w, err, http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
	w.WriteHeader(http.StatusNotFound)
}

//getGroups handles GET /groups
func (handler *restHandler) getGroups(w http.ResponseWriter, r *http.Request) {
	groups, err := models.PlanetGroups(qm.Load(models.PlanetGroupRels.UserPlanetUsers)).All(r.Context(), handler.db)
	if err != nil {
		writeErr(w, err, http.StatusInternalServerError)
		return
	}
	writeJSON(w, fromDbGroups(r.Context(), groups, handler.db))
}

//getGroups handles GET /groups/group_name
func (handler *restHandler) getGroup(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	group, _ := models.FindPlanetGroup(r.Context(), handler.db, id)
	if group != nil {
		writeJSON(w, fromDbGroups(r.Context(), models.PlanetGroupSlice{group}, handler.db)[0])
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

//postGroup handles POST /groups
func (handler *restHandler) postGroup(w http.ResponseWriter, r *http.Request) {
	var group models.PlanetGroup
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		writeErr(w, err, http.StatusBadRequest)
		return
	}
	if err := group.Insert(r.Context(), handler.db, boil.Infer()); err != nil {
		writeErr(w, err, http.StatusInternalServerError)
	}

}

//putGroup handles PUT /groups/group_name
func (handler *restHandler) putGroup(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var groupmembers GroupMembers
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&groupmembers)
	if err != nil {
		writeErr(w, err, http.StatusBadRequest)
		return
	}
	groupmembers.GroupName = id // if they are different sync them
	group, err := models.FindPlanetGroup(r.Context(), handler.db, id)
	if err != nil {
		writeErr(w, err, http.StatusNotFound)
		return
	}
	for _, userid := range groupmembers.Users {
		user, err := models.FindPlanetUser(r.Context(), handler.db, userid)
		if err != nil {
			writeErr(w, errors.Wrapf(err, "invalid user %s", userid), http.StatusNotFound)
			return
		}
		group.AddUserPlanetUsers(r.Context(), handler.db, false, user)
	}

}

//deleteGroup handles DEL /groups/group_name
func (handler *restHandler) deleteGroup(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	group, _ := models.FindPlanetGroup(r.Context(), handler.db, id)

	if group != nil {
		users, err := group.UserPlanetUsers().All(r.Context(), handler.db)
		if err != nil {
			writeErr(w, err, http.StatusInternalServerError)
		}
		group.RemoveUserPlanetUsers(r.Context(), handler.db, users...)
		_, err = group.Delete(r.Context(), handler.db)
		if err != nil {
			writeErr(w, err, http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
	w.WriteHeader(http.StatusNotFound)
}
