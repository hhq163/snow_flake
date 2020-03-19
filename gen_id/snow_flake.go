package gen_id

import(
	"time"
	"fmt"
	"sync"
	"errors"
)

type Generator struct {
	mu        sync.Mutex 
	timestamp int64      
	nodeId  int64      
	id    int64  
}

func New(nodeId int64) (*Generator, error) {
	if nodeId < 0 || nodeId > NodeIDMax {
		return nil, errors.New(fmt.Sprintf("Node ID out of limit %d", NodeIDMax))
	}
	return &Generator{
		timestamp: 0,
		nodeId:  nodeId,
		id:    0,
	}, nil
}

func (g *Generator) GenId() int64 {
	g.mu.Lock()
	defer g.mu.Unlock()

	now := time.Now().UnixNano() / 1e6 // nanosecond to millisecond
	if g.timestamp == now {
		g.id = (g.id + 1) & IdMask

		if g.id == 0 { //wait 1 millisecond
			for now <= g.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		g.id = 0
	}
	g.timestamp = now

	id := int64((now-Epoch)<<TimeShift | (g.nodeId << NodeIdShift) | (g.id))
	return id
}
//解析ID
func (g *Generator) PraseId(id int64) (int64, int32, int32, error){
	if id <=0 {
		return 0, 0, 0, errors.New("id is not valid")
	}
	timestamp := id>>TimeShift
	timestamp += Epoch

	nodeMask := NodeIDMax << NodeIdShift
	nodeId := (id & nodeMask)>>IdBits

	idNumber := id&IdMask
	return timestamp, int32(nodeId), int32(idNumber), nil
}