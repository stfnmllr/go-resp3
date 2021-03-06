// Code generated by "converter"; DO NOT EDIT.

package client

func (n _null) Attr() *Map                  { return nil }
func (n _null) ToBool() (bool, error)       { return false, newConversionError("ToBool", n) }
func (n _null) ToFloat64() (float64, error) { return 0, newConversionError("ToFloat64", n) }
func (n _null) ToInt64() (int64, error)     { return 0, newConversionError("ToInt64", n) }
func (n _null) ToString() (string, error)   { return "", newConversionError("ToString", n) }
func (n _null) ToVerbatimString() (VerbatimString, error) {
	return "", newConversionError("ToVerbatimString", n)
}

func (s _string) Attr() *Map                     { return nil }
func (s _string) ToInt64Slice() ([]int64, error) { return nil, newConversionError("ToInt64Slice", s) }
func (s _string) ToIntfSlice() ([]interface{}, error) {
	return nil, newConversionError("ToIntfSlice", s)
}
func (s _string) ToIntfSlice2() ([][]interface{}, error) {
	return nil, newConversionError("ToIntfSlice2", s)
}
func (s _string) ToIntfSlice3() ([][][]interface{}, error) {
	return nil, newConversionError("ToIntfSlice3", s)
}
func (s _string) ToMap() (Map, error)     { return nil, newConversionError("ToMap", s) }
func (s _string) ToSet() (Set, error)     { return nil, newConversionError("ToSet", s) }
func (s _string) ToSlice() (Slice, error) { return nil, newConversionError("ToSlice", s) }
func (s _string) ToStringInt64Map() (map[string]int64, error) {
	return nil, newConversionError("ToStringInt64Map", s)
}
func (s _string) ToStringMap() (map[string]interface{}, error) {
	return nil, newConversionError("ToStringMap", s)
}
func (s _string) ToStringMapSlice() ([]map[string]interface{}, error) {
	return nil, newConversionError("ToStringMapSlice", s)
}
func (s _string) ToStringSet() (map[string]bool, error) {
	return nil, newConversionError("ToStringSet", s)
}
func (s _string) ToStringSlice() ([]string, error) {
	return nil, newConversionError("ToStringSlice", s)
}
func (s _string) ToStringStringMap() (map[string]string, error) {
	return nil, newConversionError("ToStringStringMap", s)
}
func (s _string) ToStringValueMap() (map[string]RedisValue, error) {
	return nil, newConversionError("ToStringValueMap", s)
}
func (s _string) ToTree() ([]interface{}, error) { return nil, newConversionError("ToTree", s) }
func (s _string) ToVerbatimString() (VerbatimString, error) {
	return "", newConversionError("ToVerbatimString", s)
}
func (s _string) ToXrange() ([]XItem, error)           { return nil, newConversionError("ToXrange", s) }
func (s _string) ToXread() (map[string][]XItem, error) { return nil, newConversionError("ToXread", s) }

func (n _number) Attr() *Map                     { return nil }
func (n _number) ToInt64Slice() ([]int64, error) { return nil, newConversionError("ToInt64Slice", n) }
func (n _number) ToIntfSlice() ([]interface{}, error) {
	return nil, newConversionError("ToIntfSlice", n)
}
func (n _number) ToIntfSlice2() ([][]interface{}, error) {
	return nil, newConversionError("ToIntfSlice2", n)
}
func (n _number) ToIntfSlice3() ([][][]interface{}, error) {
	return nil, newConversionError("ToIntfSlice3", n)
}
func (n _number) ToMap() (Map, error)     { return nil, newConversionError("ToMap", n) }
func (n _number) ToSet() (Set, error)     { return nil, newConversionError("ToSet", n) }
func (n _number) ToSlice() (Slice, error) { return nil, newConversionError("ToSlice", n) }
func (n _number) ToStringInt64Map() (map[string]int64, error) {
	return nil, newConversionError("ToStringInt64Map", n)
}
func (n _number) ToStringMap() (map[string]interface{}, error) {
	return nil, newConversionError("ToStringMap", n)
}
func (n _number) ToStringMapSlice() ([]map[string]interface{}, error) {
	return nil, newConversionError("ToStringMapSlice", n)
}
func (n _number) ToStringSet() (map[string]bool, error) {
	return nil, newConversionError("ToStringSet", n)
}
func (n _number) ToStringSlice() ([]string, error) {
	return nil, newConversionError("ToStringSlice", n)
}
func (n _number) ToStringStringMap() (map[string]string, error) {
	return nil, newConversionError("ToStringStringMap", n)
}
func (n _number) ToStringValueMap() (map[string]RedisValue, error) {
	return nil, newConversionError("ToStringValueMap", n)
}
func (n _number) ToTree() ([]interface{}, error) { return nil, newConversionError("ToTree", n) }
func (n _number) ToVerbatimString() (VerbatimString, error) {
	return "", newConversionError("ToVerbatimString", n)
}
func (n _number) ToXrange() ([]XItem, error)           { return nil, newConversionError("ToXrange", n) }
func (n _number) ToXread() (map[string][]XItem, error) { return nil, newConversionError("ToXread", n) }

