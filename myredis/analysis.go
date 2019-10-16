/*
* @Author: lvtongbin
* @Date:   2018-10-07 09:44:44
* @Last Modified by:   lvtongbin
* @Last Modified time: 2018-10-07 09:44:44
 */

package myredis

import (
	"trism/cache"

	"github.com/gomodule/redigo/redis"
)

// GetTotolScore 获取Key的score总和
func GetTotolScore(key string) (int, error) {
	members, err := redis.Strings(redisCache.Do("ZRANGE", key, 0, -1, "WITHSCORES"))
	if err != nil || members == nil || len(members) == 0 {
		return 0, err
	}
	total := 0
	for k, member := range members {
		if k%2 == 1 {
			total = total + cache.GetInt(member)
		}
	}
	return total, nil
}

// GetMemberScore 获取key某个member的score
func GetMemberScore(key, member string) (int, error) {
	return redis.Int(redisCache.Do("ZSCORE", key, member))
}

// EachMember is ...
type EachMember struct {
	Score int `json:"score"`
	Total int `json:"total"`
}

// GetEachScore 获取key的每一个score
func GetEachScore(key string) ([]EachMember, error) {
	members, err := redis.Values(redisCache.Do("ZREVRANGE", key, 0, -1, "WITHSCORES"))
	if err != nil || members == nil || len(members) == 0 {
		memberList := make([]EachMember, 0)
		return memberList, err
	}
	count := len(members) / 2
	memberList := make([]EachMember, 0, count)
	for i := 0; i < count; i++ {
		member := EachMember{cache.GetInt(members[2*i]), cache.GetInt(members[2*i+1])}
		memberList = append(memberList, member)
	}
	return memberList, nil
}
