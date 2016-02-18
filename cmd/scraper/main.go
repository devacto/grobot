package main

import (
	"log"
	"strings"

	"github.com/devacto/grobot/data"
)

const baseUrl = "http://www.myfitnesspal.com/food/calories/"

func main() {
	scrape("179990009")
}

func scrape(id string) {
	log.Printf("Visiting %s%s\n", baseUrl, id)

	scraper := NewScraper(baseUrl + id)

	if foodExists := data.FoodWithIdExists(id); foodExists == false {
		titleTokens := processTitle(scraper)
		foodName := strings.TrimSpace(titleTokens[0])
		foodCompany := ""
		if (len(titleTokens) > 1) {
			foodCompany = strings.TrimSpace(titleTokens[1])
		}

		f := data.NewFood(id, foodName, foodCompany, getNutritions(scraper))
		data.InsertFood(f)
	}

	otherFoodIds := getOtherIds(scraper)
	for _, v := range otherFoodIds {
		scrape(v)
	}
}

func processTitle(s *Scraper) []string {
	selection := s.Find("#wrap #content #main .main-title")
	title := returnFirstValue(selection)
	title = strings.Replace(title, "Calories in ", "", 1)
	titleTokens := strings.SplitN(title, "-", 1)
	return titleTokens
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

func getOtherIds(s *Scraper) []string {
	selection := s.FindLink("#wrap #content #main #other-items ul li a")
	var otherLinks []string

	for _, v := range selection {
		if v!= "" {
			otherLinks = append(otherLinks, getIdFromPath(v))
		}
	}

	return otherLinks
}

func returnFirstValue(s []string) string {
	for _, v := range s {
		if v != "" {
			return v
		}
	}
	return ""
}

func getIdFromPath(path string) string {
	tokenised := strings.Split(path, "/")
	id := tokenised[len(tokenised)-1]
	return id
}