func (d _double) Attr() *Map                     { return nil }
func (d _double) ToInt64() (int64, error)        { return 0, newConversionError("ToInt64", d) }
func (d _double) ToInt64Slice() ([]int64, error) { return nil, newConversionError("ToInt64Slice", d) }
func (d _double) ToIntfSlice() ([]interface{}, error) {
	return nil, newConversionError("ToIntfSlice", d)
}
func (d _double) ToIntfSlice2() ([][]interface{}, error) {
	return nil, newConversionError("ToIntfSlice2", d)
}
func (d _double) ToIntfSlice3() ([][][]interface{}, error) {
	return nil, newConversionError("ToIntfSlice3", d)
}
func (d _double) ToMap() (Map, error)     { return nil, newConversionError("ToMap", d) }
func (d _double) ToSet() (Set, error)     { return nil, newConversionError("ToSet", d) }
func (d _double) ToSlice() (Slice, error) { return nil, newConversionError("ToSlice", d) }
func (d _double) ToStringInt64Map() (map[string]int64, error) {
	return nil, newConversionError("ToStringInt64Map", d)
}
func (d _double) ToStringMap() (map[string]interface{}, error) {
	return nil, newConversionError("ToStringMap", d)
}
func (d _double) ToStringMapSlice() ([]map[string]interface{}, error) {
	return nil, newConversionError("ToStringMapSlice", d)
}
func (d _double) ToStringSet() (map[string]bool, error) {
	return nil, newConversionError("ToStringSet", d)
}
func (d _double) ToStringSlice() ([]string, error) {
	return nil, newConversionError("ToStringSlice", d)
}
func (d _double) ToStringStringMap() (map[string]string, error) {
	return nil, newConversionError("ToStringStringMap", d)
}
func (d _double) ToStringValueMap() (map[string]RedisValue, error) {
	return nil, newConversionError("ToStringValueMap", d)
}
func (d _double) ToTree() ([]interface{}, error) { return nil, newConversionError("ToTree", d) }
func (d _double) ToVerbatimString() (VerbatimString, error) {
	return "", newConversionError("ToVerbatimString", d)
}
func (d _double) ToXrange() ([]XItem, error)           { return nil, newConversionError("ToXrange", d) }
func (d _double) ToXread() (map[string][]XItem, error) { return nil, newConversionError("ToXread", d) }

