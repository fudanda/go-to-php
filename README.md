# go-to-php

用Golang封装PHP常用的函数，为了区分PHP与Go中的函数，go-to-php封装的函数，都是在对应的PHP函数的前面加上大写的`P`。`Pallord`函数除外

#### 安装

1. 字符串相关包

`go get -u github.com/Echo-Mr-Pengw/go-to-php/str`

2. 数组相关包

`开发中`

.......

#### 字符串相关函数

| go-to-php封装函数                                           | 对应的PHP函数             |
| ----------------------------------------------------------- | ------------------------- |
| Pimplode(`glue` string , `pieces` interface{})  string      | implode                   |
| Pjoin(`glue` string , `pieces` interface{})  string         | implode的别名，同Pimplode |
| Pexplode(`delimiter` , `str` string , `limit` int) []string | explode                   |
| Pstrlen(`str` string) int                                   | strlen                    |
| Pmb_strlen(`str` string) int                                | mb_strlen                 |
| Plcfirst(`str` string) string                               | lcfirst                   |
| Pucfirst(`str` string) string                               | ucfirst                   |
| Pstrtoupper(`str` string) string                            | strtoupper                |
| Pstrtolower(`str` string) string                            | strtolower                |
| Pucword(`str` string) string                                | ucword                    |
| Ptrim(`str` , `character_mask` string) string               | trim                      |
| Pltrim(`str` , `character_mask` string) string              | ltrim                     |
| Prtrim(`str` , `character_mask` string) string              | rtrim                     |
| Pchop(`str` , `character_mask` string) string               | rtrim的别名，同Prtrim     |
| Pmd5(`str` string) string                                   | md5                       |
| Psha1(`str` string) string                                  | sha1                      |
| Pord(`str` string) int                                      | ord                       |
| Pallord(`str` string) interface{}                           | 转化全部字符，ord的升级版 |
| Pchr(`ascii` int32) string                                  | chr                       |
| Pecho(`a ...`interface{})                                   | echo                      |
| Pvar_dump(`a ...`interface{})                               | var_dump                  |
| Pprint(`a` interface{})                                     | print                     |
| Pstr_repeat(`input` string , `multiplier` int) string       | str_repeat                |
|                                                             |                           |





