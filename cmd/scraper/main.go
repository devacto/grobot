package main

import (
	"fmt"
	"strings"

	"github.com/devacto/grobot/data"
)

func main() {
	scrape("http://www.myfitnesspal.com/food/calories/179990009")
}

func scrape(url string) {
	fmt.Printf("URL: %s", url)
	scraper := NewScraper(url)

	f := data.NewFood(getFoodName(scraper), getNutritions(scraper))
	data.InsertFood(f)
}

func getFoodName(s *Scraper) string {
	selection := s.Find("#wrap #content #main .main-title")
	foodName := returnFirstValue(selection)
	return foodName
}

func getNutritions(s *Scraper) []data.Nutrition {
	selection := s.Find("#wrap #content #main #nutrition-facts tr")
	var nutritionArray []data.Nutrition

	for _, v := range selection {
		if trimmed := strings.TrimSpace(v); trimmed != "" {
			tokenised := strings.Split(trimmed, "\n\t\t\t")

			for k, _ := range tokenised {
				if (k == 0 || k == 2) {
					n := data.NewNutrition(tokenised[k], tokenised[k + 1])
					nutritionArray = append(nutritionArray, n)
				}
			}

		}
	}

	return nutritionArray
}

func returnFirstValue(s []string) string {
	for _, v := range s {
		if v != "" {
			return v
		}
	}
	return ""
}
