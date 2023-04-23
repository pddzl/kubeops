import { request } from "@/utils/service"

// type status = "True" | "Unknown"

export interface NodeData {
  name: string
  internalIP: string
  role: string[]
  status: string
  cpu: string
  memory: string
  creationTimestamp: string
}

export interface NodeDataPageInfo {
  list: NodeData[]
  total: number
  page: number
  pageSize: number
}

// 获取所有节点（分页）
export function getNodesApi(data: PageInfo) {
  return request<IApiResponseData<NodeDataPageInfo>>({
    url: "/k8s/node/getNodes",
    method: "post",
    data
  })
}

// export interface NodeDetail {
//   metadata: {
//     name: string
//     label: object
//     annotations: object
//     creationTimestamp: string
//     uid: string
//   }
//   spec: {
//     podCIDR: string
//     taints: object[]
//   }
//   status: {
//     conditions: object[]
//     address: object[]
//     port: {
//       kubeletEndpoint: {
//         Port: number
//       }
//     }
//     nodeInfo: object
//   }
//   allocatedResources: object
// }

// 获取某个节点详情
export function getNodeDetail(params: { name: string }) {
  return request<IApiResponseData<any>>({
    url: "/k8s/node/getNodeDetail",
    method: "get",
    params
  })
}
