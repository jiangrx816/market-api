package service

import "log"

type CacheService struct {
}

func (cs *CacheService) GetCateInfo()  {
	log.Printf("这里是缓存的时间")
}