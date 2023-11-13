package app

import "testing"

func TestIsUrl(t *testing.T) {
	testParameters := []struct {
		url      string
		expected bool
	}{
		{
			url:      "something",
			expected: false,
		},
		{
			url:      "http://localhost:8080",
			expected: true,
		},
		{
			url:      "",
			expected: false,
		},
		{
			url:      "https://github.com/ylerby",
			expected: true,
		},
		{
			url:      "https://job.ozon.ru/fintech/",
			expected: true,
		},
	}

	for testIndex, testCase := range testParameters {
		result := IsUrl(testCase.url)

		if result != testCase.expected {
			t.Errorf("Тест %d не пройден. Ожидалось %t. Получено %t", testIndex, testCase.expected, result)
		}
	}
}
