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
