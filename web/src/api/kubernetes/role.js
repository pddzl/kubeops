import service from '@/utils/request'

// 获取role list
export const getRoleList = (data) => {
  return service({
    url: '/role/getRoleList',
    method: 'post',
    data
  })
}

// 获取role编排
export const getRoleRaw = (data) => {
  return service({
    url: '/role/getRoleRaw',
    method: 'post',
    data
  })
}

// 获取role详情
export const getRoleDetail = (data) => {
  return service({
    url: '/role/getRoleDetail',
    method: 'post',
    data
  })
}

// 删除role
export const deleteRole = (data) => {
  return service({
    url: '/role/deleteRole',
    method: 'post',
    data
  })
}
