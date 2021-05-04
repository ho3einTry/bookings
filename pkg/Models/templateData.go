package Models

//TemplateDate holds data send from handler to template
type TemplateDate struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMp   map[string]float32
	Date      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
