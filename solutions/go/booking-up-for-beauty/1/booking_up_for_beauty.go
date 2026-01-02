package booking

import (
    "time"
	"fmt"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	result, _ := time.Parse("1/2/2006 15:04:05", date)
    
    return result
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    getDate, _ := time.Parse("January 2, 2006 15:04:05", date)
	return getDate.UnixNano() <= time.Now().UnixNano()
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	getDate, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
    return getDate.Hour() >= 12 && getDate.Hour() < 18 
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	convertedDate := Schedule(date)
    return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", convertedDate.Weekday(), convertedDate.Month(), convertedDate.Day(), convertedDate.Year(), convertedDate.Hour(), convertedDate.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
    return Schedule(fmt.Sprintf("9/15/%d 00:00:00", time.Now().Year()))
}
