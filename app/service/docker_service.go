package service

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

type DockerPsStatus struct {
	Name   string
	Status string
	Path   string
}

var dockerPath = "/opt/docker/"

func GetDockerComposePs() []DockerPsStatus {

	// docker项目的根目录

	// 看下有哪些docker服务
	cmd := exec.Command("bash", "-c", fmt.Sprintf("cd %s && ls", dockerPath))
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("cmd.Run() failed with %s\n", err)
	}
	//fmt.Printf("Name:\n%s\n", string(out))
	outSlice := strings.Split(string(out), "\n")

	// 定义结果对象
	result := []DockerPsStatus{}

	for _, name := range outSlice {

		if name == "" {
			continue
		}

		path := dockerPath + name + "/"
		//fmt.Println("xxxxxx" + path)

		// 当前容器的状态
		cmd := exec.Command("bash", "-c", fmt.Sprintf("cd %s && docker-compose ps", path))
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("cmd.Run() failed with %s\n", err)
		}
		//fmt.Printf("combined out:\n%s\n", string(out))

		outSlice := strings.Split(string(out), "\n")

		// 返回结果异常
		if len(outSlice) <= 1 {
			return nil
		}

		statusLine := outSlice[1]

		// 匹配返回的结果，找到z
		re := regexp.MustCompile(`^(.*?)\s+"(.*?)"\s+(.*?)\s+(.*?)\s+(.*?)$`)
		match := re.FindStringSubmatch(string(statusLine))
		if len(match) > 4 {
			//fmt.Println(match[1])
			//fmt.Println(match[2])
			//fmt.Println(match[3])
			status := match[4]
			//fmt.Println(match[4])

			result = append(result, DockerPsStatus{Name: name, Status: status, Path: path})
		} else {

			// 服务没启动g
			result = append(result, DockerPsStatus{Name: name, Status: "stopping", Path: path})
		}
	}

	return result
}

func DockerComposeUp(name string) string {

	path := dockerPath + name + "/"

	// 当前容器的状态
	cmd := exec.Command("bash", "-c", fmt.Sprintf("cd %s && docker-compose up -d", path))
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("cmd.Run() failed with %s\n", err)
	}

	return string(out)
}

func DockerComposeDown(name string) string {

	path := dockerPath + name + "/"

	// 当前容器的状态
	cmd := exec.Command("bash", "-c", fmt.Sprintf("cd %s && docker-compose down", path))
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("cmd.Run() failed with %s\n", err)
	}

	return string(out)
}