func (n *_bignumber) Attr() *Map { return nil }
func (n *_bignumber) ToInt64Slice() ([]int64, error) {
	return nil, newConversionError("ToInt64Slice", n)
}
func (n *_bignumber) ToIntfSlice() ([]interface{}, error) {
	return nil, newConversionError("ToIntfSlice", n)
}
func (n *_bignumber) ToIntfSlice2() ([][]interface{}, error) {
	return nil, newConversionError("ToIntfSlice2", n)
}
func (n *_bignumber) ToIntfSlice3() ([][][]interface{}, error) {
	return nil, newConversionError("ToIntfSlice3", n)
}
func (n *_bignumber) ToMap() (Map, error)     { return nil, newConversionError("ToMap", n) }
func (n *_bignumber) ToSet() (Set, error)     { return nil, newConversionError("ToSet", n) }
func (n *_bignumber) ToSlice() (Slice, error) { return nil, newConversionError("ToSlice", n) }
func (n *_bignumber) ToStringInt64Map() (map[string]int64, error) {
	return nil, newConversionError("ToStringInt64Map", n)
}
func (n *_bignumber) ToStringMap() (map[string]interface{}, error) {
	return nil, newConversionError("ToStringMap", n)
}
func (n *_bignumber) ToStringMapSlice() ([]map[string]interface{}, error) {
	return nil, newConversionError("ToStringMapSlice", n)
}
func (n *_bignumber) ToStringSet() (map[string]bool, error) {
	return nil, newConversionError("ToStringSet", n)
}
func (n *_bignumber) ToStringSlice() ([]string, error) {
	return nil, newConversionError("ToStringSlice", n)
}
func (n *_bignumber) ToStringStringMap() (map[string]string, error) {
	return nil, newConversionError("ToStringStringMap", n)
}
func (n *_bignumber) ToStringValueMap() (map[string]RedisValue, error) {
	return nil, newConversionError("ToStringValueMap", n)
}
func (n *_bignumber) ToTree() ([]interface{}, error) { return nil, newConversionError("ToTree", n) }
func (n *_bignumber) ToVerbatimString() (VerbatimString, error) {
	return "", newConversionError("ToVerbatimString", n)
}
func (n *_bignumber) ToXrange() ([]XItem, error) { return nil, newConversionError("ToXrange", n) }
func (n *_bignumber) ToXread() (map[string][]XItem, error) {
	return nil, newConversionError("ToXread", n)
}

func (b _boolean) Attr() *Map                     { return nil }
func (b _boolean) ToInt64Slice() ([]int64, error) { return nil, newConversionError("ToInt64Slice", b) }
func (b _boolean) ToIntfSlice() ([]interface{}, error) {
	return nil, newConversionError("ToIntfSlice", b)
}
func (b _boolean) ToIntfSlice2() ([][]interface{}, error) {
	return nil, newConversionError("ToIntfSlice2", b)
}
func (b _boolean) ToIntfSlice3() ([][][]interface{}, error) {
	return nil, newConversionError("ToIntfSlice3", b)
}
func (b _boolean) ToMap() (Map, error)     { return nil, newConversionError("ToMap", b) }
func (b _boolean) ToSet() (Set, error)     { return nil, newConversionError("ToSet", b) }
func (b _boolean) ToSlice() (Slice, error) { return nil, newConversionError("ToSlice", b) }
func (b _boolean) ToStringInt64Map() (map[string]int64, error) {
	return nil, newConversionError("ToStringInt64Map", b)
}
func (b _boolean) ToStringMap() (map[string]interface{}, error) {
	return nil, newConversionError("ToStringMap", b)
}
func (b _boolean) ToStringMapSlice() ([]map[string]interface{}, error) {
	return nil, newConversionError("ToStringMapSlice", b)
}
func (b _boolean) ToStringSet() (map[string]bool, error) {
	return nil, newConversionError("ToStringSet", b)
}
func (b _boolean) ToStringSlice() ([]string, error) {
	return nil, newConversionError("ToStringSlice", b)
}
func (b _boolean) ToStringStringMap() (map[string]string, error) {
	return nil, newConversionError("ToStringStringMap", b)
}
func (b _boolean) ToStringValueMap() (map[string]RedisValue, error) {
	return nil, newConversionError("ToStringValueMap", b)
}
func (b _boolean) ToTree() ([]interface{}, error) { return nil, newConversionError("ToTree", b) }
func (b _boolean) ToVerbatimString() (VerbatimString, error) {
	return "", newConversionError("ToVerbatimString", b)
}
func (b _boolean) ToXrange() ([]XItem, error)           { return nil, newConversionError("ToXrange", b) }
func (b _boolean) ToXread() (map[string][]XItem, error) { return nil, newConversionError("ToXread", b) }

