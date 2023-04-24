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
