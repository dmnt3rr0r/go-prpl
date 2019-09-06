package go_prpl

import (
	"testing"
)

func cmpArrays(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func testArrays(t *testing.T, result, expected []int) {
	if !cmpArrays(result, expected) {
		t.Errorf("result did not equal expected; want %v, got %v", expected, result)
	}
}

func TestNormalizeCap(t *testing.T) {
	result := normalizeCap(nil, MaxInt)
	expected := []int{ MaxInt, 0 }
	testArrays(t, result, expected)

	result = normalizeCap([]int{ 42 }, MaxInt)
	expected = []int{ 42, 0 }
	testArrays(t, result, expected)

	result = normalizeCap([]int{ 6, 3 }, MaxInt)
	expected = []int{ 6, 3 }
	testArrays(t, result, expected) 
}


func TestNormalizeCapSince(t *testing.T) {

	value := CapSince{
		[]int{ 10 },
		[]int{ 10, 3 },
		[]int{ 10, 3 },
		nil,
		[]int{ 9, 2 },
		[]int{ 11, 3 },
		[]int{ 11, 3 }}
	result := normalizeCapSince(&value, MaxInt)
	expected := CapSince{
		[]int{ 10, 0 },
		[]int{ 10, 3 },
		[]int{ 10, 3 },
		[]int{ MaxInt, 0 },
		[]int{ 9, 2 },
		[]int{ 11, 3 },
		[]int{ 11, 3 }}

	testArrays(t, result.ES2015, expected.ES2015)
	testArrays(t, result.ES2016, expected.ES2016)
	testArrays(t, result.ES2017, expected.ES2017)
	testArrays(t, result.ES2018, expected.ES2018)
	testArrays(t, result.Push, expected.Push)
	testArrays(t, result.ServiceWorker, expected.ServiceWorker)
	testArrays(t, result.Modules, expected.Modules)
}

func TestUACaps(t *testing.T) {

	capMap, err := initCapMap("browscap.ini", "capmap.yaml")
	if err != nil {
		t.Errorf("Initilization failed")
	}

	browserUA := `Mozilla/5.0 (X11; CrOS x86_64 8172.45.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.64 Safari/537.36`

	testCaps, err := UACaps(browserUA, capMap)
	if (err != nil) || (testCaps == nil) {
		t.Errorf("General error: %v", err)
	}

	if !testCaps.ES2015 || !testCaps.ES2016 || !testCaps.ES2017 || !testCaps.ES2018 || !testCaps.Push || !testCaps.ServiceWorker || !testCaps.Modules {
		t.Errorf("Did not detect the correct features")
	}

}
