package plugins

// ElasticsearchOutput CRD name
const CustomFilter = "customfilter"

// ElasticsearchDefaultValues for Elasticsearch output plugin
var CustomFilterDefaultValues = map[string]string{
	"type":         "",
	"filter":       "",
}

// ElasticsearchTemplate for Elasticsearch output plugin
const CustomFilterTemplate = `
<filter {{ .pattern }}.**>
  @type {{ .type }}
  {{ .filter }}
</filter>`
