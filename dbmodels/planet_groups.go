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

// PlanetGroup is an object representing the database table.
type PlanetGroup struct {
	GroupName string `boil:"group_name" json:"group_name" toml:"group_name" yaml:"group_name"`

	R *planetGroupR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L planetGroupL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PlanetGroupColumns = struct {
	GroupName string
}{
	GroupName: "group_name",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var PlanetGroupWhere = struct {
	GroupName whereHelperstring
}{
	GroupName: whereHelperstring{field: "`planet_groups`.`group_name`"},
}

// PlanetGroupRels is where relationship names are stored.
var PlanetGroupRels = struct {
	UserPlanetUsers string
}{
	UserPlanetUsers: "UserPlanetUsers",
}

// planetGroupR is where relationships are stored.
type planetGroupR struct {
	UserPlanetUsers PlanetUserSlice
}

// NewStruct creates a new relationship struct
func (*planetGroupR) NewStruct() *planetGroupR {
	return &planetGroupR{}
}

// planetGroupL is where Load methods for each relationship are stored.
type planetGroupL struct{}

var (
	planetGroupAllColumns            = []string{"group_name"}
	planetGroupColumnsWithoutDefault = []string{"group_name"}
	planetGroupColumnsWithDefault    = []string{}
	planetGroupPrimaryKeyColumns     = []string{"group_name"}
)

type (
	// PlanetGroupSlice is an alias for a slice of pointers to PlanetGroup.
	// This should generally be used opposed to []PlanetGroup.
	PlanetGroupSlice []*PlanetGroup
	// PlanetGroupHook is the signature for custom PlanetGroup hook methods
	PlanetGroupHook func(context.Context, boil.ContextExecutor, *PlanetGroup) error

	planetGroupQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	planetGroupType                 = reflect.TypeOf(&PlanetGroup{})
	planetGroupMapping              = queries.MakeStructMapping(planetGroupType)
	planetGroupPrimaryKeyMapping, _ = queries.BindMapping(planetGroupType, planetGroupMapping, planetGroupPrimaryKeyColumns)
	planetGroupInsertCacheMut       sync.RWMutex
	planetGroupInsertCache          = make(map[string]insertCache)
	planetGroupUpdateCacheMut       sync.RWMutex
	planetGroupUpdateCache          = make(map[string]updateCache)
	planetGroupUpsertCacheMut       sync.RWMutex
	planetGroupUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var planetGroupBeforeInsertHooks []PlanetGroupHook
var planetGroupBeforeUpdateHooks []PlanetGroupHook
var planetGroupBeforeDeleteHooks []PlanetGroupHook
var planetGroupBeforeUpsertHooks []PlanetGroupHook

var planetGroupAfterInsertHooks []PlanetGroupHook
var planetGroupAfterSelectHooks []PlanetGroupHook
var planetGroupAfterUpdateHooks []PlanetGroupHook
var planetGroupAfterDeleteHooks []PlanetGroupHook
var planetGroupAfterUpsertHooks []PlanetGroupHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *PlanetGroup) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range planetGroupBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *PlanetGroup) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range planetGroupBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *PlanetGroup) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range planetGroupBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *PlanetGroup) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range planetGroupBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *PlanetGroup) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range planetGroupAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *PlanetGroup) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range planetGroupAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *PlanetGroup) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range planetGroupAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *PlanetGroup) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range planetGroupAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *PlanetGroup) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range planetGroupAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPlanetGroupHook registers your hook function for all future operations.
func AddPlanetGroupHook(hookPoint boil.HookPoint, planetGroupHook PlanetGroupHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		planetGroupBeforeInsertHooks = append(planetGroupBeforeInsertHooks, planetGroupHook)
	case boil.BeforeUpdateHook:
		planetGroupBeforeUpdateHooks = append(planetGroupBeforeUpdateHooks, planetGroupHook)
	case boil.BeforeDeleteHook:
		planetGroupBeforeDeleteHooks = append(planetGroupBeforeDeleteHooks, planetGroupHook)
	case boil.BeforeUpsertHook:
		planetGroupBeforeUpsertHooks = append(planetGroupBeforeUpsertHooks, planetGroupHook)
	case boil.AfterInsertHook:
		planetGroupAfterInsertHooks = append(planetGroupAfterInsertHooks, planetGroupHook)
	case boil.AfterSelectHook:
		planetGroupAfterSelectHooks = append(planetGroupAfterSelectHooks, planetGroupHook)
	case boil.AfterUpdateHook:
		planetGroupAfterUpdateHooks = append(planetGroupAfterUpdateHooks, planetGroupHook)
	case boil.AfterDeleteHook:
		planetGroupAfterDeleteHooks = append(planetGroupAfterDeleteHooks, planetGroupHook)
	case boil.AfterUpsertHook:
		planetGroupAfterUpsertHooks = append(planetGroupAfterUpsertHooks, planetGroupHook)
	}
}

