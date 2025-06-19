package schedule

type Server struct {
	Name     string
	Capacity int      // 最大容量
	Load     int      // 当前负载
	Prefs    []string // 服务器对用户的偏好排序
}

type User struct {
	Name     string
	Priority int
	//Weight   int // 后续扩展成用户权重
	Prefs   []string // 用户对服务器的偏好排序
	Matched bool     // 是否已匹配
}

// man first
func StableMarriage(manPref map[string][]string, womanPref map[string][]string) map[string]string {
	freeMan := make([]string, 0, len(manPref))
	for man := range manPref {
		freeMan = append(freeMan, man)
	}
	engaged := make(map[string]string)
	womanRank := make(map[string]map[string]int)
	for woman, prefs := range womanPref { //small first
		rank := make(map[string]int)
		for i, man := range prefs {
			rank[man] = i
		}
		womanRank[woman] = rank
	}

	for len(freeMan) > 0 {
		man := freeMan[0]
		freeMan = freeMan[1:]
		for _, woman := range manPref[man] {
			if ori, ok := engaged[woman]; !ok {
				engaged[woman] = man
				break
			} else if womanRank[woman][man] < womanRank[woman][ori] {
				engaged[woman] = man
				freeMan = append(freeMan, ori)
			}
		}
	}

	rets := make(map[string]string)
	for woman, man := range engaged {
		rets[man] = woman
	}

	return rets
}

// user first, N != M
func StableTrafficMatch(users []*User, servers []*Server) (map[string]string, []string) {
	freeUsers := make([]*User, 0, len(users))
	copy(freeUsers, users)
	serversMap := make(map[string]*Server)
	for _, server := range servers {
		serversMap[server.Name] = server
	}
	matched := make(map[string]string)
	serverRank := make(map[string]map[string]int)
	for _, server := range servers {
		rank := make(map[string]int)
		for i, userName := range server.Prefs {
			rank[userName] = i
		}
		serverRank[server.Name] = rank
	}

	for len(freeUsers) > 0 {
		user := freeUsers[0]
		freeUsers = freeUsers[1:]
		for _, serverName := range user.Prefs {
			if server, ok := serversMap[serverName]; ok {
				if _, ok = matched[user.Name]; !ok || server.Load < server.Capacity {
					server.Load++
					matched[user.Name] = server.Name
					user.Matched = true
					break
				} else {
					for curUserName, curServerName := range matched {
						if curServerName == server.Name && serverRank[serverName][user.Name] < serverRank[serverName][curUserName] {
							delete(matched, curUserName)
							matched[user.Name] = server.Name
							user.Matched = true
							break
						}
					}
				}
			}
		}
	}
	return nil, nil
}

// 稳定流量匹配（支持 N ≠ M 和容量限制），用户优先
func StableTrafficMatch2(users []*User, servers []*Server) (map[string]string, []string) {
	matches := make(map[string]string) // 匹配结果 {用户名: 服务器名}
	unmatchedUsers := []string{}       // 未匹配的用户列表

	// 预处理：为服务器建立用户偏好排名
	serverRank := make(map[string]map[string]int)
	for _, s := range servers {
		rank := make(map[string]int)
		for i, userName := range s.Prefs {
			rank[userName] = i
		}
		serverRank[s.Name] = rank
	}

	// 主循环：直到所有用户完成匹配或尝试完毕
	for {
		allMatchedOrExhausted := true
		for _, user := range users {
			if user.Matched {
				continue
			}

			allMatchedOrExhausted = false
			// 遍历用户的服务器偏好
			for _, serverName := range user.Prefs {
				targetServer := findServer(servers, serverName)
				if targetServer == nil {
					continue // 服务器不存在，跳过
				}

				// 检查服务器容量
				if targetServer.Load < targetServer.Capacity {
					// 容量足够：直接匹配
					matches[user.Name] = targetServer.Name
					targetServer.Load++
					user.Matched = true
					break
				} else {
					// 容量不足：检查是否比已匹配用户更优
					for matchedUserName, matchedServerName := range matches {
						if matchedServerName == targetServer.Name {
							currentRank := serverRank[targetServer.Name][matchedUserName]
							newRank := serverRank[targetServer.Name][user.Name]
							if newRank < currentRank {
								// 替换匹配
								delete(matches, matchedUserName)
								matches[user.Name] = targetServer.Name
								user.Matched = true
								// 将原用户标记为未匹配
								for _, u := range users {
									if u.Name == matchedUserName {
										u.Matched = false
										break
									}
								}
								break
							}
						}
					}
				}
			}
		}

		// 终止条件：所有用户已匹配或无法进一步匹配
		if allMatchedOrExhausted {
			break
		}
	}

	// 收集未匹配用户
	for _, user := range users {
		if !user.Matched {
			unmatchedUsers = append(unmatchedUsers, user.Name)
		}
	}

	return matches, unmatchedUsers
}

// 根据名称查找服务器
func findServer(servers []*Server, name string) *Server {
	for _, s := range servers {
		if s.Name == name {
			return s
		}
	}
	return nil
}
