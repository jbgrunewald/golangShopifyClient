package shopify

type Collection struct {
	BodyHtml                    string       `json:"body_html"`
	Handle                      string       `json:"handle"`
	Image                       Image        `json:"image"`
	Id                          int          `json:"id"`
	MetaFields                  []MetaFields `json:"metafields"`
	Published                   bool         `json:"published"`
	PublishedAt                 string       `json:"published_at"`
	PublishedScope              string       `json:"published_scope"` // TODO this could be an enum value
	SortOrder                   string       `json:"sort_order"`      // TODO this could be an enum value
	TemplateSuffix              string       `json:"template_suffix"`
	Title                       string       `json:"title"`
	UpdatedAt                   string       `json:"updated_at"`
	Rules                       Rules        `json:"rules,omitempty"`
	Disjunctive                 bool         `json:"disjunctive,omitempty"`
	ProductsManuallySortedCount int          `json:"products_manually_sorted_count,omitempty"`
}

type Image struct {
	Src      string `json:"src"`
	Alt      string `json:"alt"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	CreateAt string `json:"create_at"`
}

type MetaFields struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	ValueType string `json:"value_type"`
	Namespace string `json:"namespace"`
}

// TODO these could be enum values
type Rules struct {
	Column    string `json:"column"`
	Relation  string `json:"relation"`
	Condition string `json:"condition"`
}