// One returns a single planetGroup record from the query.
func (q planetGroupQuery) One(ctx context.Context, exec boil.ContextExecutor) (*PlanetGroup, error) {
	o := &PlanetGroup{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for planet_groups")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all PlanetGroup records from the query.
func (q planetGroupQuery) All(ctx context.Context, exec boil.ContextExecutor) (PlanetGroupSlice, error) {
	var o []*PlanetGroup

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to PlanetGroup slice")
	}

	if len(planetGroupAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all PlanetGroup records in the query.
func (q planetGroupQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count planet_groups rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q planetGroupQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if planet_groups exists")
	}

	return count > 0, nil
}

// UserPlanetUsers retrieves all the planet_user's PlanetUsers with an executor via user_id column.
func (o *PlanetGroup) UserPlanetUsers(mods ...qm.QueryMod) planetUserQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.InnerJoin("`plant_group_members` on `planet_users`.`user_id` = `plant_group_members`.`user_id`"),
		qm.Where("`plant_group_members`.`group_id`=?", o.GroupName),
	)

	query := PlanetUsers(queryMods...)
	queries.SetFrom(query.Query, "`planet_users`")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"`planet_users`.*"})
	}

	return query
}

// LoadUserPlanetUsers allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (planetGroupL) LoadUserPlanetUsers(ctx context.Context, e boil.ContextExecutor, singular bool, maybePlanetGroup interface{}, mods queries.Applicator) error {
	var slice []*PlanetGroup
	var object *PlanetGroup

	if singular {
		object = maybePlanetGroup.(*PlanetGroup)
	} else {
		slice = *maybePlanetGroup.(*[]*PlanetGroup)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &planetGroupR{}
		}
		args = append(args, object.GroupName)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &planetGroupR{}
			}

			for _, a := range args {
				if a == obj.GroupName {
					continue Outer
				}
			}

			args = append(args, obj.GroupName)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.Select("`planet_users`.*, `a`.`group_id`"),
		qm.From("`planet_users`"),
		qm.InnerJoin("`plant_group_members` as `a` on `planet_users`.`user_id` = `a`.`user_id`"),
		qm.WhereIn("`a`.`group_id` in ?", args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load planet_users")
	}

	var resultSlice []*PlanetUser

	var localJoinCols []string
	for results.Next() {
		one := new(PlanetUser)
		var localJoinCol string

		err = results.Scan(&one.UserID, &one.FirstName, &one.LastName, &localJoinCol)
		if err != nil {
			return errors.Wrap(err, "failed to scan eager loaded results for planet_users")
		}
		if err = results.Err(); err != nil {
			return errors.Wrap(err, "failed to plebian-bind eager loaded slice planet_users")
		}

		resultSlice = append(resultSlice, one)
		localJoinCols = append(localJoinCols, localJoinCol)
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on planet_users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for planet_users")
	}

	if len(planetUserAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.UserPlanetUsers = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &planetUserR{}
			}
			foreign.R.GroupPlanetGroups = append(foreign.R.GroupPlanetGroups, object)
		}
		return nil
	}

	for i, foreign := range resultSlice {
		localJoinCol := localJoinCols[i]
		for _, local := range slice {
			if local.GroupName == localJoinCol {
				local.R.UserPlanetUsers = append(local.R.UserPlanetUsers, foreign)
				if foreign.R == nil {
					foreign.R = &planetUserR{}
				}
				foreign.R.GroupPlanetGroups = append(foreign.R.GroupPlanetGroups, local)
				break
			}
		}
	}

	return nil
}

