package dto

type ResponseDto struct {
	NumFound int      `json:"numFound"`
	Docs     ItemsDto `json:"docs"`
}

type SolrResponseDto struct {
	Response ResponseDto `json:"response"`
}

type DocDto struct {
	Doc ItemDto `json:"doc"`
}

type AddDto struct {
	Add DocDto `json:"add"`
}

type SolrResponsesDto []SolrResponseDto
