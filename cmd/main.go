package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/antihax/optional"
	"log"
	"net/http"

	gdsclient "github.com/xu-hongyi/go-sdk/go-client-generated"
)

var token, poolId string
var volumeIds, snapshotsIds, clientIds, gatewayIds, qosIds []string
var luns []gdsclient.ClientUpdateLunProperty
var nodes []gdsclient.GatewayAddNodeProperty

func PrepareRequest() (context.Context, *gdsclient.APIClient) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	apiCtx := context.WithValue(context.Background(), gdsclient.ContextAccessToken, token)
	config := gdsclient.NewConfiguration()
	config.BasePath = "https://10.1.18.185"
	config.HTTPClient = &http.Client{Transport: tr}
	return apiCtx, gdsclient.NewAPIClient(config)
}

// Login 登录
func Login() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	config := gdsclient.NewConfiguration()
	config.BasePath = "https://10.1.18.185"
	config.HTTPClient = &http.Client{Transport: tr}
	apiClient := gdsclient.NewAPIClient(config)

	r1, _, err := apiClient.LoginApi.ApiV1LoginPost(nil, gdsclient.LoginRequestView{
		Name:     "operator",
		Password: "Admin123",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("login success, AccessToken: %s\n", r1.AccessToken)
	token = r1.AccessToken
}

// Logout 退出登录
func Logout() {
	apiCtx, apiClient := PrepareRequest()
	_, _, err := apiClient.LogoutApi.ApiV1LogoutPost(apiCtx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("logout success!!")
}

// PoolList 获取存储池列表
func PoolList() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.PoolsApi.ApiV1PoolsGet(apiCtx, &gdsclient.PoolsApiApiV1PoolsGetOpts{
		Index:       optional.NewInt32(0),
		Offset:      optional.NewInt32(50),
		Application: optional.NewString("rbd_data"),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("pool count: %d\n", r1.Total)
	for _, pool := range r1.PoolsList {
		fmt.Printf("id: %d, uuid: %s, name: %s, application: %s, healthStatus: %s\n", pool.PoolId, pool.Id, pool.Name, pool.Application, pool.HealthStatus)
	}
	poolId = r1.PoolsList[0].Id
}

// CreateVolume 创建卷
func CreateVolume() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.VolumesApi.ApiV1BlockVolumesPost(apiCtx, gdsclient.CreateVolumeRequestViews{
		PoolId: poolId,
		Qos: &gdsclient.QosPropertyView{
			BpsBurst:  10485760,
			IopsLimit: 1000000,
			IopsBurst: 1000000,
			BpsLimit:  10485760,
		},
		VerifyEnabled: true,
		VolumeSize:    10737418240,
		IoPriority:    "default",
		Name:          "csi_volume_1",
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("create success: %v\n", r1)
}

// GetVolumeList 获取卷列表
func GetVolumeList() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.VolumesApi.ApiV1BlockVolumesGet(apiCtx, &gdsclient.VolumesApiApiV1BlockVolumesGetOpts{
		Index:            optional.NewInt32(0),
		Offset:           optional.NewInt32(50),
		SortBy:           optional.EmptyString(),
		OrderBy:          optional.NewString("desc"),
		Name:             optional.EmptyString(),
		Id:               optional.EmptyString(),
		PoolName:         optional.EmptyString(),
		PoolId:           optional.EmptyString(),
		RunStatus:        optional.EmptyString(),
		CapHealthStatus:  optional.EmptyString(),
		PerfHealthStatus: optional.EmptyString(),
		SyncStatus:       optional.EmptyString(),
		QosStatus:        optional.EmptyBool(),
		VerifyEnabled:    optional.EmptyBool(),
		DataProtection:   optional.EmptyBool(),
		CgName:           optional.EmptyString(),
		CgId:             optional.EmptyString(),
		VolumeType:       optional.EmptyString(),
		OwnerPlatform:    optional.EmptyString(),
		IoPriority:       optional.EmptyString(),
		SnapNum:          optional.EmptyInt32(),
		ClientNum:        optional.EmptyInt32(),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("total volumes:%d, cur volume: %d, page: %d\n", r1.Total, r1.Count, r1.Index)
	//fmt.Printf("%v\n", r2)
	for _, volume := range r1.Volumes {
		fmt.Printf("volume name: %s, id: %s\n", volume.Name, volume.Id)
		volumeIds = append(volumeIds, volume.Id)
	}

}

// BatchCreationVolume 批量创建卷
func BatchCreationVolume() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.VolumesApi.ApiV1BlockVolumesBatchCreationPost(apiCtx, gdsclient.BatchCreateVolumeRequestView{
		Count:  6,
		PoolId: poolId,
		Qos: &gdsclient.QosPropertyView{
			BpsBurst:  10485760,
			IopsLimit: 1000000,
			IopsBurst: 1000000,
			BpsLimit:  10485760,
		},
		StartNo:       11,
		VerifyEnabled: false,
		VolumeSize:    10737418240,
		IoPriority:    "default",
		Prefix:        "vst_volume_",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", r1)
}

// GetVolumeDetail 获取卷详情
func GetVolumeDetail() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.VolumesApi.ApiV1BlockVolumesVolumeIdGet(apiCtx, volumeIds[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", r1)
}

// DeleteVolume 根据uuid删除卷
func DeleteVolume() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.VolumesApi.ApiV1BlockVolumesVolumeIdDelete(apiCtx, volumeIds[0], &gdsclient.VolumesApiApiV1BlockVolumesVolumeIdDeleteOpts{
		Force: optional.NewBool(true),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", r1)
}

// UpdateVolume 更新卷
func UpdateVolume() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.VolumesApi.ApiV1BlockVolumesVolumeIdPut(apiCtx, gdsclient.UpdateVolumeRequestView{
		//VerifyEnabled: true,
		IoPriority: "priority",
		//}, volumeIds[0])
	}, "ba3758b5-90c5-43aa-8d2c-4d8f9b2beaba")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", r1)
}

// UpdateVolumeVerifyEnabled 修改卷数据校验
func UpdateVolumeVerifyEnabled() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.VolumesApi.ApiV1BlockVolumesVolumeIdVerifyEnabledPut(apiCtx, gdsclient.UpdateVolumeVerifyEnableRequestView{
		VerifyEnabled: true,
	}, volumeIds[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", r1)
}

// UpdateVolumeIoPriority 修改卷数据优先级
func UpdateVolumeIoPriority() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.VolumesApi.ApiV1BlockVolumesVolumeIdIoPriorityPut(apiCtx, gdsclient.UpdateVolumeIoPriorityRequestView{
		IoPriority: "priority",
	}, volumeIds[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", r1)
}

// UpdateVolumeName 修改卷名称
func UpdateVolumeName() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.VolumesApi.ApiV1BlockVolumesVolumeIdNamePut(apiCtx, gdsclient.UpdateVolumeNameRequestView{
		Name: "update_11",
	}, volumeIds[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", r1)
}

// BatchDeleteVolume 批量删除卷
func BatchDeleteVolume() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.VolumesApi.ApiV1BlockVolumesBatchDeletionPost(apiCtx, gdsclient.VolumeBatchDeleteRequestView{
		Force:        false,
		VolumeIdList: []string{volumeIds[1], volumeIds[2], volumeIds[3]},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", r1)
}

// VolumeFlatten 将链接克隆卷从快照树上断开关系链
// todo 未实现
func VolumeFlatten() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.VolumesApi.ApiV1BlockVolumesVolumeIdFlattenPut(apiCtx, volumeIds[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", r1)
}

// ExpandVolume 卷扩容
func ExpandVolume() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.VolumesApi.ApiV1BlockVolumesVolumeIdExpandPut(apiCtx, gdsclient.VolumeExpandRequestView{
		Size: 15032385536,
	}, volumeIds[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", r1)
}

// ShrinkVolume 卷缩容
func ShrinkVolume() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.VolumesApi.ApiV1BlockVolumesVolumeIdShrinkPut(apiCtx, gdsclient.VolumeShrinkRequestView{
		Size: 10737418240,
	}, volumeIds[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", r1)
}

// UpdateVolumeQos 修改卷Qos
func UpdateVolumeQos() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.VolumesApi.ApiV1BlockVolumesVolumeIdQosPut(apiCtx, gdsclient.SetVolumeQosRequestView{
		Qos: &gdsclient.QosPropertyView{
			BpsLimit:  10,
			IopsLimit: 100,
			BpsBurst:  20,
			IopsBurst: 150,
		},
	}, volumeIds[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", r1)
}

// GetClientInVolume 获取卷挂载的客户端
func GetClientInVolume() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.VolumesApi.ApiV1BlockVolumesVolumeIdClientsGet(apiCtx, "2d9ebee3-a2d8-4fd3-8a86-3d1b33db053a", &gdsclient.VolumesApiApiV1BlockVolumesVolumeIdClientsGetOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", r1)
}

// CreateSnapshots 创建快照
func CreateSnapshots() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.SnapsApi.ApiV1BlockSnapsPost(apiCtx, gdsclient.CreateSnapshotRequestView{
		Name:        "csi_volume_snapshot5",
		VolumeId:    volumeIds[0],
		Description: "snapshot",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", r1)
}

// GetSnapList 获取快照列表
func GetSnapList() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.SnapsApi.ApiV1BlockSnapsGet(apiCtx, &gdsclient.SnapsApiApiV1BlockSnapsGetOpts{
		Index:           optional.NewInt32(0),
		Offset:          optional.NewInt32(0),
		SortBy:          optional.NewString("name"),
		OrderBy:         optional.NewString("desc"),
		Name:            optional.EmptyString(),
		Id:              optional.EmptyString(),
		PoolName:        optional.EmptyString(),
		PoolId:          optional.EmptyString(),
		VolumeId:        optional.EmptyString(),
		VolumeName:      optional.EmptyString(),
		CgSnapId:        optional.EmptyString(),
		CgSnapName:      optional.EmptyString(),
		CloneVolNum:     optional.EmptyInt32(),
		RunStatus:       optional.EmptyString(),
		CapHealthStatus: optional.EmptyString(),
		SyncStatus:      optional.EmptyString(),
		CreateFrom:      optional.EmptyString(),
		Description:     optional.EmptyString(),
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, snap := range r1.Snaps {
		fmt.Printf("snapshot name: %s, uuid: %s\n", snap.Name, snap.Id)
		snapshotsIds = append(snapshotsIds, snap.Id)
	}

}

// DeleteSnapshots 删除快照
func DeleteSnapshots() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.SnapsApi.ApiV1BlockSnapsSnapshotIdDelete(apiCtx, snapshotsIds[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", r1)
}

// UpdateSnapshots 修改快照名称，描述信息
func UpdateSnapshots() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.SnapsApi.ApiV1BlockSnapsSnapshotIdPut(apiCtx, gdsclient.RenameSnapRequestView{
		Name:        "aaaa",
		Description: "test rename",
	}, snapshotsIds[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", r1)
}

// GetSnapshotsDetail 获取快照详情
func GetSnapshotsDetail() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.SnapsApi.ApiV1BlockSnapsSnapshotIdGet(apiCtx, snapshotsIds[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", r1)
}

// CloneSnapshots 克隆快照
func CloneSnapshots() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.SnapsApi.ApiV1BlockSnapsSnapshotIdClonePost(apiCtx, gdsclient.SnapCloneRequestView{
		IoPriority:    "priority",
		VerifyEnabled: false,
		CloneType:     "independent", //dependent
		Qos: &gdsclient.QosPropertyView{
			BpsLimit:  0,
			IopsLimit: 0,
			BpsBurst:  0,
			IopsBurst: 0,
		},
		Name:   "clone2",
		PoolId: poolId,
	}, snapshotsIds[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", r1)
}

// BatchDeleteSnapshots 批量删除快照
func BatchDeleteSnapshots() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.SnapsApi.ApiV1BlockSnapsBatchDeletionPost(apiCtx, gdsclient.DeleteBatchSnapshotRequestView{
		Snaps: []string{snapshotsIds[1], snapshotsIds[2]},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", r1)
}

// CreateQos 创建Qos
func CreateQos() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.QosApi.ApiV1BlockQosPost(apiCtx, gdsclient.CreateQosRequestView{
		BpsLimit:  52428800,
		IopsBurst: 800,
		Name:      "csi_qos_4",
		IopsLimit: 500,
		BpsBurst:  83886080,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", r1)
}

// GetQosList 获取qos列表
func GetQosList() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.QosApi.ApiV1BlockQosGet(apiCtx, &gdsclient.QosApiApiV1BlockQosGetOpts{
		Index:     optional.NewInt32(0),
		Offset:    optional.NewInt32(0),
		SortBy:    optional.NewString("name"),
		OrderBy:   optional.NewString("desc"),
		Name:      optional.EmptyString(),
		Id:        optional.EmptyString(),
		IopsLimit: optional.EmptyInt64(),
		IopsBurst: optional.EmptyInt64(),
		BpsLimit:  optional.EmptyInt64(),
		BpsBurst:  optional.EmptyInt64(),
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, qos := range r1.Qoss {
		fmt.Printf("qos name: %s, qos id: %s\n", qos.Name, qos.Id)
		qosIds = append(qosIds, qos.Id)
	}
}

// DeleteQos 删除qos
func DeleteQos() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.QosApi.ApiV1BlockQosQosPolicyIdDelete(apiCtx, qosIds[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", r1)
}

// UpdateQos 修改qos策略
func UpdateQos() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.QosApi.ApiV1BlockQosQosPolicyIdPut(apiCtx, gdsclient.UpdateQosRequestView{
		BpsLimit:  524288000,
		IopsBurst: 100,
		Name:      "csi_qos_test",
		IopsLimit: 50,
		BpsBurst:  838860800,
	}, qosIds[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", r1)
}

// GetQosDetail 获取快照详情
func GetQosDetail() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.QosApi.ApiV1BlockQosQosPolicyIdGet(apiCtx, qosIds[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", r1)
}

// CreateClient 创建块网关
func CreateClient() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.ClientApi.ApiV1BlockIscsiClientsPost(apiCtx, gdsclient.CreateIscsiClientRequestView{
		Name:         "ucsi_client_5",
		IsChap:       false,
		Description:  "csi",
		Hosts:        []string{"iqn.2021-06.com.domain:796b8192hh", "iqn.2021-06.com.domain:796b81929jj"},
		ChapUsername: "chapchap",
		ChapPassword: "chapchapchapchap",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", r1)
}

// GetClientList 获取客户端列表
func GetClientList() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.ClientApi.ApiV1BlockIscsiClientsGet(apiCtx, &gdsclient.ClientApiApiV1BlockIscsiClientsGetOpts{
		Name:        optional.EmptyString(),
		Id:          optional.EmptyString(),
		Description: optional.EmptyString(),
		RunStatus:   optional.EmptyString(),
		IsChap:      optional.EmptyBool(),
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, client := range r1.Clients {
		fmt.Printf("client name: %s, client id: %s\n", client.Name, client.Id)
		clientIds = append(clientIds, client.Id)
	}
}

// GetClientDetail 获取客户端详情
func GetClientDetail() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.ClientApi.ApiV1BlockIscsiClientsClientIdGet(apiCtx, clientIds[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", r1)
}

// DeleteClient 删除客户端
func DeleteClient() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.ClientApi.ApiV1BlockIscsiClientsClientIdDelete(apiCtx, clientIds[0], &gdsclient.ClientApiApiV1BlockIscsiClientsClientIdDeleteOpts{
		Force: optional.NewBool(true),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", r1)
}

// UpdateClient 修改客户端
func UpdateClient() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.ClientApi.ApiV1BlockIscsiClientsClientIdPut(apiCtx, gdsclient.ClientUpdateRequestView{
		ChapPassword: "eeeeeeeeeeeeee",
		ChapUsername: "eeeeeeeeeeeeee",
		IsChap:       true,
	}, "a70f7a0e-cd69-4077-b3c3-e7c88c23bc15")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", r1)
}

// UpdateClientHost 修改客户端host
func UpdateClientHost() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.ClientApi.ApiV1BlockIscsiClientsClientIdHostsPut(apiCtx, gdsclient.ClientUpdateHostsRequestView{
		Hosts:  []string{"iqn.2021-05.com.domain:796b81929e6", "iqn.2021-05.com.domain:796b81929e7"},
		Action: "add",
	}, clientIds[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", r1)
}

// GetClientHostList 获取客户端host列表
func GetClientHostList() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.ClientApi.ApiV1BlockIscsiClientsClientIdHostsGet(apiCtx, clientIds[0], &gdsclient.ClientApiApiV1BlockIscsiClientsClientIdHostsGetOpts{
		Node:      optional.EmptyString(),
		UseStatus: optional.EmptyString(),
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, node := range r1.Nodes {
		fmt.Printf("node id: %s, node name: %s\n", node.Id, node.Node)
	}
}

// AddLunToClient 挂载/卸载卷
func AddLunToClient() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.ClientApi.ApiV1BlockIscsiClientsClientIdLunsPut(apiCtx, gdsclient.ClientUpdateLunsRequestView{
		Luns:   luns[:3],
		Action: "add",
	}, clientIds[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", r1)
}

// GetLunInClient 获取客户端中挂在的卷
func GetLunInClient() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.ClientApi.ApiV1BlockIscsiClientsClientIdLunsGet(apiCtx, clientIds[1], &gdsclient.ClientApiApiV1BlockIscsiClientsClientIdLunsGetOpts{})
	if err != nil {
		log.Fatal(err)
	}
	for _, lun := range r1.Luns {
		fmt.Printf("client id: %s lun name: %s, lun type: %s\n", clientIds[1], lun.VolumeName, lun.VolumeType)
	}
}

// GetLuns 获取可以向客户端挂在的卷
func GetLuns() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.ClientApi.ApiV1BlockIscsiClientsLunsGet(apiCtx, &gdsclient.ClientApiApiV1BlockIscsiClientsLunsGetOpts{
		Name:     optional.EmptyString(),
		PoolName: optional.NewString("block"),
		//LunType:  optional.NewString("normal_volume"),
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, lun := range r1.Luns {
		luns = append(luns, gdsclient.ClientUpdateLunProperty{
			LunId:      lun.Id,
			IsReadonly: false,
			LunType:    lun.LunType,
		})
		fmt.Printf("client can use lun: lun name: %s, lun id: %s, lun type: %s\n", lun.Name, lun.Id, lun.LunType)
	}
}

// GetGatewayInClient 获取客户端关联的网关列表
func GetGatewayInClient() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.ClientApi.ApiV1BlockIscsiClientsClientIdGatewaysGet(apiCtx, clientIds[0], &gdsclient.ClientApiApiV1BlockIscsiClientsClientIdGatewaysGetOpts{
		Name:        optional.EmptyString(),
		Iqn:         optional.EmptyString(),
		Port:        optional.EmptyInt32(),
		RunStatus:   optional.EmptyString(),
		Description: optional.EmptyString(),
		IsChap:      optional.EmptyBool(),
		IsEnabled:   optional.EmptyBool(),
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, gateway := range r1.Gateways {
		fmt.Printf("client gatewayid: %s, gateway name: %s, gateway id: %s\n", clientIds[0], gateway.Name, gateway.Id)
	}
}

// GetCanUseNodes 获取可用的节点
func GetCanUseNodes() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.GatewayApi.ApiV1BlockIscsiGatewaysNodesGet(apiCtx, &gdsclient.GatewayApiApiV1BlockIscsiGatewaysNodesGetOpts{})
	if err != nil {
		log.Fatal(err)
	}
	for _, node := range r1.Nodes {
		nodes = append(nodes, gdsclient.GatewayAddNodeProperty{
			ServerId:  node.ServerId,
			GatewayIp: node.GatewayIp,
		})
		fmt.Printf("node id: %s, node ip: %s\n", node.ServerId, node.GatewayIp)
	}

}

// CreateGateway 创建网关
func CreateGateway() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.GatewayApi.ApiV1BlockIscsiGatewaysPost(apiCtx, gdsclient.CreateIscsiGwRequestView{
		Iqn:          "iqn.2021-06.com.expontech.gds:csi5",
		GatewayType:  "iSCSI",
		Name:         "tcsi_gateway_5",
		Port:         3260,
		IsChap:       false,
		Nodes:        nodes,
		ChapUsername: "sfdssfds",
		ChapPassword: "sdfsdfsfdsfd",
		Description:  "",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", r1)
}

// GetGatewayList 获取网关列表
func GetGatewayList() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.GatewayApi.ApiV1BlockIscsiGatewaysGet(apiCtx, &gdsclient.GatewayApiApiV1BlockIscsiGatewaysGetOpts{
		Id:           optional.EmptyString(),
		Name:         optional.EmptyString(),
		Iqn:          optional.EmptyString(),
		Port:         optional.EmptyInt32(),
		RunStatus:    optional.EmptyString(),
		HealthStatus: optional.EmptyString(),
		IsEnabled:    optional.EmptyBool(),
		IsChap:       optional.EmptyBool(),
		ChapUsername: optional.EmptyString(),
		Description:  optional.EmptyString(),
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, gateway := range r1.Gateways {
		fmt.Printf("gateway name: %s, gateway id: %s\n", gateway.Name, gateway.Id)
		gatewayIds = append(gatewayIds, gateway.Id)
	}
}

// DeleteGateway 删除网关
func DeleteGateway() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.GatewayApi.ApiV1BlockIscsiGatewaysGatewayIdDelete(apiCtx, gatewayIds[0], &gdsclient.GatewayApiApiV1BlockIscsiGatewaysGatewayIdDeleteOpts{
		Force: optional.NewString("true"),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", r1)
}

// GetGatewayDetail 获取块网关详情
func GetGatewayDetail() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.GatewayApi.ApiV1BlockIscsiGatewaysGatewayIdGet(apiCtx, gatewayIds[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", r1)
}

// UpdateGateway 更新快网关
func UpdateGateway() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.GatewayApi.ApiV1BlockIscsiGatewaysGatewayIdPut(apiCtx, gdsclient.UpdateBlockGatewayRequestView{
		ChapPassword: "ssssssssssss",
		//Name:         "csi_gateway_test",
		IsChap:       true,
		ChapUsername: "ssssssss",
		//Description:  "",
	}, gatewayIds[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", r1)
}

// GetNodesInGateway 获取块网关中的节点列表
func GetNodesInGateway() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.GatewayApi.ApiV1BlockIscsiGatewaysGatewayIdNodesGet(apiCtx, gatewayIds[0], &gdsclient.GatewayApiApiV1BlockIscsiGatewaysGatewayIdNodesGetOpts{
		Name:         optional.EmptyString(),
		NodeId:       optional.EmptyString(),
		ServerId:     optional.EmptyString(),
		IscsiGwId:    optional.EmptyString(),
		ManagerIp:    optional.EmptyString(),
		GatewayIp:    optional.EmptyString(),
		HealthStatus: optional.EmptyString(),
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, node := range r1.Nodes {
		fmt.Printf("server id: %s, server gatewayIp: %s\n", node.ServerId, node.GatewayIp)
	}
}

// UpdateNodesInGateway 添加删除网关中的节点
func UpdateNodesInGateway() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.GatewayApi.ApiV1BlockIscsiGatewaysGatewayIdNodesPut(apiCtx, gdsclient.UpdateBlockGatewayNodesRequestView{
		Nodes:  nodes[3:],
		Action: "remove",
	}, gatewayIds[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", r1)
}

// GetCanUseClient 获取可向网关中添加的客户端列表
func GetCanUseClient() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.GatewayApi.ApiV1BlockIscsiGatewaysClientsGet(apiCtx, &gdsclient.GatewayApiApiV1BlockIscsiGatewaysClientsGetOpts{
		Id:        optional.EmptyString(),
		Name:      optional.EmptyString(),
		RunStatus: optional.EmptyString(),
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, client := range r1.Clients {
		fmt.Printf("client name: %s, client id: %s\n", client.Name, client.Id)
		clientIds = append(clientIds, client.Id)
	}
}

// UpdateGatewayInClient 更新网关中的客户端
func UpdateGatewayInClient() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.GatewayApi.ApiV1BlockIscsiGatewaysGatewayIdClientsPut(apiCtx, gdsclient.UpdateBlockGatewayClientsRequestView{
		Clients: clientIds,
		Action:  "add",
	}, gatewayIds[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", r1)
}

// GetClientInGateway 获取网关中的客户端
func GetClientInGateway() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.GatewayApi.ApiV1BlockIscsiGatewaysGatewayIdClientsGet(apiCtx, gatewayIds[0], &gdsclient.GatewayApiApiV1BlockIscsiGatewaysGatewayIdClientsGetOpts{
		Name:      optional.EmptyString(),
		RunStatus: optional.EmptyString(),
		IscsiGwId: optional.EmptyString(),
		IsChap:    optional.EmptyBool(),
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, client := range r1.Clients {
		fmt.Printf("gateway client name: %s, client id: %s\n", client.Name, client.Id)
	}
}

// GetLunsInGateway 获取块网关中的卷
func GetLunsInGateway() {
	apiCtx, apiClient := PrepareRequest()
	r1, _, err := apiClient.GatewayApi.ApiV1BlockIscsiGatewaysGatewayIdLunsGet(apiCtx, gatewayIds[0], &gdsclient.GatewayApiApiV1BlockIscsiGatewaysGatewayIdLunsGetOpts{
		Name:             optional.String{},
		VolumeName:       optional.String{},
		Description:      optional.String{},
		PoolName:         optional.String{},
		VolumeType:       optional.String{},
		CapHealthStatus:  optional.String{},
		PerfHealthStatus: optional.String{},
		RunStatus:        optional.String{},
		SyncStatus:       optional.String{},
		CreateFrom:       optional.String{},
		UseStatus:        optional.String{},
		IsReadonly:       optional.Bool{},
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, lun := range r1.Luns {
		fmt.Printf("lun name: %s\n", lun.VolumeName)
	}
}

func main() {
	Login()
	//PoolList()
	//CreateVolume()
	//BatchCreationVolume()
	GetVolumeList()
	GetClientInVolume()
	//GetVolumeDetail()
	//DeleteVolume()
	//BatchDeleteVolume()
	//UpdateVolumeIoPriority()
	//UpdateVolumeVerifyEnabled()
	//UpdateVolumeName()
	//ExpandVolume()
	//ShrinkVolume()
	//UpdateVolumeQos()
	//CreateSnapshots()
	//GetSnapList()
	//DeleteSnapshots()
	//UpdateSnapshots()
	//GetSnapshotsDetail()
	//CloneSnapshots()
	//VolumeFlatten()
	//BatchDeleteSnapshots()
	//CreateQos()
	//GetQosList()
	//DeleteQos()
	//UpdateQos()
	//GetQosDetail()
	//GetCanUseNodes()
	//CreateGateway()
	//GetGatewayList()
	//GetGatewayDetail()
	//DeleteGateway()
	//UpdateGateway()
	//GetNodesInGateway()
	//UpdateNodesInGateway()
	//GetCanUseClient()
	//UpdateGatewayInClient()
	//GetClientInGateway()
	//GetLunsInGateway()
	//CreateClient()
	//GetClientList()
	//GetClientDetail()
	//DeleteClient()
	//UpdateClient()
	//UpdateClientHost()
	//GetClientHostList()
	//GetLuns()
	//AddLunToClient()
	//GetLunInClient()
	//GetGatewayInClient()
	//Logout()
}
