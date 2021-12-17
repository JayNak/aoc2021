package aoc

import (
	"math"
	"strconv"

	"github.com/jaynak/aoc2021/pkg/util"
)

func Day16(path string) (int, int) {

	lines := util.ReadToStrings(path)

	pkt_string := PacketToBinaryString(lines[0])

	pkts := ReadPackets(pkt_string, 1)

	cnt := 0
	val := 0
	for _, p := range pkts {
		cnt += p.ver

		if p.parent == nil {
			val = int(p.Evaluate())
		}
	}

	return cnt, val
}

type packet struct {
	ver      int
	typ      int
	val      int64
	parent   *packet
	children []*packet
	source   string
}

func (p *packet) Evaluate() int64 {
	var ret int64

	switch p.typ {
	case 0:
		for _, c := range p.children {
			ret += c.Evaluate()
		}
	case 1:
		for i, c := range p.children {
			if i == 0 {
				ret = c.Evaluate()
			} else {
				ret *= c.Evaluate()
			}
		}
	case 2:
		ret = math.MaxInt64
		for _, c := range p.children {
			v := c.Evaluate()
			if v < ret {
				ret = v
			}
		}
	case 3:
		ret = math.MinInt64
		for _, c := range p.children {
			v := c.Evaluate()
			if v > ret {
				ret = v
			}
		}
	case 4:
		ret = p.val
	case 5:
		ret = 0
		if p.children[0].Evaluate() > p.children[1].Evaluate() {
			ret = 1
		}
	case 6:
		ret = 0
		if p.children[0].Evaluate() < p.children[1].Evaluate() {
			ret = 1
		}
	case 7:
		ret = 0
		if p.children[0].Evaluate() == p.children[1].Evaluate() {
			ret = 1
		}
	}

	return ret
}

func PacketToBinaryString(source string) string {
	ret := ""

	ugly_map := map[rune]string{
		'0': "0000", '1': "0001", '2': "0010", '3': "0011",
		'4': "0100", '5': "0101", '6': "0110", '7': "0111",
		'8': "1000", '9': "1001", 'A': "1010", 'B': "1011",
		'C': "1100", 'D': "1101", 'E': "1110", 'F': "1111",
	}

	for _, r := range source {
		ret += ugly_map[r]
	}

	return ret
}

func ReadPackets(binary_packets string, num_packets int64) []*packet {
	ptr := 0

	var p *packet

	packets := []*packet{}
	pkt_count := 0

	for ptr < len(binary_packets) && int64(pkt_count) < num_packets {

		pkt_start := ptr
		pkt_count++

		v, err := strconv.ParseInt(binary_packets[ptr:ptr+3], 2, 64)
		if err != nil {
			panic(err)
		}

		t, err := strconv.ParseInt(binary_packets[ptr+3:ptr+6], 2, 64)
		if err != nil {
			panic(err)
		}

		p = &packet{
			ver: int(v),
			typ: int(t),
		}

		// Advance pointer beyond the version and type
		ptr = ptr + 6

		if p.typ == 4 {
			// Binary number
			num_string := ""
			for p.val == 0 {
				num_string += binary_packets[ptr+1 : ptr+5]
				if binary_packets[ptr] == '0' {
					// This is the exit case - calculate remaining bits to ignore if instructed
					p.source = binary_packets[pkt_start : ptr+5]
					v, err := strconv.ParseInt(num_string, 2, 64)
					if err != nil {
						panic(err)
					}
					p.val = v
				}

				// Advance pointer beyond this 5 bits
				ptr = ptr + 5
			}

		} else {
			// Operators
			offset := 0

			if binary_packets[ptr] == '0' {
				// 15 bits beyond the indicator is the # of bits that comprise the subpackets
				num_bits, err := strconv.ParseInt(binary_packets[ptr+1:ptr+16], 2, 64)
				if err != nil {
					panic(err)
				}

				// Advance beyond the # bits
				ptr = ptr + 16

				// Recurse, rely on the length of the sub-packets to know when to quit
				children := ReadPackets(binary_packets[ptr:ptr+int(num_bits)], math.MaxInt64)

				// In this case the offset is the # of bits
				offset = int(num_bits)

				// Set the parent-child relationships
				for _, c := range children {
					packets = append(packets, c)
					if c.parent != nil {
						continue
					}
					c.parent = p
					p.children = append(p.children, c)

				}

			} else {

				// 11 bits beyond the header is the # of sub-packets
				num_pkts, err := strconv.ParseInt(binary_packets[ptr+1:ptr+12], 2, 64)
				if err != nil {
					panic(err)
				}

				// Advance beyond the # of sub-packets
				ptr = ptr + 12

				// Recurse, using the # of packets to know when to stop
				children := ReadPackets(binary_packets[ptr:], num_pkts)

				for _, c := range children {
					packets = append(packets, c)

					if c.parent != nil {
						continue
					}

					c.parent = p
					p.children = append(p.children, c)

					// offset is the sum of all children
					offset += len(c.source)
				}

			}

			// Advance the pointer beyond this set of sub-packets
			ptr = ptr + offset

			// Set the source for this packet in case we need it
			p.source = binary_packets[pkt_start:ptr]

		}

		// Add this packet to the response
		packets = append(packets, p)
	}

	return packets
}

// XXXYYY
// XXX = Version
// YYY = Type -> 4 = literal (binary number)
//		 Anything else is an operator
