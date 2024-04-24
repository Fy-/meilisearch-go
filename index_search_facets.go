package meilisearch

import (
	"errors"
	"net/http"
)

type FacetSearchRequest struct {
	FacetName            string   `json:"facetName"`                      // Optional: Facet name to search values on
	FacetQuery           string   `json:"facetQuery,omitempty"`           // Optional: Search query for a given facet value
	Filter               string   `json:"filter,omitempty"`               // Optional: Filter queries by an attribute's value
	MatchingStrategy     string   `json:"matchingStrategy"`               // Strategy used to match query terms within documents, defaults to "last"
	AttributesToSearchOn []string `json:"attributesToSearchOn,omitempty"` // Optional: Restrict search to the specified attributes
}
type FacetHit struct {
	Value string `json:"value"` // Facet value matching the facetQuery
	Count int    `json:"count"` // Number of documents with a facet value matching value
}

type FacetSearchResponse struct {
	FacetHits        []FacetHit `json:"facetHits"`        // Array of facet hits
	FacetQuery       string     `json:"facetQuery"`       // The original facetQuery
	ProcessingTimeMs int        `json:"processingTimeMs"` // Processing time of the query in milliseconds
}

var ErrNoFacetName = errors.New("no search request provided")

func (i Index) SearchFacets(query string, request *FacetSearchRequest) (*FacetSearchResponse, error) {
	if request == nil {
		return nil, ErrNoSearchRequest
	}
	if request.FacetName == "" {
		return nil, ErrNoFacetName
	}

	searchPostRequestParams := facetSearchPostRequestParams(query, request)
	resp := &FacetSearchResponse{}
	req := internalRequest{
		endpoint:            "/indexes/" + i.UID + "/facet-search",
		method:              http.MethodPost,
		contentType:         contentTypeJSON,
		withRequest:         searchPostRequestParams,
		withResponse:        resp,
		acceptedStatusCodes: []int{http.StatusOK},
		functionName:        "Search",
	}

	if err := i.client.executeRequest(req); err != nil {
		return nil, err
	}

	return resp, nil
}

func facetSearchPostRequestParams(query string, request *FacetSearchRequest) map[string]interface{} {
	params := make(map[string]interface{}, 6)

	if query != "" {
		params["q"] = query
	}

	if request.FacetName != "" {
		params["facetName"] = request.FacetName
	}
	if request.FacetQuery != "" {
		params["facetQuery"] = request.FacetQuery
	}
	if request.Filter != "" {
		params["filter"] = request.Filter
	}
	if request.MatchingStrategy != "" {
		params["matchingStrategy"] = request.MatchingStrategy
	}
	if len(request.AttributesToSearchOn) > 0 {
		params["attributesToSearchOn"] = request.AttributesToSearchOn
	}

	return params
}
