package main

import (
	"fmt"
	"strings"

	"github.com/devacto/grobot/data"
)

const baseUrl = "http://www.myfitnesspal.com"

func main() {
	scrape("/food/calories/179990009")
}

func scrape(path string) {
	fmt.Printf("URL: %s%s\n", baseUrl, path)
	scraper := NewScraper(baseUrl + path)

	f := data.NewFood(getFoodName(scraper), getNutritions(scraper))
	data.InsertFood(f)
	getOtherLinks(scraper)
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

func getOtherLinks(s *Scraper) {
	selection := s.FindLink("#wrap #content #main #other-items ul li a")
	for _, v := range selection {
		if v!= "" {
			fmt.Printf("next link: %s\n", v)
		}
	}
}

func returnFirstValue(s []string) string {
	for _, v := range s {
		if v != "" {
			return v
		}
	}
	return ""
}
