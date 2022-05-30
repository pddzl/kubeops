import service from '@/utils/request'

// 获取namespace list
export const getNamespaceList = (data) => {
  return service({
    url: '/namespace/getNamespaceList',
    method: 'post',
    data
  })
}

// 获取namespace list(only name)
export const getNamespaceOnlyName = (data) => {
  return service({
    url: '/namespace/getNamespaceOnlyName',
    method: 'get'
  })
}

// 获取namespace detail
export const getNamespaceDetail = (data) => {
  return service({
    url: '/namespace/getNamespaceDetail',
    method: 'post',
    data
  })
}

// 获取namespace detail
export const getNamespaceRaw = (data) => {
  return service({
    url: '/namespace/getNamespaceRaw',
    method: 'post',
    data
  })
}

// 删除namespace
export const deleteNamespace = (data) => {
  return service({
    url: '/namespace/deleteNamespace',
    method: 'post',
    data
  })
}

// 获取namespace状态
export const getNamespaceStatus = () => {
  return service({
    url: '/namespace/getNamespaceStatus',
    method: 'get'
  })
}
