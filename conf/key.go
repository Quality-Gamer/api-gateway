package conf

const GateWay = "gt"
const MServices = "ms"
const Count = "ct"
const Stop = "st"

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