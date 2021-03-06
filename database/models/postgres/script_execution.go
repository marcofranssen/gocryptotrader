// Code generated by SQLBoiler 3.5.0-gct (https://github.com/thrasher-corp/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package postgres

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
	"github.com/thrasher-corp/sqlboiler/boil"
	"github.com/thrasher-corp/sqlboiler/queries"
	"github.com/thrasher-corp/sqlboiler/queries/qm"
	"github.com/thrasher-corp/sqlboiler/queries/qmhelper"
	"github.com/thrasher-corp/sqlboiler/strmangle"
	"github.com/volatiletech/null"
)

// ScriptExecution is an object representing the database table.
type ScriptExecution struct {
	ID              string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	ScriptID        null.String `boil:"script_id" json:"script_id,omitempty" toml:"script_id" yaml:"script_id,omitempty"`
	ExecutionType   string      `boil:"execution_type" json:"execution_type" toml:"execution_type" yaml:"execution_type"`
	ExecutionStatus string      `boil:"execution_status" json:"execution_status" toml:"execution_status" yaml:"execution_status"`
	ExecutionTime   time.Time   `boil:"execution_time" json:"execution_time" toml:"execution_time" yaml:"execution_time"`

	R *scriptExecutionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L scriptExecutionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ScriptExecutionColumns = struct {
	ID              string
	ScriptID        string
	ExecutionType   string
	ExecutionStatus string
	ExecutionTime   string
}{
	ID:              "id",
	ScriptID:        "script_id",
	ExecutionType:   "execution_type",
	ExecutionStatus: "execution_status",
	ExecutionTime:   "execution_time",
}

// Generated where

var ScriptExecutionWhere = struct {
	ID              whereHelperstring
	ScriptID        whereHelpernull_String
	ExecutionType   whereHelperstring
	ExecutionStatus whereHelperstring
	ExecutionTime   whereHelpertime_Time
}{
	ID:              whereHelperstring{field: "\"script_execution\".\"id\""},
	ScriptID:        whereHelpernull_String{field: "\"script_execution\".\"script_id\""},
	ExecutionType:   whereHelperstring{field: "\"script_execution\".\"execution_type\""},
	ExecutionStatus: whereHelperstring{field: "\"script_execution\".\"execution_status\""},
	ExecutionTime:   whereHelpertime_Time{field: "\"script_execution\".\"execution_time\""},
}

// ScriptExecutionRels is where relationship names are stored.
var ScriptExecutionRels = struct {
	Script string
}{
	Script: "Script",
}

// scriptExecutionR is where relationships are stored.
type scriptExecutionR struct {
	Script *Script
}

// NewStruct creates a new relationship struct
func (*scriptExecutionR) NewStruct() *scriptExecutionR {
	return &scriptExecutionR{}
}

// scriptExecutionL is where Load methods for each relationship are stored.
type scriptExecutionL struct{}

var (
	scriptExecutionAllColumns            = []string{"id", "script_id", "execution_type", "execution_status", "execution_time"}
	scriptExecutionColumnsWithoutDefault = []string{"script_id", "execution_type", "execution_status"}
	scriptExecutionColumnsWithDefault    = []string{"id", "execution_time"}
	scriptExecutionPrimaryKeyColumns     = []string{"id"}
)

type (
	// ScriptExecutionSlice is an alias for a slice of pointers to ScriptExecution.
	// This should generally be used opposed to []ScriptExecution.
	ScriptExecutionSlice []*ScriptExecution
	// ScriptExecutionHook is the signature for custom ScriptExecution hook methods
	ScriptExecutionHook func(context.Context, boil.ContextExecutor, *ScriptExecution) error

	scriptExecutionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	scriptExecutionType                 = reflect.TypeOf(&ScriptExecution{})
	scriptExecutionMapping              = queries.MakeStructMapping(scriptExecutionType)
	scriptExecutionPrimaryKeyMapping, _ = queries.BindMapping(scriptExecutionType, scriptExecutionMapping, scriptExecutionPrimaryKeyColumns)
	scriptExecutionInsertCacheMut       sync.RWMutex
	scriptExecutionInsertCache          = make(map[string]insertCache)
	scriptExecutionUpdateCacheMut       sync.RWMutex
	scriptExecutionUpdateCache          = make(map[string]updateCache)
	scriptExecutionUpsertCacheMut       sync.RWMutex
	scriptExecutionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var scriptExecutionBeforeInsertHooks []ScriptExecutionHook
var scriptExecutionBeforeUpdateHooks []ScriptExecutionHook
var scriptExecutionBeforeDeleteHooks []ScriptExecutionHook
var scriptExecutionBeforeUpsertHooks []ScriptExecutionHook

var scriptExecutionAfterInsertHooks []ScriptExecutionHook
var scriptExecutionAfterSelectHooks []ScriptExecutionHook
var scriptExecutionAfterUpdateHooks []ScriptExecutionHook
var scriptExecutionAfterDeleteHooks []ScriptExecutionHook
var scriptExecutionAfterUpsertHooks []ScriptExecutionHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ScriptExecution) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range scriptExecutionBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ScriptExecution) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range scriptExecutionBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ScriptExecution) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range scriptExecutionBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ScriptExecution) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range scriptExecutionBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ScriptExecution) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range scriptExecutionAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ScriptExecution) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range scriptExecutionAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ScriptExecution) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range scriptExecutionAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ScriptExecution) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range scriptExecutionAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ScriptExecution) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range scriptExecutionAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddScriptExecutionHook registers your hook function for all future operations.
func AddScriptExecutionHook(hookPoint boil.HookPoint, scriptExecutionHook ScriptExecutionHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		scriptExecutionBeforeInsertHooks = append(scriptExecutionBeforeInsertHooks, scriptExecutionHook)
	case boil.BeforeUpdateHook:
		scriptExecutionBeforeUpdateHooks = append(scriptExecutionBeforeUpdateHooks, scriptExecutionHook)
	case boil.BeforeDeleteHook:
		scriptExecutionBeforeDeleteHooks = append(scriptExecutionBeforeDeleteHooks, scriptExecutionHook)
	case boil.BeforeUpsertHook:
		scriptExecutionBeforeUpsertHooks = append(scriptExecutionBeforeUpsertHooks, scriptExecutionHook)
	case boil.AfterInsertHook:
		scriptExecutionAfterInsertHooks = append(scriptExecutionAfterInsertHooks, scriptExecutionHook)
	case boil.AfterSelectHook:
		scriptExecutionAfterSelectHooks = append(scriptExecutionAfterSelectHooks, scriptExecutionHook)
	case boil.AfterUpdateHook:
		scriptExecutionAfterUpdateHooks = append(scriptExecutionAfterUpdateHooks, scriptExecutionHook)
	case boil.AfterDeleteHook:
		scriptExecutionAfterDeleteHooks = append(scriptExecutionAfterDeleteHooks, scriptExecutionHook)
	case boil.AfterUpsertHook:
		scriptExecutionAfterUpsertHooks = append(scriptExecutionAfterUpsertHooks, scriptExecutionHook)
	}
}

