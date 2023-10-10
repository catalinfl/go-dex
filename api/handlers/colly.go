package handlers

import (
	"fmt"
	"strings"

	"github.com/catalinfl/godex/utils"
	"github.com/gocolly/colly/v2"
	"github.com/gofiber/fiber/v2"
)

type GetHTMLStruct struct {
	Words       []string
	Definitions []string
}

type Definition struct {
	Meaning string `json:"meaning"`
	Source  string `json:"source"`
}

type SimpleWord struct {
	Word       string      `json:"word"`
	TypeOfWord string      `json:"typeOfWord"`
	Definition *Definition `json:"definition"`
}

func CollyHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	var rawDefinition GetHTMLStruct

	col := colly.NewCollector()

	col.OnHTML("div#tab_2.tab-pane.show.active", func(e *colly.HTMLElement) {
		rawDefinition = GetHTMLStruct{
			Words:       e.ChildTexts("h3.tree-heading"),
			Definitions: e.ChildTexts("div.tree-body"),
		}
	})

	col.Visit(fmt.Sprintf("https://dexonline.ro/definitie/%s", id))

	wordsToJSON := make([]SimpleWord, len(rawDefinition.Words))

	for i, val := range rawDefinition.Words {

		contain := false
		for wordType := range utils.WordTypes {
			if strings.Contains(val, wordType) {
				wordsToJSON[i].Word = val[:strings.Index(val, wordType)]
				wordsToJSON[i].TypeOfWord = val[strings.Index(val, wordType):]
				contain = true
				break
			}
		}

		var source string = ""
		for definitionType := range utils.DefintionSources {
			found := -1
			if ok := strings.Contains(rawDefinition.Definitions[i], definitionType); ok {

				if found < strings.Index(rawDefinition.Definitions[i], definitionType) {
					found = strings.Index(rawDefinition.Definitions[i], definitionType)
				}

				source += definitionType + " "
			}
			if found != -1 {
				wordsToJSON[i].Definition = &Definition{
					Meaning: rawDefinition.Definitions[i][:found],
					Source:  source,
				}
			} else {
				wordsToJSON[i].Definition = &Definition{
					Meaning: rawDefinition.Definitions[i],
					Source:  source,
				}
			}
		}

		if !contain {
			wordsToJSON[i].Word = val
		}

	}

	return c.JSON(wordsToJSON)
}
