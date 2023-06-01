package template

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/anonnice1999/questx-script/internal/template/resource"
	"github.com/anonnice1999/questx-script/pkg/api"
)

var templates = []Template{
	&resource.DailyConnectTemplate{},            // 01
	&resource.FollowTwitter{},                   // 02
	&resource.InviteFriendToDiscord{},           // 03
	&resource.SetDiscordAvatarToCommunityLogo{}, // 04
	&resource.DailyKnowFrens{},                  // 05
	&resource.ReachDiscordLevel{},               // 06
	&resource.LeaveStarOnProductHunt{},          // 07
	&resource.CreateProductMeme{},               // 08
	&resource.ShareCommunity{},                  // 09
	&resource.InviteFollowCommunity{},           // 10
	&resource.TwitterReaction{},                 // 11
	&resource.SetTwitterAvatarToCommunityLogo{}, // 12
	&resource.WriteMedium{},                     // 13
	&resource.CreateCommunityVideo{},            // 14
	&resource.ParticipateInCommunityEvent{},     // 15
	&resource.TryProductDemo{},                  // 16
}

func Run(apiEndpoint, accessToken string) error {
	apiGenerator := api.NewGenerator()
	client := http.DefaultClient

	ctx := context.Background()
	ctx = context.WithValue(ctx, "httpclient", client)

	categoryMap, err := getExistingCategories(ctx, apiEndpoint, accessToken)
	if err != nil {
		return err
	}

	newCategoryMap, err := createNewCategory(ctx, apiEndpoint, accessToken, categoryMap)
	if err != nil {
		return err
	}

	for name, id := range newCategoryMap {
		categoryMap[name] = id
	}

	existingTemplate, err := getExistingTemplate(ctx, apiEndpoint, accessToken)
	if err != nil {
		return err
	}

	for _, template := range templates {
		if _, ok := existingTemplate[template.Title()]; ok {
			log.Println("Template", template.Title(), "was already existed")
			continue
		}

		category_id := ""
		if template.Category() != "" {
			var ok bool
			category_id, ok = categoryMap[template.Category()]
			if !ok {
				return errors.New("not found category")
			}
		}

		resp, err := apiGenerator.New(apiEndpoint, "/createQuest").
			Body(api.JSON{
				"type":            template.Type(),
				"title":           template.Title(),
				"description":     template.Description(),
				"status":          "draft",
				"category_id":     category_id,
				"recurrence":      template.Recurrence(),
				"condition_op":    "or",
				"points":          template.Points(),
				"awards":          template.Rewards(),
				"validation_data": template.ValidationData(),
			}).
			POST(ctx, api.OAuth2("Bearer", accessToken))
		if err != nil {
			return err
		}

		body := resp.Body.(api.JSON)
		code, err := body.GetInt("code")
		if err != nil {
			return err
		}

		if code != 0 {
			errormsg, err := body.GetString("error")
			if err != nil {
				return err
			}

			log.Println("Cannot create quest ", template.Title())
			log.Println(errormsg)

			return errors.New("cannot create quest")
		}

		log.Println("Created template", template.Title())
	}

	return nil
}

func getExistingCategories(ctx context.Context, apiEndpoint, accessToken string) (map[string]string, error) {
	apiGenerator := api.NewGenerator()
	resp, err := apiGenerator.New(apiEndpoint, "/getTemplateCategories").
		GET(ctx, api.OAuth2("Bearer", accessToken))
	if err != nil {
		return nil, err
	}

	body := resp.Body.(api.JSON)
	code, err := body.GetInt("code")
	if err != nil {
		return nil, err
	}

	if code != 0 {
		return nil, fmt.Errorf("cannot list category, %d", code)
	}

	data, err := body.GetJSON("data")
	if err != nil {
		return nil, err
	}

	categories, err := data.GetArray("categories")
	if err != nil {
		return nil, err
	}

	categoryMap := map[string]string{}
	for _, c := range categories {
		id, err := c.GetString("id")
		if err != nil {
			return nil, err
		}

		name, err := c.GetString("name")
		if err != nil {
			return nil, err
		}

		categoryMap[name] = id
	}

	return categoryMap, nil
}

func createNewCategory(
	ctx context.Context, apiEndpoint, accessToken string, excludes map[string]string,
) (map[string]string, error) {
	apiGenerator := api.NewGenerator()
	categoryMap := map[string]string{}

	for _, c := range resource.Categories {
		if _, ok := excludes[c]; ok {
			continue
		}

		resp, err := apiGenerator.New(apiEndpoint, "/createCategory").
			Body(api.JSON{"name": c}).
			POST(ctx, api.OAuth2("Bearer", accessToken))
		if err != nil {
			return nil, err
		}

		body := resp.Body.(api.JSON)
		code, err := body.GetInt("code")
		if err != nil {
			return nil, err
		}

		if code != 0 {
			return nil, fmt.Errorf("cannot create category %s", c)
		}

		data, err := body.GetJSON("data")
		if err != nil {
			return nil, err
		}

		id, err := data.GetString("id")
		if err != nil {
			return nil, err
		}

		categoryMap[c] = id
	}

	return categoryMap, nil
}

func getExistingTemplate(ctx context.Context, apiEndpoint, accessToken string) (map[string]any, error) {
	apiGenerator := api.NewGenerator()

	resp, err := apiGenerator.New(apiEndpoint, "/getTemplates").
		GET(ctx, api.OAuth2("Bearer", accessToken))
	if err != nil {
		return nil, err
	}

	body := resp.Body.(api.JSON)
	code, err := body.GetInt("code")
	if err != nil {
		return nil, err
	}

	if code != 0 {
		return nil, fmt.Errorf("cannot list category, %d", code)
	}

	data, err := body.GetJSON("data")
	if err != nil {
		return nil, err
	}

	templates, err := data.GetArray("templates")
	if err != nil {
		return nil, err
	}

	existingTemplate := map[string]any{}
	for _, t := range templates {
		title, err := t.GetString("title")
		if err != nil {
			return nil, err
		}

		existingTemplate[title] = nil
	}

	return existingTemplate, nil
}
