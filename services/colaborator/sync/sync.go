package sync

import (
	"albo/colaborator"
	c "albo/config"
	"albo/domain"
	"albo/pkg/albohttp"
	"albo/utils"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Marvel(srv colaborator.Service) error {
	// Initial sync to fill our database
	if err := syncColaborators(srv); err != nil {
		return err
	}

	for {
		<-time.After(time.Duration(c.Config.DURATION) * time.Minute)
		go syncColaborators(srv)
	}
}

func syncColaborators(srv colaborator.Service) error {
	client := albohttp.NewClient(10)

	var colaboratorsToSave []*domain.Colaborator
	characters := map[string]string{
		"ironman":    "Iron Man",
		"capamerica": "Captain America",
	}

	for key, character := range characters {
		// GET IDS FOR EACH CHARACTER
		characterByte, err := getCharacterByName(character, client)
		if err != nil {
			return err
		}

		var characterResponse CharacterResponse
		if err = json.Unmarshal(characterByte, &characterResponse); err != nil {
			return err
		}

		if len(characterResponse.Data.Results) == 0 {
			return errors.New("character not found")
		}
		characterID := characterResponse.Data.Results[0].ID

		// GET STORIES BY CHARACTER
		storiesByte, err := getStoriesByCharacterId(characterID, client)
		if err != nil {
			return err
		}

		var storiesResponse StoryResponse
		if err = json.Unmarshal(storiesByte, &storiesResponse); err != nil {
			return err
		}

		// GET COMICS BY STORY
		colaborator := &domain.Colaborator{
			ID:        primitive.NewObjectID(),
			CreatedAt: time.Now(),
			Character: key,
		}

		for _, story := range storiesResponse.Data.Results {
			comicsByte, err := getComicsByStori(story.Comics.CollectionURI, client)
			if err != nil {
				return err
			}

			var comicResponse ComicResponse
			if err = json.Unmarshal(comicsByte, &comicResponse); err != nil {
				return err
			}

			for _, comic := range comicResponse.Data.Results {
				for _, creator := range comic.Creators.Items {
					switch role := creator.Role; role {
					case "colorist":
						// Check duplicates
						if !utils.Contains(colaborator.Colorist, creator.Name) {
							colaborator.Colorist = append(colaborator.Colorist, creator.Name)
						}
					case "writer":
						// Check duplicates
						if !utils.Contains(colaborator.Writer, creator.Name) {
							colaborator.Writer = append(colaborator.Writer, creator.Name)
						}
					case "editor":
						// Check duplicates
						if !utils.Contains(colaborator.Editor, creator.Name) {
							colaborator.Editor = append(colaborator.Editor, creator.Name)
						}
					}
				}
			}
		}

		colaboratorsToSave = append(colaboratorsToSave, colaborator)
	}

	if err := srv.Save(colaboratorsToSave); err != nil {
		return err
	}

	fmt.Println("===Sync for characters was successful===")
	return nil
}
