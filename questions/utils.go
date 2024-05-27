package questions

import (
	"regexp"
	"time"
)

func getCurrentTime() string {
	currentTime := time.Now()
	formattedTime := currentTime.Format("02-01-2006 15:04:05")
	return formattedTime
}

func validateLink(link string) bool {
    // Regular expression pattern to match a valid URL
    pattern := `^(https?|ftp)://[^\s/$.?#].[^\s]*$`
    
    // Compile the regular expression pattern
    re := regexp.MustCompile(pattern)
    
    // Match the link against the pattern
    isValid := re.MatchString(link)
    
    return isValid
}

