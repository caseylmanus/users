// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// PlanetUser is an object representing the database table.
type PlanetUser struct {
	UserID    string `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	FirstName string `boil:"first_name" json:"first_name" toml:"first_name" yaml:"first_name"`
	LastName  string `boil:"last_name" json:"last_name" toml:"last_name" yaml:"last_name"`

	R *planetUserR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L planetUserL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PlanetUserColumns = struct {
	UserID    string
	FirstName string
	LastName  string
}{
	UserID:    "user_id",
	FirstName: "first_name",
	LastName:  "last_name",
}

// Generated where

var PlanetUserWhere = struct {
	UserID    whereHelperstring
	FirstName whereHelperstring
	LastName  whereHelperstring
}{
	UserID:    whereHelperstring{field: "`planet_users`.`user_id`"},
	FirstName: whereHelperstring{field: "`planet_users`.`first_name`"},
	LastName:  whereHelperstring{field: "`planet_users`.`last_name`"},
}

// PlanetUserRels is where relationship names are stored.
var PlanetUserRels = struct {
	GroupPlanetGroups string
}{
	GroupPlanetGroups: "GroupPlanetGroups",
}

// planetUserR is where relationships are stored.
type planetUserR struct {
	GroupPlanetGroups PlanetGroupSlice
}

// NewStruct creates a new relationship struct
func (*planetUserR) NewStruct() *planetUserR {
	return &planetUserR{}
}

// planetUserL is where Load methods for each relationship are stored.
type planetUserL struct{}

var (
	planetUserAllColumns            = []string{"user_id", "first_name", "last_name"}
	planetUserColumnsWithoutDefault = []string{"user_id", "first_name", "last_name"}
	planetUserColumnsWithDefault    = []string{}
	planetUserPrimaryKeyColumns     = []string{"user_id"}
)

type (
	// PlanetUserSlice is an alias for a slice of pointers to PlanetUser.
	// This should generally be used opposed to []PlanetUser.
	PlanetUserSlice []*PlanetUser
	// PlanetUserHook is the signature for custom PlanetUser hook methods
	PlanetUserHook func(context.Context, boil.ContextExecutor, *PlanetUser) error

	planetUserQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	planetUserType                 = reflect.TypeOf(&PlanetUser{})
	planetUserMapping              = queries.MakeStructMapping(planetUserType)
	planetUserPrimaryKeyMapping, _ = queries.BindMapping(planetUserType, planetUserMapping, planetUserPrimaryKeyColumns)
	planetUserInsertCacheMut       sync.RWMutex
	planetUserInsertCache          = make(map[string]insertCache)
	planetUserUpdateCacheMut       sync.RWMutex
	planetUserUpdateCache          = make(map[string]updateCache)
	planetUserUpsertCacheMut       sync.RWMutex
	planetUserUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var planetUserBeforeInsertHooks []PlanetUserHook
var planetUserBeforeUpdateHooks []PlanetUserHook
var planetUserBeforeDeleteHooks []PlanetUserHook
var planetUserBeforeUpsertHooks []PlanetUserHook

var planetUserAfterInsertHooks []PlanetUserHook
var planetUserAfterSelectHooks []PlanetUserHook
var planetUserAfterUpdateHooks []PlanetUserHook
var planetUserAfterDeleteHooks []PlanetUserHook
var planetUserAfterUpsertHooks []PlanetUserHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *PlanetUser) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range planetUserBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *PlanetUser) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range planetUserBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *PlanetUser) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range planetUserBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *PlanetUser) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range planetUserBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *PlanetUser) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range planetUserAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *PlanetUser) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range planetUserAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *PlanetUser) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range planetUserAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *PlanetUser) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range planetUserAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *PlanetUser) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range planetUserAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPlanetUserHook registers your hook function for all future operations.
func AddPlanetUserHook(hookPoint boil.HookPoint, planetUserHook PlanetUserHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		planetUserBeforeInsertHooks = append(planetUserBeforeInsertHooks, planetUserHook)
	case boil.BeforeUpdateHook:
		planetUserBeforeUpdateHooks = append(planetUserBeforeUpdateHooks, planetUserHook)
	case boil.BeforeDeleteHook:
		planetUserBeforeDeleteHooks = append(planetUserBeforeDeleteHooks, planetUserHook)
	case boil.BeforeUpsertHook:
		planetUserBeforeUpsertHooks = append(planetUserBeforeUpsertHooks, planetUserHook)
	case boil.AfterInsertHook:
		planetUserAfterInsertHooks = append(planetUserAfterInsertHooks, planetUserHook)
	case boil.AfterSelectHook:
		planetUserAfterSelectHooks = append(planetUserAfterSelectHooks, planetUserHook)
	case boil.AfterUpdateHook:
		planetUserAfterUpdateHooks = append(planetUserAfterUpdateHooks, planetUserHook)
	case boil.AfterDeleteHook:
		planetUserAfterDeleteHooks = append(planetUserAfterDeleteHooks, planetUserHook)
	case boil.AfterUpsertHook:
		planetUserAfterUpsertHooks = append(planetUserAfterUpsertHooks, planetUserHook)
	}
}

// One returns a single planetUser record from the query.
func (q planetUserQuery) One(ctx context.Context, exec boil.ContextExecutor) (*PlanetUser, error) {
	o := &PlanetUser{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for planet_users")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all PlanetUser records from the query.
func (q planetUserQuery) All(ctx context.Context, exec boil.ContextExecutor) (PlanetUserSlice, error) {
	var o []*PlanetUser

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to PlanetUser slice")
	}

	if len(planetUserAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all PlanetUser records in the query.
func (q planetUserQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count planet_users rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q planetUserQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if planet_users exists")
	}

	return count > 0, nil
}

// GroupPlanetGroups retrieves all the planet_group's PlanetGroups with an executor via group_name column.
func (o *PlanetUser) GroupPlanetGroups(mods ...qm.QueryMod) planetGroupQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.InnerJoin("`plant_group_members` on `planet_groups`.`group_name` = `plant_group_members`.`group_id`"),
		qm.Where("`plant_group_members`.`user_id`=?", o.UserID),
	)

	query := PlanetGroups(queryMods...)
	queries.SetFrom(query.Query, "`planet_groups`")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"`planet_groups`.*"})
	}

	return query
}

// LoadGroupPlanetGroups allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (planetUserL) LoadGroupPlanetGroups(ctx context.Context, e boil.ContextExecutor, singular bool, maybePlanetUser interface{}, mods queries.Applicator) error {
	var slice []*PlanetUser
	var object *PlanetUser

	if singular {
		object = maybePlanetUser.(*PlanetUser)
	} else {
		slice = *maybePlanetUser.(*[]*PlanetUser)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &planetUserR{}
		}
		args = append(args, object.UserID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &planetUserR{}
			}

			for _, a := range args {
				if a == obj.UserID {
					continue Outer
				}
			}

			args = append(args, obj.UserID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.Select("`planet_groups`.*, `a`.`user_id`"),
		qm.From("`planet_groups`"),
		qm.InnerJoin("`plant_group_members` as `a` on `planet_groups`.`group_name` = `a`.`group_id`"),
		qm.WhereIn("`a`.`user_id` in ?", args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load planet_groups")
	}

	var resultSlice []*PlanetGroup

	var localJoinCols []string
	for results.Next() {
		one := new(PlanetGroup)
		var localJoinCol string

		err = results.Scan(&one.GroupName, &localJoinCol)
		if err != nil {
			return errors.Wrap(err, "failed to scan eager loaded results for planet_groups")
		}
		if err = results.Err(); err != nil {
			return errors.Wrap(err, "failed to plebian-bind eager loaded slice planet_groups")
		}

		resultSlice = append(resultSlice, one)
		localJoinCols = append(localJoinCols, localJoinCol)
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on planet_groups")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for planet_groups")
	}

	if len(planetGroupAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.GroupPlanetGroups = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &planetGroupR{}
			}
			foreign.R.UserPlanetUsers = append(foreign.R.UserPlanetUsers, object)
		}
		return nil
	}

	for i, foreign := range resultSlice {
		localJoinCol := localJoinCols[i]
		for _, local := range slice {
			if local.UserID == localJoinCol {
				local.R.GroupPlanetGroups = append(local.R.GroupPlanetGroups, foreign)
				if foreign.R == nil {
					foreign.R = &planetGroupR{}
				}
				foreign.R.UserPlanetUsers = append(foreign.R.UserPlanetUsers, local)
				break
			}
		}
	}

	return nil
}

// AddGroupPlanetGroups adds the given related objects to the existing relationships
// of the planet_user, optionally inserting them as new records.
// Appends related to o.R.GroupPlanetGroups.
// Sets related.R.UserPlanetUsers appropriately.
func (o *PlanetUser) AddGroupPlanetGroups(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*PlanetGroup) error {
	var err error
	for _, rel := range related {
		if insert {
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		}
	}

	for _, rel := range related {
		query := "insert into `plant_group_members` (`user_id`, `group_id`) values (?, ?)"
		values := []interface{}{o.UserID, rel.GroupName}

		if boil.DebugMode {
			fmt.Fprintln(boil.DebugWriter, query)
			fmt.Fprintln(boil.DebugWriter, values)
		}

		_, err = exec.ExecContext(ctx, query, values...)
		if err != nil {
			return errors.Wrap(err, "failed to insert into join table")
		}
	}
	if o.R == nil {
		o.R = &planetUserR{
			GroupPlanetGroups: related,
		}
	} else {
		o.R.GroupPlanetGroups = append(o.R.GroupPlanetGroups, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &planetGroupR{
				UserPlanetUsers: PlanetUserSlice{o},
			}
		} else {
			rel.R.UserPlanetUsers = append(rel.R.UserPlanetUsers, o)
		}
	}
	return nil
}

// SetGroupPlanetGroups removes all previously related items of the
// planet_user replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.UserPlanetUsers's GroupPlanetGroups accordingly.
// Replaces o.R.GroupPlanetGroups with related.
// Sets related.R.UserPlanetUsers's GroupPlanetGroups accordingly.
func (o *PlanetUser) SetGroupPlanetGroups(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*PlanetGroup) error {
	query := "delete from `plant_group_members` where `user_id` = ?"
	values := []interface{}{o.UserID}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	removeGroupPlanetGroupsFromUserPlanetUsersSlice(o, related)
	if o.R != nil {
		o.R.GroupPlanetGroups = nil
	}
	return o.AddGroupPlanetGroups(ctx, exec, insert, related...)
}

// RemoveGroupPlanetGroups relationships from objects passed in.
// Removes related items from R.GroupPlanetGroups (uses pointer comparison, removal does not keep order)
// Sets related.R.UserPlanetUsers.
func (o *PlanetUser) RemoveGroupPlanetGroups(ctx context.Context, exec boil.ContextExecutor, related ...*PlanetGroup) error {
	var err error
	query := fmt.Sprintf(
		"delete from `plant_group_members` where `user_id` = ? and `group_id` in (%s)",
		strmangle.Placeholders(dialect.UseIndexPlaceholders, len(related), 2, 1),
	)
	values := []interface{}{o.UserID}
	for _, rel := range related {
		values = append(values, rel.GroupName)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}
	removeGroupPlanetGroupsFromUserPlanetUsersSlice(o, related)
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.GroupPlanetGroups {
			if rel != ri {
				continue
			}

			ln := len(o.R.GroupPlanetGroups)
			if ln > 1 && i < ln-1 {
				o.R.GroupPlanetGroups[i] = o.R.GroupPlanetGroups[ln-1]
			}
			o.R.GroupPlanetGroups = o.R.GroupPlanetGroups[:ln-1]
			break
		}
	}

	return nil
}

func removeGroupPlanetGroupsFromUserPlanetUsersSlice(o *PlanetUser, related []*PlanetGroup) {
	for _, rel := range related {
		if rel.R == nil {
			continue
		}
		for i, ri := range rel.R.UserPlanetUsers {
			if o.UserID != ri.UserID {
				continue
			}

			ln := len(rel.R.UserPlanetUsers)
			if ln > 1 && i < ln-1 {
				rel.R.UserPlanetUsers[i] = rel.R.UserPlanetUsers[ln-1]
			}
			rel.R.UserPlanetUsers = rel.R.UserPlanetUsers[:ln-1]
			break
		}
	}
}

// PlanetUsers retrieves all the records using an executor.
func PlanetUsers(mods ...qm.QueryMod) planetUserQuery {
	mods = append(mods, qm.From("`planet_users`"))
	return planetUserQuery{NewQuery(mods...)}
}

// FindPlanetUser retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPlanetUser(ctx context.Context, exec boil.ContextExecutor, userID string, selectCols ...string) (*PlanetUser, error) {
	planetUserObj := &PlanetUser{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `planet_users` where `user_id`=?", sel,
	)

	q := queries.Raw(query, userID)

	err := q.Bind(ctx, exec, planetUserObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from planet_users")
	}

	return planetUserObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *PlanetUser) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no planet_users provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(planetUserColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	planetUserInsertCacheMut.RLock()
	cache, cached := planetUserInsertCache[key]
	planetUserInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			planetUserAllColumns,
			planetUserColumnsWithDefault,
			planetUserColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(planetUserType, planetUserMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(planetUserType, planetUserMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `planet_users` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `planet_users` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `planet_users` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, planetUserPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into planet_users")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.UserID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for planet_users")
	}

CacheNoHooks:
	if !cached {
		planetUserInsertCacheMut.Lock()
		planetUserInsertCache[key] = cache
		planetUserInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the PlanetUser.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *PlanetUser) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	planetUserUpdateCacheMut.RLock()
	cache, cached := planetUserUpdateCache[key]
	planetUserUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			planetUserAllColumns,
			planetUserPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update planet_users, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `planet_users` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, planetUserPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(planetUserType, planetUserMapping, append(wl, planetUserPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update planet_users row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for planet_users")
	}

	if !cached {
		planetUserUpdateCacheMut.Lock()
		planetUserUpdateCache[key] = cache
		planetUserUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q planetUserQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for planet_users")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for planet_users")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PlanetUserSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), planetUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `planet_users` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, planetUserPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in planetUser slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all planetUser")
	}
	return rowsAff, nil
}

var mySQLPlanetUserUniqueColumns = []string{
	"user_id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *PlanetUser) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no planet_users provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(planetUserColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLPlanetUserUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	planetUserUpsertCacheMut.RLock()
	cache, cached := planetUserUpsertCache[key]
	planetUserUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			planetUserAllColumns,
			planetUserColumnsWithDefault,
			planetUserColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			planetUserAllColumns,
			planetUserPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert planet_users, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "planet_users", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `planet_users` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(planetUserType, planetUserMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(planetUserType, planetUserMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for planet_users")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(planetUserType, planetUserMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for planet_users")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for planet_users")
	}

CacheNoHooks:
	if !cached {
		planetUserUpsertCacheMut.Lock()
		planetUserUpsertCache[key] = cache
		planetUserUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single PlanetUser record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *PlanetUser) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no PlanetUser provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), planetUserPrimaryKeyMapping)
	sql := "DELETE FROM `planet_users` WHERE `user_id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from planet_users")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for planet_users")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q planetUserQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no planetUserQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from planet_users")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for planet_users")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PlanetUserSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(planetUserBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), planetUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `planet_users` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, planetUserPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from planetUser slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for planet_users")
	}

	if len(planetUserAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *PlanetUser) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPlanetUser(ctx, exec, o.UserID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PlanetUserSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PlanetUserSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), planetUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `planet_users`.* FROM `planet_users` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, planetUserPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PlanetUserSlice")
	}

	*o = slice

	return nil
}

// PlanetUserExists checks if the PlanetUser row exists.
func PlanetUserExists(ctx context.Context, exec boil.ContextExecutor, userID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `planet_users` where `user_id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, userID)
	}

	row := exec.QueryRowContext(ctx, sql, userID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if planet_users exists")
	}

	return exists, nil
}