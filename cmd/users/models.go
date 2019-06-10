package main

import (
	"context"

	"github.com/caseylmanus/users/dbmodels"
	"github.com/volatiletech/sqlboiler/boil"
)

type User struct {
	UserID    string   `json:"user_id"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Groups    []string `json:"groups"`
}
type Group struct {
	GroupName string   `json:"group_name"`
	Users     []string `json:"users"`
}
type GroupMembers struct {
	GroupName string   `json:"group_name"`
	Users     []string `json:"users"`
}

func fromDbUsers(ctx context.Context, dbusers models.PlanetUserSlice, db boil.ContextExecutor) []User {
	users := make([]User, len(dbusers))
	for i := range dbusers {
		users[i].UserID = dbusers[i].UserID
		users[i].FirstName = dbusers[i].FirstName
		users[i].LastName = dbusers[i].LastName
		users[i].Groups = []string{}
		groups, _ := dbusers[i].GroupPlanetGroups().All(ctx, db) //can't error due to previous eagar loading
		for _, group := range groups {
			users[i].Groups = append(users[i].Groups, group.GroupName)
		}
	}
	return users
}

func fromDbGroups(ctx context.Context, dbgroups models.PlanetGroupSlice, db boil.ContextExecutor) []Group {
	groups := make([]Group, len(dbgroups))
	for i := range dbgroups {
		groups[i].GroupName = dbgroups[i].GroupName
		groups[i].Users = []string{}
		users, _ := dbgroups[i].UserPlanetUsers().All(ctx, db)
		for _, user := range users {
			groups[i].Users = append(groups[i].Users, user.UserID)
		}
	}
	return groups
}
