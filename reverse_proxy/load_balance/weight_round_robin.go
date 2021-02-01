package load_balance

import (
	"errors"
	"strconv"
	"strings"
)

type WeightRoundRobinBalance struct {
	curIndex int
	rss      []*WeightNode
	rsw      []int
	//观察主体
	conf LoadBalanceConf
}

type WeightNode struct {
	addr            string
	weight          int //权重值
	currentWeight   int //节点当前权重
	effectiveWeight int //有效权重
}

func (r *WeightRoundRobinBalance) Add(params ...string) error {
	if len(params) != 2 {
		return errors.New("param len need 2")
	}
	parInt, err := strconv.ParseInt(params[1], 10, 64)
	if err != nil {
		return err
	}
	node := &WeightNode{addr: params[0], weight: int(parInt)}
	node.effectiveWeight = node.weight
	r.rss = append(r.rss, node)
	return nil
}

func (r *WeightRoundRobinBalance) Next() string {
	total := 0
	var best *WeightNode
	for i := 0; i < len(r.rss); i++ {
		w := r.rss[i]
		//step 1 sum the total of effective weights
		total += w.effectiveWeight
		//step 2 update current weight by adding effective weight
		w.currentWeight += w.effectiveWeight
		//step 3 recover effective weight
		if w.effectiveWeight < w.weight {
			w.effectiveWeight++
		}
		//step 4 select maximum current weight
		if best == nil || w.currentWeight > best.currentWeight {
			best = w
		}
	}
	if best == nil {
		return ""
	}
	//step 5 update current weight by subtracting total effective weights
	best.currentWeight -= total
	return best.addr
}

func (r *WeightRoundRobinBalance) Get(key string) (string, error) {
	return r.Next(), nil
}

func (r *WeightRoundRobinBalance) SetConf(conf LoadBalanceConf) {
	r.conf = conf
}

func (r *WeightRoundRobinBalance) Update() {
	//if conf, ok := r.conf.(*LoadBalanceZkConf); ok {
	//	fmt.Println("WeightRoundRobinBalance get zk conf:", conf.GetConf())
	//	r.rss = nil
	//	for _, ip := range conf.GetConf() {
	//		r.Add(strings.Split(ip, ",")...)
	//	}
	//}
	if conf, ok := r.conf.(*LoadBalanceCheckConf); ok {
		//fmt.Println("WeightRoundRobinBalance get check conf:", conf.GetConf())
		r.rss = nil
		for _, ip := range conf.GetConf() {
			r.Add(strings.Split(ip, ",")...)
		}
	}
}
