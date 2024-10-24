package main
import "fmt"
import "strconv"
import "encoding/json"

func Season(month int) string {
  var season string
  switch month {
    case 1,2,12:
      season = "Winter"
    case 3,4,5:
      season = "Spring"
    case 6,7,8:
      season = "Summer"
    case 9,10,11:
      season = "Autumn"
    default:
      season = "Season unknown"
    }
  return season
}

// optimized code

func SeasonOptimized(month int) string {
  switch month{
    case 1,2,12: return "Winter"
    case 3,4,5:  return "Spring"
    case 6,7,8:  return "Summer"
    case 9,10,11:return "Autumn"
    default: return "Season Unknown"
  }
}