func (s _verbatimString) Attr() *Map { return nil }
func (s _verbatimString) ToInt64Slice() ([]int64, error) {
	return nil, newConversionError("ToInt64Slice", s)
}
func (s _verbatimString) ToIntfSlice() ([]interface{}, error) {
	return nil, newConversionError("ToIntfSlice", s)
}
func (s _verbatimString) ToIntfSlice2() ([][]interface{}, error) {
	return nil, newConversionError("ToIntfSlice2", s)
}
func (s _verbatimString) ToIntfSlice3() ([][][]interface{}, error) {
	return nil, newConversionError("ToIntfSlice3", s)
}
func (s _verbatimString) ToMap() (Map, error)     { return nil, newConversionError("ToMap", s) }
func (s _verbatimString) ToSet() (Set, error)     { return nil, newConversionError("ToSet", s) }
func (s _verbatimString) ToSlice() (Slice, error) { return nil, newConversionError("ToSlice", s) }
func (s _verbatimString) ToStringInt64Map() (map[string]int64, error) {
	return nil, newConversionError("ToStringInt64Map", s)
}
func (s _verbatimString) ToStringMap() (map[string]interface{}, error) {
	return nil, newConversionError("ToStringMap", s)
}
func (s _verbatimString) ToStringMapSlice() ([]map[string]interface{}, error) {
	return nil, newConversionError("ToStringMapSlice", s)
}
func (s _verbatimString) ToStringSet() (map[string]bool, error) {
	return nil, newConversionError("ToStringSet", s)
}
func (s _verbatimString) ToStringSlice() ([]string, error) {
	return nil, newConversionError("ToStringSlice", s)
}
func (s _verbatimString) ToStringStringMap() (map[string]string, error) {
	return nil, newConversionError("ToStringStringMap", s)
}
func (s _verbatimString) ToStringValueMap() (map[string]RedisValue, error) {
	return nil, newConversionError("ToStringValueMap", s)
}
func (s _verbatimString) ToTree() ([]interface{}, error) { return nil, newConversionError("ToTree", s) }
func (s _verbatimString) ToXrange() ([]XItem, error)     { return nil, newConversionError("ToXrange", s) }
func (s _verbatimString) ToXread() (map[string][]XItem, error) {
	return nil, newConversionError("ToXread", s)
}

func (s _slice) Attr() *Map                  { return nil }
func (s _slice) ToBool() (bool, error)       { return false, newConversionError("ToBool", s) }
func (s _slice) ToFloat64() (float64, error) { return 0, newConversionError("ToFloat64", s) }
func (s _slice) ToInt64() (int64, error)     { return 0, newConversionError("ToInt64", s) }
func (s _slice) ToMap() (Map, error)         { return nil, newConversionError("ToMap", s) }
func (s _slice) ToSet() (Set, error)         { return nil, newConversionError("ToSet", s) }
func (s _slice) ToString() (string, error)   { return "", newConversionError("ToString", s) }
func (s _slice) ToStringInt64Map() (map[string]int64, error) {
	return nil, newConversionError("ToStringInt64Map", s)
}
func (s _slice) ToStringMap() (map[string]interface{}, error) {
	return nil, newConversionError("ToStringMap", s)
}
func (s _slice) ToStringSet() (map[string]bool, error) {
	return nil, newConversionError("ToStringSet", s)
}
func (s _slice) ToStringStringMap() (map[string]string, error) {
	return nil, newConversionError("ToStringStringMap", s)
}
func (s _slice) ToStringValueMap() (map[string]RedisValue, error) {
	return nil, newConversionError("ToStringValueMap", s)
}
func (s _slice) ToVerbatimString() (VerbatimString, error) {
	return "", newConversionError("ToVerbatimString", s)
}
func (s _slice) ToXread() (map[string][]XItem, error) { return nil, newConversionError("ToXread", s) }

