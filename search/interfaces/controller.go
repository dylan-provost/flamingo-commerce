package interfaces

import (
	"context"
	"net/url"

	"flamingo.me/flamingo-commerce/v3/search/application"
	"flamingo.me/flamingo-commerce/v3/search/domain"
	"flamingo.me/flamingo-commerce/v3/search/utils"
	"flamingo.me/flamingo/v3/framework/web"
)

type (
	// ViewController demonstrates a search view controller
	ViewController struct {
		Responder             *web.Responder               `inject:""`
		SearchService         *application.SearchService   `inject:""`
		PaginationInfoFactory *utils.PaginationInfoFactory `inject:""`
	}

	viewData struct {
		SearchMeta     domain.SearchMeta
		SearchResult   map[string]*application.SearchResult
		PaginationInfo utils.PaginationInfo
	}
)

// Get Response for search
func (vc *ViewController) Get(c context.Context, r *web.Request) web.Result {
	query, _ := r.Query1("q")

	vd := viewData{
		SearchMeta: domain.SearchMeta{
			Query: query,
		},
	}

	searchRequest := application.SearchRequest{
		Query: query,
	}
	searchRequest.AddAdditionalFilters(domain.NewKeyValueFilters(r.QueryAll())...)

	if typ, ok := r.Params["type"]; ok {
		searchResult, err := vc.SearchService.FindBy(c, typ, searchRequest)
		if err != nil {
			if re, ok := err.(*domain.RedirectError); ok {
				u, _ := url.Parse(re.To)
				return vc.Responder.URLRedirect(u).Permanent()
			}

			return vc.Responder.ServerError(err)
		}
		vd.SearchMeta = searchResult.SearchMeta
		vd.SearchMeta.Query = query
		vd.SearchResult = map[string]*application.SearchResult{typ: searchResult}
		vd.PaginationInfo = vc.PaginationInfoFactory.Build(
			searchResult.SearchMeta.Page,
			searchResult.SearchMeta.NumResults,
			searchRequest.PageSize,
			searchResult.SearchMeta.NumPages,
			r.Request().URL,
		)
		return vc.Responder.Render("search/"+typ, vd)
	}

	searchResult, err := vc.SearchService.Find(c, searchRequest)
	if err != nil {
		if re, ok := err.(*domain.RedirectError); ok {
			u, _ := url.Parse(re.To)
			return vc.Responder.URLRedirect(u).Permanent()
		}

		return vc.Responder.ServerError(err)
	}
	vd.SearchResult = searchResult
	return vc.Responder.Render("search/search", vd)
}
