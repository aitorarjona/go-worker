package trigger

import (
	"encoding/json"
	cloudevents "github.com/cloudevents/sdk-go"
	log "github.com/sirupsen/logrus"
	"sync/atomic"
)

type Join struct {
	Join    int    `json:"join"`
	Counter uint32 `json:"counter"`
}

func PassDataParser(rawData []byte) (interface{}, error) {
	return nil, nil
}

func JoinDataParser(rawData []byte) (interface{}, error) {
	data := Join{}
	err := json.Unmarshal(rawData, &data)
	if err != nil {
		return nil, err
	} else {
		return &data, nil
	}
}

func JoinCondition(context *Context, event cloudevents.Event) (bool, error) {
	parsedData := (*context).ConditionParsedData.(*Join)
	cnt := int(atomic.AddUint32(&parsedData.Counter, 1))
	return cnt >= parsedData.Join, nil
}

func TrueCondition(context *Context, event cloudevents.Event) (bool, error) {
	return true, nil
}

func CounterThresholdCondition(context *Context, event cloudevents.Event) (bool, error) {
	//(*context).Counter++
	//
	//totalActivations := 0
	//if val, ok := (*context).Data["threshold"]; ok {
	//	totalActivations = val.(int)
	//}
	//
	//joined := (*context).Counter >= totalActivations
	//
	//return joined
	return false, nil
}

func PassAction(context *Context, event cloudevents.Event) error {
	return nil
}

func TerminateAction(context *Context, event cloudevents.Event) error {
	// TODO implement worker halt
	//fmt.Println(context.ConditionParsedData.(Join).Counter, time.Now().UTC().UnixNano())
	//fmt.Println(time.Now().UTC().UnixNano())
	log.Infof("Terminate worker call")
	return nil
}