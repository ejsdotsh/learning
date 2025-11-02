package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var sb int = 0
    for i := 0; i <= (len(birdsPerDay) - 1); i++ {
        sb += birdsPerDay[i]
    }
    return sb
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var bw int = 0
	idx := 7 * (week - 1)
    for i := idx; i < (idx + 7); i++ {
        bw += birdsPerDay[i]
    }
	return bw
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i <= (len(birdsPerDay) - 1); i++ {
        if i % 2 == 0 {
            birdsPerDay[i] += 1
        }
    }
    return birdsPerDay
}
