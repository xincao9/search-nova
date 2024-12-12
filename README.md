# search-nova

## 文心一言

### elasticsearch索引

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