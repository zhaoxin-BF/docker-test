package container

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"strconv"
)

func GetcContainer() {
	//apiClient, err := client.NewClientWithOpts(client.FromEnv)
	apiClient, err := client.NewClientWithOpts(client.WithVersion("1.43"))
	if err != nil {
		panic(err)
	}
	defer apiClient.Close()

	// step1: container
	// 1、list
	//containers, err := apiClient.ContainerList(context.Background(), container.ListOptions{All: true})
	//if err != nil {
	//	panic(err)
	//}
	//for _, ctr := range containers {
	//	fmt.Printf("%v\n", ctr)
	//	fmt.Printf("%v\n", ctr.ID)
	//	fmt.Printf("%v\n", ctr.Labels)
	//}

	// 2、get
	//container, err := apiClient.ContainerInspect(context.Background(), "af2032d23319bed2b88b4a82271d27d44843177ccc3b2ec4536864df71e69414")
	//if err != nil {
	//	panic(err)
	//}
	//ports := container.NetworkSettings.Ports
	//fmt.Printf("%+v\n", ports)
	//for k, v := range ports {
	//	fmt.Printf(k.Port())
	//	if k.Port() == "80" {
	//		for _, p := range v {
	//			fmt.Printf(p.HostPort)
	//		}
	//	}
	//}
	//fmt.Printf("%v\n", container.Image)

	//// 3、run
	natPort := strconv.Itoa(80) + "/tcp"
	containerConfig := &container.Config{
		Image: "nginx:latest",
		Env:   []string{"k=v"},
		ExposedPorts: map[nat.Port]struct{}{
			nat.Port(natPort): {},
		},
	}

	var portBind80 []nat.PortBinding
	portBind80 = append(portBind80, nat.PortBinding{HostPort: "11022"})

	var portBind90 []nat.PortBinding
	portBind90 = append(portBind90, nat.PortBinding{HostPort: "11023"})
	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			nat.Port(natPort): []nat.PortBinding{
				{
					HostIP:   "127.0.0.1", // 可以指定特定的主机 IP，这里使用 0.0.0.0 表示监听所有接口
					HostPort: "",          // 空字符串表示让 Docker 自动分配主机端口
				},
			},
		},
		AutoRemove: true,
		Resources: container.Resources{
			NanoCPUs:   2e9,               // 限制使用2个虚拟CPU
			Memory:     512 * 1024 * 1024, // 限制使用512MB内存
			MemorySwap: -1,                // 禁用内存交换
		},
		//RestartPolicy: container.RestartPolicy{
		//	Name: container.RestartPolicyAlways,
		//},
	}
	netConfig := &network.NetworkingConfig{}
	platform := &ocispec.Platform{}

	// create container
	ret, err := apiClient.ContainerCreate(context.Background(), containerConfig, hostConfig, netConfig, platform, "")
	if err != nil {
		panic(err)
		return
	}
	fmt.Printf("%+v\n", ret)

	err = apiClient.ContainerStart(context.Background(), ret.ID, container.StartOptions{})
	if err != nil {
		panic(err)
		return
	}

	// 4、stop
	//err = apiClient.ContainerStop(context.Background(), "", container.StopOptions{})

	// 5、remove (rm) container name 可以，container id 可以
	//err = apiClient.ContainerRemove(context.Background(), "0965132d0567bc7995396848598b124f0dd9849d6a3843eb0f91954b9e2826ed", container.RemoveOptions{
	//	RemoveVolumes: false,
	//	RemoveLinks:   false,
	//	Force:         true,
	//})
	//if err != nil {
	//	panic(err)
	//}

	// 6、kill
	//err = apiClient.ContainerKill(context.Background(), "ce4c251d403afb4e216808feed100f23232b50c47d00d23cfb3bbe52867dff33", "")
	//if err != nil {
	//	panic(err)
	//}

	// 7、restart
	//err = apiClient.ContainerRestart(context.Background(), "", container.StopOptions{})

	/////////////////////////////////////////////////////////////////////////////////////////////////////////////image
	// step2: image
	// 1、list
	//var keyValuePair []filters.KeyValuePair
	//keyValuePair = append(keyValuePair, filters.Arg("reference", "nginx:latest")) // 筛选镜像
	////keyValuePair = append(keyValuePair, filters.Arg("dangling", "true")) // 筛选悬空镜像
	//args := filters.NewArgs(keyValuePair...)
	//
	//images, err := apiClient.ImageList(context.Background(), types.ImageListOptions{
	//	Filters: args,
	//})
	//if err != nil {
	//	panic(err)
	//}
	//
	//for _, img := range images {
	//	fmt.Printf("%+v\n", img)
	//	fmt.Printf("%v\n", img.ID)
	//	//fmt.Printf("%v\n", img.Labels)
	//	//if len(img.RepoDigests) == 0 {
	//	//	fmt.Printf("%+v", img)
	//	//}
	//	//fmt.Printf("%v\n", img.RepoDigests)
	//	//
	//	//if len(img.RepoTags) == 0 {
	//	//	fmt.Printf("%+v\n", img)
	//	//}
	//	//fmt.Printf("%+v\n", img.RepoTags)
	//}

	//2、pull
	//jsonBytes, _ := json.Marshal(map[string]string{
	//	"username": "my username goes here",
	//	"password": "my password goes here",
	//})
	//registryAuth := base64.StdEncoding.EncodeToString(jsonBytes)

	//_, err = apiClient.ImagePull(context.Background(), "nginx@sha256:6db391d1c0cfb30588ba0bf72ea999404f2764febf0f1f196acd5867ac7efa7e", types.ImagePullOptions{
	//	All: false,
	//	//RegistryAuth: registryAuth,
	//})
	//if err != nil {
	//	panic(err)
	//}

	// 3、remove image
	//_, err = apiClient.ImageRemove(context.Background(), "nginx", types.ImageRemoveOptions{Force: true})
	//if err != nil {
	//	panic(err)
	//}

	//4、get image
	//image, _, err := apiClient.ImageInspectWithRaw(context.Background(), "sha256:07002dd7a3cbe09ac697570e31174acc1699701bd0626d2cf71e01623f41a10f53")
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Printf("%v", image.RepoDigests)

	//// 5、set tag
	//err = apiClient.ImageTag(context.Background(), "nginx:latest", "user_id:aaaaaa")
	//if err != nil {
	//	fmt.Printf(err.Error())
	//}
	////
	// 6、delete tag
	//_, err = apiClient.ImageRemove(context.Background(), "user_id:aaaaaa", types.ImageRemoveOptions{})
	//if err != nil {
	//	panic(err)
	//}
}
