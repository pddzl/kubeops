import service from '@/utils/request'

// 获取serviceAccount list
export const getServiceAccountList = (data) => {
  return service({
    url: '/serviceAccount/getServiceAccountList',
    method: 'post',
    data
  })
}

// 获取serviceAccount编排
export const getServiceAccountRaw = (data) => {
  return service({
    url: '/serviceAccount/getServiceAccountRaw',
    method: 'post',
    data
  })
}

// 获取serviceAccount in detail
export const getServiceAccountDetail = (data) => {
  return service({
    url: '/serviceAccount/getServiceAccountDetail',
    method: 'post',
    data
  })
}

// 删除serviceAccount
export const deleteServiceAccount = (data) => {
  return service({
    url: '/serviceAccount/deleteServiceAccount',
    method: 'post',
    data
  })
}
