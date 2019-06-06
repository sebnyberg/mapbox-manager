package style

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getSourcesFromSourcesURL(t *testing.T) {
	testCases := []struct {
		testName  string
		input     string
		output    []string
		shouldErr bool
	}{
		{"EmptyString", "", nil, true},
		{"MissingURL", "test1,test2", nil, true},
		{"MissingSources", "mapbox://", nil, true},
		{"ValidURLTwoSources", "mapbox://test1,test2", []string{"test1", "test2"}, false},
		{"ValidURLTwoSourcesWithUser", "mapbox://user1.layer1,user2.layer2", []string{"user1.layer1", "user2.layer2"}, false},
		{"ValidURLOneSource", "mapbox://user1.layer1", []string{"user1.layer1"}, false},
	}

	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			sources, err := getSourcesFromSourcesURL(tt.input)
			if !tt.shouldErr {
				require.Nil(t, err)
			}
			require.Equal(t, tt.output, sources)
		})
	}
}

func Test_getMapboxSourcesFromURL(t *testing.T) {
	testCases := []struct {
		testName  string
		input     string
		output    []string
		shouldErr bool
	}{
		{"EmptyString", "", nil, true},
		{"MissingURL", "test1,test2", nil, true},
		{"MissingSources", "mapbox://", nil, true},
		{"TwoUnscopedNonMapboxSources", "mapbox://test1,test2", nil, true},
		{"TwoScopedNonMapboxSources", "mapbox://user1.layer1,user2.layer2", []string{}, false},
		{"ValidUrlWithMapbox", "mapbox://user1.layer1,mapbox.layer2,mapbox.layer3,user2.layer4", []string{"mapbox.layer2", "mapbox.layer3"}, false},
	}

	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			sources, err := getMapboxSourcesFromURL(tt.input)
			if !tt.shouldErr {
				require.Nil(t, err)
			}
			require.Equal(t, tt.output, sources)
		})
	}
}
