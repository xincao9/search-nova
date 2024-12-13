# search-nova


## Install the browsers and OS dependencies

```bash
go install github.com/playwright-community/playwright-go/cmd/playwright@latest
playwright install --with-deps
```


## elasticsearch

```text
PUT /search_nova
{
  "mappings": {
    "properties": {
      "id": {
        "type": "long"
      },
      "url": {
        "type": "keyword"
      },
      "content": {
        "type": "text",
        "analyzer":"ik_max_word"
      }
    }
  }
}

```