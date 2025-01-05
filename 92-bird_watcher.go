package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	count := 0
    for i := 0; i < len(birdsPerDay); i++ {
        count += birdsPerDay[i]
    }
    return count
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	endDayI := week * 7 // 14 
	count := 0
    for i := endDayI - 7; i < endDayI; i++ {
        count += birdsPerDay[i]
    }
    return count
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    birdsPerDay[0]++
	for i := 1; i < len(birdsPerDay); i++ {
        if i % 2 == 0 {
            birdsPerDay[i]++
        }
    }
    return birdsPerDay
}
