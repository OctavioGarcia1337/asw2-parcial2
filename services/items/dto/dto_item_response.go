package dto

type ItemResponseDto struct {
	ItemId      string `json:"id"`
	Titulo      string `json:"titulo"`
	Tipo        string `json:"tipo"`
	Ubicacion   string `json:"ubicacion"`
	PrecioBase  int    `json:"precio_base"`
	Vendedor    string `json:"vendedor"`
	Barrio      string `json:"barrio"`
	Descripcion string `json:"descripcion"`
	Dormitorios int    `json:"dormitorios"`
	Banos       int    `json:"banos"`
	Mts2        int    `json:"mts2"`
	Ambientes   int    `json:"ambientes"`
	UrlImg      string `json:"url_img"`
	Expensas    int    `json:"expensas"`
	UsuarioId   int    `json:"usuario_id"`
	Usuario     string `json:"usuario"`
	UNombre     string `json:"usuario_nombre"`
	UApellido   string `json:"usuario_apellido"`
	UEmail      string `json:"usuario_email"`
}
type ItemsResponseDto []ItemResponseDto
