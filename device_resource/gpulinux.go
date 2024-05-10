package device_resource

//type GpuInfo struct {
//}
//
//type GPUCollector struct {
//	GPU GpuInfo
//}
//
//func NewGPUCollector() *GPUCollector {
//	return &GPUCollector{}
//}
//
//func (g *GPUCollector) GetGPUInfo() (*GpuInfo, error) {
//	ret := nvml.Init()
//	if ret != nvml.SUCCESS {
//		log.Fatalf("Unable to initialize NVML: %v", nvml.ErrorString(ret))
//	}
//	defer func() {
//		ret := nvml.Shutdown()
//		if ret != nvml.SUCCESS {
//			log.Fatalf("Unable to shutdown NVML: %v", nvml.ErrorString(ret))
//		}
//	}()
//
//	count, ret := nvml.DeviceGetCount()
//	if ret != nvml.SUCCESS {
//		log.Fatalf("Unable to get device count: %v", nvml.ErrorString(ret))
//	}
//
//	for i := 0; i < count; i++ {
//		device, ret := nvml.DeviceGetHandleByIndex(i)
//		if ret != nvml.SUCCESS {
//			log.Fatalf("Unable to get device at index %d: %v", i, nvml.ErrorString(ret))
//		}
//
//		uuid, ret := device.GetUUID()
//		if ret != nvml.SUCCESS {
//			log.Fatalf("Unable to get uuid of device at index %d: %v", i, nvml.ErrorString(ret))
//		}
//
//		fmt.Printf("%v\n", uuid)
//	}
//	return &g.GPU, nil
//}
