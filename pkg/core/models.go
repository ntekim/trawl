package core

// MediaItem represents the structured data we want to find
type MediaItem struct {
	ID          string            `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	ImageURL    string            `json:"image_url"`
	SourceURL   string            `json:"source_url"`
	Metadata    map[string]string `json:"metadata"` // For things like "Year", "Genre", etc.
}

// ScrapeConfig defines the YAML structure for a scraping job
type ScrapeConfig struct {
	Name           string            `yaml:"name"`
	URL            string            `yaml:"url"`
	Driver         string            `yaml:"driver"` // "colly" or "chromedp"
	ItemContainer  string            `yaml:"item_container"`
	Selectors      map[string]string `yaml:"selectors"`
}