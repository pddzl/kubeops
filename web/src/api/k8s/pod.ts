import { request } from "@/utils/service"

export interface PodData {
  name: string
  namespace: string
  status: string
  node: string
  creationTimestamp: string
}

interface PodListData extends PageInfo {
  list: PodData[]
  total: number
}

interface reqData extends PageInfo {
  namespace: string
}

// 获取所有pod（分页）
export function getPodsApi(data: reqData) {
  return request<IApiResponseData<PodListData>>({
    url: "/k8s/pod/getPods",
    method: "post",
    data
  })
}
