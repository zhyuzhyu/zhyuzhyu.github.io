package schedule

import (
	"fmt"
)

// 稳定流量匹配算法（带容量限制）
func stableTrafficMatch(users []*User, servers []*Server) map[string]string {
	// 初始化
	freeUsers := make([]*User, len(users))
	copy(freeUsers, users)
	matches := make(map[string]string) // 匹配结果 {用户名: 服务器名}

	// 预处理：为服务器建立用户偏好排名（值越小优先级越高）
	serverRank := make(map[string]map[string]int)
	for _, s := range servers {
		rank := make(map[string]int)
		for i, userName := range s.Prefs {
			rank[userName] = i
		}
		serverRank[s.Name] = rank
	}

	// 主循环：直到所有用户完成匹配
	for len(freeUsers) > 0 {
		user := freeUsers[0] // 取出一个未匹配用户
		freeUsers = freeUsers[1:]

		// 遍历用户的服务器偏好列表
		for _, serverName := range user.Prefs {
			// 找到目标服务器
			var targetServer *Server
			for _, s := range servers {
				if s.Name == serverName {
					targetServer = s
					break
				}
			}
			if targetServer == nil {
				continue // 服务器不存在，跳过
			}

			// 检查服务器容量
			if targetServer.Load < targetServer.Capacity {
				// 容量足够：直接匹配
				matches[user.Name] = targetServer.Name
				targetServer.Load++
				break
			} else {
				// 容量不足：检查是否比当前已匹配用户更优
				for matchedUserName, matchedServerName := range matches {
					if matchedServerName == targetServer.Name {
						// 比较服务器对这两个用户的偏好
						currentRank := serverRank[targetServer.Name][matchedUserName]
						newRank := serverRank[targetServer.Name][user.Name]
						if newRank < currentRank {
							// 服务器更偏好新用户：替换匹配
							delete(matches, matchedUserName)
							matches[user.Name] = targetServer.Name
							freeUsers = append(freeUsers, &User{Name: matchedUserName}) // 原用户重新加入队列
							break
						}
					}
				}
			}
		}
	}
	return matches
}

func main() {
	// 示例：初始化用户和服务器
	users := []*User{
		{Name: "User1", Prefs: []string{"ServerA", "ServerB", "ServerC"}},
		{Name: "User2", Prefs: []string{"ServerA", "ServerC", "ServerB"}},
		{Name: "User3", Prefs: []string{"ServerB", "ServerA", "ServerC"}},
	}

	servers := []*Server{
		{
			Name:     "ServerA",
			Capacity: 1, // 容量为1
			Load:     0,
			Prefs:    []string{"User1", "User2", "User3"}, // 服务器偏好用户顺序
		},
		{
			Name:     "ServerB",
			Capacity: 2,
			Load:     0,
			Prefs:    []string{"User3", "User1", "User2"},
		},
		{
			Name:     "ServerC",
			Capacity: 1,
			Load:     0,
			Prefs:    []string{"User2", "User3", "User1"},
		},
	}

	// 执行匹配
	result := stableTrafficMatch(users, servers)

	// 打印匹配结果
	fmt.Println("用户-服务器匹配结果（带容量限制）:")
	for userName, serverName := range result {
		fmt.Printf("用户 %s → 服务器 %s\n", userName, serverName)
	}

	// 打印服务器最终负载
	fmt.Println("\n服务器负载状态:")
	for _, s := range servers {
		fmt.Printf("服务器 %s: 负载 %d/%d\n", s.Name, s.Load, s.Capacity)
	}
}
