// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package model

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/strmangle"
	"gopkg.in/volatiletech/null.v6"
)

// Output is an object representing the database table.
type Output struct {
	ID                 uint64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	TransactionID      uint64       `boil:"transaction_id" json:"transaction_id" toml:"transaction_id" yaml:"transaction_id"`
	TransactionHash    string       `boil:"transaction_hash" json:"transaction_hash" toml:"transaction_hash" yaml:"transaction_hash"`
	Value              null.Float64 `boil:"value" json:"value,omitempty" toml:"value" yaml:"value,omitempty"`
	Vout               uint         `boil:"vout" json:"vout" toml:"vout" yaml:"vout"`
	Type               null.String  `boil:"type" json:"type,omitempty" toml:"type" yaml:"type,omitempty"`
	ScriptPubKeyAsm    null.String  `boil:"script_pub_key_asm" json:"script_pub_key_asm,omitempty" toml:"script_pub_key_asm" yaml:"script_pub_key_asm,omitempty"`
	ScriptPubKeyHex    null.String  `boil:"script_pub_key_hex" json:"script_pub_key_hex,omitempty" toml:"script_pub_key_hex" yaml:"script_pub_key_hex,omitempty"`
	RequiredSignatures null.Uint    `boil:"required_signatures" json:"required_signatures,omitempty" toml:"required_signatures" yaml:"required_signatures,omitempty"`
	AddressList        null.String  `boil:"address_list" json:"address_list,omitempty" toml:"address_list" yaml:"address_list,omitempty"`
	IsSpent            bool         `boil:"is_spent" json:"is_spent" toml:"is_spent" yaml:"is_spent"`
	SpentByInputID     null.Uint64  `boil:"spent_by_input_id" json:"spent_by_input_id,omitempty" toml:"spent_by_input_id" yaml:"spent_by_input_id,omitempty"`
	CreatedAt          time.Time    `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	ModifiedAt         time.Time    `boil:"modified_at" json:"modified_at" toml:"modified_at" yaml:"modified_at"`
	ClaimID            null.String  `boil:"claim_id" json:"claim_id,omitempty" toml:"claim_id" yaml:"claim_id,omitempty"`

	R *outputR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L outputL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OutputColumns = struct {
	ID                 string
	TransactionID      string
	TransactionHash    string
	Value              string
	Vout               string
	Type               string
	ScriptPubKeyAsm    string
	ScriptPubKeyHex    string
	RequiredSignatures string
	AddressList        string
	IsSpent            string
	SpentByInputID     string
	CreatedAt          string
	ModifiedAt         string
	ClaimID            string
}{
	ID:                 "id",
	TransactionID:      "transaction_id",
	TransactionHash:    "transaction_hash",
	Value:              "value",
	Vout:               "vout",
	Type:               "type",
	ScriptPubKeyAsm:    "script_pub_key_asm",
	ScriptPubKeyHex:    "script_pub_key_hex",
	RequiredSignatures: "required_signatures",
	AddressList:        "address_list",
	IsSpent:            "is_spent",
	SpentByInputID:     "spent_by_input_id",
	CreatedAt:          "created_at",
	ModifiedAt:         "modified_at",
	ClaimID:            "claim_id",
}

// outputR is where relationships are stored.
type outputR struct {
	Transaction    *Transaction
	AbnormalClaims AbnormalClaimSlice
}

// outputL is where Load methods for each relationship are stored.
type outputL struct{}

var (
	outputColumns               = []string{"id", "transaction_id", "transaction_hash", "value", "vout", "type", "script_pub_key_asm", "script_pub_key_hex", "required_signatures", "address_list", "is_spent", "spent_by_input_id", "created_at", "modified_at", "claim_id"}
	outputColumnsWithoutDefault = []string{"transaction_id", "transaction_hash", "value", "vout", "type", "script_pub_key_asm", "script_pub_key_hex", "required_signatures", "address_list", "spent_by_input_id", "claim_id"}
	outputColumnsWithDefault    = []string{"id", "is_spent", "created_at", "modified_at"}
	outputPrimaryKeyColumns     = []string{"id"}
)

type (
	// OutputSlice is an alias for a slice of pointers to Output.
	// This should generally be used opposed to []Output.
	OutputSlice []*Output

	outputQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	outputType                 = reflect.TypeOf(&Output{})
	outputMapping              = queries.MakeStructMapping(outputType)
	outputPrimaryKeyMapping, _ = queries.BindMapping(outputType, outputMapping, outputPrimaryKeyColumns)
	outputInsertCacheMut       sync.RWMutex
	outputInsertCache          = make(map[string]insertCache)
	outputUpdateCacheMut       sync.RWMutex
	outputUpdateCache          = make(map[string]updateCache)
	outputUpsertCacheMut       sync.RWMutex
	outputUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)

// OneP returns a single output record from the query, and panics on error.
func (q outputQuery) OneP() *Output {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single output record from the query.
func (q outputQuery) One() (*Output, error) {
	o := &Output{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: failed to execute a one query for output")
	}

	return o, nil
}

// AllP returns all Output records from the query, and panics on error.
func (q outputQuery) AllP() OutputSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Output records from the query.
func (q outputQuery) All() (OutputSlice, error) {
	var o []*Output

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "model: failed to assign all query results to Output slice")
	}

	return o, nil
}

// CountP returns the count of all Output records in the query, and panics on error.
func (q outputQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Output records in the query.
func (q outputQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to count output rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q outputQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q outputQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "model: failed to check if output exists")
	}

	return count > 0, nil
}

// TransactionG pointed to by the foreign key.
func (o *Output) TransactionG(mods ...qm.QueryMod) transactionQuery {
	return o.Transaction(boil.GetDB(), mods...)
}

// Transaction pointed to by the foreign key.
func (o *Output) Transaction(exec boil.Executor, mods ...qm.QueryMod) transactionQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.TransactionID),
	}

	queryMods = append(queryMods, mods...)

	query := Transactions(exec, queryMods...)
	queries.SetFrom(query.Query, "`transaction`")

	return query
}

// AbnormalClaimsG retrieves all the abnormal_claim's abnormal claim.
func (o *Output) AbnormalClaimsG(mods ...qm.QueryMod) abnormalClaimQuery {
	return o.AbnormalClaims(boil.GetDB(), mods...)
}

// AbnormalClaims retrieves all the abnormal_claim's abnormal claim with an executor.
func (o *Output) AbnormalClaims(exec boil.Executor, mods ...qm.QueryMod) abnormalClaimQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`abnormal_claim`.`output_id`=?", o.ID),
	)

	query := AbnormalClaims(exec, queryMods...)
	queries.SetFrom(query.Query, "`abnormal_claim`")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"`abnormal_claim`.*"})
	}

	return query
}

// LoadTransaction allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (outputL) LoadTransaction(e boil.Executor, singular bool, maybeOutput interface{}) error {
	var slice []*Output
	var object *Output

	count := 1
	if singular {
		object = maybeOutput.(*Output)
	} else {
		slice = *maybeOutput.(*[]*Output)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &outputR{}
		}
		args[0] = object.TransactionID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &outputR{}
			}
			args[i] = obj.TransactionID
		}
	}

	query := fmt.Sprintf(
		"select * from `transaction` where `id` in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Transaction")
	}
	defer results.Close()

	var resultSlice []*Transaction
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Transaction")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		object.R.Transaction = resultSlice[0]
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.TransactionID == foreign.ID {
				local.R.Transaction = foreign
				break
			}
		}
	}

	return nil
}

// LoadAbnormalClaims allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (outputL) LoadAbnormalClaims(e boil.Executor, singular bool, maybeOutput interface{}) error {
	var slice []*Output
	var object *Output

	count := 1
	if singular {
		object = maybeOutput.(*Output)
	} else {
		slice = *maybeOutput.(*[]*Output)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &outputR{}
		}
		args[0] = object.ID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &outputR{}
			}
			args[i] = obj.ID
		}
	}

	query := fmt.Sprintf(
		"select * from `abnormal_claim` where `output_id` in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load abnormal_claim")
	}
	defer results.Close()

	var resultSlice []*AbnormalClaim
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice abnormal_claim")
	}

	if singular {
		object.R.AbnormalClaims = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.OutputID {
				local.R.AbnormalClaims = append(local.R.AbnormalClaims, foreign)
				break
			}
		}
	}

	return nil
}

// SetTransactionG of the output to the related item.
// Sets o.R.Transaction to related.
// Adds o to related.R.Outputs.
// Uses the global database handle.
func (o *Output) SetTransactionG(insert bool, related *Transaction) error {
	return o.SetTransaction(boil.GetDB(), insert, related)
}

// SetTransactionP of the output to the related item.
// Sets o.R.Transaction to related.
// Adds o to related.R.Outputs.
// Panics on error.
func (o *Output) SetTransactionP(exec boil.Executor, insert bool, related *Transaction) {
	if err := o.SetTransaction(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetTransactionGP of the output to the related item.
// Sets o.R.Transaction to related.
// Adds o to related.R.Outputs.
// Uses the global database handle and panics on error.
func (o *Output) SetTransactionGP(insert bool, related *Transaction) {
	if err := o.SetTransaction(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetTransaction of the output to the related item.
// Sets o.R.Transaction to related.
// Adds o to related.R.Outputs.
func (o *Output) SetTransaction(exec boil.Executor, insert bool, related *Transaction) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `output` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"transaction_id"}),
		strmangle.WhereClause("`", "`", 0, outputPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TransactionID = related.ID

	if o.R == nil {
		o.R = &outputR{
			Transaction: related,
		}
	} else {
		o.R.Transaction = related
	}

	if related.R == nil {
		related.R = &transactionR{
			Outputs: OutputSlice{o},
		}
	} else {
		related.R.Outputs = append(related.R.Outputs, o)
	}

	return nil
}

// AddAbnormalClaimsG adds the given related objects to the existing relationships
// of the output, optionally inserting them as new records.
// Appends related to o.R.AbnormalClaims.
// Sets related.R.Output appropriately.
// Uses the global database handle.
func (o *Output) AddAbnormalClaimsG(insert bool, related ...*AbnormalClaim) error {
	return o.AddAbnormalClaims(boil.GetDB(), insert, related...)
}

// AddAbnormalClaimsP adds the given related objects to the existing relationships
// of the output, optionally inserting them as new records.
// Appends related to o.R.AbnormalClaims.
// Sets related.R.Output appropriately.
// Panics on error.
func (o *Output) AddAbnormalClaimsP(exec boil.Executor, insert bool, related ...*AbnormalClaim) {
	if err := o.AddAbnormalClaims(exec, insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddAbnormalClaimsGP adds the given related objects to the existing relationships
// of the output, optionally inserting them as new records.
// Appends related to o.R.AbnormalClaims.
// Sets related.R.Output appropriately.
// Uses the global database handle and panics on error.
func (o *Output) AddAbnormalClaimsGP(insert bool, related ...*AbnormalClaim) {
	if err := o.AddAbnormalClaims(boil.GetDB(), insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddAbnormalClaims adds the given related objects to the existing relationships
// of the output, optionally inserting them as new records.
// Appends related to o.R.AbnormalClaims.
// Sets related.R.Output appropriately.
func (o *Output) AddAbnormalClaims(exec boil.Executor, insert bool, related ...*AbnormalClaim) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.OutputID = o.ID
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `abnormal_claim` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"output_id"}),
				strmangle.WhereClause("`", "`", 0, abnormalClaimPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.OutputID = o.ID
		}
	}

	if o.R == nil {
		o.R = &outputR{
			AbnormalClaims: related,
		}
	} else {
		o.R.AbnormalClaims = append(o.R.AbnormalClaims, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &abnormalClaimR{
				Output: o,
			}
		} else {
			rel.R.Output = o
		}
	}
	return nil
}

// OutputsG retrieves all records.
func OutputsG(mods ...qm.QueryMod) outputQuery {
	return Outputs(boil.GetDB(), mods...)
}

// Outputs retrieves all the records using an executor.
func Outputs(exec boil.Executor, mods ...qm.QueryMod) outputQuery {
	mods = append(mods, qm.From("`output`"))
	return outputQuery{NewQuery(exec, mods...)}
}

// FindOutputG retrieves a single record by ID.
func FindOutputG(id uint64, selectCols ...string) (*Output, error) {
	return FindOutput(boil.GetDB(), id, selectCols...)
}

// FindOutputGP retrieves a single record by ID, and panics on error.
func FindOutputGP(id uint64, selectCols ...string) *Output {
	retobj, err := FindOutput(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindOutput retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOutput(exec boil.Executor, id uint64, selectCols ...string) (*Output, error) {
	outputObj := &Output{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `output` where `id`=?", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(outputObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: unable to select from output")
	}

	return outputObj, nil
}

// FindOutputP retrieves a single record by ID with an executor, and panics on error.
func FindOutputP(exec boil.Executor, id uint64, selectCols ...string) *Output {
	retobj, err := FindOutput(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Output) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Output) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Output) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Output) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("model: no output provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(outputColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	outputInsertCacheMut.RLock()
	cache, cached := outputInsertCache[key]
	outputInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			outputColumns,
			outputColumnsWithDefault,
			outputColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(outputType, outputMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(outputType, outputMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `output` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `output` () VALUES ()"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `output` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, outputPrimaryKeyColumns))
		}

		if len(wl) != 0 {
			cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "model: unable to insert into output")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == outputMapping["ID"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for output")
	}

CacheNoHooks:
	if !cached {
		outputInsertCacheMut.Lock()
		outputInsertCache[key] = cache
		outputInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single Output record. See Update for
// whitelist behavior description.
func (o *Output) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Output record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Output) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Output, and panics on error.
// See Update for whitelist behavior description.
func (o *Output) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Output.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Output) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	key := makeCacheKey(whitelist, nil)
	outputUpdateCacheMut.RLock()
	cache, cached := outputUpdateCache[key]
	outputUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			outputColumns,
			outputPrimaryKeyColumns,
			whitelist,
		)

		if len(wl) == 0 {
			return errors.New("model: unable to update output, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `output` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, outputPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(outputType, outputMapping, append(wl, outputPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "model: unable to update output row")
	}

	if !cached {
		outputUpdateCacheMut.Lock()
		outputUpdateCache[key] = cache
		outputUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q outputQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q outputQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "model: unable to update all for output")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o OutputSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o OutputSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o OutputSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OutputSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("model: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), outputPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `output` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, outputPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "model: unable to update all in output slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Output) UpsertG(updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Output) UpsertGP(updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Output) UpsertP(exec boil.Executor, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Output) Upsert(exec boil.Executor, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("model: no output provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(outputColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
	buf := strmangle.GetBuffer()
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	outputUpsertCacheMut.RLock()
	cache, cached := outputUpsertCache[key]
	outputUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			outputColumns,
			outputColumnsWithDefault,
			outputColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			outputColumns,
			outputPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("model: unable to upsert output, could not build update column list")
		}

		cache.query = queries.BuildUpsertQueryMySQL(dialect, "output", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `output` WHERE `id`=?",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
		)

		cache.valueMapping, err = queries.BindMapping(outputType, outputMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(outputType, outputMapping, ret)
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

	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "model: unable to upsert for output")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == outputMapping["ID"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for output")
	}

CacheNoHooks:
	if !cached {
		outputUpsertCacheMut.Lock()
		outputUpsertCache[key] = cache
		outputUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteP deletes a single Output record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Output) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Output record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Output) DeleteG() error {
	if o == nil {
		return errors.New("model: no Output provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Output record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Output) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Output record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Output) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("model: no Output provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), outputPrimaryKeyMapping)
	sql := "DELETE FROM `output` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "model: unable to delete from output")
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q outputQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q outputQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("model: no outputQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "model: unable to delete all from output")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o OutputSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o OutputSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("model: no Output slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o OutputSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OutputSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("model: no Output slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), outputPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `output` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, outputPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "model: unable to delete all from output slice")
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Output) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Output) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Output) ReloadG() error {
	if o == nil {
		return errors.New("model: no Output provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Output) Reload(exec boil.Executor) error {
	ret, err := FindOutput(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *OutputSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *OutputSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OutputSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("model: empty OutputSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OutputSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	outputs := OutputSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), outputPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `output`.* FROM `output` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, outputPrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&outputs)
	if err != nil {
		return errors.Wrap(err, "model: unable to reload all in OutputSlice")
	}

	*o = outputs

	return nil
}

// OutputExists checks if the Output row exists.
func OutputExists(exec boil.Executor, id uint64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `output` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "model: unable to check if output exists")
	}

	return exists, nil
}

// OutputExistsG checks if the Output row exists.
func OutputExistsG(id uint64) (bool, error) {
	return OutputExists(boil.GetDB(), id)
}

// OutputExistsGP checks if the Output row exists. Panics on error.
func OutputExistsGP(id uint64) bool {
	e, err := OutputExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// OutputExistsP checks if the Output row exists. Panics on error.
func OutputExistsP(exec boil.Executor, id uint64) bool {
	e, err := OutputExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
