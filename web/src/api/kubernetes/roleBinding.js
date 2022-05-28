import service from '@/utils/request'

// 获取roleBinding list
export const getRoleBindingList = (data) => {
  return service({
    url: '/roleBinding/getRoleBindingList',
    method: 'post',
    data
  })
}

// 获取roleBinding编排
export const getRoleBindingRaw = (data) => {
  return service({
    url: '/roleBinding/getRoleBindingRaw',
    method: 'post',
    data
  })
}

// 获取roleBinding详情
export const getRoleBindingDetail = (data) => {
  return service({
    url: '/roleBinding/getRoleBindingDetail',
    method: 'post',
    data
  })
}

// 删除roleBinding
export const deleteRoleBinding = (data) => {
  return service({
    url: '/roleBinding/deleteRoleBinding',
    method: 'post',
    data
  })
}