// One returns a single scriptExecution record from the query.
func (q scriptExecutionQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ScriptExecution, error) {
	o := &ScriptExecution{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "postgres: failed to execute a one query for script_execution")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ScriptExecution records from the query.
func (q scriptExecutionQuery) All(ctx context.Context, exec boil.ContextExecutor) (ScriptExecutionSlice, error) {
	var o []*ScriptExecution

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "postgres: failed to assign all query results to ScriptExecution slice")
	}

	if len(scriptExecutionAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ScriptExecution records in the query.
func (q scriptExecutionQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "postgres: failed to count script_execution rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q scriptExecutionQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "postgres: failed to check if script_execution exists")
	}

	return count > 0, nil
}

// Script pointed to by the foreign key.
func (o *ScriptExecution) Script(mods ...qm.QueryMod) scriptQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ScriptID),
	}

	queryMods = append(queryMods, mods...)

	query := Scripts(queryMods...)
	queries.SetFrom(query.Query, "\"script\"")

	return query
}

// LoadScript allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (scriptExecutionL) LoadScript(ctx context.Context, e boil.ContextExecutor, singular bool, maybeScriptExecution interface{}, mods queries.Applicator) error {
	var slice []*ScriptExecution
	var object *ScriptExecution

	if singular {
		object = maybeScriptExecution.(*ScriptExecution)
	} else {
		slice = *maybeScriptExecution.(*[]*ScriptExecution)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &scriptExecutionR{}
		}
		if !queries.IsNil(object.ScriptID) {
			args = append(args, object.ScriptID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &scriptExecutionR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ScriptID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.ScriptID) {
				args = append(args, obj.ScriptID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`script`), qm.WhereIn(`script.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Script")
	}

	var resultSlice []*Script
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Script")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for script")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for script")
	}

	if len(scriptExecutionAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Script = foreign
		if foreign.R == nil {
			foreign.R = &scriptR{}
		}
		foreign.R.ScriptExecutions = append(foreign.R.ScriptExecutions, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.ScriptID, foreign.ID) {
				local.R.Script = foreign
				if foreign.R == nil {
					foreign.R = &scriptR{}
				}
				foreign.R.ScriptExecutions = append(foreign.R.ScriptExecutions, local)
				break
			}
		}
	}

	return nil
}

// SetScript of the scriptExecution to the related item.
// Sets o.R.Script to related.
// Adds o to related.R.ScriptExecutions.
func (o *ScriptExecution) SetScript(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Script) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"script_execution\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"script_id"}),
		strmangle.WhereClause("\"", "\"", 2, scriptExecutionPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.ScriptID, related.ID)
	if o.R == nil {
		o.R = &scriptExecutionR{
			Script: related,
		}
	} else {
		o.R.Script = related
	}

	if related.R == nil {
		related.R = &scriptR{
			ScriptExecutions: ScriptExecutionSlice{o},
		}
	} else {
		related.R.ScriptExecutions = append(related.R.ScriptExecutions, o)
	}

	return nil
}

// RemoveScript relationship.
// Sets o.R.Script to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *ScriptExecution) RemoveScript(ctx context.Context, exec boil.ContextExecutor, related *Script) error {
	var err error

	queries.SetScanner(&o.ScriptID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("script_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Script = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.ScriptExecutions {
		if queries.Equal(o.ScriptID, ri.ScriptID) {
			continue
		}

		ln := len(related.R.ScriptExecutions)
		if ln > 1 && i < ln-1 {
			related.R.ScriptExecutions[i] = related.R.ScriptExecutions[ln-1]
		}
		related.R.ScriptExecutions = related.R.ScriptExecutions[:ln-1]
		break
	}
	return nil
}

// ScriptExecutions retrieves all the records using an executor.
func ScriptExecutions(mods ...qm.QueryMod) scriptExecutionQuery {
	mods = append(mods, qm.From("\"script_execution\""))
	return scriptExecutionQuery{NewQuery(mods...)}
}

// FindScriptExecution retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindScriptExecution(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*ScriptExecution, error) {
	scriptExecutionObj := &ScriptExecution{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"script_execution\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, scriptExecutionObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "postgres: unable to select from script_execution")
	}

	return scriptExecutionObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ScriptExecution) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("postgres: no script_execution provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(scriptExecutionColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	scriptExecutionInsertCacheMut.RLock()
	cache, cached := scriptExecutionInsertCache[key]
	scriptExecutionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			scriptExecutionAllColumns,
			scriptExecutionColumnsWithDefault,
			scriptExecutionColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(scriptExecutionType, scriptExecutionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(scriptExecutionType, scriptExecutionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"script_execution\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"script_execution\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "postgres: unable to insert into script_execution")
	}

	if !cached {
		scriptExecutionInsertCacheMut.Lock()
		scriptExecutionInsertCache[key] = cache
		scriptExecutionInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ScriptExecution.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ScriptExecution) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	scriptExecutionUpdateCacheMut.RLock()
	cache, cached := scriptExecutionUpdateCache[key]
	scriptExecutionUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			scriptExecutionAllColumns,
			scriptExecutionPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("postgres: unable to update script_execution, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"script_execution\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, scriptExecutionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(scriptExecutionType, scriptExecutionMapping, append(wl, scriptExecutionPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "postgres: unable to update script_execution row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "postgres: failed to get rows affected by update for script_execution")
	}

	if !cached {
		scriptExecutionUpdateCacheMut.Lock()
		scriptExecutionUpdateCache[key] = cache
		scriptExecutionUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q scriptExecutionQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "postgres: unable to update all for script_execution")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "postgres: unable to retrieve rows affected for script_execution")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ScriptExecutionSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("postgres: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), scriptExecutionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"script_execution\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, scriptExecutionPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "postgres: unable to update all in scriptExecution slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "postgres: unable to retrieve rows affected all in update all scriptExecution")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ScriptExecution) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("postgres: no script_execution provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(scriptExecutionColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
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
	key := buf.String()
	strmangle.PutBuffer(buf)

	scriptExecutionUpsertCacheMut.RLock()
	cache, cached := scriptExecutionUpsertCache[key]
	scriptExecutionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			scriptExecutionAllColumns,
			scriptExecutionColumnsWithDefault,
			scriptExecutionColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			scriptExecutionAllColumns,
			scriptExecutionPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("postgres: unable to upsert script_execution, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(scriptExecutionPrimaryKeyColumns))
			copy(conflict, scriptExecutionPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"script_execution\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(scriptExecutionType, scriptExecutionMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(scriptExecutionType, scriptExecutionMapping, ret)
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

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "postgres: unable to upsert script_execution")
	}

	if !cached {
		scriptExecutionUpsertCacheMut.Lock()
		scriptExecutionUpsertCache[key] = cache
		scriptExecutionUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ScriptExecution record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ScriptExecution) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("postgres: no ScriptExecution provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), scriptExecutionPrimaryKeyMapping)
	sql := "DELETE FROM \"script_execution\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "postgres: unable to delete from script_execution")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "postgres: failed to get rows affected by delete for script_execution")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q scriptExecutionQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("postgres: no scriptExecutionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "postgres: unable to delete all from script_execution")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "postgres: failed to get rows affected by deleteall for script_execution")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ScriptExecutionSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(scriptExecutionBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), scriptExecutionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"script_execution\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, scriptExecutionPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "postgres: unable to delete all from scriptExecution slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "postgres: failed to get rows affected by deleteall for script_execution")
	}

	if len(scriptExecutionAfterDeleteHooks) != 0 {
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
func (o *ScriptExecution) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindScriptExecution(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ScriptExecutionSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ScriptExecutionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), scriptExecutionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"script_execution\".* FROM \"script_execution\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, scriptExecutionPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "postgres: unable to reload all in ScriptExecutionSlice")
	}

	*o = slice

	return nil
}

// ScriptExecutionExists checks if the ScriptExecution row exists.
func ScriptExecutionExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"script_execution\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "postgres: unable to check if script_execution exists")
	}

	return exists, nil
}
