/*
* @Author: lvtongbin
* @Date:   2018-10-07 09:44:44
* @Last Modified by:   lvtongbin
* @Last Modified time: 2018-10-07 09:44:44
 */

package myredis

import (
	"encoding/json"
	"trism/cache"
	_ "trism/cache/redis" // import redis driver

	"github.com/astaxie/beego/logs"
)

var redisCache cache.Cache

func init() {
	// Connect redis
	cf := map[string]string{"conn": "127.0.0.1:6379"}
	b, err := json.Marshal(cf)
	redisCache, err = cache.NewCache("redis", string(b))
	if err != nil {
		logs.Error("Connect redis faild!", err)
	}
}
