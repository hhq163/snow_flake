package gen_id

const (
	MachineIdBits uint8 = 10 // the biggest is 1024
	IdBits uint8 = 12 // the maximum every millisecond generate is 4096
	NodeIDMax   int64 = -1 ^ (-1 << workerBits) // the maximum of node id
	IdMask  int64 = -1 ^ (-1 << numberBits) // the maximum of id, as mask off code
	TimeShift   uint8 = workerBits + numberBits // timestamp left offset
	NodeIdShift uint8 = numberBits              // node id left offset
	Epoch int64 = 1584583903000 // this is the time I write this lib(millisecond)
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

func (g *Generator) GetId() int64 {
	g.mu.Lock()
	defer g.mu.Unlock()

	now := time.Now().UnixNano() / 1e6 // 纳秒转毫秒
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