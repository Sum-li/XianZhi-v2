闲智（闲置）微信小程序接口测试

GET 获取分类 {.pm-h2 .docs-desc-title .docs-desc-title--request}
------------

127.0.0.1:8081/goods/category

获取商品的所有分类

``` {.pm-snippet-body}
curl --location --request GET "127.0.0.1:8081/goods/category"
```

GET 商品首页 {.pm-h2 .docs-desc-title .docs-desc-title--request}
------------

127.0.0.1:8081/goods/index?page\_count=1

商品首页获取商品的列表

#### Params {.pm-h4}

------------- ---
  page\_count   1
------------- ---



``` {.pm-snippet-body}
curl --location --request GET "127.0.0.1:8081/goods/index?page_count=1"
```

GET 分类商品页 {.pm-h2 .docs-desc-title .docs-desc-title--request}
--------------

127.0.0.1:8081/goods/goodscategory?category\_id=4&page\_count=1

获取分类商品页面的商品列表

#### Params {.pm-h4}

-------------- ---
  category\_id   4
  page\_count    1
-------------- ---



``` {.pm-snippet-body}
curl --location --request GET "127.0.0.1:8081/goods/goodscategory?category_id=4&page_count=1"
```

GET 商品详情 {.pm-h2 .docs-desc-title .docs-desc-title--request}
------------

127.0.0.1:8081/goods/detail?goods\_id=54&user\_id=5

获取商品的详细信息以及卖家的基本信息

#### Params {.pm-h4}

----------- ----
  goods\_id   54
  user\_id    5
----------- ----



``` {.pm-snippet-body}
curl --location --request GET "127.0.0.1:8081/goods/detail?goods_id=54&user_id=5"
```

GET 已买商品 {.pm-h2 .docs-desc-title .docs-desc-title--request}
------------

127.0.0.1:8081/goods/havebuy?user\_id=775&page\_count=1

获取用户买到的商品列表

#### Params {.pm-h4}

------------- -----
  user\_id      775
  page\_count   1
------------- -----



``` {.pm-snippet-body}
curl --location --request GET "127.0.0.1:8081/goods/havebuy?user_id=775&page_count=1"
```

GET 收藏商品 {.pm-h2 .docs-desc-title .docs-desc-title--request}
------------

127.0.0.1:8081/goods/wantbuy?user\_id=775&page\_count=1

获取用户想买的商品列表

#### Params {.pm-h4}

------------- -----
  user\_id      775
  page\_count   1
------------- -----



``` {.pm-snippet-body}
curl --location --request GET "127.0.0.1:8081/goods/wantbuy?user_id=775&page_count=1"
```

GET 发布的商品 {.pm-h2 .docs-desc-title .docs-desc-title--request}
--------------

127.0.0.1:8081/goods/mypub?user\_id=2&page\_count=1

获取用户发布的商品列表

#### Params {.pm-h4}

------------- ---
  user\_id      2
  page\_count   1
------------- ---



``` {.pm-snippet-body}
curl --location --request GET "127.0.0.1:8081/goods/mypub?user_id=2&page_count=1"
```

GET 卖出的商品 {.pm-h2 .docs-desc-title .docs-desc-title--request}
--------------

127.0.0.1:8081/goods/sold?user\_id=2&page\_count=1

获取用户卖出的商品列表

#### Params {.pm-h4}

------------- ---
  user\_id      2
  page\_count   1
------------- ---



``` {.pm-snippet-body}
curl --location --request GET "127.0.0.1:8081/goods/sold?user_id=2&page_count=1"
```

GET 活动 {.pm-h2 .docs-desc-title .docs-desc-title--request}
--------

127.0.0.1:8081/goods/active?active=gd&page\_count=1

获取在活动中的商品列表

#### Params {.pm-h4}

------------- ----
  active        gd
  page\_count   1
------------- ----



``` {.pm-snippet-body}
curl --location --request GET "127.0.0.1:8081/goods/active?active=gd&page_count=1"
```
