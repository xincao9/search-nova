# NLP Project

### 部署

```shell
python3 install -r requirements.txt
python3 run.py
```

### 文本分析

```text
curl -X POST -H 'content-type:application/json;charset=utf-8' 'http://localhost:5000/analysis' -d '{"text":"自然语言处理是计算机科学领域与人工智能领域中的一个重要方向"}'
```

*结果*

```json
{
    "keyword": [
        "领域",
        "智能",
        "人工",
        "科学",
        "计算机"
    ],
    "summary": [
        "自然语言处理是计算机科学领域与人工智能领域中的一个重要方向"
    ]
}
```