package resource

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTrackerAPI(t *testing.T) {
	parent, ok := TrackerParentName("users/3134441396375598")
	require.True(t, ok)
	assert.Equal(t, "users", parent.collection)
	assert.Equal(t, int64(3134441396375598), parent.id)

	tracker, ok := TrackerName("users/3134441396375598/trackers/6097286400577570")
	require.True(t, ok)
	assert.Equal(t, "trackers", tracker.collection)
	assert.Equal(t, int64(6097286400577570), tracker.id)
	assert.Equal(t, "users", tracker.Parent.collection)
	assert.Equal(t, int64(3134441396375598), tracker.Parent.id)

	trackerFull, ok := TrackerFullName("//tracker.animeapis.com/users/3134441396375598/trackers/6097286400577570")
	require.True(t, ok)
	assert.Equal(t, tracker.String(), trackerFull.String())

	activity, ok := ActivityName("audiences/3134441396375598/trackers/6097286400577570/activities/6097611928899618")
	require.True(t, ok)
	assert.Equal(t, "activities", activity.collection)
	assert.Equal(t, int64(6097611928899618), activity.id)
	assert.Equal(t, "trackers", activity.Parent.collection)
	assert.Equal(t, int64(6097286400577570), activity.Parent.id)
	assert.Equal(t, "audiences", activity.Parent.Parent.collection)
	assert.Equal(t, int64(3134441396375598), activity.Parent.Parent.id)

	activityFull, ok := ActivityFullName("//tracker.animeapis.com/audiences/3134441396375598/trackers/6097286400577570/activities/6097611928899618")
	require.True(t, ok)
	assert.Equal(t, activity.String(), activityFull.String())
}