// AddUserPlanetUsers adds the given related objects to the existing relationships
// of the planet_group, optionally inserting them as new records.
// Appends related to o.R.UserPlanetUsers.
// Sets related.R.GroupPlanetGroups appropriately.
func (o *PlanetGroup) AddUserPlanetUsers(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*PlanetUser) error {
	var err error
	for _, rel := range related {
		if insert {
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		}
	}

	for _, rel := range related {
		query := "insert into `plant_group_members` (`group_id`, `user_id`) values (?, ?)"
		values := []interface{}{o.GroupName, rel.UserID}

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
		o.R = &planetGroupR{
			UserPlanetUsers: related,
		}
	} else {
		o.R.UserPlanetUsers = append(o.R.UserPlanetUsers, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &planetUserR{
				GroupPlanetGroups: PlanetGroupSlice{o},
			}
		} else {
			rel.R.GroupPlanetGroups = append(rel.R.GroupPlanetGroups, o)
		}
	}
	return nil
}

// SetUserPlanetUsers removes all previously related items of the
// planet_group replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.GroupPlanetGroups's UserPlanetUsers accordingly.
// Replaces o.R.UserPlanetUsers with related.
// Sets related.R.GroupPlanetGroups's UserPlanetUsers accordingly.
func (o *PlanetGroup) SetUserPlanetUsers(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*PlanetUser) error {
	query := "delete from `plant_group_members` where `group_id` = ?"
	values := []interface{}{o.GroupName}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	removeUserPlanetUsersFromGroupPlanetGroupsSlice(o, related)
	if o.R != nil {
		o.R.UserPlanetUsers = nil
	}
	return o.AddUserPlanetUsers(ctx, exec, insert, related...)
}

// RemoveUserPlanetUsers relationships from objects passed in.
// Removes related items from R.UserPlanetUsers (uses pointer comparison, removal does not keep order)
// Sets related.R.GroupPlanetGroups.
func (o *PlanetGroup) RemoveUserPlanetUsers(ctx context.Context, exec boil.ContextExecutor, related ...*PlanetUser) error {
	var err error
	query := fmt.Sprintf(
		"delete from `plant_group_members` where `group_id` = ? and `user_id` in (%s)",
		strmangle.Placeholders(dialect.UseIndexPlaceholders, len(related), 2, 1),
	)
	values := []interface{}{o.GroupName}
	for _, rel := range related {
		values = append(values, rel.UserID)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}
	removeUserPlanetUsersFromGroupPlanetGroupsSlice(o, related)
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.UserPlanetUsers {
			if rel != ri {
				continue
			}

			ln := len(o.R.UserPlanetUsers)
			if ln > 1 && i < ln-1 {
				o.R.UserPlanetUsers[i] = o.R.UserPlanetUsers[ln-1]
			}
			o.R.UserPlanetUsers = o.R.UserPlanetUsers[:ln-1]
			break
		}
	}

	return nil
}

func removeUserPlanetUsersFromGroupPlanetGroupsSlice(o *PlanetGroup, related []*PlanetUser) {
	for _, rel := range related {
		if rel.R == nil {
			continue
		}
		for i, ri := range rel.R.GroupPlanetGroups {
			if o.GroupName != ri.GroupName {
				continue
			}

			ln := len(rel.R.GroupPlanetGroups)
			if ln > 1 && i < ln-1 {
				rel.R.GroupPlanetGroups[i] = rel.R.GroupPlanetGroups[ln-1]
			}
			rel.R.GroupPlanetGroups = rel.R.GroupPlanetGroups[:ln-1]
			break
		}
	}
}

// PlanetGroups retrieves all the records using an executor.
func PlanetGroups(mods ...qm.QueryMod) planetGroupQuery {
	mods = append(mods, qm.From("`planet_groups`"))
	return planetGroupQuery{NewQuery(mods...)}
}