func (m _map) Attr() *Map                          { return nil }
func (m _map) ToBool() (bool, error)               { return false, newConversionError("ToBool", m) }
func (m _map) ToFloat64() (float64, error)         { return 0, newConversionError("ToFloat64", m) }
func (m _map) ToInt64() (int64, error)             { return 0, newConversionError("ToInt64", m) }
func (m _map) ToInt64Slice() ([]int64, error)      { return nil, newConversionError("ToInt64Slice", m) }
func (m _map) ToIntfSlice() ([]interface{}, error) { return nil, newConversionError("ToIntfSlice", m) }
func (m _map) ToIntfSlice2() ([][]interface{}, error) {
	return nil, newConversionError("ToIntfSlice2", m)
}
func (m _map) ToIntfSlice3() ([][][]interface{}, error) {
	return nil, newConversionError("ToIntfSlice3", m)
}
func (m _map) ToSet() (Set, error)       { return nil, newConversionError("ToSet", m) }
func (m _map) ToSlice() (Slice, error)   { return nil, newConversionError("ToSlice", m) }
func (m _map) ToString() (string, error) { return "", newConversionError("ToString", m) }
func (m _map) ToStringMapSlice() ([]map[string]interface{}, error) {
	return nil, newConversionError("ToStringMapSlice", m)
}
func (m _map) ToStringSet() (map[string]bool, error) {
	return nil, newConversionError("ToStringSet", m)
}
func (m _map) ToStringSlice() ([]string, error) { return nil, newConversionError("ToStringSlice", m) }
func (m _map) ToTree() ([]interface{}, error)   { return nil, newConversionError("ToTree", m) }
func (m _map) ToVerbatimString() (VerbatimString, error) {
	return "", newConversionError("ToVerbatimString", m)
}
func (m _map) ToXrange() ([]XItem, error) { return nil, newConversionError("ToXrange", m) }

func (s _set) Attr() *Map                          { return nil }
func (s _set) ToBool() (bool, error)               { return false, newConversionError("ToBool", s) }
func (s _set) ToFloat64() (float64, error)         { return 0, newConversionError("ToFloat64", s) }
func (s _set) ToInt64() (int64, error)             { return 0, newConversionError("ToInt64", s) }
func (s _set) ToInt64Slice() ([]int64, error)      { return nil, newConversionError("ToInt64Slice", s) }
func (s _set) ToIntfSlice() ([]interface{}, error) { return nil, newConversionError("ToIntfSlice", s) }
func (s _set) ToIntfSlice2() ([][]interface{}, error) {
	return nil, newConversionError("ToIntfSlice2", s)
}
func (s _set) ToIntfSlice3() ([][][]interface{}, error) {
	return nil, newConversionError("ToIntfSlice3", s)
}
func (s _set) ToMap() (Map, error)       { return nil, newConversionError("ToMap", s) }
func (s _set) ToSlice() (Slice, error)   { return nil, newConversionError("ToSlice", s) }
func (s _set) ToString() (string, error) { return "", newConversionError("ToString", s) }
func (s _set) ToStringInt64Map() (map[string]int64, error) {
	return nil, newConversionError("ToStringInt64Map", s)
}
func (s _set) ToStringMap() (map[string]interface{}, error) {
	return nil, newConversionError("ToStringMap", s)
}
func (s _set) ToStringMapSlice() ([]map[string]interface{}, error) {
	return nil, newConversionError("ToStringMapSlice", s)
}
func (s _set) ToStringSlice() ([]string, error) { return nil, newConversionError("ToStringSlice", s) }
func (s _set) ToStringStringMap() (map[string]string, error) {
	return nil, newConversionError("ToStringStringMap", s)
}
func (s _set) ToStringValueMap() (map[string]RedisValue, error) {
	return nil, newConversionError("ToStringValueMap", s)
}
func (s _set) ToTree() ([]interface{}, error) { return nil, newConversionError("ToTree", s) }
func (s _set) ToVerbatimString() (VerbatimString, error) {
	return "", newConversionError("ToVerbatimString", s)
}
func (s _set) ToXrange() ([]XItem, error)           { return nil, newConversionError("ToXrange", s) }
func (s _set) ToXread() (map[string][]XItem, error) { return nil, newConversionError("ToXread", s) }
