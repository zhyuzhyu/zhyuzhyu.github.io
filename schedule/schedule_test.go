package schedule

import (
	"fmt"
	"testing"
)

func TestStableMarriage(t *testing.T) {
	// 示例：男性和女性的偏好列表
	menPref := map[string][]string{
		"Alex":  {"Eva", "Lena", "Sophia"},
		"Bob":   {"Lena", "Sophia", "Eva"},
		"Chris": {"Sophia", "Eva", "Lena"},
	}

	womenPref := map[string][]string{
		"Eva":    {"Alex", "Bob", "Chris"},
		"Lena":   {"Bob", "Chris", "Alex"},
		"Sophia": {"Chris", "Alex", "Bob"},
	}

	// 执行匹配
	matches := StableMarriage(menPref, womenPref)

	// 打印结果
	fmt.Println("稳定婚姻匹配结果:")
	for man, woman := range matches {
		fmt.Printf("%s \u2194 %s\n", man, woman)
	}
}

func TestStableTrafficMatch(t *testing.T) {
	// 示例1：用户数 > 服务器数 (N > M)
	users := []*User{
		{Name: "User1", Prefs: []string{"ServerA", "ServerB"}},
		{Name: "User2", Prefs: []string{"ServerA", "ServerB"}},
		{Name: "User3", Prefs: []string{"ServerB", "ServerA"}},
		{Name: "User4", Prefs: []string{"ServerB", "ServerA"}},
	}

	servers := []*Server{
		{Name: "ServerA", Capacity: 2, Prefs: []string{"User1", "User2", "User3", "User4"}},
		{Name: "ServerB", Capacity: 1, Prefs: []string{"User3", "User1", "User4", "User2"}},
	}

	matches, unmatched := StableTrafficMatch(users, servers)
	fmt.Println("案例1: 用户数 > 服务器数 (N > M)")
	fmt.Println("匹配结果:")
	for user, server := range matches {
		fmt.Printf("  %s → %s\n", user, server)
	}
	fmt.Println("未匹配用户:", unmatched)
	fmt.Printf("服务器负载: A=%d/%d, B=%d/%d\n\n",
		servers[0].Load, servers[0].Capacity,
		servers[1].Load, servers[1].Capacity)

	// 示例2：服务器数 > 用户数 (M > N)
	users = []*User{
		{Name: "User1", Prefs: []string{"ServerA", "ServerB", "ServerC"}},
		{Name: "User2", Prefs: []string{"ServerB", "ServerA", "ServerC"}},
	}

	servers = []*Server{
		{Name: "ServerA", Capacity: 1, Prefs: []string{"User1", "User2"}},
		{Name: "ServerB", Capacity: 1, Prefs: []string{"User2", "User1"}},
		{Name: "ServerC", Capacity: 1, Prefs: []string{"User1", "User2"}},
	}

	matches, unmatched = StableTrafficMatch(users, servers)
	fmt.Println("案例2: 服务器数 > 用户数 (M > N)")
	fmt.Println("匹配结果:")
	for user, server := range matches {
		fmt.Printf("  %s → %s\n", user, server)
	}
	fmt.Println("未匹配用户:", unmatched)
	fmt.Printf("服务器负载: A=%d/%d, B=%d/%d, C=%d/%d\n",
		servers[0].Load, servers[0].Capacity,
		servers[1].Load, servers[1].Capacity,
		servers[2].Load, servers[2].Capacity)
}
