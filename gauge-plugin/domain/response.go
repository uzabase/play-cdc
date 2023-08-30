package domain

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Response struct {
	Status   int             `json:"status"`
	Headers  ResponseHeaders `json:"headers"`
	JsonBody any             `json:"jsonBody"`
}

type ResponseHeaders map[string]string

type KeyChain []string

func (r *Response) toStatusCodeStep() Step {
	return Step(fmt.Sprintf(`レスポンスステータスコードが"%d"である`, r.Status))
}

func (r *Response) toHeaderAssertions() []Step {
	var assertions []Step
	for k, v := range r.Headers {
		assertions = append(assertions, Step(fmt.Sprintf(`レスポンスヘッダーに"%s"が存在し、その値が"%s"である`, k, v)))
	}

	sort.Slice(assertions, func(i, j int) bool { return assertions[i] < assertions[j] })

	return assertions
}

func (r *Response) toBodyAssertions() []Step {
	return toAssertions(r.JsonBody, []string{"$"})
}

func toAssertions(v any, keyChain KeyChain) []Step {
	var steps []Step

	switch v := v.(type) {
	case map[string]any:
		steps = objectToAssertions(v, keyChain)
	case []any:
		steps = arrayToAssertions(v, keyChain)
	case string:
		steps = []Step{Step(toStringStep(v, keyChain))}
	case float64:
		steps = []Step{Step(toNumberStep(v, keyChain))}
	case bool:
		steps = []Step{Step(toBoolStep(v, keyChain))}
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

func toStringStep(s string, keyChain KeyChain) string {
	return fmt.Sprintf(`レスポンスのJSONの"%s"が文字列の"%s"である`, keyChain.toPath(), s)
}

func toNumberStep(n float64, keyChain KeyChain) string {
	if n == math.Trunc(n) {
		return fmt.Sprintf(`レスポンスのJSONの"%s"が整数の"%d"である`, keyChain.toPath(), int64(n))
	} else {
		formatted := strconv.FormatFloat(n, 'f', -1, 64)
		return fmt.Sprintf(`レスポンスのJSONの"%s"が小数の"%s"である`, keyChain.toPath(), formatted)
	}
}

func toBoolStep(b bool, keyChain KeyChain) string {
	return fmt.Sprintf(`レスポンスのJSONの"%s"が真偽値の"%t"である`, keyChain.toPath(), b)
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
