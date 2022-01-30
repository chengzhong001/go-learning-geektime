package objpool

import (
	"errors"
	"time"
)

type ReusableObj struct {
}

type ObjPool struct {
	buffChan chan *ReusableObj
}

func NewOjbPool(numOfObj int) *ObjPool {
	objPool := ObjPool{}
	objPool.buffChan = make(chan *ReusableObj, numOfObj)
	for i := 0; i < numOfObj; i++ {
		objPool.buffChan <- &ReusableObj{}
	}
	return &objPool
}

func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <-p.buffChan:
		return ret, nil
	case <-time.After(timeout):
		return nil, errors.New("time out")
	}
}

func (p *ObjPool) RelaseObj(obj *ReusableObj) error {
	select {
	case p.buffChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}
