import service from '@/utils/request'

// 获取clusterRoleBinding list
export const getClusterRoleBindingList = (data) => {
  return service({
    url: '/clusterRoleBinding/getClusterRoleBindingList',
    method: 'post',
    data
  })
}

// 获取clusterRoleBinding编排
export const getClusterRoleBindingRaw = (data) => {
  return service({
    url: '/clusterRoleBinding/getClusterRoleBindingRaw',
    method: 'post',
    data
  })
}

// 获取clusterRoleBinding详情
export const getClusterRoleBindingDetail = (data) => {
  return service({
    url: '/clusterRoleBinding/getClusterRoleBindingDetail',
    method: 'post',
    data
  })
}
