package domain

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type JsonResponseBody struct {
	value any
}

func NewJsonResponseBody(value any) JsonResponseBody {
	return JsonResponseBody{
		value,
	}
}

type KeyChain []string

func (j JsonResponseBody) toAssertions() []Step {
	return toAssertions(j.value, []string{"$"})
}

func toAssertions(v any, keyChain KeyChain) []Step {
	var steps []Step

	switch v := v.(type) {
	case map[string]any:
		steps = objectToAssertions(v, keyChain)
	case []any:
		steps = arrayToAssertions(v, keyChain)
	case string:
		steps = []Step{toStringStep(v, keyChain)}
	case float64:
		steps = []Step{toNumberStep(v, keyChain)}
	case bool:
		steps = []Step{toBoolStep(v, keyChain)}
	case nil:
		steps = []Step{toNullStep(keyChain)}
	default:
		fmt.Printf("Warning: toAssertions - I don't know about type %T!\n", v)
	}

	sortSteps(steps)

	return steps
}

func objectToAssertions(object map[string]any, keyChain KeyChain) []Step {
	var assertions []Step
	for k, v := range object {
		keyChain := append(keyChain, k)
		assertions = append(assertions, toAssertions(v, keyChain)...)
	}

	return assertions
}

func arrayToAssertions(array []any, keyChain KeyChain) []Step {
	var assertions []Step
	for i, v := range array {
		copiedKeyChain := make(KeyChain, len(keyChain))
		copy(copiedKeyChain, keyChain)

		copiedKeyChain[len(keyChain)-1] += fmt.Sprintf("[%d]", i)
		assertions = append(assertions, toAssertions(v, copiedKeyChain)...)
	}
	return assertions
}

func toStringStep(s string, keyChain KeyChain) Step {
	return Step(fmt.Sprintf(`レスポンスのJSONの"%s"が文字列の"%s"である`, keyChain.toPath(), s))
}

func toNumberStep(n float64, keyChain KeyChain) Step {
	if n == math.Trunc(n) {
		return Step(fmt.Sprintf(`レスポンスのJSONの"%s"が整数の"%d"である`, keyChain.toPath(), int64(n)))
	} else {
		formatted := strconv.FormatFloat(n, 'f', -1, 64)
		return Step(fmt.Sprintf(`レスポンスのJSONの"%s"が小数の"%s"である`, keyChain.toPath(), formatted))
	}
}

func toBoolStep(b bool, keyChain KeyChain) Step {
	return Step(fmt.Sprintf(`レスポンスのJSONの"%s"が真偽値の"%t"である`, keyChain.toPath(), b))
}

func toNullStep(keyChain KeyChain) Step {
	return Step(fmt.Sprintf(`レスポンスのJSONの"%s"がnullである`, keyChain.toPath()))
}

func (k KeyChain) toPath() string {
	return strings.Join(k, ".")
}

func sortSteps(steps []Step) {
	replaceArrayIndexesWithZeroPaddedOne := func(step string) string {
		r := regexp.MustCompile(`\[\d+\]`)
		indexes := r.FindAllStringSubmatch(step, -1)

		var indexNumbers []int
		for _, v := range indexes {
			i, _ := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(v[0], "["), "]"))
			indexNumbers = append(indexNumbers, i)
		}

		var paddedIndexes []string
		for _, v := range indexNumbers {
			padded := fmt.Sprintf("%03d", v)
			paddedIndexes = append(paddedIndexes, padded)
		}

		var replaced = step
		for i, _ := range indexes {
			replaced = strings.Replace(replaced, indexes[i][0], paddedIndexes[i], -1)
		}

		return replaced
	}

	sort.Slice(steps, func(i, j int) bool {
		return replaceArrayIndexesWithZeroPaddedOne(string(steps[i])) < replaceArrayIndexesWithZeroPaddedOne(string(steps[j]))
	})
}
