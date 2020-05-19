package conf

const GateWay = "gt"
const MServices = "ms"
const Count = "ct"
const Stop = "st"
const Cache = "cc"
const Key = "ky"
const QualityGamer = "qg"

func getDefaultKey() string {
	return GateWay + ":" + MServices
}

func GetMicroserviceKeyCount(ms string) string {
	return getDefaultKey() + ":" + ms + ":" + Count
}

func GetMicroserviceActionKeyCount(ms,action string) string {
	return getDefaultKey() + ":" + ms + ":" + action + ":" + Count
}

func GetStoppedMicroserviceKey() string {
	return getDefaultKey() + ":" + Stop
}

func GetCacheKey(ms,hash string) string {
	return getDefaultKey() + ":" + Cache + ":" + ms + ":" + hash
}

func GetCacheCountKey(ms,hash string) string {
	return getDefaultKey() + ":" + Cache + ":" + ms + ":" + hash + ":" + Count
}

func GetAuthKey() string {
	return getDefaultKey() + ":" + Key + ":" + QualityGamer
}