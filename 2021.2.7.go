package leetcode_go


/**
 * 990. 等式方程的可满足性
 * 给定一个由表示变量之间关系的字符串方程组成的数组，每个字符串方程 equations[i] 的长度为 4，并采用两种不同的形式之一："a==b" 或 "a!=b"。在这里，a 和 b 是小写字母（不一定不同），表示单字母变量名。
 * 只有当可以将整数分配给变量名，以便满足所有给定的方程时才返回 true，否则返回 false。 
 */

func equationsPossible(equations []string) bool {
	eqm := make(map[byte]byte)
	var neql []string

	for _, e := range equations {
		if e[1] == '=' {
			union(eqm, e[0], e[3])
		} else {
			if e[0] == e[3]{
				return false
			}
			neql = append(neql, e)
		}
	}
	for _, neq := range neql {
		p1, p2 := find(eqm, neq[0]), find(eqm, neq[3])
		if !(p1 == 0 || p2 == 0 || p1 != p2)  {
			return false
		}
	}
	return true
}

func union(eqm map[byte]byte, v1, v2 byte) {
	p1 := find(eqm, v1)
	p2 := find(eqm, v2)
	if p1 == 0 && p2 == 0 {
		eqm[v1] = v1
		eqm[v2] = v1
	} else if p1 != 0 && p2 == 0 {
		eqm[v2] = p1
	} else if p1 == 0 && p2 != 0 {
		eqm[v1] = p2
	} else if p1 != 0 && p2 !=0 {
		eqm[p2] = p1
		//这里还需要路径压缩；路径压缩的时候必须知道所有以p为父的节点
	}
}

func find(eqm map[byte]byte, v byte) byte {
	for {
		if p, ok := eqm[v]; ok {
			if p == v {
				return p
			} else {
				v = p
			}
		} else {
			return 0
		}
	}
}

