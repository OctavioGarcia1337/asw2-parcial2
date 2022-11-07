package dto

type ResponseDto struct {
	NumFound int      `json:"numFound"`
	Docs     ItemsDto `json:"docs"`
}

type SolrResponseDto struct {
	Response ResponseDto `json:"response"`
}

type SolrResponsesDto []SolrResponseDto
