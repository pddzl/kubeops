import { request } from "@/utils/service"

interface namespacePageInfo extends PageInfo {
  list: any[]
  total: number
}

export function getNamespacesApi(data: PageInfo) {
  return request<IApiResponseData<namespacePageInfo>>({
    url: "/k8s/namespace/getNamespaces",
    method: "post",
    data
  })
}

// 获取某个命名空间详情
export function getNamespaceDetailApi(params: { name: string }) {
  return request<IApiResponseData<any>>({
    url: "/k8s/namespace/getNamespaceDetail",
    method: "get",
    params
  })
}

// 删除命名空间
// 获取某个命名空间详情
export function deleteNamespaceApi(data: { name: string }) {
  return request<IApiResponseData<null>>({
    url: "/k8s/namespace/deleteNamespace",
    method: "post",
    data
  })
}