// FindPlanetGroup retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPlanetGroup(ctx context.Context, exec boil.ContextExecutor, groupName string, selectCols ...string) (*PlanetGroup, error) {
	planetGroupObj := &PlanetGroup{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `planet_groups` where `group_name`=?", sel,
	)

	q := queries.Raw(query, groupName)

	err := q.Bind(ctx, exec, planetGroupObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from planet_groups")
	}

	return planetGroupObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *PlanetGroup) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no planet_groups provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(planetGroupColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	planetGroupInsertCacheMut.RLock()
	cache, cached := planetGroupInsertCache[key]
	planetGroupInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			planetGroupAllColumns,
			planetGroupColumnsWithDefault,
			planetGroupColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(planetGroupType, planetGroupMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(planetGroupType, planetGroupMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `planet_groups` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `planet_groups` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `planet_groups` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, planetGroupPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into planet_groups")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.GroupName,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for planet_groups")
	}

CacheNoHooks:
	if !cached {
		planetGroupInsertCacheMut.Lock()
		planetGroupInsertCache[key] = cache
		planetGroupInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the PlanetGroup.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *PlanetGroup) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	planetGroupUpdateCacheMut.RLock()
	cache, cached := planetGroupUpdateCache[key]
	planetGroupUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			planetGroupAllColumns,
			planetGroupPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update planet_groups, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `planet_groups` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, planetGroupPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(planetGroupType, planetGroupMapping, append(wl, planetGroupPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update planet_groups row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for planet_groups")
	}

	if !cached {
		planetGroupUpdateCacheMut.Lock()
		planetGroupUpdateCache[key] = cache
		planetGroupUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q planetGroupQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for planet_groups")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for planet_groups")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PlanetGroupSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), planetGroupPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `planet_groups` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, planetGroupPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in planetGroup slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all planetGroup")
	}
	return rowsAff, nil
}

var mySQLPlanetGroupUniqueColumns = []string{
	"group_name",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *PlanetGroup) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no planet_groups provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(planetGroupColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLPlanetGroupUniqueColumns, o)

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

	planetGroupUpsertCacheMut.RLock()
	cache, cached := planetGroupUpsertCache[key]
	planetGroupUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			planetGroupAllColumns,
			planetGroupColumnsWithDefault,
			planetGroupColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			planetGroupAllColumns,
			planetGroupPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert planet_groups, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "planet_groups", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `planet_groups` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(planetGroupType, planetGroupMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(planetGroupType, planetGroupMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for planet_groups")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(planetGroupType, planetGroupMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for planet_groups")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for planet_groups")
	}

CacheNoHooks:
	if !cached {
		planetGroupUpsertCacheMut.Lock()
		planetGroupUpsertCache[key] = cache
		planetGroupUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single PlanetGroup record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *PlanetGroup) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no PlanetGroup provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), planetGroupPrimaryKeyMapping)
	sql := "DELETE FROM `planet_groups` WHERE `group_name`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from planet_groups")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for planet_groups")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q planetGroupQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no planetGroupQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from planet_groups")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for planet_groups")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PlanetGroupSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(planetGroupBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), planetGroupPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `planet_groups` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, planetGroupPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from planetGroup slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for planet_groups")
	}

	if len(planetGroupAfterDeleteHooks) != 0 {
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
func (o *PlanetGroup) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPlanetGroup(ctx, exec, o.GroupName)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PlanetGroupSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PlanetGroupSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), planetGroupPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `planet_groups`.* FROM `planet_groups` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, planetGroupPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PlanetGroupSlice")
	}

	*o = slice

	return nil
}

// PlanetGroupExists checks if the PlanetGroup row exists.
func PlanetGroupExists(ctx context.Context, exec boil.ContextExecutor, groupName string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `planet_groups` where `group_name`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, groupName)
	}

	row := exec.QueryRowContext(ctx, sql, groupName)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if planet_groups exists")
	}

	return exists, nil
}