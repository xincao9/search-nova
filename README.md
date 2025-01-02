# search-nova

### “Search”直接表明了项目的搜索引擎属性，而“Nova”在拉丁语中意为“新的”，象征着创新和前沿技术。

## 部署

### 依赖以下中间件

1. mysql
    ```sql
    CREATE
    DATABASE `search_nova` /*!40100 DEFAULT CHARACTER SET utf8 */;
    
    USE `search_nova`;
    DROP TABLE IF EXISTS `user`;
    CREATE TABLE `user`
    (
        `id`         bigint(11) unsigned NOT NULL AUTO_INCREMENT,
        `username`   char(32)  NOT NULL DEFAULT '',
        `password`   char(32)  NOT NULL DEFAULT '',
        `expire`     timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        `token`      char(128) NOT NULL DEFAULT '',
        `role`       int(11) NOT NULL DEFAULT '0',
        `created_at` timestamp NULL DEFAULT NULL,
        `deleted_at` timestamp NULL DEFAULT NULL,
        `updated_at` timestamp NULL DEFAULT NULL,
        PRIMARY KEY (`id`),
        UNIQUE KEY `username_idx` (`username`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8;
    CREATE TABLE `page`
    (
        `id`         bigint(11) unsigned NOT NULL AUTO_INCREMENT,
        `md5`        char(32)  NOT NULL DEFAULT '',
        `url`        text,
        `title`      text,
        `describe`   text,
        `keywords`   text,
        `status`     int(11) NOT NULL DEFAULT 0,
        `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        `deleted_at` timestamp NULL DEFAULT NULL,
        `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        PRIMARY KEY (`id`),
        UNIQUE KEY `md5_idx` (`md5`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
    
    insert into user (`username`, `password`, `role`)
    values ("admin", "21232f297a57a5a743894a0e4a801fc3", 3);
    
    insert into `page` (`url`, `md5`)
    values ("https://www.hao123.com", "2b90da495b2633aae6487c0f4c536adb");
    ```

2. elasticsearch 【推荐目前最新】
    ```json
    {
      "mappings": {
        "properties": {
          "id": {
            "type": "long"
          },
          "md5": {
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
3. [playwright](github.com/playwright-community/playwright-go/cmd/playwright@latest) 【浏览器驱动】

### 启动nlp服务
    
    ```shell
    cd nlp
    python3 install -r requirements.txt
    python3 run.py
    ```

### 启动爬虫服务

    ```shell
    go run cmd/crawler/main.go
    ```

### 启动查询服务

    ```shell
    go run cmd/main.go
    ```

### 相关配置

    ```yaml
    db:
      datasource: root:{替换你的密码}@tcp(localhost:3306)/search_nova?charset=utf8&parseTime=true&loc=Local
    logger:
      dir: /tmp/search-nova/log
      level: debug
    manager:
      server:
        port: 8090
    server:
      port: 8080
    public:
      dir: ./front/dist
    elasticsearch:
      password: "{替换你的密码}"
    nlp:
      endpoint: http://localhost:5000/analysis
    
    ```

## 依赖外部技术

* [https://github.com/isnowfy/snownlp](https://github.com/isnowfy/snownlp)
* [github.com/playwright-community/playwright-go/cmd/playwright@latest](github.com/playwright-community/playwright-go/cmd/playwright@latest)
* [https://flask.palletsprojects.com/en/stable/](https://flask.palletsprojects.com/en/stable/)
* [https://element-plus.org/zh-CN/](https://element-plus.org/zh-CN/)

