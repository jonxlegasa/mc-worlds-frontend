// Code generated by ent, DO NOT EDIT.

package emojis

import (
	"entgo.io/ent/dialect/sql"
	"github.com/mikestefanello/pagoda/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Emojis {
	return predicate.Emojis(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Emojis {
	return predicate.Emojis(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Emojis {
	return predicate.Emojis(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Emojis {
	return predicate.Emojis(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Emojis {
	return predicate.Emojis(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Emojis {
	return predicate.Emojis(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Emojis {
	return predicate.Emojis(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Emojis {
	return predicate.Emojis(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Emojis {
	return predicate.Emojis(sql.FieldLTE(FieldID, id))
}

// UnifiedCode applies equality check predicate on the "unified_code" field. It's identical to UnifiedCodeEQ.
func UnifiedCode(v string) predicate.Emojis {
	return predicate.Emojis(sql.FieldEQ(FieldUnifiedCode, v))
}

// Shortcode applies equality check predicate on the "shortcode" field. It's identical to ShortcodeEQ.
func Shortcode(v string) predicate.Emojis {
	return predicate.Emojis(sql.FieldEQ(FieldShortcode, v))
}

// UnifiedCodeEQ applies the EQ predicate on the "unified_code" field.
func UnifiedCodeEQ(v string) predicate.Emojis {
	return predicate.Emojis(sql.FieldEQ(FieldUnifiedCode, v))
}

// UnifiedCodeNEQ applies the NEQ predicate on the "unified_code" field.
func UnifiedCodeNEQ(v string) predicate.Emojis {
	return predicate.Emojis(sql.FieldNEQ(FieldUnifiedCode, v))
}

// UnifiedCodeIn applies the In predicate on the "unified_code" field.
func UnifiedCodeIn(vs ...string) predicate.Emojis {
	return predicate.Emojis(sql.FieldIn(FieldUnifiedCode, vs...))
}

// UnifiedCodeNotIn applies the NotIn predicate on the "unified_code" field.
func UnifiedCodeNotIn(vs ...string) predicate.Emojis {
	return predicate.Emojis(sql.FieldNotIn(FieldUnifiedCode, vs...))
}

// UnifiedCodeGT applies the GT predicate on the "unified_code" field.
func UnifiedCodeGT(v string) predicate.Emojis {
	return predicate.Emojis(sql.FieldGT(FieldUnifiedCode, v))
}

// UnifiedCodeGTE applies the GTE predicate on the "unified_code" field.
func UnifiedCodeGTE(v string) predicate.Emojis {
	return predicate.Emojis(sql.FieldGTE(FieldUnifiedCode, v))
}

// UnifiedCodeLT applies the LT predicate on the "unified_code" field.
func UnifiedCodeLT(v string) predicate.Emojis {
	return predicate.Emojis(sql.FieldLT(FieldUnifiedCode, v))
}

// UnifiedCodeLTE applies the LTE predicate on the "unified_code" field.
func UnifiedCodeLTE(v string) predicate.Emojis {
	return predicate.Emojis(sql.FieldLTE(FieldUnifiedCode, v))
}

// UnifiedCodeContains applies the Contains predicate on the "unified_code" field.
func UnifiedCodeContains(v string) predicate.Emojis {
	return predicate.Emojis(sql.FieldContains(FieldUnifiedCode, v))
}

// UnifiedCodeHasPrefix applies the HasPrefix predicate on the "unified_code" field.
func UnifiedCodeHasPrefix(v string) predicate.Emojis {
	return predicate.Emojis(sql.FieldHasPrefix(FieldUnifiedCode, v))
}

// UnifiedCodeHasSuffix applies the HasSuffix predicate on the "unified_code" field.
func UnifiedCodeHasSuffix(v string) predicate.Emojis {
	return predicate.Emojis(sql.FieldHasSuffix(FieldUnifiedCode, v))
}

// UnifiedCodeEqualFold applies the EqualFold predicate on the "unified_code" field.
func UnifiedCodeEqualFold(v string) predicate.Emojis {
	return predicate.Emojis(sql.FieldEqualFold(FieldUnifiedCode, v))
}

// UnifiedCodeContainsFold applies the ContainsFold predicate on the "unified_code" field.
func UnifiedCodeContainsFold(v string) predicate.Emojis {
	return predicate.Emojis(sql.FieldContainsFold(FieldUnifiedCode, v))
}

// ShortcodeEQ applies the EQ predicate on the "shortcode" field.
func ShortcodeEQ(v string) predicate.Emojis {
	return predicate.Emojis(sql.FieldEQ(FieldShortcode, v))
}

// ShortcodeNEQ applies the NEQ predicate on the "shortcode" field.
func ShortcodeNEQ(v string) predicate.Emojis {
	return predicate.Emojis(sql.FieldNEQ(FieldShortcode, v))
}

// ShortcodeIn applies the In predicate on the "shortcode" field.
func ShortcodeIn(vs ...string) predicate.Emojis {
	return predicate.Emojis(sql.FieldIn(FieldShortcode, vs...))
}

// ShortcodeNotIn applies the NotIn predicate on the "shortcode" field.
func ShortcodeNotIn(vs ...string) predicate.Emojis {
	return predicate.Emojis(sql.FieldNotIn(FieldShortcode, vs...))
}

// ShortcodeGT applies the GT predicate on the "shortcode" field.
func ShortcodeGT(v string) predicate.Emojis {
	return predicate.Emojis(sql.FieldGT(FieldShortcode, v))
}

// ShortcodeGTE applies the GTE predicate on the "shortcode" field.
func ShortcodeGTE(v string) predicate.Emojis {
	return predicate.Emojis(sql.FieldGTE(FieldShortcode, v))
}

// ShortcodeLT applies the LT predicate on the "shortcode" field.
func ShortcodeLT(v string) predicate.Emojis {
	return predicate.Emojis(sql.FieldLT(FieldShortcode, v))
}

// ShortcodeLTE applies the LTE predicate on the "shortcode" field.
func ShortcodeLTE(v string) predicate.Emojis {
	return predicate.Emojis(sql.FieldLTE(FieldShortcode, v))
}

// ShortcodeContains applies the Contains predicate on the "shortcode" field.
func ShortcodeContains(v string) predicate.Emojis {
	return predicate.Emojis(sql.FieldContains(FieldShortcode, v))
}

// ShortcodeHasPrefix applies the HasPrefix predicate on the "shortcode" field.
func ShortcodeHasPrefix(v string) predicate.Emojis {
	return predicate.Emojis(sql.FieldHasPrefix(FieldShortcode, v))
}

// ShortcodeHasSuffix applies the HasSuffix predicate on the "shortcode" field.
func ShortcodeHasSuffix(v string) predicate.Emojis {
	return predicate.Emojis(sql.FieldHasSuffix(FieldShortcode, v))
}

// ShortcodeEqualFold applies the EqualFold predicate on the "shortcode" field.
func ShortcodeEqualFold(v string) predicate.Emojis {
	return predicate.Emojis(sql.FieldEqualFold(FieldShortcode, v))
}

// ShortcodeContainsFold applies the ContainsFold predicate on the "shortcode" field.
func ShortcodeContainsFold(v string) predicate.Emojis {
	return predicate.Emojis(sql.FieldContainsFold(FieldShortcode, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Emojis) predicate.Emojis {
	return predicate.Emojis(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Emojis) predicate.Emojis {
	return predicate.Emojis(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Emojis) predicate.Emojis {
	return predicate.Emojis(sql.NotPredicates(p))
}
