package cmd

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/zclconf/go-cty/cty"
)

// decodeVarFile parses a flat HCL var-file (top-level key = value pairs only)
// into a map[string]interface{}, preserving native Go types (string, int, float64, bool, []interface{}, map[string]interface{})
// for compatibility with the rest of the parsing pipeline.
func decodeVarFile(content []byte, filename string) (map[string]interface{}, error) {
	file, diags := hclsyntax.ParseConfig(content, filename, hcl.InitialPos)
	if diags.HasErrors() {
		return nil, diags
	}

	body, ok := file.Body.(*hclsyntax.Body)
	if !ok {
		return nil, fmt.Errorf("unexpected HCL body type")
	}

	decoded := map[string]interface{}{}

	for name, attribute := range body.Attributes {
		value, diags := attribute.Expr.Value(nil)
		if diags.HasErrors() {
			return nil, diags
		}

		converted, err := ctyValueToGo(value)
		if err != nil {
			return nil, err
		}

		decoded[name] = converted
	}

	return decoded, nil
}

func ctyValueToGo(value cty.Value) (interface{}, error) {
	if value.IsNull() {
		return nil, nil
	}

	if !value.IsKnown() {
		return nil, fmt.Errorf("unsupported var-file value: value is not known")
	}

	ty := value.Type()

	switch {
	case ty == cty.String:
		return value.AsString(), nil
	case ty == cty.Bool:
		return value.True(), nil
	case ty == cty.Number:
		bigFloat := value.AsBigFloat()
		if bigFloat.IsInt() {
			intValue, _ := bigFloat.Int64()
			return int(intValue), nil
		}
		floatValue, _ := bigFloat.Float64()
		return floatValue, nil
	case ty.IsListType() || ty.IsTupleType() || ty.IsSetType():
		list := make([]interface{}, 0, value.LengthInt())
		for it := value.ElementIterator(); it.Next(); {
			_, element := it.Element()
			converted, err := ctyValueToGo(element)
			if err != nil {
				return nil, err
			}
			list = append(list, converted)
		}
		return list, nil
	case ty.IsMapType() || ty.IsObjectType():
		mapped := make(map[string]interface{}, value.LengthInt())
		for it := value.ElementIterator(); it.Next(); {
			key, element := it.Element()
			converted, err := ctyValueToGo(element)
			if err != nil {
				return nil, err
			}
			mapped[key.AsString()] = converted
		}
		return mapped, nil
	default:
		return nil, fmt.Errorf("unsupported var-file value type: %s", ty.FriendlyName())
	}
}
